package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
	"github.com/filecoin-project/lotus/chain/types"

	"github.com/filecoin-project/oni/tvx/lotus"
	"github.com/filecoin-project/oni/tvx/schema"
	"github.com/filecoin-project/oni/tvx/state"
)

var extractMsgFlags struct {
	cid  string
	file string
}

var extractMsgCmd = &cli.Command{
	Name:        "extract-message",
	Description: "generate a message-class test vector by extracting it from a network",
	Action:      runExtractMsg,
	Flags: []cli.Flag{
		&apiFlag,
		&cli.StringFlag{
			Name:        "cid",
			Usage:       "message CID to generate test vector from",
			Required:    true,
			Destination: &extractMsgFlags.cid,
		},
		&cli.StringFlag{
			Name:        "file",
			Usage:       "output file",
			Required:    true,
			Destination: &extractMsgFlags.file,
		},
	},
}

func runExtractMsg(c *cli.Context) error {
	// LOTUS_DISABLE_VM_BUF disables what's called "VM state tree buffering",
	// which stashes write operations in a BufferedBlockstore
	// (https://github.com/filecoin-project/lotus/blob/b7a4dbb07fd8332b4492313a617e3458f8003b2a/lib/bufbstore/buf_bstore.go#L21)
	// such that they're not written until the VM is actually flushed.
	//
	// For some reason, the standard behaviour was not working for me (raulk),
	// and disabling it (such that the state transformations are written immediately
	// to the blockstore) worked.
	_ = os.Setenv("LOTUS_DISABLE_VM_BUF", "iknowitsabadidea")

	ctx := context.Background()

	// get the output file.
	if extractMsgFlags.file == "" {
		return fmt.Errorf("output file required")
	}

	mid, err := cid.Decode(extractMsgFlags.cid)
	if err != nil {
		return err
	}

	// Make the client.
	api, err := makeClient(c)
	if err != nil {
		return err
	}

	// locate the message.
	msgInfo, err := api.StateSearchMsg(ctx, mid)
	if err != nil {
		return fmt.Errorf("failed to locate message: %w", err)
	}

	// Extract the serialized message.
	msg, err := api.ChainGetMessage(ctx, mid)
	if err != nil {
		return err
	}

	// create a read through store that uses ChainGetObject to fetch unknown CIDs.
	pst := state.NewProxyingStore(ctx, api)

	g := state.NewSurgeon(ctx, api, pst)

	// Get actors accessed by message.
	retain, err := g.GetAccessedActors(ctx, api, mid)
	if err != nil {
		return err
	}

	retain = append(retain, builtin.RewardActorAddr)
	retain = append(retain, builtin.BurntFundsActorAddr)

	fmt.Println("accessed actors:")
	for _, k := range retain {
		fmt.Println("\t", k.String())
	}

	// get the tipset on which this message was "executed".
	// https://github.com/filecoin-project/lotus/issues/2847
	execTs, err := api.ChainGetTipSet(ctx, msgInfo.TipSet)
	if err != nil {
		return err
	}

	// get the previous tipset, on which this message was mined.
	includedTs, err := api.ChainGetTipSet(ctx, execTs.Parents())
	if err != nil {
		return err
	}

	neededPrecursorMsgs := make([]*types.Message, 0)
	for _, b := range includedTs.Blocks() {
		messages, err := api.ChainGetBlockMessages(ctx, b.Cid())
		if err != nil {
			return err
		}
		for _, other := range messages.BlsMessages {
			if other.Cid() == mid {
				break
			}
			if other.From == msg.From && other.Nonce < msg.Nonce {
				included := false
				for _, o := range neededPrecursorMsgs {
					if o.Cid() == other.Cid() {
						included = true
					}
				}
				if !included {
					neededPrecursorMsgs = append(neededPrecursorMsgs, other)
				}
			}
		}
		for _, m := range messages.SecpkMessages {
			if m.Message.Cid() == mid {
				break
			}
			if m.Message.From == msg.From && m.Message.Nonce < msg.Nonce {
				included := false
				for _, o := range neededPrecursorMsgs {
					if o.Cid() == m.Message.Cid() {
						included = true
					}
				}
				if !included {
					neededPrecursorMsgs = append(neededPrecursorMsgs, &m.Message)
				}	
			}
		}
	}

	fmt.Println("getting the _before_ filtered state tree")
	tree, err := g.GetStateTreeRootFromTipset(includedTs.Parents())

	driver := lotus.NewDriver(ctx)

	for _, pm := range neededPrecursorMsgs {
		_, tree, err = driver.ExecuteMessage(pm, tree, pst.Blockstore, execTs.Height())
		if err != nil {
			return fmt.Errorf("Failed to execute preceding message: %w", err)
		}
	}

	preroot, err := g.GetMaskedStateTree(tree, retain)
	if err != nil {
		return err
	}

	_, postroot, err := driver.ExecuteMessage(msg, preroot, pst.Blockstore, execTs.Height())
	if err != nil {
		return fmt.Errorf("failed to execute message: %w", err)
	}

	msgBytes, err := msg.Serialize()
	if err != nil {
		return err
	}

	out := new(bytes.Buffer)
	gw := gzip.NewWriter(out)
	if err := g.WriteCAR(gw, preroot, postroot); err != nil {
		return err
	}
	if err = gw.Flush(); err != nil {
		return err
	}
	if err = gw.Close(); err != nil {
		return err
	}

	version, err := api.Version(ctx)
	if err != nil {
		return err
	}

	// Write out the test vector.
	vector := schema.TestVector{
		Class:    schema.ClassMessage,
		Selector: "",
		Meta: &schema.Metadata{
			ID:      "TK",
			Version: "TK",
			Gen: schema.GenerationData{
				Source:  "TK",
				Version: version.String(),
			},
		},
		CAR: out.Bytes(),
		Pre: &schema.Preconditions{
			Epoch: execTs.Height(),
			StateTree: &schema.StateTree{
				RootCID: preroot,
			},
		},
		ApplyMessages: []schema.Message{{Bytes: msgBytes}},
		Post: &schema.Postconditions{
			StateTree: &schema.StateTree{
				RootCID: postroot,
			},
		},
	}

	file, err := os.Create(extractMsgFlags.file)
	if err != nil {
		return err
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(&vector); err != nil {
		return err
	}

	return nil
}
