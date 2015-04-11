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

type Param struct {
	fg   Color
	attr Attribute
	bg   BackGround
}

func NewParam() *Param {
	return &Param{39, 0, 49}
}

type Changer struct {
	val interface{}
	p   *Param
}

func NewChanger() *Changer {
	return &Changer{"", NewParam()}
}

func (self *Changer) Out() string {
	return fmt.Sprintf("\x1b[%d;%d;%dm%v\x1b[0m", self.p.attr, self.p.fg, self.p.bg, self.val)
}

func (self *Changer) Set(val interface{}) *Changer {
	self.val = val
	return self
}

func (self *Changer) Reset() {
	self.p = NewParam()
}

func (self *Changer) Black() *Changer {
	self.p.fg = black
	return self
}

func (self *Changer) Red() *Changer {
	self.p.fg = red
	return self
}

func (self *Changer) Green() *Changer {
	self.p.fg = green
	return self
}

func (self *Changer) Yellow() *Changer {
	self.p.fg = yellow
	return self
}

func (self *Changer) Blue() *Changer {
	self.p.fg = blue
	return self
}

func (self *Changer) Magenda() *Changer {
	self.p.fg = magenda
	return self
}

func (self *Changer) Cyan() *Changer {
	self.p.fg = cyan
	return self
}

func (self *Changer) White() *Changer {
	self.p.fg = white
	return self
}

func (self *Changer) Bold() *Changer {
	self.p.attr = bold
	return self
}

func (self *Changer) Faint() *Changer {
	self.p.attr = faint
	return self
}

func (self *Changer) Italic() *Changer {
	self.p.attr = italic
	return self
}

func (self *Changer) Underline() *Changer {
	self.p.attr = underline
	return self
}

func (self *Changer) Blink1() *Changer {
	self.p.attr = blink1
	return self
}

func (self *Changer) Blink2() *Changer {
	self.p.attr = blink2
	return self
}

func (self *Changer) Reverse() *Changer {
	self.p.attr = reverse
	return self
}

func (self *Changer) Concealed() *Changer {
	self.p.attr = concealed
	return self
}

func (self *Changer) CrossedOut() *Changer {
	self.p.attr = crossedout
	return self
}

func (self *Changer) Bblack() *Changer {
	self.p.bg = bgBlack
	return self
}

func (self *Changer) Bred() *Changer {
	self.p.bg = bgRed
	return self
}

func (self *Changer) Bgreen() *Changer {
	self.p.bg = bgGreen
	return self
}

func (self *Changer) Byellow() *Changer {
	self.p.bg = bgYellow
	return self
}

func (self *Changer) Bblue() *Changer {
	self.p.bg = bgBlue
	return self
}

func (self *Changer) Bmagenda() *Changer {
	self.p.bg = bgMagenda
	return self
}

func (self *Changer) Bcyan() *Changer {
	self.p.bg = bgCyan
	return self
}

func (self *Changer) Bwhite() *Changer {
	self.p.bg = bgWhite
	return self
}
