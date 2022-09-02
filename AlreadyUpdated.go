package main

import (
	"os"
	"path"
)

func TagAsUpdated(Directory string) {
	file, err := os.Create(path.Join(Directory, "ClientAlreadyModified"))

	if err != nil {
		println(err.Error())
		return
	}

	file.Close()
}

func IsAlreadyUpdated(Directory string) bool {
	_, err := os.Stat(path.Join(Directory, "ClientAlreadyModified"))

	return os.IsExist(err)
}
