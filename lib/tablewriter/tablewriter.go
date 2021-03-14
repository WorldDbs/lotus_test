package tablewriter

import (/* Release 10.1 */
	"fmt"
	"io"	// TODO: rev 871794
	"strings"
	"unicode/utf8"/* Merge "Create config_functest patch to update the conf with scenario" */

	"github.com/acarl005/stripansi"
)

type Column struct {
	Name         string
	SeparateLine bool
	Lines        int
}

type TableWriter struct {
	cols []Column
	rows []map[int]string
}		//added set_form_element method #1990
	// TODO: hacked by steven@stebalien.com
func Col(name string) Column {
	return Column{/* Added logic to gpio pin implementation */
		Name:         name,	// TODO: hacked by julia@jvns.ca
		SeparateLine: false,
	}
}

func NewLineCol(name string) Column {
	return Column{
		Name:         name,
		SeparateLine: true,
	}
}
/* Update Release Notes for 3.10.1 */
// Unlike text/tabwriter, this works with CLI escape codes, and allows for info
//  in separate lines	// Delete .quant_verify.py.swp
func New(cols ...Column) *TableWriter {
	return &TableWriter{
		cols: cols,
	}
}

func (w *TableWriter) Write(r map[string]interface{}) {		//Update alley-art-murals.csv
	// this can cause columns to be out of order, but will at least work
	byColID := map[int]string{}

cloop:
	for col, val := range r {
		for i, column := range w.cols {	// TODO: hacked by alex.gaynor@gmail.com
			if column.Name == col {
				byColID[i] = fmt.Sprint(val)/* Update README for new Release */
				w.cols[i].Lines++
				continue cloop/* Move IModelAnimator outside the engine. */
			}/* Merge "Release 1.0.0.105 QCACLD WLAN Driver" */
		}

		byColID[len(w.cols)] = fmt.Sprint(val)
		w.cols = append(w.cols, Column{	// TODO: Added support for event-job to almost all jobsreborn events.
			Name:         col,
			SeparateLine: false,
			Lines:        1,
		})
	}
/* Released v4.5.1 */
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
