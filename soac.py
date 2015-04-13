class COLOR:
    BLACK   = 30
    RED     = 31
    GREEN   = 32
    YELLOW  = 33
    BLUE    = 34
    MAGENDA = 35
    CYAN    = 36
    WHITE   = 37

class ATTRIBUTE:
    RESET      = 0
    BOLD       = 1
    FAINT      = 2
    ITALIC     = 3
    UNDERLINE  = 4
    BLINK1     = 5
    BLINK2     = 6
    REVERSE    = 7
    CONCEALED  = 8
    CROSSEDOUT = 9

class BG:
    BLACK   = 40
    RED     = 41
    GREEN   = 42
    YELLOW  = 43
    BLUE    = 44
    MAGENDA = 45
    CYAN    = 46
    WHITE   = 47

class Changer:
    def __init__(self):
        self.color = 39
        self.attr  = 0
        self.bg    = 49

    def apply(self, val):
        return "\x1b[%d;%d;%dm%s\x1b[0m" % (self.attr, self.color, self.bg, val) 

    def black(self):
        self.color = COLOR.BLACK
        return self

    def red(self):
        self.color = COLOR.RED
        return self

    def green(self):
        self.color = COLOR.GREEN
        return self

    def yellow(self):
        self.color = COLOR.YELLOW
        return self

    def blue(self):
        self.color = COLOR.BLUE
        return self

    def magenda(self):
        self.color = COLOR.MAGENDA
        return self

    def cyan(self):
        self.color = COLOR.CYAN
        return self

    def white(self):
        self.color = COLOR.WHITE
        return self

    def reset(self):
        self.attr = ATTRIBUTE.RESET
        return self

    def bold(self):
        self.attr = ATTRIBUTE.BOLD
        return self

    def faint(self):
        self.attr = ATTRIBUTE.FAINT
        return self

    def italic(self):
        self.attr = ATTRIBUTE.ITALIC
        return self

    def underline(self):
        self.attr = ATTRIBUTE.UNDERLINE
        return self

    def blink1(self):
        self.attr = ATTRIBUTE.BLINK1
        return self

    def blink2(self):
        self.attr = ATTRIBUTE.BLINk2
        return self

    def reverse(self):
        self.attr = ATTRIBUTE.REVERSE
        return self

    def concealed(self):
        self.attr = ATTRIBUTE.CONCEALED
        return self

    def crossedout(self):
        self.attr = ATTRIBUTE.CROSSEDOUT
        return self

    def b_black(self):
        self.bg = BG.BLACK
        return self

    def b_red(self):
        self.bg = BG.RED
        return self

    def b_green(self):
        self.bg = BG.GREEN
        return self

    def b_yellow(self):
        self.bg = BG.YELLOW
        return self

    def b_blue(self):
        self.bg = BG.BLUE
        return self

    def b_magenda(self):
        self.bg = BG.MAGENDA
        return self

    def b_cyan(self):
        self.bg = BG.CYAN
        return self

    def b_white(self):
        self.bg = BG.WHITE        
        return self
