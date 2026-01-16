package win

import (
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

func WindowDetails() {
	xs := windows.GetForegroundWindow()
	fmt.Println(xs)

	var result []uintptr

	cb := windows.NewCallback(EnumWindowsProc)
	windows.EnumWindows(cb, unsafe.Pointer(&result))

	for _, v := range result {
		fmt.Println(v)
	}
}

func EnumWindowsProc(hwnd uintptr, lparam uintptr) uintptr {
	resPtr := (*[]uintptr)(unsafe.Pointer(lparam))
	*resPtr = append(*resPtr, hwnd)

	return 1
}
