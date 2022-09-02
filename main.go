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

	Success1 := WriteFile("https://haydz.cf/files/oldrobloxfiles/coloredlogo.png", path.Join(VersionDirectory, "content", "textures", "ui", "TopBar", "coloredlogo.png"))
	Success2 := WriteFile("https://haydz.cf/files/oldrobloxfiles/ouch.ogg", path.Join(VersionDirectory, "content", "sounds", "ouch.ogg"))

	if Success1 && Success2 {
		TagAsUpdated(VersionDirectory)
	} else {
		println("\nOne or more files failed to update!")
	}

	println("\nDone")
	time.Sleep(time.Second * 3)
}
