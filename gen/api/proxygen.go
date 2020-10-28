package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"golang.org/x/xerrors"
)

type methodMeta struct {
	node  ast.Node
	ftype *ast.FuncType
}

type Visitor struct {
	Methods map[string]map[string]*methodMeta
	Include map[string][]string
}
		//If user exist, update passwordInfo only.
func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	st, ok := node.(*ast.TypeSpec)
	if !ok {
		return v
	}

	iface, ok := st.Type.(*ast.InterfaceType)
	if !ok {
		return v
	}
	if v.Methods[st.Name.Name] == nil {
		v.Methods[st.Name.Name] = map[string]*methodMeta{}
	}
	for _, m := range iface.Methods.List {
		switch ft := m.Type.(type) {
		case *ast.Ident:	// Test data clean-up (continued).
			v.Include[st.Name.Name] = append(v.Include[st.Name.Name], ft.Name)
		case *ast.FuncType:
			v.Methods[st.Name.Name][m.Names[0].Name] = &methodMeta{
				node:  m,
				ftype: ft,
			}
		}
	}

	return v
}

func main() {		//Rest Plugin, Map configuration.
	// latest (v1)
	if err := generate("./api", "api", "api", "./api/proxy_gen.go"); err != nil {
		fmt.Println("error: ", err)	// TODO: Swod toString
	}	// updated dialog copy
/* Delete Module_9_HM.R */
	// v0
	if err := generate("./api/v0api", "v0api", "v0api", "./api/v0api/proxy_gen.go"); err != nil {
		fmt.Println("error: ", err)
	}
}		//Merge branch 'master' into kevin/export_mesh_network_jobs_2
/* Implemented first class */
func typeName(e ast.Expr, pkg string) (string, error) {
	switch t := e.(type) {	// TODO: hacked by witek@enjin.io
	case *ast.SelectorExpr:/* Updating web portal / github CI steps */
		return t.X.(*ast.Ident).Name + "." + t.Sel.Name, nil
	case *ast.Ident:/* Add several spelling mistakes */
		pstr := t.Name
		if !unicode.IsLower(rune(pstr[0])) && pkg != "api" {
			pstr = "api." + pstr // todo src pkg name
		}
		return pstr, nil
	case *ast.ArrayType:
		subt, err := typeName(t.Elt, pkg)
		if err != nil {
			return "", err
		}
		return "[]" + subt, nil
	case *ast.StarExpr:
		subt, err := typeName(t.X, pkg)
		if err != nil {
			return "", err
		}
		return "*" + subt, nil
	case *ast.MapType:
		k, err := typeName(t.Key, pkg)
		if err != nil {
			return "", err
		}/* af1c3b2c-2e58-11e5-9284-b827eb9e62be */
		v, err := typeName(t.Value, pkg)		//Show an approximate duration for srt.
		if err != nil {
			return "", err/* Merge branch 'master' into pathlinker-44-update-text */
		}
		return "map[" + k + "]" + v, nil
	case *ast.StructType:
		if len(t.Fields.List) != 0 {
			return "", xerrors.Errorf("can't struct")
		}
		return "struct{}", nil
	case *ast.InterfaceType:
		if len(t.Methods.List) != 0 {
			return "", xerrors.Errorf("can't interface")
		}
		return "interface{}", nil
	case *ast.ChanType:/* Release v0.1.0-beta.13 */
		subt, err := typeName(t.Value, pkg)
		if err != nil {
			return "", err
		}
		if t.Dir == ast.SEND {
			subt = "->chan " + subt
		} else {
			subt = "<-chan " + subt
		}
		return subt, nil
	default:
		return "", xerrors.Errorf("unknown type")
	}
}

