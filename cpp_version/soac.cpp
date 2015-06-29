#include "soac.h"

std::string Changer::Apply(std::string val) {
    tmp << "\x1b[" << attr << ";" << fg << ";" << bg << "m%v\x1b[0m";
    std::string tmptmp = tmp.str();
    tmp.str("");
    return tmptmp;
}

std::string Changer::Apply(int val) {
    tmp << "\x1b[" << attr << ";" << fg << ";" << bg << "m%v\x1b[0m";
    std::string tmptmp = tmp.str();
    tmp.str("");
    return tmptmp;
}

int main() {
    return 0;
}
