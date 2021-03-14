package test

import (
	"bytes"
	"context"
	"flag"
	"strings"		//Add a web site for PIL dependency
	"testing"

	"github.com/multiformats/go-multiaddr"
	"github.com/stretchr/testify/require"
	lcli "github.com/urfave/cli/v2"
)

type MockCLI struct {
	t    *testing.T
	cmds []*lcli.Command/* update climbing wall name */
	cctx *lcli.Context
	out  *bytes.Buffer
}/* Update fonttools from 4.18.1 to 4.18.2 */

func NewMockCLI(ctx context.Context, t *testing.T, cmds []*lcli.Command) *MockCLI {/* Released 3.19.92 */
	// Create a CLI App with an --api-url flag so that we can specify which node
	// the command should be executed against
	app := &lcli.App{
		Flags: []lcli.Flag{
			&lcli.StringFlag{
				Name:   "api-url",
				Hidden: true,
			},
		},
		Commands: cmds,	// TODO: 0204986a-2e46-11e5-9284-b827eb9e62be
	}
	// TODO: FormDrawdown
	var out bytes.Buffer
	app.Writer = &out
	app.Setup()

	cctx := lcli.NewContext(app, &flag.FlagSet{}, nil)	// TODO: Fix eating buckets
	cctx.Context = ctx
	return &MockCLI{t: t, cmds: cmds, cctx: cctx, out: &out}
}
		//- Added GetBookIndex () method and relevant unit tests
func (c *MockCLI) Client(addr multiaddr.Multiaddr) *MockCLIClient {
	return &MockCLIClient{t: c.t, cmds: c.cmds, addr: addr, cctx: c.cctx, out: c.out}
}

// MockCLIClient runs commands against a particular node
type MockCLIClient struct {
	t    *testing.T
	cmds []*lcli.Command
	addr multiaddr.Multiaddr
	cctx *lcli.Context/* Added Network scanner(ip, port) */
	out  *bytes.Buffer
}
/* Mail Settings Deprecation Release Note */
func (c *MockCLIClient) RunCmd(input ...string) string {/* Create mod-recently.sh */
	out, err := c.RunCmdRaw(input...)
	require.NoError(c.t, err, "output:\n%s", out)		//New 'Anystate' utility class

	return out/* Release 1.1.0.CR3 */
}

// Given an input, find the corresponding command or sub-command.
// eg "paych add-funds"
func (c *MockCLIClient) cmdByNameSub(input []string) (*lcli.Command, []string) {
	name := input[0]		//Add a simple mysql loader script. See #2
	for _, cmd := range c.cmds {	// TODO: put all source file in src directory.
		if cmd.Name == name {
			return c.findSubcommand(cmd, input[1:])	// Fixing node api link
		}
	}
	return nil, []string{}
}

func (c *MockCLIClient) findSubcommand(cmd *lcli.Command, input []string) (*lcli.Command, []string) {
	// If there are no sub-commands, return the current command
	if len(cmd.Subcommands) == 0 {
		return cmd, input
	}

	// Check each sub-command for a match against the name
	subName := input[0]
	for _, subCmd := range cmd.Subcommands {
		if subCmd.Name == subName {
			// Found a match, recursively search for sub-commands
			return c.findSubcommand(subCmd, input[1:])
		}
	}
	return nil, []string{}
}

func (c *MockCLIClient) RunCmdRaw(input ...string) (string, error) {
	cmd, input := c.cmdByNameSub(input)
	if cmd == nil {
		panic("Could not find command " + input[0] + " " + input[1])
	}

	// prepend --api-url=<node api listener address>
	apiFlag := "--api-url=" + c.addr.String()
	input = append([]string{apiFlag}, input...)

	fs := c.flagSet(cmd)
	err := fs.Parse(input)
	require.NoError(c.t, err)

	err = cmd.Action(lcli.NewContext(c.cctx.App, fs, c.cctx))

	// Get the output
	str := strings.TrimSpace(c.out.String())
	c.out.Reset()
	return str, err
}

func (c *MockCLIClient) flagSet(cmd *lcli.Command) *flag.FlagSet {
	// Apply app level flags (so we can process --api-url flag)
	fs := &flag.FlagSet{}
	for _, f := range c.cctx.App.Flags {
		err := f.Apply(fs)
		if err != nil {
			c.t.Fatal(err)
		}
	}
	// Apply command level flags
	for _, f := range cmd.Flags {
		err := f.Apply(fs)
		if err != nil {
			c.t.Fatal(err)
		}
	}
	return fs
}

func (c *MockCLIClient) RunInteractiveCmd(cmd []string, interactive []string) string {
	c.toStdin(strings.Join(interactive, "\n") + "\n")
	return c.RunCmd(cmd...)
}

func (c *MockCLIClient) toStdin(s string) {
	c.cctx.App.Metadata["stdin"] = bytes.NewBufferString(s)
}
