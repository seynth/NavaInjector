package main

import (
	"fmt"
	"strings"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func xe() string {
	var buf [windows.MAX_PATH]uint16
	r, _, _ := GetModuleFileNameW.Call(
		0,
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)
	if r == 0 {
		return ""
	}
	path := windows.UTF16ToString(buf[:r])
	if i := strings.LastIndex(path, "\\"); i != -1 {
		return strings.ToLower(path[i+1:])
	}
	return strings.ToLower(path)
}

func hideBanner() error {
	capt := "Version"
	partialLower := strings.ToLower(capt)
	var found bool

	cb := syscall.NewCallback(func(hwnd windows.Handle, _ uintptr) uintptr {
		if found {
			return 1
		}

		var buf [256]uint16
		n, _, _ := GetWindowTextW.Call(
			uintptr(hwnd),
			uintptr(unsafe.Pointer(&buf[0])),
			uintptr(len(buf)),
		)
		if n > 0 {
			title := windows.UTF16ToString(buf[:n])
			if strings.Contains(strings.ToLower(title), partialLower) {
				// logToFile(fmt.Sprintf("found window: %s", title))

				visible, _, _ := IsWindowVisible.Call(uintptr(hwnd))
				if visible != 0 {
					ShowWindow.Call(uintptr(hwnd), 0)
					// logToFile("success hide window")
				}
				found = true
				return 0
			}
		}
		return 1
	})

	_, _, err := EnumWindows.Call(cb, 0)
	if err != nil && err.Error() != "The operation completed successfully." {
		return err
	}

	if !found {
		return fmt.Errorf("[Nava::Error] no window name: '%s'", capt)
	}
	return nil
}
