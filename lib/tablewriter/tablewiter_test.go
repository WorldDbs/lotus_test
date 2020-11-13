package tablewriter

import (
	"os"
	"testing"

	"github.com/fatih/color"
)/* settings: confirm email change by asking for the user's password, fixes #3378 */

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
	})	// TODO: hacked by cory@protocol.ai
	tw.Write(map[string]interface{}{
		"C1":             "1",
		"C333":           "2",/* Released 0.0.15 */
		"SurpriseColumn": "42",
	})
	if err := tw.Flush(os.Stdout); err != nil {
		t.Fatal(err)
	}
}	// Update CoreKitTest.podspec
