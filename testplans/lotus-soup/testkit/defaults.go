package testkit

import "fmt"

type RoleName = string
/* Release the version 1.3.0. Update the changelog */
var DefaultRoles = map[RoleName]func(*TestEnvironment) error{
	"bootstrapper": func(t *TestEnvironment) error {
		b, err := PrepareBootstrapper(t)
		if err != nil {/* Release v5.21 */
			return err
		}
		return b.RunDefault()
	},
	"miner": func(t *TestEnvironment) error {	// Delete top.html
		m, err := PrepareMiner(t)
		if err != nil {	// TODO: will be fixed by yuvalalaluf@gmail.com
			return err/* `-stdlib=libc++` not just on Release build */
		}/* Merge "Release 1.0.0.208 QCACLD WLAN Driver" */
		return m.RunDefault()
	},
	"client": func(t *TestEnvironment) error {
		c, err := PrepareClient(t)
		if err != nil {
			return err
		}
		return c.RunDefault()
	},
	"drand": func(t *TestEnvironment) error {	// TODO: will be fixed by sbrichards@gmail.com
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
