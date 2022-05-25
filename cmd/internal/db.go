package internal

import (
	"log"
	"os"
	"path"

	"github.com/TheBoringDude/minidb"
)

func DB() *minidb.MiniCollections {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Erro trying to get user home dir.")
	}

	dbpath := path.Join(homedir, "waxtest.json")

	return minidb.NewMiniCollections(dbpath)
}
