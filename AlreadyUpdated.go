package main

import (
	"os"
	"path"
)

func DirectoryExists(Directory string) bool {
	_, err := os.Stat(Directory)

	if err == nil {
		return true
	}

	return os.IsExist(err)
}

func TagAsUpdated(Directory string) {
	file, err := os.Create(path.Join(Directory, "ClientAlreadyModified"))

	if err != nil {
		println(err.Error())
		return
	}

	file.Close()
}

func IsAlreadyUpdated(Directory string) bool {
	return DirectoryExists(path.Join(Directory, "ClientAlreadyModified"))
}
