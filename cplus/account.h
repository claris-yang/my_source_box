//
// Created by yangtao on 18-3-31.
//

#ifndef CPLUSPLUS_ACCOUNT_H
#define CPLUSPLUS_ACCOUNT_H

#include <iostream>
#include <cstdlib>
#include <cstring>
class Account {
private:
    std::string _name;
    std::string _acc_id;
    char *_sync_name;
    double _del;
    enum in_acc:short {count = 10};
    double in_acc_num[count];
public:

    char *operator[](int i) const;
    ~Account() {

        delete [] _sync_name;
        std::cout << "~Account()" << std::endl;
    };

    Account(std::string name, std::string acc_id, double del, char *sync_name) {
        _name = name;
        _acc_id = acc_id;
        _del  = del;
        int sync_name_len = strlen(sync_name);
        _sync_name = (char *)malloc(sync_name_len+1);
        strcpy(_sync_name, sync_name);
    };

    Account(const Account & acc) {
        _name = acc._name;
        _acc_id =  acc._acc_id;
        _del =  acc._del;
        int sync_name_len = strlen(acc._sync_name);
        _sync_name = (char *)malloc(sync_name_len);
        strcpy(_sync_name, acc._sync_name);
    }

     Account(double del) {
        _name = "default";
        _acc_id = "default_acc_id";
        _del = del;
    }

    void set_del(double del) ;

    void show_info() const;

    void sort_acc() ;

    bool operator>(const Account & aac) const;

    Account operator+(const Account &acc) const;

    operator double () const;

    friend Account operator+(double del, const Account &acc) ;

};

#endif //CPLUSPLUS_ACCOUNT_H
