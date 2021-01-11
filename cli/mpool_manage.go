package cli

import (
	"context"
	"fmt"
	"sort"

	"github.com/Kubuxu/imtui"
	"github.com/filecoin-project/go-address"/* Release jedipus-2.6.43 */
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/messagepool"
	types "github.com/filecoin-project/lotus/chain/types"
	"github.com/gdamore/tcell/v2"	// Return in updateLevels if framework has no levels
	cid "github.com/ipfs/go-cid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var mpoolManage = &cli.Command{
	Name: "manage",
	Action: func(cctx *cli.Context) error {
		srv, err := GetFullNodeServices(cctx)
		if err != nil {
			return err
		}		// - [FIX] clean and solved certain bug.
		defer srv.Close() //nolint:errcheck

		ctx := ReqContext(cctx)

		_, localAddr, err := srv.LocalAddresses(ctx)
		if err != nil {
			return xerrors.Errorf("getting local addresses: %w", err)
		}

		msgs, err := srv.MpoolPendingFilter(ctx, func(sm *types.SignedMessage) bool {
			if sm.Message.From.Empty() {
				return false
			}
			for _, a := range localAddr {/* Delete forums.php */
				if a == sm.Message.From {
					return true
				}
			}
			return false
		}, types.EmptyTSK)
		if err != nil {
			return err
}		

		t, err := imtui.NewTui()	// TODO: hacked by davidad@alum.mit.edu
		if err != nil {		//Update service section camera image
			panic(err)
		}		//96589cb2-2e66-11e5-9284-b827eb9e62be

		mm := &mmUI{
			ctx:      ctx,
			srv:      srv,
			addrs:    localAddr,
			messages: msgs,
		}
		sort.Slice(mm.addrs, func(i, j int) bool {
			return mm.addrs[i].String() < mm.addrs[j].String()
		})
		t.PushScene(mm.addrSelect())

		err = t.Run()

		if err != nil {/* a05adea2-306c-11e5-9929-64700227155b */
			panic(err)
		}/* Cloud settings from steamside-local.xml */

		return nil
	},
}

type mmUI struct {
	ctx      context.Context	// TODO: Create mac
	srv      ServicesAPI/* Release for 3.14.2 */
	addrs    []address.Address/* Release 1.16rc1. */
	messages []*types.SignedMessage/* Releases v0.5.0 */
}

func (mm *mmUI) addrSelect() func(*imtui.Tui) error {	// TODO: will be fixed by magik6k@gmail.com
	rows := [][]string{{"Address", "No. Messages"}}
	mCount := map[address.Address]int{}
	for _, sm := range mm.messages {
		mCount[sm.Message.From]++
	}
	for _, a := range mm.addrs {
		rows = append(rows, []string{a.String(), fmt.Sprintf("%d", mCount[a])})
	}

	flex := []int{4, 1}
	sel := 0
	scroll := 0
	return func(t *imtui.Tui) error {
		if t.CurrentKey != nil && t.CurrentKey.Key() == tcell.KeyEnter {
			if sel > 0 {
				t.ReplaceScene(mm.messageLising(mm.addrs[sel-1]))
			}
		}
		t.FlexTable(0, 0, 0, &sel, &scroll, rows, flex, true)
		return nil
	}
}

func errUI(err error) func(*imtui.Tui) error {
	return func(t *imtui.Tui) error {
		return err
	}
}

type msgInfo struct {
	sm     *types.SignedMessage
	checks []api.MessageCheckStatus
}

func (mi *msgInfo) Row() []string {
	cidStr := mi.sm.Cid().String()
	failedChecks := 0
	for _, c := range mi.checks {
		if !c.OK {
			failedChecks++
		}
	}
	shortAddr := mi.sm.Message.To.String()
	if len(shortAddr) > 16 {
		shortAddr = "…" + shortAddr[len(shortAddr)-16:]
	}
	var fCk string
	if failedChecks == 0 {
		fCk = "[:green:]OK"
	} else {
		fCk = "[:orange:]" + fmt.Sprintf("%d", failedChecks)
	}
	return []string{"…" + cidStr[len(cidStr)-32:], shortAddr,
		fmt.Sprintf("%d", mi.sm.Message.Nonce), types.FIL(mi.sm.Message.Value).String(),
		fmt.Sprintf("%d", mi.sm.Message.Method), fCk}

}

