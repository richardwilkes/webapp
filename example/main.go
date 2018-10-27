package main

import (
	"fmt"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
)

func main() {
	webapp.WillFinishStartupCallback = func() {
		wnd := webapp.NewWindow(webapp.StdWindowMask, geom.Rect{
			Point: geom.Point{X: 20, Y: 20},
			Size:  geom.Size{Width: 1024, Height: 768},
		}, "https://youtube.com")
		wnd.SetTitle("webapp example")
		bar := webapp.MenuBarForWindow(wnd)
		_, aboutItem, prefsItem := bar.InstallAppMenu()
		aboutItem.Handler = func() { fmt.Println("About menu item selected") }
		prefsItem.Handler = func() { fmt.Println("Preferences menu item selected") }
		bar.InstallEditMenu()
		bar.InstallWindowMenu()
		bar.InstallHelpMenu()
		wnd.ToFront()
	}
	webapp.Start()
}
