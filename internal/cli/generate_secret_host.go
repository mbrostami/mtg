package cli

import (
	"fmt"

	"github.com/9seconds/mtg/v2/internal/utils"
)

type GenerateSecretHost struct {
	ConfigPath string `kong:"arg,required,type='existingfile',help='Path to the configuration file.',name='config-path'"` //nolint: lll
	HostName   string `kong:"arg,required,help='Hostname to use for domain fronting.',name='hostname'"`
	Hex        bool   `kong:"help='Print secret in hex encoding.',short='x'"`
}

func (g *GenerateSecretHost) Run(cli *CLI, _ string) error {
	conf, err := utils.ReadConfig(g.ConfigPath)
	if err != nil {
		return fmt.Errorf("cannot init config: %w", err)
	}

	conf.Secret.Host = cli.GenerateSecretHost.HostName

	if g.Hex {
		fmt.Println(conf.Secret.Hex()) //nolint: forbidigo
	} else {
		fmt.Println(conf.Secret.Base64()) //nolint: forbidigo
	}

	return nil
}