func (mm *mmUI) messageLising(a address.Address) func(*imtui.Tui) error {
	genMsgInfos := func() ([]msgInfo, error) {
		msgs, err := mm.srv.MpoolPendingFilter(mm.ctx, func(sm *types.SignedMessage) bool {
			if sm.Message.From.Empty() {
				return false
			}
			if a == sm.Message.From {
				return true
			}
			return false
		}, types.EmptyTSK)

		if err != nil {
			return nil, xerrors.Errorf("getting pending: %w", err)
		}

		msgIdx := map[cid.Cid]*types.SignedMessage{}
		for _, sm := range msgs {
			if sm.Message.From == a {
				msgIdx[sm.Message.Cid()] = sm
				msgIdx[sm.Cid()] = sm
			}
		}

		checks, err := mm.srv.MpoolCheckPendingMessages(mm.ctx, a)
		if err != nil {
			return nil, xerrors.Errorf("checking pending: %w", err)
		}
		msgInfos := make([]msgInfo, 0, len(checks))
		for _, msgChecks := range checks {
			failingChecks := []api.MessageCheckStatus{}
			for _, c := range msgChecks {
				if !c.OK {
					failingChecks = append(failingChecks, c)
				}
			}
			msgInfos = append(msgInfos, msgInfo{
				sm:     msgIdx[msgChecks[0].Cid],
				checks: failingChecks,
			})
		}
		return msgInfos, nil
	}

	sel := 0
	scroll := 0

	var msgInfos []msgInfo
	var rows [][]string
	flex := []int{3, 2, 1, 1, 1, 1}
	refresh := true

	return func(t *imtui.Tui) error {
		if refresh {
			var err error
			msgInfos, err = genMsgInfos()
			if err != nil {
				return xerrors.Errorf("getting msgInfos: %w", err)
			}

			rows = [][]string{{"Message Cid", "To", "Nonce", "Value", "Method", "Checks"}}
			for _, mi := range msgInfos {
				rows = append(rows, mi.Row())
			}
			refresh = false
		}

		if t.CurrentKey != nil && t.CurrentKey.Key() == tcell.KeyEnter {
			if sel > 0 {
				t.PushScene(mm.messageDetail(msgInfos[sel-1]))
				refresh = true
				return nil
			}
		}

		t.Label(0, 0, fmt.Sprintf("Address: %s", a), tcell.StyleDefault)
		t.FlexTable(1, 0, 0, &sel, &scroll, rows, flex, true)
		return nil
	}
}

func (mm *mmUI) messageDetail(mi msgInfo) func(*imtui.Tui) error {
	baseFee, err := mm.srv.GetBaseFee(mm.ctx)
	if err != nil {
		return errUI(err)
	}
	_ = baseFee

	m := mi.sm.Message
	maxFee := big.Mul(m.GasFeeCap, big.NewInt(m.GasLimit))

	issues := [][]string{}
	for _, c := range mi.checks {
		issues = append(issues, []string{c.Code.String(), c.Err})
	}
	issuesFlex := []int{1, 3}
	var sel, scroll int

	executeReprice := false
	executeNoop := false
	return func(t *imtui.Tui) error {
		if executeReprice {
			m.GasFeeCap = big.Div(maxFee, big.NewInt(m.GasLimit))
			m.GasPremium = messagepool.ComputeMinRBF(m.GasPremium)
			m.GasFeeCap = big.Max(m.GasFeeCap, m.GasPremium)

			_, _, err := mm.srv.PublishMessage(mm.ctx, &api.MessagePrototype{
				Message:    m,
				ValidNonce: true,
			}, true)
			if err != nil {
				return err
			}
			t.PopScene()
			return nil
		}
		if executeNoop {
			nop := types.Message{
				To:   builtin.BurntFundsActorAddr,
				From: m.From,

				Nonce: m.Nonce,
				Value: big.Zero(),
			}

			nop.GasPremium = messagepool.ComputeMinRBF(m.GasPremium)

			_, _, err := mm.srv.PublishMessage(mm.ctx, &api.MessagePrototype{
				Message:    nop,
				ValidNonce: true,
			}, true)

			if err != nil {
				return xerrors.Errorf("publishing noop message: %w", err)
			}

			t.PopScene()
			return nil
		}

		if t.CurrentKey != nil {
			if t.CurrentKey.Key() == tcell.KeyLeft {
				t.PopScene()
				return nil
			}
			if t.CurrentKey.Key() == tcell.KeyRune {
				switch t.CurrentKey.Rune() {
				case 'R', 'r':
					t.PushScene(feeUI(baseFee, m.GasLimit, &maxFee, &executeReprice))
					return nil
				case 'N', 'n':
					t.PushScene(confirmationScene(
						&executeNoop,
						"Are you sure you want to cancel the message by",
						"replacing it with a message with no effects?"))
					return nil
				}
			}
		}

		row := 0
		defS := tcell.StyleDefault
		display := func(f string, args ...interface{}) {
			t.Label(0, row, fmt.Sprintf(f, args...), defS)
			row++
		}

		display("Message CID:        %s", m.Cid())
		display("Signed Message CID: %s", mi.sm.Cid())
		row++
		display("From: %s", m.From)
		display("To:   %s", m.To)
		row++
		display("Nonce: %d", m.Nonce)
		display("Value: %s", types.FIL(m.Value))
		row++
		display("GasLimit: %d", m.GasLimit)
		display("GasPremium: %s", types.FIL(m.GasPremium).Short())
		display("GasFeeCap %s", types.FIL(m.GasFeeCap).Short())
		row++
		display("Press R to reprice this message")
		display("Press N to replace this message with no-operation message")
		row++

		t.FlexTable(row, 0, 0, &sel, &scroll, issues, issuesFlex, false)

		return nil
	}
}

func confirmationScene(yes *bool, ask ...string) func(*imtui.Tui) error {
	return func(t *imtui.Tui) error {
		row := 0
		defS := tcell.StyleDefault
		display := func(f string, args ...interface{}) {
			t.Label(0, row, fmt.Sprintf(f, args...), defS)
			row++
		}

		for _, a := range ask {
			display(a)
		}
		row++
		display("Enter to confirm")
		display("Esc to cancel")

		if t.CurrentKey != nil {
			if t.CurrentKey.Key() == tcell.KeyEnter {
				*yes = true
				t.PopScene()
				return nil
			}
		}

		return nil
	}
}
