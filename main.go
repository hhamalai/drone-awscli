package main

import (
  "os"

  "github.com/sirupsen/logrus"
  "github.com/urfave/cli"
)

var revision string // build number set at compile-time

func main() {
  app := cli.NewApp()
  app.Name = "awscli plugin"
  app.Usage = "awscli plugin"
  app.Action = run
  app.Version = revision
  app.Flags = []cli.Flag{

    //
    // plugin args
    //

    cli.StringSliceFlag{
      Name:   "actions",
      Usage:  "a list of actions to have terraform perform",
      EnvVar: "PLUGIN_ACTIONS",
      Value:  &cli.StringSlice{"validate", "plan", "apply"},
    },
    cli.StringFlag{
      Name:   "assume_role",
      Usage:  "A role to assume before running the awscli commands",
      EnvVar: "PLUGIN_ASSUME_ROLE",
    },
    cli.StringFlag{
      Name:   "awscli_version",
      Usage:  "AWSCli version number",
      EnvVar: "PLUGIN_AWSCLI_VERSION",
    },
    cli.StringFlag{
      Name:    "awscli_command",
      Usage:   "AWSCli command to be run",
      EnvVar:  "PLUGIN_COMMAND",
    },
  }

  if err := app.Run(os.Args); err != nil {
    logrus.Fatal(err)
  }
}

func run(c *cli.Context) error {
  plugin := Plugin{
    Config: Config{
      RoleARN:          c.String("assume_role"),
    },
    AWSCli: AWSCli{
      Version:          c.String("awscli_version"),
      Command:          c.String("awscli_command"),
    },
  }

  return plugin.Exec()
}