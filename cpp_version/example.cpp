#include <iostream>
#include "soac.h"
#include <string>

int main() {
    Changer* c1 = new Changer();
    Changer* c2 = new Changer();    
    c1->red()->bgBlue()->underline();
    c2->blue()->bgRed()->bold();
    std::cout << c1->Apply("This is") << c2->Apply(" a pen.") << std::endl;
    delete c1;
    delete c2;
    return 0;
}
