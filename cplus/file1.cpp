//
// Created by yangtao on 18-3-20.
//

#include <iostream>

namespace  HELLO {
    static int extern_int = 0;
    void Hello() {
        std::cout << "hello world" << std::endl;
        std::cout << extern_int << std::endl;
    }
}
