package main

import (
	"fmt"
	"os"
	"runtime"
	"sfml-test/engine"
	"sfml-test/game"
	"sfml-test/graphics"
	"sync"
	"time"

	. "sfml-test/common"

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
		tickStart := time.Now()
		elapsed := tickStart.Sub(lastTick)
		lastTick = tickStart
		sprmgr.Animate(elapsed)

		// Draw
		sfmlgraphics.SfRenderWindow_clear(w, sfmlgraphics.GetSfRed())
		sprmgr.Draw(w, Vec2{})
		sfmlgraphics.SfRenderWindow_display(w)

		// Wait for the next render
		frameDuration := time.Now().Sub(tickStart)
		if frameDuration.Microseconds() > 16667 {
			time.Sleep(time.Duration((16667 - frameDuration.Microseconds())) * time.Microsecond)
		}
	}
	sfmlgraphics.SfRenderWindow_setActive(w, 0)
}

func main() {
	vm := sfmlwindow.NewSfVideoMode()
	defer sfmlwindow.DeleteSfVideoMode(vm)

	vm.SetWidth(800)
	vm.SetHeight(600)
	vm.SetBitsPerPixel(32)

	//* Create the main window
	cs := sfmlwindow.NewSfContextSettings()
	defer sfmlwindow.DeleteSfContextSettings(cs)
	w := sfmlgraphics.SfRenderWindow_create(vm, "SFML window", uint(sfmlwindow.SfResize|sfmlwindow.SfClose), cs)
	defer sfmlwindow.SfWindow_destroy(w)
	sfmlwindow.SfWindow_setVerticalSyncEnabled(w, 1)

	//* Load textures
	txtrmgr := graphics.NewTextureManager()
	defer txtrmgr.Cleanup()

	texturesToLoad := map[string]string{
		"explosion": "resources/images/explosion_48.png",
		"lpc":       "resources/images/BODY_male.png",
	}
	for i := range texturesToLoad {
		if err := txtrmgr.LoadTexture(i, texturesToLoad[i]); err != nil {
			fmt.Printf("[FATAL] Error loading texture %s: %s\n", i, err)
			os.Exit(1)
		}
	}

	//* Load sprites
	sprmgr := graphics.NewSpriteManager(txtrmgr)
	defer sprmgr.Cleanup()

	/*if _, err := sprmgr.LoadBasicSprite("explosionspr_static", "explosion", Rect{X: 0, Y: 0, W: 48, H: 48}, Vec2{X: 50, Y: 20}); err != nil {
		fmt.Printf("[FATAL] Error loading sprite %s: %s", "explosionspr", err)
		os.Exit(1)
	}*/

	explosionIdleFrames := []graphics.AnimatedSpriteFrame{
		{Rect: Rect{X: 48 * 0, Y: 0, W: 48, H: 48}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 48 * 1, Y: 0, W: 48, H: 48}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 48 * 2, Y: 0, W: 48, H: 48}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 48 * 3, Y: 0, W: 48, H: 48}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 48 * 4, Y: 0, W: 48, H: 48}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 48 * 5, Y: 0, W: 48, H: 48}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 48 * 6, Y: 0, W: 48, H: 48}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 48 * 7, Y: 0, W: 48, H: 48}, Duration: time.Millisecond * 100},
	}
	if _, err := sprmgr.LoadAnimatedSprite("explosionspr", "explosion", PlaybackMode_LOOP, map[AnimationType]graphics.Animation{AnimationType_IDLES: explosionIdleFrames}, AnimationType_IDLES, Vec2{24, 40}, Vec2{X: 350, Y: 150}); err != nil {
		fmt.Printf("[FATAL] Error loading sprite %s: %s", "explosionspr", err)
		os.Exit(1)
	}

	lpcWalkNFrames := []graphics.AnimatedSpriteFrame{
		{Rect: Rect{X: 64 * 1, Y: 64 * 0, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 2, Y: 64 * 0, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 3, Y: 64 * 0, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 4, Y: 64 * 0, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 5, Y: 64 * 0, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 6, Y: 64 * 0, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 7, Y: 64 * 0, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 8, Y: 64 * 0, W: 64, H: 64}, Duration: time.Millisecond * 100},
	}
	lpcWalkWFrames := []graphics.AnimatedSpriteFrame{
		{Rect: Rect{X: 64 * 1, Y: 64 * 1, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 2, Y: 64 * 1, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 3, Y: 64 * 1, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 4, Y: 64 * 1, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 5, Y: 64 * 1, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 6, Y: 64 * 1, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 7, Y: 64 * 1, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 8, Y: 64 * 1, W: 64, H: 64}, Duration: time.Millisecond * 100},
	}
	lpcWalkSFrames := []graphics.AnimatedSpriteFrame{
		{Rect: Rect{X: 64 * 1, Y: 64 * 2, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 2, Y: 64 * 2, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 3, Y: 64 * 2, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 4, Y: 64 * 2, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 5, Y: 64 * 2, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 6, Y: 64 * 2, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 7, Y: 64 * 2, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 8, Y: 64 * 2, W: 64, H: 64}, Duration: time.Millisecond * 100},
	}
	lpcWalkEFrames := []graphics.AnimatedSpriteFrame{
		{Rect: Rect{X: 64 * 1, Y: 64 * 3, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 2, Y: 64 * 3, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 3, Y: 64 * 3, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 4, Y: 64 * 3, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 5, Y: 64 * 3, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 6, Y: 64 * 3, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 7, Y: 64 * 3, W: 64, H: 64}, Duration: time.Millisecond * 100},
		{Rect: Rect{X: 64 * 8, Y: 64 * 3, W: 64, H: 64}, Duration: time.Millisecond * 100},
	}
	lpcIdleNFrames := []graphics.AnimatedSpriteFrame{{Rect: Rect{X: 0, Y: 64 * 0, W: 64, H: 64}, Duration: time.Second}}
	lpcIdleWFrames := []graphics.AnimatedSpriteFrame{{Rect: Rect{X: 0, Y: 64 * 1, W: 64, H: 64}, Duration: time.Second}}
	lpcIdleSFrames := []graphics.AnimatedSpriteFrame{{Rect: Rect{X: 0, Y: 64 * 2, W: 64, H: 64}, Duration: time.Second}}
	lpcIdleEFrames := []graphics.AnimatedSpriteFrame{{Rect: Rect{X: 0, Y: 64 * 3, W: 64, H: 64}, Duration: time.Second}}

	lpcAnimations := map[AnimationType]graphics.Animation{
		AnimationType_IDLEN: lpcIdleNFrames,
		AnimationType_IDLEW: lpcIdleWFrames,
		AnimationType_IDLES: lpcIdleSFrames,
		AnimationType_IDLEE: lpcIdleEFrames,
		AnimationType_WALKN: lpcWalkNFrames,
		AnimationType_WALKW: lpcWalkWFrames,
		AnimationType_WALKS: lpcWalkSFrames,
		AnimationType_WALKE: lpcWalkEFrames,
	}

	playersprite, err := sprmgr.LoadAnimatedSprite("playerspr", "lpc", PlaybackMode_LOOP, lpcAnimations, AnimationType_IDLES, Vec2{32, 55}, Vec2{X: 250, Y: 150})
	if err != nil {
		fmt.Printf("[FATAL] Error loading player sprite %s", err)
		os.Exit(1)
	}

	//* Load actors
	gmgr := game.NewGameManager()
	player := gmgr.AddActor(Vec2{X: 250, Y: 150}, playersprite, playersprite)

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
	lastTick := time.Now()
	for sfmlwindow.SfWindow_isOpen(w) > 0 {
		tickStart := time.Now()
		elapsed := tickStart.Sub(lastTick)
		lastTick = tickStart

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

		// Get keyboard state
		playerMovement := engine.GetInput()
		playerMovement.X *= float32(elapsed.Milliseconds() / 10)
		playerMovement.Y *= float32(elapsed.Milliseconds() / 10)

		player.Move(playerMovement)
		time.Sleep(time.Millisecond * 10)
	}

	fmt.Println("Waiting for render thread to finish...")
	wg.Wait()

}
