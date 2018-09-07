// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package misc

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/BurntSushi/xdg"

	"github.com/Anima-OS/Wonderland/logger"
)

var ConfigPaths = xdg.Paths{
	Override:     "",
	XDGSuffix:    "Wonderland",
	GoImportPath: "github.com/Anima-OS/Wonderland/config",
}

var DataPaths = xdg.Paths{
	Override:     "",
	XDGSuffix:    "Wonderland",
	GoImportPath: "github.com/Anima-OS/Wonderland/data",
}

var ScriptPaths = xdg.Paths{
	Override:     "",
	XDGSuffix:    "Wonderland",
	GoImportPath: "github.com/Anima-OS/Wonderland/config",
}

func ConfigFile(name string) string {
	fpath, err := ConfigPaths.ConfigFile(name)
	if err != nil {
		logger.Error.Fatalln(err)
	}
	return fpath
}

func DataFile(name string) []byte {
	fpath, err := DataPaths.DataFile(name)
	if err != nil {
		logger.Error.Fatalln(err)
	}
	bs, err := ioutil.ReadFile(fpath)
	if err != nil {
		logger.Error.Fatalf("Could not read %s: %s", fpath, err)
	}
	return bs
}

func ScriptPath(name string) string {
	fpath, err := ScriptPaths.ConfigFile(path.Join("scripts", name, name))
	if err != nil {
		fpath, err = ScriptPaths.ConfigFile(path.Join("scripts", name))
		if err != nil {
			logger.Warning.Println(err)
			return ""
		}
	}
	return fpath
}

func ScriptConfigPath(name string) string {
	fname := fmt.Sprintf("%s.cfg", name)
	fp, err := ScriptPaths.ConfigFile(path.Join("scripts", name, fname))
	if err != nil {
		logger.Warning.Println(err)
		return ""
	}
	return fp
}
