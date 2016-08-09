package main

import (
	"os"
	"path"
	"strings"

	bp "github.com/nexustix/boilerplate"
)

func GetLocalForgeVersions(minecraftPath string) []string {

	var forgeVersions []string

	handle, err := os.Open(path.Join(minecraftPath, "versions"))
	bp.FailError(err)

	dir, err := handle.Stat()
	bp.FailError(err)

	if dir.IsDir() {
		//dirNames, err := dir.Readdirnames(dir)
		//bp.FailError(err)
		items, err := handle.Readdir(-1)
		bp.FailError(err)

		for _, v := range items {

			if strings.Contains(v.Name(), "forge") {
				//fmt.Printf("%v : %v\n", k, v.Name())
				forgeVersions = append(forgeVersions, v.Name())
			} else {
				//fmt.Printf("(%v : %v)\n", k, v.Name())
			}
		}
	}
	return forgeVersions

}
