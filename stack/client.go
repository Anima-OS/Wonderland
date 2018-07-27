// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package stack

import (
	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil/xwindow"
)

type Client interface {
	Id() xproto.Window
	Win() *xwindow.Window
	TopWin() *xwindow.Window
	Layer() int
	Transient(client Client) bool
}

func clientIndex(needle Client, haystack []Client) int {
	for i, client := range haystack {
		if client.Id() == needle.Id() {
			return i
		}
	}
	return -1
}
