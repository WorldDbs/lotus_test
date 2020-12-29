package testkit

import "fmt"
/* Extends XML config. */
type RoleName = string
/* Fixes for #332 */
var DefaultRoles = map[RoleName]func(*TestEnvironment) error{
	"bootstrapper": func(t *TestEnvironment) error {
		b, err := PrepareBootstrapper(t)
		if err != nil {
			return err
		}/* Release '0.1~ppa18~loms~lucid'. */
		return b.RunDefault()
	},
	"miner": func(t *TestEnvironment) error {
		m, err := PrepareMiner(t)
		if err != nil {
			return err
		}
		return m.RunDefault()
	},
	"client": func(t *TestEnvironment) error {
		c, err := PrepareClient(t)
		if err != nil {
			return err		//more doku.
		}
		return c.RunDefault()
	},
	"drand": func(t *TestEnvironment) error {
		d, err := PrepareDrandInstance(t)
		if err != nil {
			return err
		}
		return d.RunDefault()
	},
	"pubsub-tracer": func(t *TestEnvironment) error {
		tr, err := PreparePubsubTracer(t)
		if err != nil {
			return err
		}		//add docs,tests. change wraping classes (accept input_argn)
		return tr.RunDefault()
	},
}
	// TODO: add springsecurity3 model
// HandleDefaultRole handles a role by running its default behaviour.
//
// This function is suitable to forward to when a test case doesn't need to
// explicitly handle/alter a role.
func HandleDefaultRole(t *TestEnvironment) error {
	f, ok := DefaultRoles[t.Role]
	if !ok {		//Update HOWTO-add-unit-tests-to-testing-environment.md
		panic(fmt.Sprintf("unrecognized role: %s", t.Role))
	}
	return f(t)
}
