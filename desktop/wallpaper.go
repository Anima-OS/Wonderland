// Copyright (c) 2019, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package desktop

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"os"
	"strings"

	"github.com/BurntSushi/xgb/xproto"

	"github.com/Anima-OS/Wonderland/misc"
	"github.com/Anima-OS/Wonderland/wini"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xwindow"
)

type Wallpaper struct {
	Path string
}

func Initialize() {
	var wallpaper Wallpaper
	var img image.Image

	tdata, err := wini.Parse(misc.ConfigFile("wallpaper.wini"))
	if err != nil {
	}

	for _, section := range tdata.Sections() {
		switch section {
		case "global":
			for _, key := range tdata.Keys(section) {
				switch key.Name() {
				case "path":
					wallpaper.Path = strings.Join(key.Strings(), "")
				}
			}
		}
	}

	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	// A bit hackish. First we try to read the file given in wallpaper.wini.
	// If this fails, use default wallpaper. No scaling, thus far.
	reader, err := os.Open(wallpaper.Path)
	if err != nil {
		img, _, _ = image.Decode(bytes.NewBuffer(misc.Wallpaper))
	}
	if err == nil {
		img, _, err = image.Decode(reader)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer reader.Close()

	// Convert decoded image into an X image.
	ximg := xgraphics.NewConvert(X, img)

	// Now show it in a new window.
	win := showImage(ximg, "The Go Gopher!", true)

	// Listen for key press events.
	win.Listen(xproto.EventMaskKeyPress)
}

// This is a slightly modified version of xgraphics.XShowExtra that does
// not set any resize constraints on the window (so that it can go
// fullscreen).
func showImage(im *xgraphics.Image, name string, quit bool) *xwindow.Window {
	if len(name) == 0 {
		name = "xgbutil Image Window"
	}
	w, h := im.Rect.Dx(), im.Rect.Dy()

	win, err := xwindow.Generate(im.X)
	if err != nil {
		xgbutil.Logger.Printf("Could not generate new window id: %s", err)
		return nil
	}

	// Create a very simple window with dimensions equal to the image.
	win.Create(im.X.RootWin(), 0, 0, w, h, 0)

	if err := ewmh.WmWindowTypeSet(im.X, win.Id, []string{
		"_NET_WM_WINDOW_TYPE_DESKTOP"}); err != nil {
	}

	if err := ewmh.WmStateSet(im.X, win.Id, []string{
		"_NET_WM_STATE_STICKY"}); err != nil {
	}

	// Paint our image before mapping.
	im.XSurfaceSet(win.Id)
	im.XDraw()
	im.XPaint(win.Id)

	// Now we can map, since we've set all our properties.
	// (The initial map is when the window manager starts managing.)
	win.Map()

	return win
}
