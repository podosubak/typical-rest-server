package typrest

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/typical-go/typical-go/pkg/utility/runn"
	"github.com/typical-go/typical-go/pkg/utility/runner"
	"github.com/urfave/cli/v2"
)

// Module of typrest
type Module struct{}

// BuildCommands is commands to exectuce from Build-Tool
func (m *Module) BuildCommands(c *typcli.BuildCli) []*cli.Command {
	return []*cli.Command{
		{
			Name:  "rest",
			Usage: "rest application utility",
			Subcommands: []*cli.Command{
				{Name: "scaffold", Aliases: []string{"s"}, Usage: "Scaffold the MVC", Action: m.scaffold},
			},
		},
	}
}

func (m *Module) scaffold(ctx *cli.Context) (err error) {
	e := Entity{
		Name:  "Music",
		Table: "musics",
	}
	return runn.Execute(
		runner.WriteTemplate{
			Target:   fmt.Sprintf("app/repository/%s_repo.go", strings.ToLower(e.Name)),
			Template: repoTmpl,
			Data:     e,
		},
	)
}

type Entity struct {
	Name  string
	Table string
}