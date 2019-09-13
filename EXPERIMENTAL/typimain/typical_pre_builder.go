package typimain

import (
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typigen"
	"github.com/urfave/cli"
)

// TypicalPreBuilder represent typical generator
type TypicalPreBuilder struct {
	*typictx.Context
}

// NewTypicalPreBuilder return new instance of TypicalCli
func NewTypicalPreBuilder(context *typictx.Context) *TypicalPreBuilder {
	return &TypicalPreBuilder{
		Context: context,
	}
}

// Cli return the command line interface
func (g *TypicalPreBuilder) Cli() *cli.App {
	app := cli.NewApp()
	app.Action = g.run
	return app
}

func (g *TypicalPreBuilder) run(ctx *cli.Context) error {
	return typigen.Generate(g.Context)
}