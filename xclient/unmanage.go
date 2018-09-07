// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package xclient

import (
	"github.com/BurntSushi/xgb/xproto"

	"github.com/BurntSushi/xgbutil/icccm"
	"github.com/BurntSushi/xgbutil/xevent"

	"github.com/Anima-OS/Wonderland/event"
	"github.com/Anima-OS/Wonderland/focus"
	"github.com/Anima-OS/Wonderland/logger"
	"github.com/Anima-OS/Wonderland/stack"
	"github.com/Anima-OS/Wonderland/wm"
)

func (c *Client) unmanage() {
	wm.X.Grab()
	defer wm.X.Ungrab()

	go func() {
		c.frames.destroy()
		c.prompts.destroy()
	}()

	logger.Message.Printf("Unmanaging client: %s", c)

	infoWorkspace := c.workspace.String()
	infoClass := c.Class().Class
	infoInstance := c.Class().Instance
	infoName := c.Name()

	c.frame.Unmap()
	c.win.Detach()
	icccm.WmStateSet(wm.X, c.Id(), &icccm.WmState{State: icccm.StateWithdrawn})
	focus.Remove(c)
	wm.FocusFallback()
	stack.Remove(c)
	c.workspace.Remove(c)
	wm.RemoveClient(c)
	c.attnStop()
	xproto.ChangeSaveSetChecked(
		wm.X.Conn(), xproto.SetModeDelete, c.Id()).Check()

	if c.hadStruts {
		wm.Heads.ApplyStruts(wm.Clients)
	}

	event.Notify(event.UnmanagedClient{
		Id:        c.Id(),
		Name:      infoName,
		Workspace: infoWorkspace,
		Class:     infoClass,
		Instance:  infoInstance,
	})
}

func (c *Client) ImminentDestruction() bool {
	toIgnore := c.unmapIgnore
	for _, evOrErr := range xevent.Peek(wm.X) {
		ev := evOrErr.Event
		if ev == nil {
			continue
		}

		evUnmap, ok := ev.(xproto.UnmapNotifyEvent)
		if !ok {
			continue
		}

		if evUnmap.Window == c.Id() {
			if toIgnore <= 0 {
				return true
			}
			toIgnore--
		}
	}
	return false
}
