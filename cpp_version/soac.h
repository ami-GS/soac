#include <iostream>
#include <string.h>
#include <sstream>

typedef uint8_t COLOR;
typedef uint8_t ATTRIBUTE;
typedef uint8_t BACKGROUND;

enum {
    Black = 30,
    Red,
    Green,
    Yellow,
    Blue,
    Magenda,
    Cyan,
    White,
};

enum {
    Reset,
    Bold,
    Failt,
    Italic,
    Underline,
    Blink1,
    Blink2,
    Reverse,
    ConCealed,
    Crossedout,
};

enum {
    BgBlack = 40,
    BgRed,
    BgGreen,
    BgYellow,
    BgBlue,
    BgMagenda,
    BgCyan,
    BgWhite,
};
    

class Changer {
public:
    COLOR fg;
    ATTRIBUTE attr;
    BACKGROUND bg;
    std::stringstream tmp;

    Changer() : fg(39), attr(0), bg(49) {}
    std::string Apply(std::string str) {
        tmp << "\x1b[" << attr << ";" << fg << ";" << bg << "m%v\x1b[0m";
        std::string tmptmp = tmp.str();
        tmp.str("");
        return tmptmp;
    }
    std::string Apply(int val) {
        tmp << "\x1b[" << attr << ";" << fg << ";" << bg << "m%v\x1b[0m";
        std::string tmptmp = tmp.str();
        tmp.str("");
        return tmptmp;
    }
};
