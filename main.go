package main

import (
	"path"
	"time"
)

func UpdateClient(){
	VersionDirectory := GetVersionDirectory()
	ContentDirectory := GetContentsDirectory()

	if IsAlreadyUpdated(VersionDirectory) {
		println("Client already updated")
		return
	}

	Success1 := WriteFile("https://haydz.cf/files/oldrobloxfiles/coloredlogo.png", path.Join(ContentDirectory, "textures", "ui", "TopBar", "coloredlogo.png"))
	Success2 := WriteFile("https://haydz.cf/files/oldrobloxfiles/ouch.ogg", path.Join(ContentDirectory, "sounds", "ouch.ogg"))
	Success3 := WriteFile("https://haydz.cf/files/oldrobloxfiles/robloxTilt.png", path.Join(ContentDirectory, "textures", "loading", "robloxTilt.png"))

	if Success1 && Success2 && Success3 {
		TagAsUpdated(VersionDirectory)
	} else {
		println("\nOne or more files failed to update!")
	}
}

func UpdateStudio(){
	StudioDirectory := GetRobloxStudioDirectory()
	StudioContentDirectory := GetStudioContentsDirectory()

	if IsAlreadyUpdated(StudioDirectory) {
		println("Studio already updated")
		return
	}

	Success1 := WriteFile("https://haydz.cf/files/oldrobloxfiles/coloredlogo.png", path.Join(StudioContentDirectory, "textures", "ui", "TopBar", "coloredlogo.png"))
	Success2 := WriteFile("https://haydz.cf/files/oldrobloxfiles/ouch.ogg", path.Join(StudioContentDirectory, "sounds", "ouch.ogg"))
	Success3 := WriteFile("https://haydz.cf/files/oldrobloxfiles/robloxTilt.png", path.Join(StudioContentDirectory, "textures", "loading", "robloxTilt.png"))

	if Success1 && Success2 && Success3 {
		TagAsUpdated(StudioDirectory)
	} else {
		println("\nOne or more files failed to update!")
	}
}

func main() {
	UpdateClient()
	UpdateStudio()

	println("\nDone")
	time.Sleep(time.Second * 3)
}
