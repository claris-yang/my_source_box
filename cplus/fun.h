//
// Created by yangtao on 18-4-28.
//

#ifndef CPLUSPLUS_FUN_H
#define CPLUSPLUS_FUN_H
namespace fun {
    int half(int x);

    struct third_t {
        int operator() (int x) {return x/3;}
    };

    struct MyValue {
        int value;
        int fifth() {return value / 5;}
    };
}
#endif //CPLUSPLUS_FUN_H
