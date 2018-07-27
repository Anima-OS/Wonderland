// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package wm

type CommandHacks struct {
	MouseResizeDirection     func(cmdStr string) (string, error)
	CycleClientRunWithKeyStr func(keyStr, cmdStr string) (func(), error)
}
