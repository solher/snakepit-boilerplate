package cmd

import (
	"github.com/solher/snakepit/root"
	"github.com/solher/snakepit/run"
	"git.wid.la/versatile/versatile-server/app"
	"git.wid.la/versatile/versatile-server/constants"
)

func init() {
	run.Builder = app.Builder

	// SERVICES
	run.Cmd.PersistentFlags().String("authServerUrl", "", "auth server URL")
	root.Viper.BindPFlag(constants.AuthServerURL, run.Cmd.PersistentFlags().Lookup("authServerUrl"))
	root.Viper.RegisterAlias(constants.AuthServerURL, "AUTH_SERVER_PORT")
}