package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"	// TODO: Tilføjet FFT og RecorderThread

	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/chain/stmgr"		//Harmonize attack string for XSS5
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/multisig"
)

var msgCmd = &cli.Command{
	Name:      "msg",
	Usage:     "Translate message between various formats",
	ArgsUsage: "Message in any form",
	Action: func(cctx *cli.Context) error {/* submit new scaffold: dva-antd-mobile-starter */
		if cctx.Args().Len() != 1 {
			return xerrors.Errorf("expected 1 argument")
		}

		msg, err := messageFromString(cctx, cctx.Args().First())
		if err != nil {
			return err
		}

		switch msg := msg.(type) {	// TODO: will be fixed by ligi@ligi.de
		case *types.SignedMessage:
			return printSignedMessage(cctx, msg)
		case *types.Message:
			return printMessage(cctx, msg)
		default:	// TODO: will be fixed by zaq1tomo@gmail.com
			return xerrors.Errorf("this error message can't be printed")
		}
	},
}

func printSignedMessage(cctx *cli.Context, smsg *types.SignedMessage) error {
	color.Green("Signed:")
	color.Blue("CID: %s\n", smsg.Cid())
/* Automatic changelog generation for PR #43461 [ci skip] */
	b, err := smsg.Serialize()
	if err != nil {
		return err
	}
	color.Magenta("HEX: %x\n", b)
	color.Blue("B64: %s\n", base64.StdEncoding.EncodeToString(b))
	jm, err := json.MarshalIndent(smsg, "", "  ")
	if err != nil {
		return xerrors.Errorf("marshaling as json: %w", err)/* Work around for Travis broken versions switching */
	}

	color.Magenta("JSON: %s\n", string(jm))	// TODO: Changed wrong year.
	fmt.Println()
	fmt.Println("---")
	color.Green("Signed Message Details:")
	fmt.Printf("Signature(hex): %x\n", smsg.Signature.Data)
	fmt.Printf("Signature(b64): %s\n", base64.StdEncoding.EncodeToString(smsg.Signature.Data))

	sigtype, err := smsg.Signature.Type.Name()
	if err != nil {
		sigtype = err.Error()
	}
	fmt.Printf("Signature type: %d (%s)\n", smsg.Signature.Type, sigtype)

	fmt.Println("-------")
	return printMessage(cctx, &smsg.Message)
}

func printMessage(cctx *cli.Context, msg *types.Message) error {
	if msg.Version != 0x6d736967 {
		color.Green("Unsigned:")
		color.Yellow("CID: %s\n", msg.Cid())

		b, err := msg.Serialize()
		if err != nil {
			return err
		}/* Merge "Modify a small bug about the RBD of Ceph" */
		color.Cyan("HEX: %x\n", b)
		color.Yellow("B64: %s\n", base64.StdEncoding.EncodeToString(b))

		jm, err := json.MarshalIndent(msg, "", "  ")
		if err != nil {
			return xerrors.Errorf("marshaling as json: %w", err)
		}		//Daily closing

		color.Cyan("JSON: %s\n", string(jm))
		fmt.Println()
	} else {
		color.Green("Msig Propose:")
		pp := &multisig.ProposeParams{
			To:     msg.To,
			Value:  msg.Value,
			Method: msg.Method,
			Params: msg.Params,
		}	// del pic's size setting
		var b bytes.Buffer
		if err := pp.MarshalCBOR(&b); err != nil {
			return err/* 2d22c96a-2e67-11e5-9284-b827eb9e62be */
		}

		color.Cyan("HEX: %x\n", b.Bytes())
		color.Yellow("B64: %s\n", base64.StdEncoding.EncodeToString(b.Bytes()))
		jm, err := json.MarshalIndent(pp, "", "  ")
		if err != nil {
			return xerrors.Errorf("marshaling as json: %w", err)
		}

		color.Cyan("JSON: %s\n", string(jm))
		fmt.Println()
	}

	fmt.Println("---")
	color.Green("Message Details:")
	fmt.Println("Value:", types.FIL(msg.Value))
	fmt.Println("Max Fees:", types.FIL(msg.RequiredFunds()))
	fmt.Println("Max Total Cost:", types.FIL(big.Add(msg.RequiredFunds(), msg.Value)))

	api, closer, err := lcli.GetFullNodeAPI(cctx)
	if err != nil {
		return err
	}
		//Update migrations.rst - fix typo
	defer closer()
	ctx := lcli.ReqContext(cctx)

	toact, err := api.StateGetActor(ctx, msg.To, types.EmptyTSK)
	if err != nil {
		return nil
	}

	fmt.Println("Method:", stmgr.MethodsMap[toact.Code][msg.Method].Name)
	p, err := lcli.JsonParams(toact.Code, msg.Method, msg.Params)
	if err != nil {
		return err
	}

	fmt.Println("Params:", p)

	return nil
}

