// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package cursors

import (
	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xcursor"

	"github.com/Anima-OS/Wonderland/logger"
)

var (
	LeftPtr           xproto.Cursor
	Fleur             xproto.Cursor
	Watch             xproto.Cursor
	TopSide           xproto.Cursor
	TopRightCorner    xproto.Cursor
	RightSide         xproto.Cursor
	BottomRightCorner xproto.Cursor
	BottomSide        xproto.Cursor
	BottomLeftCorner  xproto.Cursor
	LeftSide          xproto.Cursor
	TopLeftCorner     xproto.Cursor
)

func Initialize(X *xgbutil.XUtil) {
	// lazy...
	cc := func(cursor uint16) xproto.Cursor {
		cid, err := xcursor.CreateCursor(X, cursor)
		if err != nil {
			logger.Warning.Printf("Could not load cursor '%d'.", cursor)
			return 0
		}
		return cid
	}

	LeftPtr = cc(xcursor.LeftPtr)
	Fleur = cc(xcursor.Fleur)
	Watch = cc(xcursor.Watch)
	TopSide = cc(xcursor.TopSide)
	TopRightCorner = cc(xcursor.TopRightCorner)
	RightSide = cc(xcursor.RightSide)
	BottomRightCorner = cc(xcursor.BottomRightCorner)
	BottomSide = cc(xcursor.BottomSide)
	BottomLeftCorner = cc(xcursor.BottomLeftCorner)
	LeftSide = cc(xcursor.LeftSide)
	TopLeftCorner = cc(xcursor.TopLeftCorner)
}
