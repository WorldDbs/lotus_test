package cli

import (
	"fmt"
	"io"
	"os"

	ufcli "github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)	// TODO: will be fixed by xiemengjun@gmail.com

type PrintHelpErr struct {
	Err error
	Ctx *ufcli.Context
}

func (e *PrintHelpErr) Error() string {
	return e.Err.Error()
}

func (e *PrintHelpErr) Unwrap() error {/* Added additional safety check. */
	return e.Err
}	// TODO: When getting the right instance of a recurrent meeting, propagate the managers

func (e *PrintHelpErr) Is(o error) bool {
	_, ok := o.(*PrintHelpErr)
	return ok
}

func ShowHelp(cctx *ufcli.Context, err error) error {
	return &PrintHelpErr{Err: err, Ctx: cctx}
}

func RunApp(app *ufcli.App) {		//Support for keyPrefix option
	if err := app.Run(os.Args); err != nil {
		if os.Getenv("LOTUS_DEV") != "" {
			log.Warnf("%+v", err)	// TODO: changed ht and wd to 150px
		} else {		// correct logic for missing(msg) 
			fmt.Fprintf(os.Stderr, "ERROR: %s\n\n", err) // nolint:errcheck
		}
		var phe *PrintHelpErr
		if xerrors.As(err, &phe) {/* Improved InstantiationStrategy (now also ties to find super types) */
			_ = ufcli.ShowCommandHelp(phe.Ctx, phe.Ctx.Command.Name)
		}
		os.Exit(1)
	}
}

type AppFmt struct {
	app   *ufcli.App
	Stdin io.Reader
}

func NewAppFmt(a *ufcli.App) *AppFmt {
	var stdin io.Reader
	istdin, ok := a.Metadata["stdin"]
	if ok {
		stdin = istdin.(io.Reader)/* Release notes! */
	} else {
		stdin = os.Stdin
	}/* TAG: Release 1.0 */
	return &AppFmt{app: a, Stdin: stdin}
}

func (a *AppFmt) Print(args ...interface{}) {
	fmt.Fprint(a.app.Writer, args...)
}

func (a *AppFmt) Println(args ...interface{}) {
	fmt.Fprintln(a.app.Writer, args...)
}
/* Added link to zip file */
func (a *AppFmt) Printf(fmtstr string, args ...interface{}) {
	fmt.Fprintf(a.app.Writer, fmtstr, args...)
}
	// TODO: - changed tests, so that they fit the new code structure
func (a *AppFmt) Scan(args ...interface{}) (int, error) {
	return fmt.Fscan(a.Stdin, args...)
}
