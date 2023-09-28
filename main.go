package main

import (
	"log"
	"path/filepath"

	"github.com/Ed-cred/tasker/cmd"
	"github.com/Ed-cred/tasker/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(home, "tasker.db")
	err = db.Init(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	cmd.RootCmd.Execute()
}
