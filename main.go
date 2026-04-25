package main

import (
	"fmt"
	"strings"

	"github.com/tilinwindow/win32"
)

func main() {
	fmt.Println(GetFilteredWindows())
	// handles := win32.GetWindowDetails()
	// for _, i := range handles {
	// 	if win32.VisibleWindow(i) && win32.IsRealWind(i) {

	// 		// getWin := win32.GetWindowLongW(i)
	// 		// fin := win32.IsRealWind(i)
	// 		class, _ := win32.GetClassName(i)
	// 		title, _ := win32.GetWindowTextW(i)

	// 		if title == "" || class == "Windows.UI.Core.CoreWindow" || strings.HasPrefix(class, "HwndWrapper[Raycast") {
	// 			continue
	// 		}

	// 		fmt.Printf("Title: %s | Class: %s\n", title, class)

	// 		// fmt.Println(fin, "meeeeeeeeee")
	// 	}
	// 	// if class == "Chrome_WidgetWin_1" && strings.Contains(title, "Cursor") {
	// 	// 	fmt.Printf("found: hwnd=%v title=%s\n", i, title)
	// 	// 	win32.ShowWindow(i, win32.SW_RESTORE)
	// 	// 	err := win32.SetWindowPos(i, 0, 0, 800, 600)
	// 	// 	if err != nil {
	// 	// 		fmt.Printf("failed bro...%v\n", err)
	// 	// 	} else {
	// 	// 		fmt.Println("succeeded")
	// 	// 	}
	// 	// }
	// 	// // wh := win32.VisibleWindow(i)

	// 	// fmt.Printf("Title: %s | Class: %s\n", title, class)
	// 	// fmt.Println(wh)
	// }
}

func GetFilteredWindows() []uintptr {
	handles := win32.GetWindowDetails()
	filtered := []uintptr{}

	for _, i := range handles {
		if win32.VisibleWindow(i) && win32.IsRealWind(i) {
			class, _ := win32.GetClassName(i)
			title, _ := win32.GetWindowTextW(i)

			if title == "" || class == "Windows.UI.Core.CoreWindow" || strings.HasPrefix(class, "HwndWrapper[Raycast") {
				continue
			}

			filtered = append(filtered, i)
		}
	}

	return filtered // All "real" window handles
}
