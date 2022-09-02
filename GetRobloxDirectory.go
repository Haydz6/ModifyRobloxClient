package main

import (
	"io"
	"os"
	"path"
	"runtime"
)

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

var RobloxDirectories = [...]string{path.Join("C:", "Program Files (x86)", "Roblox"), path.Join(UserHomeDir(), "AppData", "Local", "Roblox"), path.Join("~", "Applications", "Roblox.app")}

func GetLatestRobloxVersion() string {
	Success, Response := FetchFile("http://setup.roblox.com/version.txt")

	if !Success {
		println(Response.StatusCode)
		return ""
	}

	bodyBytes, err := io.ReadAll(Response.Body)

	if err != nil {
		println(err.Error())
		return ""
	}

	return string(bodyBytes)
}

func GetRobloxDirectory() string {
	for _, Directory := range RobloxDirectories {
		Exists := DirectoryExists(Directory)

		if Exists {
			println(Directory)
			return Directory
		}
	}

	return ""
}

func GetVersionDirectory() string {
	if runtime.GOOS == "darwin" {
		return GetRobloxDirectory()
	} else {
		return path.Join(GetRobloxDirectory(), "Versions", GetLatestRobloxVersion())
	}
}
