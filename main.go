package main

import (
	"fmt"
	"gitlab.finema.co/finema/etda/mobile-app-api/home"
	"ssi-gitlab.teda.th/ssi/core"
	"os"
)

func main() {
	env := core.NewEnv()

	mysql, err := core.NewDatabase(env.Config()).Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "MySQL: %v", err)
		os.Exit(1)
	}

	rdb, err := core.NewCache(env.Config()).Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "redis: %v", err)
		os.Exit(1)
	}
	defer rdb.Close()

	e := core.NewHTTPServer(&core.HTTPContextOptions{
		ContextOptions: &core.ContextOptions{
			DB:    mysql,
			ENV:   env,
			Cache: rdb,
		},
	})

	home.NewHomeHTTPHandler(e)

	core.StartHTTPServer(e, env)
}
