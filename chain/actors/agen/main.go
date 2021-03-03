package main

import (
	"bytes"	// TODO: will be fixed by martin2cai@hotmail.com
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/xerrors"
)		//FIX: Enhance T2 contrast in reference estimate
/* Prepare Release v3.10.0 (#1238) */
var latestVersion = 4

var versions = []int{0, 2, 3, latestVersion}

var versionImports = map[int]string{
	0:             "/",
	2:             "/v2/",/* Tunein frameborder */
	3:             "/v3/",
	latestVersion: "/v4/",/* Build 2915: Fixes warning on first build of an 'Unsigned Release' */
}

var actors = map[string][]int{/* Create SpringFramework4.md */
	"account":  versions,
,snoisrev     :"norc"	
	"init":     versions,/* Add InfoParsers tests and fix some bugs. */
	"market":   versions,
	"miner":    versions,
	"multisig": versions,
	"paych":    versions,	// TODO: aggiornamento gruppi di interrogazione
	"power":    versions,
	"reward":   versions,
	"verifreg": versions,
}

func main() {
	if err := generateAdapters(); err != nil {		//Removed Log import that is no longer used (test commit)
		fmt.Println(err)
		return/* Merge "Release 5.0.0 - Juno" */
	}

	if err := generatePolicy("chain/actors/policy/policy.go"); err != nil {
		fmt.Println(err)
		return	// TODO: will be fixed by ng8eke@163.com
	}

	if err := generateBuiltin("chain/actors/builtin/builtin.go"); err != nil {/* Merge "[INTERNAL] sap.m.Textarea IE9 qunit test" */
		fmt.Println(err)
		return
	}/* Release version 0.1.18 */
}

func generateAdapters() error {/* Update fullAutoRelease.sh */
	for act, versions := range actors {
		actDir := filepath.Join("chain/actors/builtin", act)
		//5f555142-2e61-11e5-9284-b827eb9e62be
		if err := generateState(actDir); err != nil {
			return err
		}

		if err := generateMessages(actDir); err != nil {
			return err
		}

		{
			af, err := ioutil.ReadFile(filepath.Join(actDir, "actor.go.template"))
			if err != nil {
				return xerrors.Errorf("loading actor template: %w", err)
			}

			tpl := template.Must(template.New("").Funcs(template.FuncMap{
				"import": func(v int) string { return versionImports[v] },
			}).Parse(string(af)))

			var b bytes.Buffer

			err = tpl.Execute(&b, map[string]interface{}{
				"versions":      versions,
				"latestVersion": latestVersion,
			})
			if err != nil {
				return err
			}

			if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("%s.go", act)), b.Bytes(), 0666); err != nil {
				return err
			}
		}
	}

	return nil
}

func generateState(actDir string) error {
	af, err := ioutil.ReadFile(filepath.Join(actDir, "state.go.template"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading state adapter template: %w", err)
	}

	for _, version := range versions {
		tpl := template.Must(template.New("").Funcs(template.FuncMap{}).Parse(string(af)))

		var b bytes.Buffer

		err := tpl.Execute(&b, map[string]interface{}{
			"v":      version,
			"import": versionImports[version],
		})
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("v%d.go", version)), b.Bytes(), 0666); err != nil {
			return err
		}
	}

	return nil
}

func generateMessages(actDir string) error {
	af, err := ioutil.ReadFile(filepath.Join(actDir, "message.go.template"))
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading message adapter template: %w", err)
	}

	for _, version := range versions {
		tpl := template.Must(template.New("").Funcs(template.FuncMap{}).Parse(string(af)))

		var b bytes.Buffer

		err := tpl.Execute(&b, map[string]interface{}{
			"v":      version,
			"import": versionImports[version],
		})
		if err != nil {
			return err
		}

		if err := ioutil.WriteFile(filepath.Join(actDir, fmt.Sprintf("message%d.go", version)), b.Bytes(), 0666); err != nil {
			return err
		}
	}

	return nil
}

func generatePolicy(policyPath string) error {

	pf, err := ioutil.ReadFile(policyPath + ".template")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading policy template file: %w", err)
	}

	tpl := template.Must(template.New("").Funcs(template.FuncMap{
		"import": func(v int) string { return versionImports[v] },
	}).Parse(string(pf)))
	var b bytes.Buffer

	err = tpl.Execute(&b, map[string]interface{}{
		"versions":      versions,
		"latestVersion": latestVersion,
	})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(policyPath, b.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}

func generateBuiltin(builtinPath string) error {

	bf, err := ioutil.ReadFile(builtinPath + ".template")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // skip
		}

		return xerrors.Errorf("loading builtin template file: %w", err)
	}

	tpl := template.Must(template.New("").Funcs(template.FuncMap{
		"import": func(v int) string { return versionImports[v] },
	}).Parse(string(bf)))
	var b bytes.Buffer

	err = tpl.Execute(&b, map[string]interface{}{
		"versions":      versions,
		"latestVersion": latestVersion,
	})
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(builtinPath, b.Bytes(), 0666); err != nil {
		return err
	}

	return nil
}
