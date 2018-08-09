//
// Created by yangtao on 18-4-21.
//

#include "student.h"
int global =  1000;
void Student::set_name( std::string &n) {
    this->name = n ;
}

void Student::show() {
    std::cout << this->name << std::endl;
}

Student::Student(std::string & n) : name(n) {
    this->name = n;
}

Student::~Student() {
    std::cout << " Bye." << std::endl;
}

