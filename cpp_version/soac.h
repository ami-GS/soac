#include <iostream>

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
    COLOR color;
    ATTRIBUTE attr;
    BACKGROUND bg;

    Changer() : color(39), attr(0), bg(49) {} 
 
};
