//
// Created by yangtao on 18-4-21.
//

#ifndef CPLUSPLUS_STUDENT_H
#define CPLUSPLUS_STUDENT_H
#include <iostream>

class Student
{
private:
     std::string &name;
public:
    Student(std::string &n);
    ~Student();
    void set_name( std::string &n);
    void show();
};
#endif //CPLUSPLUS_STUDENT_H
