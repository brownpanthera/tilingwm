package main

import (
	"fmt"

	"github.com/tilinwindow/win32"
)

func main() {
	handles := win32.GetWindowDetails()
	for _, i := range handles {
		title, _ := win32.GetWindowTextW(i)
		class, _ := win32.GetClassName(i)

		fmt.Printf("Title: %s | Class: %s\n", title, class)
	}
}
