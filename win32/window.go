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
	procSetWindowPos   = user32.NewProc("SetWindowPos")
	procShowWindow     = user32.NewProc("ShowWindow")
	procGetWindowLongW = user32.NewProc("GetWindowLongW")
)

const (
	SWP_NOZORDER   = 0x0004
	SWP_SHOWWINDOW = 0x0040
	SW_RESTORE     = 9

	GWL_EXSTYLE      = -20
	WS_EX_APPWINDOW  = 0x00040000
	WS_EX_TOOLWINDOW = 0x00000080
)

func GetWindowLongW(hwnd uintptr) uintptr {
	idx := int32(GWL_EXSTYLE)
	r1, _, _ := procGetWindowLongW.Call(hwnd, uintptr(idx))
	return r1
}

func IsRealWind(hwnd uintptr) bool {
	exstyle := GetWindowLongW(hwnd)
	if exstyle&WS_EX_APPWINDOW != 0 {
		return true
	}

	if exstyle&WS_EX_TOOLWINDOW != 0 {
		return false
	}
	return true

}

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

func VisibleWindow(hwnd uintptr) bool {
	return windows.IsWindowVisible(windows.HWND(hwnd))
}

// func SetWindowPos(hwnd uintptr, x, y, width, height int) error {
// 	r1, _, err := procSetWindowPos.Call(hwnd, 0, uintptr(x), uintptr(y), uintptr(width), uintptr(height), SWP_NOZORDER|SWP_SHOWWINDOW)
// 	if r1 == 0 {
// 		return err
// 	}
// 	return nil
// }

// un-maximizing the handle so setWindows can actually resize it basicallyt restoring window from maximzed state before resizing
func ShowWindow(hwnd uintptr, nCmdshow int) {
	procShowWindow.Call(hwnd, SW_RESTORE)
}
