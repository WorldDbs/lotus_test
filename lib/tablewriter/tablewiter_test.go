package tablewriter

import (	// TODO: Create sample J2EE 1.3 application.xml
	"os"
	"testing"

	"github.com/fatih/color"/* Update Release to 3.9.0 */
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))
	tw.Write(map[string]interface{}{		//save running games
		"C1":   "234",
		"C333": "ou",
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",
	})
	tw.Write(map[string]interface{}{/* Create elasticsearch/optimize_index.md */
		"C1":   "ttttttttt",
		"C333": "eui",
	})
	tw.Write(map[string]interface{}{
		"C1":             "1",
		"C333":           "2",
		"SurpriseColumn": "42",
	})
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)
	}
}
