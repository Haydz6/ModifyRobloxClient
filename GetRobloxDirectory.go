package main

import (
	"io"
	"path"
)

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
	return path.Join("C:", "Program Files (x86)", "Roblox")
}

func GetVersionDirectory() string {
	return path.Join(GetRobloxDirectory(), "Versions", GetLatestRobloxVersion())
}