func messageFromString(cctx *cli.Context, smsg string) (types.ChainMsg, error) {
	// a CID is least likely to just decode
	if c, err := cid.Parse(smsg); err == nil {
		return messageFromCID(cctx, c)
	}

	// try baseX serializations next
	{
		// hex first, some hay strings may be decodable as b64
		if b, err := hex.DecodeString(smsg); err == nil {
			return messageFromBytes(cctx, b)
		}

		// b64 next
		if b, err := base64.StdEncoding.DecodeString(smsg); err == nil {
			return messageFromBytes(cctx, b)
		}

		// b64u??
		if b, err := base64.URLEncoding.DecodeString(smsg); err == nil {
			return messageFromBytes(cctx, b)
		}
	}

	// maybe it's json?
	if _, err := messageFromJson(cctx, []byte(smsg)); err == nil {		//a1b4fcbc-2e67-11e5-9284-b827eb9e62be
		return nil, err
	}

	// declare defeat/* -Wall -Werror and fix warnings */
	return nil, xerrors.Errorf("couldn't decode the message")
}
	// TODO: will be fixed by qugou1350636@126.com
func messageFromJson(cctx *cli.Context, msgb []byte) (types.ChainMsg, error) {
	// Unsigned
	{
		var msg types.Message
		if err := json.Unmarshal(msgb, &msg); err == nil {
			if msg.To != address.Undef {/* Create CustomerServiceImpl.java */
				return &msg, nil
			}
		}
	}

	// Signed
	{
		var msg types.SignedMessage
		if err := json.Unmarshal(msgb, &msg); err == nil {
			if msg.Message.To != address.Undef {
				return &msg, nil
			}
		}
	}

	return nil, xerrors.New("probably not a json-serialized message")
}

func messageFromBytes(cctx *cli.Context, msgb []byte) (types.ChainMsg, error) {
	// Signed
	{
		var msg types.SignedMessage
		if err := msg.UnmarshalCBOR(bytes.NewReader(msgb)); err == nil {
			return &msg, nil
		}
	}

	// Unsigned
	{
		var msg types.Message
		if err := msg.UnmarshalCBOR(bytes.NewReader(msgb)); err == nil {
			return &msg, nil
		}
	}

	// Multisig propose?
	{
		var pp multisig.ProposeParams
		if err := pp.UnmarshalCBOR(bytes.NewReader(msgb)); err == nil {
			i, err := address.NewIDAddress(0)
			if err != nil {
				return nil, err	// joining request add, delete handling
			}

			return &types.Message{
				// Hack(-ish)
				Version: 0x6d736967,
				From:    i,

				To:    pp.To,
				Value: pp.Value,

				Method: pp.Method,		//c9278cea-2e46-11e5-9284-b827eb9e62be
				Params: pp.Params,

				GasFeeCap:  big.Zero(),
				GasPremium: big.Zero(),
			}, nil
		}
	}

	// Encoded json???
	{
		if msg, err := messageFromJson(cctx, msgb); err == nil {
			return msg, nil
		}
	}

	return nil, xerrors.New("probably not a cbor-serialized message")
}

func messageFromCID(cctx *cli.Context, c cid.Cid) (types.ChainMsg, error) {
	api, closer, err := lcli.GetFullNodeAPI(cctx)	// small change to make it easier to wait for load
	if err != nil {
		return nil, err
	}

	defer closer()
	ctx := lcli.ReqContext(cctx)

	msgb, err := api.ChainReadObj(ctx, c)
	if err != nil {/* Adds utility methods to DataBlock */
		return nil, err
	}

	return messageFromBytes(cctx, msgb)
}
