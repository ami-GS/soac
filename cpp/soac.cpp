#include "soac.h"

std::string Changer::Apply(std::string val) {
    tmp << "\x1b[" << attr << ";" << fg << ";" << bg << "m" << val << "\x1b[0m";
    std::string tmptmp = tmp.str();
    tmp.str("");
    return tmptmp;
}

std::string Changer::Apply(int val) {
    tmp << "\x1b[" << attr << ";" << fg << ";" << bg << "m" << val << "\x1b[0m";
    std::string tmptmp = tmp.str();
    tmp.str("");
    return tmptmp;
}


Changer* Changer::black() {
    fg = Black;
    return this;
}

Changer* Changer::red() {
    fg = Red;
    return this;
}


Changer* Changer::green() {
    fg = Green;
    return this;
}

Changer* Changer::yellow() {
    fg = Yellow;
    return this;
}

Changer* Changer::blue() {
    fg = Blue;
    return this;
}

Changer* Changer::magenda() {
    fg = Magenda;
    return this;
}

Changer* Changer::cyan() {
    fg = Cyan;
    return this;
}

Changer* Changer::white() {
    fg = White;
    return this;
}

Changer* Changer::reset() {
    attr = Reset;
    return this;
}

Changer* Changer::bold() {
    attr = Bold;
    return this;
}

Changer* Changer::faint() {
    attr = Faint;
    return this;
}

Changer* Changer::italic() {
    attr = Italic;
    return this;
}

Changer* Changer::underline() {
    attr = Underline;
    return this;
}

Changer* Changer::blink1() {
    attr = Blink1;
    return this;
}

Changer* Changer::blink2() {
    attr = Blink2;
    return this;
}

Changer* Changer::reverse() {
    attr = Reverse;
    return this;
}

Changer* Changer::concealed() {
    attr = Concealed;
    return this;
}

Changer* Changer::crossedout() {
    attr = Crossedout;
    return this;
}

Changer* Changer::bgBlack() {
    bg = BgBlack;
    return this;
}

Changer* Changer::bgRed() {
    bg = BgRed;
    return this;
}


Changer* Changer::bgGreen() {
    bg = BgGreen;
    return this;
}

Changer* Changer::bgYellow() {
    bg = BgYellow;
    return this;
}

Changer* Changer::bgBlue() {
    bg = BgBlue;
    return this;
}

Changer* Changer::bgMagenda() {
    bg = BgMagenda;
    return this;
}

Changer* Changer::bgCyan() {
    bg = BgCyan;
    return this;
}

Changer* Changer::bgWhite() {
    bg = BgWhite;
    return this;
}
