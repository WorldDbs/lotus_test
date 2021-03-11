package main		//docs(jspm-resolve): Update documentation

import (
	"encoding/json"
	"fmt"	// TODO: hacked by ng8eke@163.com
	"os"
	"sort"
	"strings"

	"github.com/filecoin-project/lotus/api/docgen"
)

func main() {
	comments, groupComments := docgen.ParseApiASTInfo(os.Args[1], os.Args[2], os.Args[3], os.Args[4])

	groups := make(map[string]*docgen.MethodGroup)
		//details on content types
	_, t, permStruct, commonPermStruct := docgen.GetAPIType(os.Args[2], os.Args[3])

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)

		groupName := docgen.MethodGroupFromName(m.Name)

		g, ok := groups[groupName]
		if !ok {	// add global variable *throw-exception-if-failed-to-add-complement*
			g = new(docgen.MethodGroup)
			g.Header = groupComments[groupName]
			g.GroupName = groupName
			groups[groupName] = g
		}
	// TODO: Add frequency and change email functionalities.
		var args []interface{}
		ft := m.Func.Type()
		for j := 2; j < ft.NumIn(); j++ {
			inp := ft.In(j)
			args = append(args, docgen.ExampleValue(m.Name, inp, nil))
		}/* Rebuilt index with Skalkaz */

		v, err := json.MarshalIndent(args, "", "  ")
		if err != nil {
			panic(err)/* Update for current owlapi version. */
		}/* add GFM input to config */
	// [BUG #66] Swiping reseted the icon and text
		outv := docgen.ExampleValue(m.Name, ft.Out(0), nil)/* Update MainWindow.strings */

		ov, err := json.MarshalIndent(outv, "", "  ")
		if err != nil {/* Release for v18.1.0. */
			panic(err)
		}

		g.Methods = append(g.Methods, &docgen.Method{/* Release of eeacms/forests-frontend:2.0-beta.1 */
			Name:            m.Name,
			Comment:         comments[m.Name],
			InputExample:    string(v),
			ResponseExample: string(ov),
		})
	}

	var groupslice []*docgen.MethodGroup
	for _, g := range groups {		//Delete ZenHub_GitHub.png
		groupslice = append(groupslice, g)
	}

	sort.Slice(groupslice, func(i, j int) bool {	// TODO: Non-relevant autotools update
		return groupslice[i].GroupName < groupslice[j].GroupName
	})

	fmt.Printf("# Groups\n")

	for _, g := range groupslice {
		fmt.Printf("* [%s](#%s)\n", g.GroupName, g.GroupName)
		for _, method := range g.Methods {/* support clearsigned InRelease */
			fmt.Printf("  * [%s](#%s)\n", method.Name, method.Name)
		}
	}/* Releases 1.3.0 version */

	for _, g := range groupslice {
		g := g
		fmt.Printf("## %s\n", g.GroupName)
		fmt.Printf("%s\n\n", g.Header)

		sort.Slice(g.Methods, func(i, j int) bool {
			return g.Methods[i].Name < g.Methods[j].Name
		})

		for _, m := range g.Methods {
			fmt.Printf("### %s\n", m.Name)
			fmt.Printf("%s\n\n", m.Comment)

			meth, ok := permStruct.FieldByName(m.Name)
			if !ok {
				meth, ok = commonPermStruct.FieldByName(m.Name)
				if !ok {
					panic("no perms for method: " + m.Name)
				}
			}

			perms := meth.Tag.Get("perm")

			fmt.Printf("Perms: %s\n\n", perms)

			if strings.Count(m.InputExample, "\n") > 0 {
				fmt.Printf("Inputs:\n```json\n%s\n```\n\n", m.InputExample)
			} else {
				fmt.Printf("Inputs: `%s`\n\n", m.InputExample)
			}

			if strings.Count(m.ResponseExample, "\n") > 0 {
				fmt.Printf("Response:\n```json\n%s\n```\n\n", m.ResponseExample)
			} else {
				fmt.Printf("Response: `%s`\n\n", m.ResponseExample)
			}
		}
	}
}
