package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("[Nava] Welcome . . .")
	seb := "safeexambrowser.exe"
	sebClient := "safeexambrowser.client"

	fmt.Println("[Nava] Running setup")
	navaDest := setup()
	enablePriv()

	fmt.Println("[Nava] Starting . . .")
	sebPid := findProcessSync(seb)

	if sebPid != 0 {
		fmt.Println("[Nava] Found parent pid")

		time.Sleep(100 * time.Millisecond)
		if e := inject(sebPid, navaDest); e != nil {
			fmt.Println("[Nava::Parent] Failed to inject nava in parent process", e.Error())
		}

		if e := callExportedFunction(sebPid, navaDest, "Nava"); e != nil {
			fmt.Println("[Nava::Parent] Failed to call Nava", e.Error())
		}

		fmt.Println("[Nava] Call nava for client")
		clientPid := findProcessSync(sebClient)
		time.Sleep(200 * time.Millisecond)

		if e := inject(clientPid, navaDest); e != nil {
			fmt.Println("[Nava::Client] Failed to inject Nava in client", e.Error())
		}
		if e := callExportedFunction(clientPid, navaDest, "Nava"); e != nil {
			fmt.Println("[Nava::Client] Failed to call Nava function", e.Error())
		}

	}

}
