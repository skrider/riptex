/*
This is the root level command. Its function is to bring up a screen.
*/
package commands

import (
	"fmt"
	"github.com/rthornton128/goncurses"
)

func Default() {
	_, err := goncurses.Init()
	if err != nil {
		fmt.Println("init:", err)
	}
	defer goncurses.End()
}
