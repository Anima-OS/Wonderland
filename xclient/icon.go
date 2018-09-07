// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package xclient

import (
	"github.com/BurntSushi/xgbutil/xgraphics"

	"github.com/Anima-OS/Wonderland/logger"
	"github.com/Anima-OS/Wonderland/wm"
)

func (c *Client) Icon(width, height int) *xgraphics.Image {
	ximg, err := xgraphics.FindIcon(wm.X, c.Id(), width, height)
	if err != nil {
		logger.Message.Printf("Could not find icon for '%s': %s", c, err)
		ximg = xgraphics.NewConvert(wm.X, wm.Theme.DefaultIcon)
		ximg = ximg.Scale(width, height)
	}

	return ximg
}
