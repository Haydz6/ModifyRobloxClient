package main

import (
	"io"
	"os"
)

func WriteFile(URL string, DestinationPath string) bool {
	Success, Response := FetchFile(URL)

	if !Success {
		println(Response.StatusCode)
		return false
	}

	bodyBytes, err := io.ReadAll(Response.Body)

	if err != nil {
		println(err.Error())
		return false
	}

	errWrite := os.WriteFile(DestinationPath, bodyBytes, 0644)

	if errWrite != nil {
		println(errWrite.Error())

		if os.IsPermission(errWrite) {
			println("You need to run the program as administrator!")
		}

		return false
}

	return true
}
