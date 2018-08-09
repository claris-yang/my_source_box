//
// Created by yangtao on 18-4-21.
//

#include "tercher.h"
#include "student.h"

void Tercher::test_student() {
    std::string n = "name1";
    Student stu1 = Student(n);

    Student stu2(n);

    Student *stu3 = new Student(n);

    Student stu4 = {n};



    // invoke
    stu1.show();
    stu2.show();
    stu3->show();
    stu4.show();

    delete stu3;


}

