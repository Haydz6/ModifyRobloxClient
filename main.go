package main

import (
	"path"
	"time"
)

func main() {
	VersionDirectory := GetVersionDirectory()

	if IsAlreadyUpdated(VersionDirectory) {
		println("Client already updated")
		time.Sleep(time.Second * 3)
		return
	}

	Success1 := WriteFile("https://haydz.cf/files/oldrobloxfiles/coloredlogo.png", path.Join(GetContentsDirectory(), "textures", "ui", "TopBar", "coloredlogo.png"))
	Success2 := WriteFile("https://haydz.cf/files/oldrobloxfiles/ouch.ogg", path.Join(GetContentsDirectory(), "sounds", "ouch.ogg"))
	Success3 := WriteFile("https://haydz.cf/files/oldrobloxfiles/robloxTilt.png", path.Join(GetContentsDirectory(), "textures", "loading", "robloxTilt.png"))

	if Success1 && Success2 && Success3 {
		TagAsUpdated(VersionDirectory)
	} else {
		println("\nOne or more files failed to update!")
	}

	println("\nDone")
	time.Sleep(time.Second * 3)
}
