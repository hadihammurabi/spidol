package spidol

import (
	"strings"
)

const (
	colorFormat string = "\033[00;000;0000m"
	nobold      string = "[00;"
	bold        string = "[1;"

	nobg             string = "000;"
	bgColorBlack     string = "40;"
	bgColorRed       string = "41;"
	bgColorGreen     string = "42;"
	bgColorYellow    string = "43;"
	bgColorBlue      string = "44;"
	bgColorPurple    string = "45;"
	bgColorDarkGreen string = "46;"
	bgColorWhite     string = "47;"

	nofg          string = "0000m"
	fgColorBlack  string = "30m"
	fgColorRed    string = "31m"
	fgColorGreen  string = "32m"
	fgColorYellow string = "33m"
	fgColorBlue   string = "34m"
	fgColorPurple string = "35m"
	fgColorCyan   string = "36m"
	fgColorWhite  string = "37m"
)

type Spidol string

func New(text string) Spidol {
	return Spidol(colorFormat + text)
}

func (s Spidol) Reset() Spidol {
	return Spidol(string(s) + colorFormat)
}

func (s Spidol) Red() Spidol {
	result := strings.ReplaceAll(string(s), nofg, fgColorRed)
	return Spidol(result)
}

func (s Spidol) Green() Spidol {
	result := strings.ReplaceAll(string(s), nofg, fgColorGreen)
	return Spidol(result)
}

func (s Spidol) Yellow() Spidol {
	result := strings.ReplaceAll(string(s), nofg, fgColorYellow)
	return Spidol(result)
}

func (s Spidol) Blue() Spidol {
	result := strings.ReplaceAll(string(s), nofg, fgColorBlue)
	return Spidol(result)
}

func (s Spidol) Purple() Spidol {
	result := strings.ReplaceAll(string(s), nofg, fgColorPurple)
	return Spidol(result)
}

func (s Spidol) White() Spidol {
	result := strings.ReplaceAll(string(s), nofg, fgColorWhite)
	return Spidol(result)
}

func (s Spidol) BgRed() Spidol {
	result := strings.ReplaceAll(string(s), nobg, bgColorRed)
	return Spidol(result)
}

func (s Spidol) BgBlack() Spidol {
	result := strings.ReplaceAll(string(s), nobg, bgColorBlack)
	return Spidol(result)
}

func (s Spidol) BgGreen() Spidol {
	result := strings.ReplaceAll(string(s), nobg, bgColorGreen)
	return Spidol(result)
}

func (s Spidol) BgYellow() Spidol {
	result := strings.ReplaceAll(string(s), nobg, bgColorYellow)
	return Spidol(result)
}

func (s Spidol) BgBlue() Spidol {
	result := strings.ReplaceAll(string(s), nobg, bgColorBlue)
	return Spidol(result)
}

func (s Spidol) BgPurple() Spidol {
	result := strings.ReplaceAll(string(s), nobg, bgColorPurple)
	return Spidol(result)
}

func (s Spidol) BgWhite() Spidol {
	result := strings.ReplaceAll(string(s), nobg, bgColorWhite)
	return Spidol(result)
}

func (s Spidol) Bold() Spidol {
	result := strings.ReplaceAll(string(s), nobold, bold)
	return Spidol(result)
}
