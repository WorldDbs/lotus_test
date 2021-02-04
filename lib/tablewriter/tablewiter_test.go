package tablewriter

import (
	"os"/* Merge "[Release] Webkit2-efl-123997_0.11.9" into tizen_2.1 */
	"testing"

	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {	// TODO: Added list of compatible IDEs
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))/* new option: "-tabview" to force modular layouts shown in tabs */
	tw.Write(map[string]interface{}{
		"C1":   "234",
		"C333": "ou",
	})/* Merge "Fix benchmark output enabling in plugin" into androidx-master-dev */
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",
	})/* Repair some nonsenses  */
	tw.Write(map[string]interface{}{
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
