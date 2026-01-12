package main

import "golang.org/x/sys/windows"

var (
	kernel32 = windows.NewLazySystemDLL("kernel32.dll")
	user32   = windows.NewLazySystemDLL("user32.dll")

	EnumWindows        = user32.NewProc("EnumWindows")
	ShowWindow         = user32.NewProc("ShowWindow")
	IsWindowVisible    = user32.NewProc("IsWindowVisible")
	GetWindowTextW     = user32.NewProc("GetWindowTextW")
	AllocConsole       = kernel32.NewProc("AllocConsole")
	GetModuleFileNameW = kernel32.NewProc("GetModuleFileNameW")
)
