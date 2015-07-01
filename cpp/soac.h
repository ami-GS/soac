#ifndef _SOAC_H_
#define _SOAC_H_

#include <iostream>
#include <string.h>
#include <sstream>

typedef uint16_t COLOR;
typedef uint16_t ATTRIBUTE;
typedef uint16_t BACKGROUND;

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
    Faint,
    Italic,
    Underline,
    Blink1,
    Blink2,
    Reverse,
    Concealed,
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
    std::string Apply(std::string str);
    std::string Apply(int val);

    Changer* black();
    Changer* red();
    Changer* green();
    Changer* yellow();
    Changer* blue();
    Changer* magenda();
    Changer* cyan();
    Changer* white();

    Changer* reset();
    Changer* bold();
    Changer* faint();
    Changer* italic();
    Changer* underline();
    Changer* blink1();
    Changer* blink2();
    Changer* reverse();
    Changer* concealed();
    Changer* crossedout();

    Changer* bgBlack();
    Changer* bgRed();
    Changer* bgGreen();
    Changer* bgYellow();
    Changer* bgBlue();
    Changer* bgMagenda();
    Changer* bgCyan();
    Changer* bgWhite();
};

#endif // _SOAC_H_
