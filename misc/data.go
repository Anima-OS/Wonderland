// Copyright (c) 2018, Mark "Happy-Ferret" Bauermeister
//
// This software may be modified and distributed under the terms
// of the BSD license.  See the LICENSE file for details.

package misc

var (
	DejavusansTtf []byte
	WingoWav      []byte
	WingoPng      []byte
	ClosePng      []byte
	MinimizePng   []byte
	MaximizePng   []byte
)

func ReadData() {
	DejavusansTtf = DataFile("DejaVuSans.ttf")
	WingoWav = DataFile("wingo.wav")
	WingoPng = DataFile("wingo.png")
	ClosePng = DataFile("close.png")
	MinimizePng = DataFile("minimize.png")
	MaximizePng = DataFile("maximize.png")
}
