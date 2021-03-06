package main

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"runtime"
	"strconv"
	"strings"

	bp "github.com/nexustix/boilerplate"
)

//mcinterface setupprofile <modpackName> <minecraftVersion> <instanceDir>

func main() {
	/*
		var zeMadness McConfig

		usr, err := user.Current()
		bp.FailError(err)
		workingDir := usr.HomeDir
		minecraftDir := path.Join(workingDir, ".minecraft")
		os.MkdirAll(minecraftDir, 0777)

		filePath := path.Join(minecraftDir, "launcher_profiles.json")


		/////LOAD

		if bp.FileExists(filePath) {
			dat, err := ioutil.ReadFile(filePath)
			bp.FailError(err)

			json.Unmarshal(dat, &zeMadness)
		}


		/////DO STUFF


		for k, v := range zeMadness.Profiles {
			fmt.Printf("%v %v\n", k, v)
		}

		tmpProfile := Profile{}

		tmpProfile.Name = "cheseCake"
		tmpProfile.LastVersionID = "1.10.2-forge1.10.2-12.18.1.2011"
		tmpProfile.GameDir = path.Join(workingDir, "cheeseCake")

		zeMadness.Profiles[tmpProfile.Name] = tmpProfile


		/////SAVE

		outFile, err := os.Create(filePath)
		bp.FailError(err)
		//tmpJSON, _ := json.Marshal(zeMadness)
		tmpJSON, _ := json.MarshalIndent(zeMadness, "", "  ")
		outFile.Write(tmpJSON)
		outFile.Close()
	*/

	/*
		var tmpProfileConfig ProfileConfig

		usr, err := user.Current()
		bp.FailError(err)
		workingDir := usr.HomeDir
		minecraftDir := path.Join(workingDir, ".minecraft")
		os.MkdirAll(minecraftDir, 0777)

		filePath := path.Join(minecraftDir, "launcher_profiles.json")

		tmpProfileConfig.LoadFromFile(filePath)

		tmpProfile := Profile{}

		forgeVersions := GetLocalForgeVersions(minecraftDir)
		if len(forgeVersions) >= 1 {
			tmpProfile.Name = "cheseCake"
			tmpProfile.LastVersionID = forgeVersions[0]
			tmpProfile.GameDir = path.Join(workingDir, "cheeseCake")

			tmpProfileConfig.AddProfile(tmpProfile)
			tmpProfileConfig.SaveToFile(filePath)
		} else {
			fmt.Printf("<!> No forge installation found please install forge\n")
		}
	*/

	//usr, err := user.Current()
	//bp.FailError(err)
	//workingDir := usr.HomeDir
	//minecraftDir := path.Join(workingDir, ".minecraft")

	var tmpProfileConfig ProfileConfig

	usr, err := user.Current()
	bp.FailError(err)
	workingDir := usr.HomeDir
	minecraftDir := path.Join(workingDir, ".minecraft")
	if runtime.GOOS == "windows" {
		minecraftDir = path.Join(workingDir, "AppData", "Roaming", ".minecraft")
	}
	os.MkdirAll(minecraftDir, 0777)

	filePath := path.Join(minecraftDir, "launcher_profiles.json")

	tmpProfileConfig.LoadFromFile(filePath)

	tmpProfile := Profile{}

	forgeVersions := GetLocalForgeVersions(minecraftDir)

	args := os.Args

	programAction := bp.StringAtIndex(1, args)
	modpackName := bp.StringAtIndex(2, args)
	//minecraftVersion := bp.StringAtIndex(3, args)
	instanceDir := bp.StringAtIndex(4, args)

	switch programAction {
	case "setupprofile":
		if len(forgeVersions) >= 1 {
			tmpProfile.Name = modpackName
			tmpProfile.LastVersionID = getLatestForgeVersionID(forgeVersions)
			tmpProfile.GameDir = instanceDir
			tmpProfile.JavaArgs = "-Xmx2G"

			tmpProfileConfig.AddProfile(tmpProfile)
			tmpProfileConfig.SelectedProfile = modpackName
			tmpProfileConfig.SaveToFile(filePath)
		} else {
			fmt.Printf("<!> No forge installation found please install forge\n")
		}

	}

}

func getLatestForgeVersionID(forgeVersions []string) string {
	largestInt := 0
	latestVersionString := ""
	for _, v := range forgeVersions {
		segments := strings.Split(v, ".")
		if len(segments) > 0 {
			i, err := strconv.ParseInt(segments[len(segments)-1], 10, 32)
			bp.FailError(err)
			if int(i) > largestInt {
				latestVersionString = v
				largestInt = int(i)
			}
		}
	}
	return latestVersionString
}
