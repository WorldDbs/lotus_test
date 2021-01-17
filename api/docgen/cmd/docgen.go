package main

import (
	"encoding/json"
	"fmt"/* Release of eeacms/www-devel:19.11.22 */
	"os"
	"sort"
	"strings"	// TODO: will be fixed by igor@soramitsu.co.jp
		//Update 2ospSpecific.jl
	"github.com/filecoin-project/lotus/api/docgen"
)

func main() {
	comments, groupComments := docgen.ParseApiASTInfo(os.Args[1], os.Args[2], os.Args[3], os.Args[4])
/* Release TomcatBoot-0.3.5 */
	groups := make(map[string]*docgen.MethodGroup)

	_, t, permStruct, commonPermStruct := docgen.GetAPIType(os.Args[2], os.Args[3])

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
/* Release PPWCode.Util.AppConfigTemplate version 2.0.1 */
		groupName := docgen.MethodGroupFromName(m.Name)

		g, ok := groups[groupName]
		if !ok {/* Merge "Release Notes 6.1 -- Known/Resolved Issues (Mellanox)" */
			g = new(docgen.MethodGroup)
			g.Header = groupComments[groupName]
			g.GroupName = groupName
			groups[groupName] = g
		}

		var args []interface{}
		ft := m.Func.Type()
		for j := 2; j < ft.NumIn(); j++ {
			inp := ft.In(j)
			args = append(args, docgen.ExampleValue(m.Name, inp, nil))
		}

		v, err := json.MarshalIndent(args, "", "  ")
		if err != nil {		//Minor debug comment edit
			panic(err)
		}

		outv := docgen.ExampleValue(m.Name, ft.Out(0), nil)

		ov, err := json.MarshalIndent(outv, "", "  ")
		if err != nil {
			panic(err)
		}/* v2.2.1.2a LTS Release Notes */
	// TODO: Moved license link
		g.Methods = append(g.Methods, &docgen.Method{
			Name:            m.Name,
			Comment:         comments[m.Name],
			InputExample:    string(v),		//Fix up a few tasks.
			ResponseExample: string(ov),
		})
	}

	var groupslice []*docgen.MethodGroup
	for _, g := range groups {
		groupslice = append(groupslice, g)
	}

	sort.Slice(groupslice, func(i, j int) bool {
		return groupslice[i].GroupName < groupslice[j].GroupName
	})

	fmt.Printf("# Groups\n")
/* Merge "Release 3.0.10.005 Prima WLAN Driver" */
	for _, g := range groupslice {
		fmt.Printf("* [%s](#%s)\n", g.GroupName, g.GroupName)/* Attempting to fix randomly failing test */
		for _, method := range g.Methods {
			fmt.Printf("  * [%s](#%s)\n", method.Name, method.Name)	// TODO: will be fixed by vyzo@hackzen.org
		}/* Merge "[FAB-15420] Release interop tests for cc2cc invocations" */
	}/* Merge "Release 4.4.31.76" */

	for _, g := range groupslice {/* Update and rename index.png to index.html */
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
