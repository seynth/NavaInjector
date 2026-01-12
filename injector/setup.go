package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func setup() string {

	tempPath := os.TempDir()

	navaDll, e := os.ReadFile(`..\Nava\nava.dll`)
	if e != nil {
		fmt.Println("failed read nava dll")
	}

	milimDll, e := os.ReadFile(`..\Milim\bin\x64\Debug\Milim.dll`)
	if e != nil {
		fmt.Println("failed read milim dll")
	}

	navaDest := filepath.Join(tempPath, "nava.dll")
	milimDest := filepath.Join(tempPath, "Milim.dll")

	if e := os.WriteFile(navaDest, navaDll, 0644); e != nil {
		fmt.Println(e.Error())
	}

	if e := os.WriteFile(milimDest, milimDll, 0644); e != nil {
		fmt.Println(e.Error())
	}

	fmt.Println("[Nava] Setup done")
	return navaDest
}
