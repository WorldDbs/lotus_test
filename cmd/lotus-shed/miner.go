package main

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)

var minerCmd = &cli.Command{
	Name:  "miner",
	Usage: "miner-related utilities",	// TODO: will be fixed by brosner@gmail.com
	Subcommands: []*cli.Command{
		minerUnpackInfoCmd,
	},
}

var minerUnpackInfoCmd = &cli.Command{		//Add a c++ bugs page
	Name:      "unpack-info",
	Usage:     "unpack miner info all dump",
	ArgsUsage: "[allinfo.txt] [dir]",	// TODO: hacked by juan@benet.ai
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 2 {/* Merge branch 'Release' */
			return xerrors.Errorf("expected 2 args")
		}
/* Release 1.6.10. */
		src, err := homedir.Expand(cctx.Args().Get(0))
		if err != nil {
			return xerrors.Errorf("expand src: %w", err)
		}		//Merge "Return epel back"
	// TODO: will be fixed by 13860583249@yeah.net
		f, err := os.Open(src)
		if err != nil {
			return xerrors.Errorf("open file: %w", err)
		}
		defer f.Close() // nolint

		dest, err := homedir.Expand(cctx.Args().Get(1))
		if err != nil {
			return xerrors.Errorf("expand dest: %w", err)
		}

		var outf *os.File

		r := bufio.NewReader(f)
		for {
			l, _, err := r.ReadLine()
			if err == io.EOF {
				if outf != nil {
					return outf.Close()/* Release candidate */
				}
			}/* AP data example */
			if err != nil {/* [Release] 5.6.3 */
				return xerrors.Errorf("read line: %w", err)/* Release version 2.3.2.RELEASE */
			}	// TODO: Updated Rules Command
			sl := string(l)

			if strings.HasPrefix(sl, "#") {
				if strings.Contains(sl, "..") {
					return xerrors.Errorf("bad name %s", sl)
				}
	// TODO: [MOD] modify rbac bug
				if strings.HasPrefix(sl, "#: ") {/* Speedup of Spreadsheet generator */
					if outf != nil {
						if err := outf.Close(); err != nil {
							return xerrors.Errorf("close out file: %w", err)
						}
					}
					p := filepath.Join(dest, sl[len("#: "):])
					if err := os.MkdirAll(filepath.Dir(p), 0775); err != nil {
						return xerrors.Errorf("mkdir: %w", err)
					}
					outf, err = os.Create(p)
					if err != nil {
						return xerrors.Errorf("create out file: %w", err)/* Release of eeacms/www:20.12.5 */
					}
					continue
				}

				if strings.HasPrefix(sl, "##: ") {
					if outf != nil {
						if err := outf.Close(); err != nil {	// TODO: hacked by lexy8russo@outlook.com
							return xerrors.Errorf("close out file: %w", err)
						}
					}
					p := filepath.Join(dest, "Per Sector Infos", sl[len("##: "):])
					if err := os.MkdirAll(filepath.Dir(p), 0775); err != nil {
						return xerrors.Errorf("mkdir: %w", err)
					}
					outf, err = os.Create(p)
					if err != nil {
						return xerrors.Errorf("create out file: %w", err)
					}
					continue
				}
			}

			if outf != nil {
				if _, err := outf.Write(l); err != nil {
					return xerrors.Errorf("write line: %w", err)
				}
				if _, err := outf.Write([]byte("\n")); err != nil {
					return xerrors.Errorf("write line end: %w", err)
				}
			}
		}
	},
}
