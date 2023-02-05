package main

import (
	"fmt"
	"os"
	"runtime"
	"sfml-test/graphics"
	"sync"
	"time"

	sfmlgraphics "github.com/telroshan/go-sfml/v2/graphics"
	sfmlwindow "github.com/telroshan/go-sfml/v2/window"
)

func init() { runtime.LockOSThread() }

func render(wg *sync.WaitGroup, w sfmlgraphics.Struct_SS_sfRenderWindow, sprmgr graphics.SpriteManager) {
	defer wg.Done()
	sfmlgraphics.SfRenderWindow_setActive(w, 1)
	lastTick := time.Now()
	for sfmlwindow.SfWindow_isOpen(w) > 0 {
		// Animate
		now := time.Now()
		elapsed := now.Sub(lastTick)
		lastTick = now
		sprmgr.Animate(elapsed)

		// Draw
		sfmlgraphics.SfRenderWindow_clear(w, sfmlgraphics.GetSfRed())
		sprmgr.Draw(w)
		sfmlgraphics.SfRenderWindow_display(w)
	}
	sfmlgraphics.SfRenderWindow_setActive(w, 0)
}

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

	if err := sprmgr.LoadBasicSprite("explosionspr", "explosion", graphics.Rect{X: 0, Y: 0, W: 48, H: 48}); err != nil {
		fmt.Printf("[FATAL] Error loading sprite %s: %s", "explosionspr", err)
		os.Exit(1)
	}

	frames := []graphics.AnimatedSpriteFrame{
		{graphics.Rect{48 * 0, 0, 48, 48}, time.Millisecond * 100},
		{graphics.Rect{48 * 1, 0, 48, 48}, time.Millisecond * 100},
		{graphics.Rect{48 * 2, 0, 48, 48}, time.Millisecond * 100},
		{graphics.Rect{48 * 3, 0, 48, 48}, time.Millisecond * 100},
		{graphics.Rect{48 * 4, 0, 48, 48}, time.Millisecond * 100},
		{graphics.Rect{48 * 5, 0, 48, 48}, time.Millisecond * 100},
		{graphics.Rect{48 * 6, 0, 48, 48}, time.Millisecond * 100},
		{graphics.Rect{48 * 7, 0, 48, 48}, time.Millisecond * 100},
	}
	if err := sprmgr.LoadAnimatedSprite("explosionspr", "explosion", graphics.PlaybackMode_LOOP, frames); err != nil {
		fmt.Printf("[FATAL] Error loading sprite %s: %s", "explosionspr", err)
		os.Exit(1)
	}

	/*
		explosion_spr, err := sprmgr.GetSprite("explosionspr")
		if err != nil {
			fmt.Printf("[FATAL] Error getting sprite explosionspr: %s", err)
			os.Exit(1)
		}
	*/

	// Start render thread
	sfmlgraphics.SfRenderWindow_setActive(w, 0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go render(&wg, w, sprmgr)

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
				sfmlwindow.SfWindow_close(w) // Note: window should be activated again before calling close, so the render thread should have already returned. Todo, use a chan to close the thread here
			}

			if ev.GetEvType() == sfmlwindow.SfEventType(sfmlwindow.SfEvtKeyPressed) {
				switch ev.GetKey().GetCode() {
				case sfmlwindow.SfKeyCode(sfmlwindow.SfKeyEscape):
					sfmlwindow.SfWindow_close(w) // Note: window should be activated again before calling close, so the render thread should have already returned. Todo, use a chan to close the thread here
				}
			}
		}
	}

	fmt.Println("Waiting for render thread to finish...")
	wg.Wait()

}