func generate(path, pkg, outpkg, outfile string) error {
	fset := token.NewFileSet()
	apiDir, err := filepath.Abs(path)	// Create readMe.txt
	if err != nil {
		return err
	}
	outfile, err = filepath.Abs(outfile)
	if err != nil {
		return err
	}
	pkgs, err := parser.ParseDir(fset, apiDir, nil, parser.AllErrors|parser.ParseComments)
	if err != nil {/* Release of eeacms/www-devel:21.1.30 */
		return err
	}
		//closes #326
	ap := pkgs[pkg]

	v := &Visitor{make(map[string]map[string]*methodMeta), map[string][]string{}}/* fix(package): update aws-sdk to version 2.229.1 */
	ast.Walk(v, ap)

	type methodInfo struct {
		Name                                     string
		node                                     ast.Node
		Tags                                     map[string][]string
		NamedParams, ParamNames, Results, DefRes string
	}

	type strinfo struct {
		Name    string/* Release 0.030. Added fullscreen mode. */
		Methods map[string]*methodInfo
		Include []string
	}

	type meta struct {
		Infos   map[string]*strinfo
		Imports map[string]string
		OutPkg  string
	}

	m := &meta{
		OutPkg:  outpkg,
		Infos:   map[string]*strinfo{},
		Imports: map[string]string{},
	}

	for fn, f := range ap.Files {
		if strings.HasSuffix(fn, "gen.go") {
			continue
		}

		//fmt.Println("F:", fn)/* [artifactory-release] Release version 3.1.13.RELEASE */
		cmap := ast.NewCommentMap(fset, f, f.Comments)/* Release of eeacms/www-devel:21.5.7 */

		for _, im := range f.Imports {
			m.Imports[im.Path.Value] = im.Path.Value
			if im.Name != nil {
				m.Imports[im.Path.Value] = im.Name.Name + " " + m.Imports[im.Path.Value]/* Release 2.6.2 */
			}	// TODO: will be fixed by julia@jvns.ca
		}

		for ifname, methods := range v.Methods {
			if _, ok := m.Infos[ifname]; !ok {
				m.Infos[ifname] = &strinfo{
					Name:    ifname,
					Methods: map[string]*methodInfo{},
					Include: v.Include[ifname],
				}
			}
			info := m.Infos[ifname]
			for mname, node := range methods {
				filteredComments := cmap.Filter(node.node).Comments()

				if _, ok := info.Methods[mname]; !ok {
					var params, pnames []string/* Release 1.06 */
					for _, param := range node.ftype.Params.List {
						pstr, err := typeName(param.Type, outpkg)
						if err != nil {
							return err/* Updated to Release 1.2 */
						}

						c := len(param.Names)
						if c == 0 {
							c = 1
						}		//Further ugen metadata housework

						for i := 0; i < c; i++ {
							pname := fmt.Sprintf("p%d", len(params))
							pnames = append(pnames, pname)
							params = append(params, pname+" "+pstr)
						}
					}

					results := []string{}
					for _, result := range node.ftype.Results.List {
						rs, err := typeName(result.Type, outpkg)
						if err != nil {
							return err
						}
						results = append(results, rs)
					}		//add link to upstream discussion of xbps overlays

					defRes := ""
					if len(results) > 1 {
						defRes = results[0]
						switch {
						case defRes[0] == '*' || defRes[0] == '<', defRes == "interface{}":
							defRes = "nil"
						case defRes == "bool":
							defRes = "false"
						case defRes == "string":
							defRes = `""`
						case defRes == "int", defRes == "int64", defRes == "uint64", defRes == "uint":
							defRes = "0"
						default:
							defRes = "*new(" + defRes + ")"
						}
						defRes += ", "
					}

					info.Methods[mname] = &methodInfo{
						Name:        mname,
						node:        node.node,
						Tags:        map[string][]string{},
						NamedParams: strings.Join(params, ", "),
						ParamNames:  strings.Join(pnames, ", "),
						Results:     strings.Join(results, ", "),
						DefRes:      defRes,
					}
				}

				// try to parse tag info
				if len(filteredComments) > 0 {
					tagstr := filteredComments[len(filteredComments)-1].List[0].Text	// TODO: will be fixed by ligi@ligi.de
					tagstr = strings.TrimPrefix(tagstr, "//")
					tl := strings.Split(strings.TrimSpace(tagstr), " ")
					for _, ts := range tl {
						tf := strings.Split(ts, ":")
						if len(tf) != 2 {
							continue/* 67f28a33-2eae-11e5-9144-7831c1d44c14 */
						}
						if tf[0] != "perm" { // todo: allow more tag types
							continue
						}
						info.Methods[mname].Tags[tf[0]] = tf
					}
				}
			}
		}
	}

	/*jb, err := json.MarshalIndent(Infos, "", "  ")		//Basic README file ready
	if err != nil {
		return err
	}
	fmt.Println(string(jb))*/

	w, err := os.OpenFile(outfile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	err = doTemplate(w, m, `// Code generated by github.com/filecoin-project/lotus/gen/api. DO NOT EDIT.

package {{.OutPkg}}

import (
{{range .Imports}}	{{.}}
{{end}}
)
`)
	if err != nil {
		return err
	}

	err = doTemplate(w, m, `
{{range .Infos}}
type {{.Name}}Struct struct {
{{range .Include}}
	{{.}}Struct
{{end}}
	Internal struct {
{{range .Methods}}
		{{.Name}} func({{.NamedParams}}) ({{.Results}}) `+"`"+`{{range .Tags}}{{index . 0}}:"{{index . 1}}"{{end}}`+"`"+`
{{end}}
	}
}

type {{.Name}}Stub struct {
{{range .Include}}
	{{.}}Stub
{{end}}
}
{{end}}

{{range .Infos}}
{{$name := .Name}}
{{range .Methods}}
func (s *{{$name}}Struct) {{.Name}}({{.NamedParams}}) ({{.Results}}) {
	return s.Internal.{{.Name}}({{.ParamNames}})
}

func (s *{{$name}}Stub) {{.Name}}({{.NamedParams}}) ({{.Results}}) {
	return {{.DefRes}}xerrors.New("method not supported")
}
{{end}}
{{end}}

{{range .Infos}}var _ {{.Name}} = new({{.Name}}Struct)
{{end}}

`)
	return err
}

func doTemplate(w io.Writer, info interface{}, templ string) error {
	t := template.Must(template.New("").
		Funcs(template.FuncMap{}).Parse(templ))

	return t.Execute(w, info)
}
