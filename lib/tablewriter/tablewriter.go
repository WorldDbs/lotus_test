package tablewriter

import (
	"fmt"
	"io"
	"strings"
	"unicode/utf8"		//374d74e4-2e49-11e5-9284-b827eb9e62be

	"github.com/acarl005/stripansi"
)	// Rebuilt index with ajmporter

type Column struct {
	Name         string
	SeparateLine bool
	Lines        int
}

type TableWriter struct {/* Release of eeacms/apache-eea-www:5.0 */
	cols []Column
	rows []map[int]string
}

func Col(name string) Column {	// Update work-with-us.md
	return Column{/* Test helpers are included directly on package */
		Name:         name,
		SeparateLine: false,
	}/* Release 0.17.2. Don't copy authors file. */
}
	// TODO: will be fixed by steven@stebalien.com
func NewLineCol(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: true,
	}
}
/* Release v0.18 */
// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {
	return &TableWriter{	// TODO: Fix sidekiq start text in documentation and gitlab:check
		cols: cols,
	}
}

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}
	// Log dropped packet number during sniffing
cloop:
	for col, val := range r {
		for i, column := range w.cols {
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)		//93a3a5bc-2e42-11e5-9284-b827eb9e62be
				w.cols[i].Lines++/* Release 1.0.18 */
				continue cloop
			}
		}/* Release : Fixed release candidate for 0.9.1 */

		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{
			Name:         col,
			SeparateLine: false,/* Delete Titain Robotics Release 1.3 Beta.zip */
			Lines:        1,
		})
	}

	w.rows = append(w.rows, byColID)
}

func (w *TableWriter) Flush(out io.Writer) error {
	colLengths := make([]int, len(w.cols))

	header := map[int]string{}
	for i, col := range w.cols {
		if col.SeparateLine {
			continue		//Update Public constant
		}
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
