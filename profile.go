package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	bp "github.com/nexustix/boilerplate"
)

/*
{
  "profiles": {
    "listings": {
      "name": "listings",
      "gameDir": "/home/foo/.minecraft",
      "lastVersionId": "1.10.2",
      "javaDir": "/usr/lib/jvm/java-8-openjdk/jre/bin/java",
      "javaArgs": "-Xmx1G -XX:+UseConcMarkSweepGC -XX:+CMSIncrementalMode -XX:-UseAdaptiveSizePolicy -Xmn128M",
      "resolution": {
        "width": 854,
        "height": 480
      },
      "allowedReleaseTypes": [
        "release",
        "old_beta",
        "snapshot",
        "old_alpha"
      ],
      "launcherVisibilityOnGameClose": "close launcher when game starts"
    },
    "foo": {
      "name": "foo"
    },
    "forge": {
      "name": "forge",
      "lastVersionId": "1.10.2-forge1.10.2-12.18.1.2011"
    }
  },
  "selectedProfile": "foo",
  "clientToken": "foo",
  "authenticationDatabase": {
    "foo": {
      "displayName": "foo",
      "accessToken": "foo",
      "userid": "foo",
      "uuid": "foo",
      "username": "foo"
    }
  },
  "selectedUser": "foo",
  "launcherVersion": {
    "name": "1.6.61",
    "format": 18
  }
}
*/

type Resolution struct {
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Profile struct {
	Name                          string         `json:"name,omitempty"`
	GameDir                       string         `json:"gameDir,omitempty"`
	LastVersionID                 string         `json:"lastVersionId,omitempty"`
	JavaDir                       string         `json:"javaDir,omitempty"`
	JavaArgs                      string         `json:"javaArgs,omitempty"`
	Resolution                    map[string]int `json:"resolution,omitempty"`
	AllowedReleaseTypes           []string       `json:"allowedReleaseTypes,omitempty"`
	LauncherVisibilityOnGameClose string         `json:"launcherVisibilityOnGameClose,omitempty"`

	//Resolution                    Resolution `json:"resolution,omitempty"`
	//Resolution map[string]interface{} `json:"resolution,omitempty"`
}

type LauncherVersion struct {
	Name   string `json:"name,omitempty"`
	Format int    `json:"format,omitempty"`
}

type ProfileConfig struct {
	//Profiles               []Profile              `json:"profiles"`
	Profiles               map[string]Profile     `json:"profiles,omitempty"`
	SelectedProfile        string                 `json:"selectedProfile,omitempty"`
	ClientToken            string                 `json:"clientToken,omitempty"`
	AuthenticationDatabase map[string]interface{} `json:"authenticationDatabase,omitempty"` // ???
	SelectedUser           string                 `json:"selectedUser,omitempty"`
	LauncherVersion        LauncherVersion        `json:"launcherVersion,omitempty"`
}

func (pc *ProfileConfig) AddProfile(profile Profile) {
	pc.Profiles[profile.Name] = profile
}

func (pc *ProfileConfig) RemoveProfile(profileName string) {
	var tmpItems map[string]Profile
	for k, v := range pc.Profiles {
		if v.Name == profileName {

		} else {
			tmpItems[k] = v
		}
	}
	pc.Profiles = tmpItems
}

func (pc *ProfileConfig) LoadFromFile(filePath string) {
	if bp.FileExists(filePath) {
		dat, err := ioutil.ReadFile(filePath)
		bp.FailError(err)

		json.Unmarshal(dat, &pc)

		return

	}
	fmt.Printf("<!> INFO File >%s< not found\n", filePath)

}

func (pc *ProfileConfig) SaveToFile(filePath string) {
	//if bp.FileExists(filePath) {
	outFile, err := os.Create(filePath)
	bp.FailError(err)
	//tmpJSON, _ := json.Marshal(pc)
	tmpJSON, _ := json.MarshalIndent(pc, "", "  ")
	outFile.Write(tmpJSON)
	outFile.Close()
	//}
}
