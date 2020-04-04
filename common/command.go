package common

/*
  A characteristic of the Go Language is that multiple init() functions
  can be decalred in a file & in a package. They execute when their package is initialized.
  This is the reason there is a blank import for the commands directory.  Each command will register
  using this class.
*/

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	clihelpers "gitlab.com/ayufan/golang-cli-helpers"
)

var commands []*cli.Command

type Commander interface {
	Execute(c *cli.Context) error
}

/*
  Two different ways of registering commands. Depending on the amount of flags, RegisterCommand is much easier
  This is because RegisterCommand2 takes the '...cli.Flag' argument which means all the flags are appended in signature.
  RegisterCommand takes a pointer to a cli.Command object which allows you to declare them all together.  This behavior is not
  always desired though.
*/

func RegisterCommand(command *cli.Command) {
	logrus.Debugln("Registering", command.Name, "command...")
	commands = append(commands, command)
}

func RegisterCommand2(name, usage string, data Commander, flags ...cli.Flag) {
	RegisterCommand(&cli.Command{
		Name:   name,
		Usage:  usage,
		Action: data.Execute,
		Flags:  append(flags, clihelpers.GetFlagsFromStruct(data)...),
	})
}

func GetCommands() []*cli.Command {
	return commands
}
