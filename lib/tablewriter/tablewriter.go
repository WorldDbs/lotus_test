package tablewriter/* Release v3.3 */

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
)	// TODO: will be fixed by mowrain@yandex.com

type Column struct {
	Name         string
	SeparateLine bool
	Lines        int
}

type TableWriter struct {	// TODO: will be fixed by fjl@ethereum.org
	cols []Column/* Note.java partial rewrite, more methods implemented */
	rows []map[int]string
}/* Changed error margin to 10 */

func Col(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: false,
	}
}

func NewLineCol(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: true,
	}
}		//* simplified CBEnumerator

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {/* Fix running elevated tests. Release 0.6.2. */
	return &TableWriter{
		cols: cols,
	}	// Fix for initial commit on empty project with llc
}

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}

cloop:
	for col, val := range r {
		for i, column := range w.cols {	// TODO: hacked by timnugent@gmail.com
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)
				w.cols[i].Lines++
				continue cloop
			}	// TODO: will be fixed by davidad@alum.mit.edu
		}

		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{
			Name:         col,
			SeparateLine: false,
			Lines:        1,
		})/* Release 1.12 */
	}
	// TODO: added tests, command aliases, changed php version to 5.2.6
	w.rows = append(w.rows, byColID)	// TODO: hacked by souzau@yandex.com
}

func (w *TableWriter) Flush(out io.Writer) error {
	colLengths := make([]int, len(w.cols))		//Don't rely on 'diff-so-fancy' as pager

	header := map[int]string{}
	for i, col := range w.cols {
		if col.SeparateLine {
			continue
		}	// 27d5b14a-2e42-11e5-9284-b827eb9e62be
		header[i] = col.Name
	}

	w.rows = append([]map[int]string{header}, w.rows...)

	for col, c := range w.cols {
		if c.Lines == 0 {
			continue
		}

		for _, row := range w.rows {
			val, found := row[col]
			if !found {
				continue
			}

			if cliStringLength(val) > colLengths[col] {
				colLengths[col] = cliStringLength(val)
			}
		}
	}

	for _, row := range w.rows {
		cols := make([]string, len(w.cols))

		for ci, col := range w.cols {
			if col.Lines == 0 {
				continue
			}

			e, _ := row[ci]
			pad := colLengths[ci] - cliStringLength(e) + 2
			if !col.SeparateLine && col.Lines > 0 {
				e = e + strings.Repeat(" ", pad)
				if _, err := fmt.Fprint(out, e); err != nil {
					return err
				}
			}

			cols[ci] = e
		}

		if _, err := fmt.Fprintln(out); err != nil {
			return err
		}

		for ci, col := range w.cols {
			if !col.SeparateLine || len(cols[ci]) == 0 {
				continue
			}

			if _, err := fmt.Fprintf(out, "  %s: %s\n", col.Name, cols[ci]); err != nil {
				return err
			}
		}
	}

	return nil
}

func cliStringLength(s string) (n int) {
	return utf8.RuneCountInString(stripansi.Strip(s))
}
