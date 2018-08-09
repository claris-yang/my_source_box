//
// Created by yangtao on 18-4-25.
//

#include "TableTennisPlayer.h"
#include <iostream>
int TableTennisPlayer::strings_num = 0;

void TableTennisPlayer::ViewAcct() const {
    std::cout << "father view Acct " << std::endl;
}

TableTennisPlayer::~TableTennisPlayer() {
    std::cout << "father ~TableTennisPlayer " << std::endl;
}
string&  TableTennisPlayer::operator[](int n) {
    return this->firstname;
}

TableTennisPlayer::TableTennisPlayer(TableTennisPlayer &t) {
    std::cout << "copy construct " << std::endl;
}

TableTennisPlayer::TableTennisPlayer(const string &fn, const string &ln, bool ht) :
firstname(fn), lastname(ln), hasTable(ht){
    std::cout << "invoke TableTennisPlayer start()" << std::endl;
}

void TableTennisPlayer::Name() const {
    std::cout << lastname << ", " << firstname;
    std::cout << strings_num << std::endl;
}



void RatedPlayer::ViewAcct() const {
    std::cout << "child viewAcct() " << std::endl;
}

RatedPlayer::~RatedPlayer() {
    std::cout << "vhild ~RatedPlayer " << std::endl;
}

void showString() {
}