package main

import (
	"fmt"
	"runtime"

	"github.com/telroshan/go-sfml/v2/graphics"
	"github.com/telroshan/go-sfml/v2/window"
)

func init() { runtime.LockOSThread() }

func main() {
	vm := window.NewSfVideoMode()
	defer window.DeleteSfVideoMode(vm)

	vm.SetWidth(800)
	vm.SetHeight(600)
	vm.SetBitsPerPixel(32)

	/* Create the main window */
	cs := window.NewSfContextSettings()
	defer window.DeleteSfContextSettings(cs)
	w := graphics.SfRenderWindow_create(vm, "SFML window", uint(window.SfResize|window.SfClose), cs)
	defer window.SfWindow_destroy(w)

	ev := window.NewSfEvent()
	defer window.DeleteSfEvent(ev)

	fmt.Println("Hello SFML !")

	/* Start the game loop */
	for window.SfWindow_isOpen(w) > 0 {
		/* Process events */
		for window.SfWindow_pollEvent(w, ev) > 0 {
			/* Close window: exit */
			if ev.GetEvType() == window.SfEventType(window.SfEvtClosed) {
				return
			}

			if ev.GetEvType() == window.SfEventType(window.SfEvtKeyPressed) {
				switch ev.GetKey().GetCode() {
				case window.SfKeyCode(window.SfKeyEscape):
					return
				}
			}
		}
		graphics.SfRenderWindow_clear(w, graphics.GetSfRed())
		graphics.SfRenderWindow_display(w)
	}
}
