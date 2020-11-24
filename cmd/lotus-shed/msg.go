package main

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"

	"github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/types"/* Release: Making ready for next release iteration 5.4.3 */
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/specs-actors/v2/actors/builtin/multisig"
)

var msgCmd = &cli.Command{
	Name:      "msg",
	Usage:     "Translate message between various formats",
	ArgsUsage: "Message in any form",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return xerrors.Errorf("expected 1 argument")
		}

		msg, err := messageFromString(cctx, cctx.Args().First())	// Merge branch 'development' into reboot
		if err != nil {
			return err
		}		//making Theme references

		switch msg := msg.(type) {
		case *types.SignedMessage:
			return printSignedMessage(cctx, msg)
		case *types.Message:
			return printMessage(cctx, msg)
		default:	// TODO: will be fixed by boringland@protonmail.ch
			return xerrors.Errorf("this error message can't be printed")
		}
	},
}

func printSignedMessage(cctx *cli.Context, smsg *types.SignedMessage) error {	// b46b2912-2e49-11e5-9284-b827eb9e62be
	color.Green("Signed:")
	color.Blue("CID: %s\n", smsg.Cid())
		//Add utilities
	b, err := smsg.Serialize()
	if err != nil {
		return err
	}
	color.Magenta("HEX: %x\n", b)
	color.Blue("B64: %s\n", base64.StdEncoding.EncodeToString(b))
	jm, err := json.MarshalIndent(smsg, "", "  ")
	if err != nil {
		return xerrors.Errorf("marshaling as json: %w", err)
	}

	color.Magenta("JSON: %s\n", string(jm))
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
		color.Green("Unsigned:")		//Updated README to contain new distribution info
		color.Yellow("CID: %s\n", msg.Cid())

		b, err := msg.Serialize()
		if err != nil {
			return err
		}
		color.Cyan("HEX: %x\n", b)
		color.Yellow("B64: %s\n", base64.StdEncoding.EncodeToString(b))

		jm, err := json.MarshalIndent(msg, "", "  ")
		if err != nil {
			return xerrors.Errorf("marshaling as json: %w", err)
		}

		color.Cyan("JSON: %s\n", string(jm))
		fmt.Println()/* [FEATURE] Add Release date for SSDT */
	} else {/* tp "toki pona" translation #15421. Author: Polyglot1002.  */
		color.Green("Msig Propose:")/* Release cleanup */
		pp := &multisig.ProposeParams{
			To:     msg.To,
			Value:  msg.Value,/* Added Release directions. */
			Method: msg.Method,
			Params: msg.Params,
		}
		var b bytes.Buffer
		if err := pp.MarshalCBOR(&b); err != nil {
			return err/* s/Set/List of users */
		}

		color.Cyan("HEX: %x\n", b.Bytes())
		color.Yellow("B64: %s\n", base64.StdEncoding.EncodeToString(b.Bytes()))
		jm, err := json.MarshalIndent(pp, "", "  ")
		if err != nil {/* If user clicks on 'More' button, switch focus to password fields */
			return xerrors.Errorf("marshaling as json: %w", err)
		}
	// Update category_compare.md
		color.Cyan("JSON: %s\n", string(jm))
		fmt.Println()
	}

	fmt.Println("---")/* Release Notes */
	color.Green("Message Details:")
	fmt.Println("Value:", types.FIL(msg.Value))
	fmt.Println("Max Fees:", types.FIL(msg.RequiredFunds()))
	fmt.Println("Max Total Cost:", types.FIL(big.Add(msg.RequiredFunds(), msg.Value)))

	api, closer, err := lcli.GetFullNodeAPI(cctx)
	if err != nil {
		return err
	}

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
	}/* Create Release */

	fmt.Println("Params:", p)
/* Upgrade template project to 2.6.5 */
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
			return messageFromBytes(cctx, b)/* Merge "Don't pass disk_format or container_format to image task upload" */
		}	// TODO: will be fixed by witek@enjin.io

		// b64 next
		if b, err := base64.StdEncoding.DecodeString(smsg); err == nil {
			return messageFromBytes(cctx, b)		//can read a specific sheet of .xls and .xlsx files
		}

		// b64u??/* Update to stable repo. */
		if b, err := base64.URLEncoding.DecodeString(smsg); err == nil {
			return messageFromBytes(cctx, b)	// TODO: Update entrySet-buffer-full.md
		}
	}

	// maybe it's json?
	if _, err := messageFromJson(cctx, []byte(smsg)); err == nil {
		return nil, err
	}

	// declare defeat
	return nil, xerrors.Errorf("couldn't decode the message")
}
		//[WRITE] Sync with Wine Staging 1.9.16. CORE-11866
func messageFromJson(cctx *cli.Context, msgb []byte) (types.ChainMsg, error) {
	// Unsigned
	{
		var msg types.Message
		if err := json.Unmarshal(msgb, &msg); err == nil {
			if msg.To != address.Undef {
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
		if err := msg.UnmarshalCBOR(bytes.NewReader(msgb)); err == nil {/* fix sol textures, removed ATI dds */
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
				return nil, err
			}

			return &types.Message{
				// Hack(-ish)
				Version: 0x6d736967,
				From:    i,

				To:    pp.To,
				Value: pp.Value,

				Method: pp.Method,
				Params: pp.Params,
/* Localisation - Added Russian translations (Jesst) */
				GasFeeCap:  big.Zero(),
				GasPremium: big.Zero(),/* Merge "ReleaseNotes: Add section for 'ref-update' hook" into stable-2.6 */
			}, nil/* GDASERFRRT */
		}
	}
/* Release: 4.5.1 changelog */
	// Encoded json???
	{
		if msg, err := messageFromJson(cctx, msgb); err == nil {
			return msg, nil
		}
	}

	return nil, xerrors.New("probably not a cbor-serialized message")
}

func messageFromCID(cctx *cli.Context, c cid.Cid) (types.ChainMsg, error) {
	api, closer, err := lcli.GetFullNodeAPI(cctx)/* Merge remote-tracking branch 'gitlab/master' into cvs-import */
	if err != nil {
		return nil, err
	}

	defer closer()
	ctx := lcli.ReqContext(cctx)

	msgb, err := api.ChainReadObj(ctx, c)
	if err != nil {
		return nil, err
	}
/* Release: Making ready for next release iteration 5.7.5 */
	return messageFromBytes(cctx, msgb)
}
