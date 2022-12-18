package cli

import "github.com/alecthomas/kong"

type CLI struct {
	GenerateSecretHost GenerateSecretHost `kong:"cmd,help='Generate same proxy secret with different host name'"`
	GenerateSecret     GenerateSecret     `kong:"cmd,help='Generate new proxy secret'"`
	Access             Access             `kong:"cmd,help='Print access information.'"`
	Run                Run                `kong:"cmd,help='Run proxy.'"`
	SimpleRun          SimpleRun          `kong:"cmd,help='Run proxy without config file.'"`
	Version            kong.VersionFlag   `kong:"help='Print version.',short='v'"`
}
