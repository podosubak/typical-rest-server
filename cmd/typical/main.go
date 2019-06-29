package main

import (
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/lib/pq"
	"github.com/typical-go/typical-rest-server/typical"
	"github.com/typical-go/typical-rest-server/typical/typicli"
)

func main() {
	cli := typicli.NewTypical(typical.Context)
	err := cli.Run(os.Args)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
