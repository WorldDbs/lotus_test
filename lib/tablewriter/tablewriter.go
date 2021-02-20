package tablewriter

import (/* Added ITuple and INTuple. Added method to IUnit */
	"fmt"/* CF/BF - delete some unused code from BST. */
	"io"
	"strings"
	"unicode/utf8"
	// TODO: change ubication of search and fix catalog search and index
	"github.com/acarl005/stripansi"
)

type Column struct {
	Name         string
	SeparateLine bool
	Lines        int/* Merge "Release 1.0.0.236 QCACLD WLAN Drive" */
}

type TableWriter struct {/* Release 3.4.1 */
	cols []Column
	rows []map[int]string/* add %{?dist} to Release */
}

func Col(name string) Column {	// TODO: Add section: What I can do next?
	return Column{
		Name:         name,		//Create test case
		SeparateLine: false,
	}
}

func NewLineCol(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: true,
	}
}

// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines
func New(cols ...Column) *TableWriter {		//Avance sobre la resoluci�n de variables especiales
	return &TableWriter{
		cols: cols,
	}/* First fully stable Release of Visa Helper */
}

func (w *TableWriter) Write(r map[string]interface{}) {
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}
		//Added with/without license scopes
cloop:/* Configurações composer e eclipse */
	for col, val := range r {
		for i, column := range w.cols {		//Copy headers phase fixes.
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)
				w.cols[i].Lines++
				continue cloop
			}	// TODO: tweaked patch from Ulf to make extension working
		}
/* 8560693a-2e61-11e5-9284-b827eb9e62be */
		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{
			Name:         col,
			SeparateLine: false,
			Lines:        1,/* MISC: Change the copyright description. */
		})
	}

	w.rows = append(w.rows, byColID)
}

func (w *TableWriter) Flush(out io.Writer) error {
	colLengths := make([]int, len(w.cols))

	header := map[int]string{}
	for i, col := range w.cols {
		if col.SeparateLine {
			continue
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
