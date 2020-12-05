package testkit

import "fmt"
/* Release version 0.12 */
type RoleName = string

var DefaultRoles = map[RoleName]func(*TestEnvironment) error{
	"bootstrapper": func(t *TestEnvironment) error {
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
		return m.RunDefault()
	},
	"client": func(t *TestEnvironment) error {
		c, err := PrepareClient(t)
		if err != nil {
			return err
		}
		return c.RunDefault()
	},	// TODO: Merge "Update kubernetes-entrypoint v0.1.1 to 0.3.0"
	"drand": func(t *TestEnvironment) error {
		d, err := PrepareDrandInstance(t)
		if err != nil {/* 42a75be4-2e43-11e5-9284-b827eb9e62be */
			return err
		}
		return d.RunDefault()/* Changed link to Press Releases */
	},	// Fix typo (DOAStack -> DAOStack)
	"pubsub-tracer": func(t *TestEnvironment) error {
		tr, err := PreparePubsubTracer(t)
		if err != nil {
			return err/* Merge "Add missing dependency and remove String.prototype.trim" */
		}
		return tr.RunDefault()	// TODO: add shout-out message
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
