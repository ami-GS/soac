package soac

import (
	"fmt"
)

type Color byte
type Attribute byte
type BackGround byte

const (
	black Color = iota + 30
	red
	green
	yellow
	blue
	magenda
	cyan
	white
)

const (
	reset Attribute = iota
	bold
	faint
	italic
	underline
	blink1
	blink2
	reverse
	concealed
	crossedout
)

const (
	bgBlack BackGround = iota + 40
	bgRed
	bgGreen
	bgYellow
	bgBlue
	bgMagenda
	bgCyan
	bgWhite
)

type Changer struct {
	fg   Color
	attr Attribute
	bg   BackGround
}

func NewChanger() *Changer {
	return &Changer{39, 0, 49}
}

func (self *Changer) Apply(val interface{}) string {
	return fmt.Sprintf("\x1b[%d;%d;%dm%v\x1b[0m", self.attr, self.fg, self.bg, val)
}

func (self *Changer) Reset() {
	self.fg = 39
	self.attr = 0
	self.bg = 49
}

func (self *Changer) Black() *Changer {
	self.fg = black
	return self
}

func (self *Changer) Red() *Changer {
	self.fg = red
	return self
}

func (self *Changer) Green() *Changer {
	self.fg = green
	return self
}

func (self *Changer) Yellow() *Changer {
	self.fg = yellow
	return self
}

func (self *Changer) Blue() *Changer {
	self.fg = blue
	return self
}

func (self *Changer) Magenda() *Changer {
	self.fg = magenda
	return self
}

func (self *Changer) Cyan() *Changer {
	self.fg = cyan
	return self
}

func (self *Changer) White() *Changer {
	self.fg = white
	return self
}

func (self *Changer) Bold() *Changer {
	self.attr = bold
	return self
}

func (self *Changer) Faint() *Changer {
	self.attr = faint
	return self
}

func (self *Changer) Italic() *Changer {
	self.attr = italic
	return self
}

func (self *Changer) Underline() *Changer {
	self.attr = underline
	return self
}

func (self *Changer) Blink1() *Changer {
	self.attr = blink1
	return self
}

func (self *Changer) Blink2() *Changer {
	self.attr = blink2
	return self
}

func (self *Changer) Reverse() *Changer {
	self.attr = reverse
	return self
}

func (self *Changer) Concealed() *Changer {
	self.attr = concealed
	return self
}

func (self *Changer) CrossedOut() *Changer {
	self.attr = crossedout
	return self
}

func (self *Changer) Bblack() *Changer {
	self.bg = bgBlack
	return self
}

func (self *Changer) Bred() *Changer {
	self.bg = bgRed
	return self
}

func (self *Changer) Bgreen() *Changer {
	self.bg = bgGreen
	return self
}

func (self *Changer) Byellow() *Changer {
	self.bg = bgYellow
	return self
}

func (self *Changer) Bblue() *Changer {
	self.bg = bgBlue
	return self
}

func (self *Changer) Bmagenda() *Changer {
	self.bg = bgMagenda
	return self
}

func (self *Changer) Bcyan() *Changer {
	self.bg = bgCyan
	return self
}

func (self *Changer) Bwhite() *Changer {
	self.bg = bgWhite
	return self
}
