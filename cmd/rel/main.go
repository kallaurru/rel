package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kallaurru/rel/cmd/rel/internal"
	"github.com/subosito/gotenv"
)

var (
	version = ""
)

/*
	Требуется добавить:
		- читать файла env, и брать имя проекта по которому будет определять проект из которого брать переменные
		- создавать шаблоны файлов миграций в указанном каталоге
*/

func main() {
	var (
		err error
		ctx = context.Background()
	)

	log.SetFlags(0)
	fmt.Printf("run programm: %s\n", os.Args[0])
	path := internal.ValidateRootEnvFile(os.Args[0])
	if path == "" {
		fmt.Println(internal.MsgBadRootEnvFile())
		os.Exit(1)
	}

	err = gotenv.Load(path)
	if err != nil {
		fmt.Printf("Env file = %s not uploaded", path)
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("Available command are: migrate, rollback, make")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "migrate", "up", "rollback", "down":
		err = internal.ExecMigrate(ctx, os.Args)
	case "make":
		err = internal.ExecMakeMigration(os.Args)
	case "version", "-v", "-version":
		fmt.Println("REL CLI " + version)
	case "-help":
		fmt.Println("Usage: rel [command] -help")
		fmt.Println("Available commands: migrate, rollback")
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		log.Fatal(err)
	}
}
