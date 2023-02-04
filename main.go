package main

import (
	"fmt"
	"os"
	"runtime"
	"sfml-test/graphics"

	sfmlgraphics "github.com/telroshan/go-sfml/v2/graphics"
	sfmlwindow "github.com/telroshan/go-sfml/v2/window"
)

func init() { runtime.LockOSThread() }

func main() {
	vm := sfmlwindow.NewSfVideoMode()
	defer sfmlwindow.DeleteSfVideoMode(vm)

	vm.SetWidth(800)
	vm.SetHeight(600)
	vm.SetBitsPerPixel(32)

	/* Create the main window */
	cs := sfmlwindow.NewSfContextSettings()
	defer sfmlwindow.DeleteSfContextSettings(cs)
	w := sfmlgraphics.SfRenderWindow_create(vm, "SFML window", uint(sfmlwindow.SfResize|sfmlwindow.SfClose), cs)
	defer sfmlwindow.SfWindow_destroy(w)

	// Load textures
	txtrmgr := graphics.NewTextureManager()
	defer txtrmgr.Cleanup()

	if err := txtrmgr.LoadTexture("explosion", "resources/images/explosion_48.png"); err != nil {
		fmt.Printf("[FATAL] Error loading texture %s: %s", "explosion", err)
		os.Exit(1)
	}

	// Load sprites
	sprmgr := graphics.NexSpriteManager(txtrmgr)
	defer sprmgr.Cleanup()

	if err := sprmgr.LoadSprite("explosionspr", "explosion", graphics.Rect{0, 0, 48, 48}); err != nil {
		fmt.Printf("[FATAL] Error loading sprite %s: %s", "explosionspr", err)
		os.Exit(1)
	}

	explosion_spr, err := sprmgr.GetSprite("explosionspr")
	if err != nil {
		fmt.Printf("[FATAL] Error getting sprite explosionspr: %s", err)
		os.Exit(1)
	}

	// Handle events
	ev := sfmlwindow.NewSfEvent()
	defer sfmlwindow.DeleteSfEvent(ev)

	fmt.Println("Hello SFML !")

	/* Start the game loop */
	for sfmlwindow.SfWindow_isOpen(w) > 0 {
		/* Process events */
		for sfmlwindow.SfWindow_pollEvent(w, ev) > 0 {
			/* Close window: exit */
			if ev.GetEvType() == sfmlwindow.SfEventType(sfmlwindow.SfEvtClosed) {
				return
			}

			if ev.GetEvType() == sfmlwindow.SfEventType(sfmlwindow.SfEvtKeyPressed) {
				switch ev.GetKey().GetCode() {
				case sfmlwindow.SfKeyCode(sfmlwindow.SfKeyEscape):
					return
				}
			}
		}
		sfmlgraphics.SfRenderWindow_clear(w, sfmlgraphics.GetSfRed())

		sfmlgraphics.SfRenderWindow_drawSprite(w, explosion_spr.GetSfSprite(), (sfmlgraphics.SfRenderStates)(sfmlgraphics.SwigcptrSfRenderStates(0)))

		sfmlgraphics.SfRenderWindow_display(w)
	}
}
