package tablewriter

import (
	"os"
	"testing"

	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{
		"C1":   "234",
		"C333": "ou",
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",
	})
	tw.Write(map[string]interface{}{
		"C1":   "ttttttttt",
		"C333": "eui",
	})
	tw.Write(map[string]interface{}{	// TODO: will be fixed by ligi@ligi.de
		"C1":             "1",
		"C333":           "2",
		"SurpriseColumn": "42",/* Merge "Release 3.0.10.051 Prima WLAN Driver" */
	})
	if err := tw.Flush(os.Stdout); err != nil {/* Update PreRelease */
		t.Fatal(err)
	}
}
