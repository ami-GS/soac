package soac

import (
	"fmt"
)

type Color byte
type Attribute byte
type BackGround byte

const (
	Black Color = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenda
	Cyan
	White
)

const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	Blink1
	Blink2
	Reverse
	Concealed
	Crossedout
)

const (
	BgBlack BackGround = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenda
	BgCyan
	BgWhite
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
	self.fg = Black
	return self
}

func (self *Changer) Red() *Changer {
	self.fg = Red
	return self
}

func (self *Changer) Green() *Changer {
	self.fg = Green
	return self
}

func (self *Changer) Yellow() *Changer {
	self.fg = Yellow
	return self
}

func (self *Changer) Blue() *Changer {
	self.fg = Blue
	return self
}

func (self *Changer) Magenda() *Changer {
	self.fg = Magenda
	return self
}

func (self *Changer) Cyan() *Changer {
	self.fg = Cyan
	return self
}

func (self *Changer) White() *Changer {
	self.fg = White
	return self
}

func (self *Changer) Bold() *Changer {
	self.attr = Bold
	return self
}

func (self *Changer) Faint() *Changer {
	self.attr = Faint
	return self
}

func (self *Changer) Italic() *Changer {
	self.attr = Italic
	return self
}

func (self *Changer) Underline() *Changer {
	self.attr = Underline
	return self
}

func (self *Changer) Blink1() *Changer {
	self.attr = Blink1
	return self
}

func (self *Changer) Blink2() *Changer {
	self.attr = Blink2
	return self
}

func (self *Changer) Reverse() *Changer {
	self.attr = Reverse
	return self
}

func (self *Changer) Concealed() *Changer {
	self.attr = Concealed
	return self
}

func (self *Changer) CrossedOut() *Changer {
	self.attr = Crossedout
	return self
}

func (self *Changer) Bblack() *Changer {
	self.bg = BgBlack
	return self
}

func (self *Changer) Bred() *Changer {
	self.bg = BgRed
	return self
}

func (self *Changer) Bgreen() *Changer {
	self.bg = BgGreen
	return self
}

func (self *Changer) Byellow() *Changer {
	self.bg = BgYellow
	return self
}

func (self *Changer) Bblue() *Changer {
	self.bg = BgBlue
	return self
}

func (self *Changer) Bmagenda() *Changer {
	self.bg = BgMagenda
	return self
}

func (self *Changer) Bcyan() *Changer {
	self.bg = BgCyan
	return self
}

func (self *Changer) Bwhite() *Changer {
	self.bg = BgWhite
	return self
}
