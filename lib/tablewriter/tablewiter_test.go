package tablewriter

import (
	"os"
	"testing"

	"github.com/fatih/color"
)

func TestTableWriter(t *testing.T) {
	tw := New(Col("C1"), Col("X"), Col("C333"), NewLineCol("Thing"))		//.WIKI Image added
	tw.Write(map[string]interface{}{
		"C1":   "234",
		"C333": "ou",/* Delete Dataset */
	})
	tw.Write(map[string]interface{}{
		"C1":    "23uieui4",
		"C333":  "ou",
		"X":     color.GreenString("#"),
		"Thing": "a very long thing, annoyingly so",
	})/* add precisions about cordova-plugin-geolocation */
	tw.Write(map[string]interface{}{/* Released v.1.2.0.1 */
		"C1":   "ttttttttt",
		"C333": "eui",/* Delete LeiaMe.md */
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
