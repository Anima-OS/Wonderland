// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package layout

import (
	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil/xrect"
)

// Client is the method set required for a particular client to be used by
// any layout.
//
// Note that since layout clients come from workspace clients, this method
// set *must* be a subset of workspace.Client. (That is, if a method is added
// here, it must also be added to workspace.Client if it hasn't been already.)
type Client interface {
	Id() xproto.Window
	String() string
	Layout() Layout
	Geom() xrect.Rect
	DragGeom() xrect.Rect
	ShouldForceFloating() bool
	Focus()
	Raise()
	IsActive() bool

	MROpt(validate bool, flags, x, y, width, height int)
	MoveResize(x, y, width, height int)
	MoveResizeValid(x, y, width, height int)
	Move(x, y int)
	Resize(validate bool, width, height int)

	FrameTile()

	HasState(name string) bool
	SaveState(name string)
	LoadState(name string)
	DeleteState(name string)
}
