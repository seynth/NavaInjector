package main

import "C"
import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"

	"github.com/ropnop/go-clr"
)

func checkOK(hr uintptr, caller string) {
	if hr != 0x0 {
		fmt.Printf("%s returned 0x%08x\n", caller, hr)
	}
}

func logToFile(msg string) {
	userProfile := os.Getenv("USERPROFILE")
	if userProfile == "" {
		fmt.Println("[Nava] USERPROFILE not found, fallback")
		userProfile = `C:\Users\Public\Desktop`
	}
	desktopPath := filepath.Join(userProfile, "Desktop")
	logPath := filepath.Join(desktopPath, "log_nava.txt")
	f, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	f.WriteString(msg + "\n")
	f.Close()
}

func AllocTerm() {
	AllocConsole.Call()

	stdout, _ := syscall.Open("CONOUT$", syscall.O_RDWR, 0)
	if stdout != syscall.InvalidHandle {
		os.Stdout = os.NewFile(uintptr(stdout), "/dev/stdout")
		os.Stderr = os.NewFile(uintptr(stdout), "/dev/stderr")
	}
	fmt.Println("[Nava] Nava loaded successfully")

}

//export Nava
func Nava() {
	// AllocTerm()
	metaHost, err := clr.GetICLRMetaHost()
	if err != nil {
		fmt.Println("failed create clr metahost")
		return
	}
	defer metaHost.Release()

	// [System.Reflection.Assembly]::LoadFrom("C:\Program Files\SafeExamBrowser\Application\SafeExamBrowser.Client.exe").ImageRuntimeVersion
	versionString := "v4.0.30319"
	var pRuntimeInfo uintptr

	pwzVersion, _ := syscall.UTF16PtrFromString(versionString)
	hr := metaHost.GetRuntime(pwzVersion, &clr.IID_ICLRRuntimeInfo, &pRuntimeInfo)
	checkOK(hr, "metaHost.GetRuntime")

	runtimeInfo := clr.NewICLRRuntimeInfoFromPtr(pRuntimeInfo)
	if runtimeInfo == nil {
		fmt.Println("runtimeInfo is nil")
		return
	}

	var pRuntimeHost uintptr
	hr = runtimeInfo.GetInterface(&clr.CLSID_CLRRuntimeHost, &clr.IID_ICLRRuntimeHost, &pRuntimeHost)

	checkOK(hr, "runtimeInfo.GetInterface")
	host := clr.NewICLRRuntimeHostFromPtr(pRuntimeHost)

	tempPath := os.TempDir()
	milimDest := filepath.Join(tempPath, "Milim.dll")

	pDLLPath, _ := syscall.UTF16PtrFromString(milimDest)
	pTypeName, _ := syscall.UTF16PtrFromString("Milim.InitMilim")
	pMethodName, _ := syscall.UTF16PtrFromString("Nava")
	pArgument, _ := syscall.UTF16PtrFromString("unused")
	var pReturnVal *uint16
	hr = host.ExecuteInDefaultAppDomain(
		pDLLPath,
		pTypeName,
		pMethodName,
		pArgument,
		pReturnVal)
	if hr != 0 {
		fmt.Printf("ExecuteInAppDomain failed (code: 0x%x)\n", hr)
		return
	}
	fmt.Println("[Nava] Loaded Nava.dll")

	if xe() == "safeexambrowser.client.exe" {
		fmt.Println("begin hiding banner")
		if e := hideBanner(); e != nil {
			fmt.Println("error hiding banner: ", e.Error())
		}
	}
}

func main() {}
