package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func callExportedFunction(pid uint32, dllPath, funcName string) error {
	hProcess, err := windows.OpenProcess(
		windows.PROCESS_CREATE_THREAD|
			windows.PROCESS_QUERY_INFORMATION|
			windows.PROCESS_VM_OPERATION|
			windows.PROCESS_VM_WRITE|
			windows.PROCESS_VM_READ,
		false,
		pid,
	)
	if err != nil {
		return err
	}
	defer windows.CloseHandle(hProcess)

	addr, err := getFunctionAddr(hProcess, dllPath, funcName)
	if err != nil {
		return err
	}

	hThread, _, err := CreateRemoteThread.Call(
		uintptr(hProcess),
		0, 0,
		addr,
		0,
		0, 0,
	)
	if hThread == 0 {
		return fmt.Errorf("[Nava::Error] CreateRemoteThread failed: %v", err)
	}
	defer windows.CloseHandle(windows.Handle(hThread))

	_, err = windows.WaitForSingleObject(windows.Handle(hThread), windows.INFINITE)
	if err != nil {
		return err
	}
	return nil
}

func getRemoteModuleBase(hProcess windows.Handle, moduleName string) uintptr {
	const maxModules = 1024
	var hMods [maxModules]windows.Handle
	var cbNeeded uint32

	bufSize := uint32(len(hMods)) * uint32(unsafe.Sizeof(windows.Handle(0)))
	if err := windows.EnumProcessModules(hProcess, &hMods[0], bufSize, &cbNeeded); err != nil {
		fmt.Printf("[Nava::Error] EnumProcessModules failed: %v\n", err)
		return 0
	}

	nMods := min(cbNeeded/uint32(unsafe.Sizeof(windows.Handle(0))), maxModules)

	for i := range nMods {
		var name [260]uint16
		if err := windows.GetModuleBaseName(hProcess, hMods[i], &name[0], uint32(len(name))); err != nil {
			continue
		}

		length := 0
		for length < len(name) && name[length] != 0 {
			length++
		}
		if length == 0 {
			continue
		}

		if strings.EqualFold(windows.UTF16ToString(name[:length]), moduleName) {
			return uintptr(hMods[i])
		}
	}
	return 0
}

func getFunctionAddr(hProcess windows.Handle, dllPath, funcName string) (uintptr, error) {
	moduleBase := getRemoteModuleBase(hProcess, filepath.Base(dllPath))
	if moduleBase == 0 {
		return 0, fmt.Errorf("[Nava::Error] Failed to find remote module base for %s", dllPath)

	}

	localDLL := syscall.NewLazyDLL(dllPath)
	proc := localDLL.NewProc(funcName)
	if proc == nil {
		return 0, fmt.Errorf("[Nava::Error] Function %s not found in local DLL", funcName)
	}

	localAddr := proc.Addr()
	if localAddr == 0 {
		return 0, fmt.Errorf("[Nava::Error] Local address of %s is zero", funcName)
	}

	dllPathUtf16, _ := syscall.UTF16PtrFromString(dllPath)
	localBase, _, _ := GetModuleHandle.Call(uintptr(unsafe.Pointer(dllPathUtf16)))
	if localBase == 0 {
		return 0, fmt.Errorf("[Nava::Error] failed to get local module base")
	}

	rva := localAddr - localBase

	remoteAddr := moduleBase + rva
	return remoteAddr, nil
}
