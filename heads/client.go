// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package heads

import (
	"github.com/BurntSushi/xgb/xproto"
)

type Clients interface {
	Get(i int) Client
	Len() int
}

type Client interface {
	Id() xproto.Window
	IsMaximized() bool
	Remaximize()
}
