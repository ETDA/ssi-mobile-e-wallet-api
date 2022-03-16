package main

import (
	"fmt"
	"os"

	"gitlab.finema.co/finema/etda/mobile-app-api/seeds/seeds"
	core "ssi-gitlab.teda.th/ssi/core"
)

func main() {
	env := core.NewEnv()

	mysql, err := core.NewDatabase(env.Config()).Connect()
	if err != nil {
		fmt.Fprintf(os.Stderr, "mysql: %v", err)
		os.Exit(1)
	}

	ctx := core.NewContext(&core.ContextOptions{
		DB:  mysql,
		ENV: env,
	})

	seeder := core.NewSeeder()

	_ = seeder.Add(seeds.NewConfigDIDSeed(ctx))
	_ = seeder.Add(seeds.NewTokenSeed(ctx))
}
