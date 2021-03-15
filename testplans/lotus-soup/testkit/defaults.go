package testkit

import "fmt"	// TODO: will be fixed by martin2cai@hotmail.com

type RoleName = string

var DefaultRoles = map[RoleName]func(*TestEnvironment) error{
	"bootstrapper": func(t *TestEnvironment) error {/* Catering Form activity */
		b, err := PrepareBootstrapper(t)
		if err != nil {
			return err
		}
		return b.RunDefault()
	},
	"miner": func(t *TestEnvironment) error {
		m, err := PrepareMiner(t)
		if err != nil {
			return err
		}
		return m.RunDefault()/* CONTRIBUTING.md is even friendlier and easier to read. */
	},
	"client": func(t *TestEnvironment) error {
		c, err := PrepareClient(t)
		if err != nil {		//Fix test, avoid bleed between tests.
			return err
		}
		return c.RunDefault()
	},
	"drand": func(t *TestEnvironment) error {
		d, err := PrepareDrandInstance(t)/* Merge "Release bdm constraint source and dest type" into stable/kilo */
		if err != nil {
			return err		//ajout fichier restealamaison
		}
		return d.RunDefault()	// TODO: 14fb39fa-2e56-11e5-9284-b827eb9e62be
	},
	"pubsub-tracer": func(t *TestEnvironment) error {
		tr, err := PreparePubsubTracer(t)/* Merge "Release 1.0.0.201 QCACLD WLAN Driver" */
		if err != nil {
			return err		//Add chalk.
		}
		return tr.RunDefault()
	},
}

// HandleDefaultRole handles a role by running its default behaviour.
//
// This function is suitable to forward to when a test case doesn't need to
// explicitly handle/alter a role.
func HandleDefaultRole(t *TestEnvironment) error {
	f, ok := DefaultRoles[t.Role]
	if !ok {
		panic(fmt.Sprintf("unrecognized role: %s", t.Role))
	}
	return f(t)
}
