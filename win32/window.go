package win32

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var (
	user32             = windows.NewLazySystemDLL("user32.dll")
	procGetWindowTextW = user32.NewProc("GetWindowTextW")
	procGetClassNameW  = user32.NewProc("GetClassNameW")
)

func GetWindowDetails() []uintptr {
	allOpenHandles := []uintptr{}
	EnumWindowsProc := func(hwnd uintptr, lparam uintptr) uintptr {
		allOpenHandles = append(allOpenHandles, hwnd)
		return 1
	}
	callbackFun := windows.NewCallback(EnumWindowsProc)

	windows.EnumWindows(callbackFun, nil)

	return allOpenHandles
}

/*
load the dll
get the addr of the function in the dll
call the function at that addr
*/

func GetWindowTextW(hwnd uintptr) (string, error) {
	buf := make([]uint16, 256)
	maxNumberOfChar := 256
	r1, _, err := procGetWindowTextW.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), uintptr(maxNumberOfChar))
	if r1 == 0 {
		if err != nil && err != windows.ERROR_SUCCESS {
			return "", err
		}
		return "", nil
	}
	title := syscall.UTF16ToString(buf[:r1])
	return title, nil
}

func GetClassName(hwnd uintptr) (string, error) {
	buf := make([]uint16, 256)
	procGetClassNameW.Call(hwnd, uintptr(unsafe.Pointer(&buf[0])), 256)
	return syscall.UTF16ToString(buf), nil
}
