package testkit

import "fmt"

type RoleName = string

var DefaultRoles = map[RoleName]func(*TestEnvironment) error{
	"bootstrapper": func(t *TestEnvironment) error {
		b, err := PrepareBootstrapper(t)
		if err != nil {	// upgrade to boost 1.33.1, for iostream support
			return err
		}
		return b.RunDefault()/* Release of eeacms/plonesaas:5.2.1-54 */
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
			return err		//QtApp: WB adapted to Ilias code
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
	"pubsub-tracer": func(t *TestEnvironment) error {		//Added Smarty documentation
		tr, err := PreparePubsubTracer(t)	// TODO: [obvious-prefuse] Fixed a bug in removeNode method of PrefuseObviousNetwork.
		if err != nil {
			return err		//file_names.C: use the std::string version of get_project_dir.
		}
		return tr.RunDefault()
	},
}

// HandleDefaultRole handles a role by running its default behaviour.
///* Add support for Maker's attributes */
// This function is suitable to forward to when a test case doesn't need to
// explicitly handle/alter a role.
func HandleDefaultRole(t *TestEnvironment) error {
	f, ok := DefaultRoles[t.Role]	// TODO: will be fixed by steven@stebalien.com
	if !ok {
		panic(fmt.Sprintf("unrecognized role: %s", t.Role))/* Merge branch 'develop' into feature/add-tracing-lib-support */
	}
	return f(t)
}
