// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package event

import (
	"github.com/BurntSushi/xgb/xproto"
)

type Event interface{}

type Noop struct{}

type Restarting struct{}

type Subscribed struct{}

type (
	ChangedWorkspace        struct{}
	ChangedVisibleWorkspace struct{}
	ChangedWorkspaceNames   struct{}

	AddedWorkspace struct {
		Name string
	}

	RemovedWorkspace struct {
		Name string
	}
)

type (
	FocusedClient struct {
		Id xproto.Window
	}
	UnfocusedClient struct {
		Id xproto.Window
	}
	MappedClient struct {
		Id xproto.Window
	}
	UnmappedClient struct {
		Id xproto.Window
	}
	ManagedClient struct {
		Id xproto.Window
	}
	UnmanagedClient struct {
		Id xproto.Window
		Name string
		Workspace string
		Class string
		Instance string
	}
	ChangedClientName struct {
		Id xproto.Window
	}
	ChangedActiveClient struct {
		Id xproto.Window
	}
)

type ChangedLayout struct {
	Workspace string
}
