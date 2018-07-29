// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package main

import (
	"fmt"

	"github.com/BurntSushi/wingo/commands"
	"github.com/BurntSushi/wingo/wm"
)

func newHacks() wm.CommandHacks {
	return wm.CommandHacks{
		MouseResizeDirection:     mouseResizeDirection,
		CycleClientRunWithKeyStr: cycleClientRunWithKeyStr,
	}
}

func mouseResizeDirection(cmdStr string) (string, error) {
	cmd, err := commands.Env.Command(cmdStr)
	if err != nil {
		return "", err
	}
	return cmd.(*commands.MouseResize).Direction, nil
}

func cycleClientRunWithKeyStr(keyStr, cmdStr string) (func(), error) {
	var run func() = nil
	cmd, err := commands.Env.Command(cmdStr)
	if err != nil {
		return nil, err
	}

	switch t := cmd.(type) {
	case *commands.CycleClientNext:
		run = func() { t.RunWithKeyStr(keyStr) }
	case *commands.CycleClientPrev:
		run = func() { t.RunWithKeyStr(keyStr) }
	default:
		panic(fmt.Sprintf("bug: unknown type %T", t))
	}
	return run, nil
}
