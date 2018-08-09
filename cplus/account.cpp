//
// Created by yangtao on 18-3-31.
//

#include "account.h"
#include <algorithm>

void Account::show_info() const {
    std::cout << "name = " << _name << std::endl;
    std::cout << "acc_id = " << _acc_id << std::endl;
    std::cout << "sync_name = " << _sync_name << std::endl;

}

char * Account::operator[](int i ) const {
    switch (i) {
        case 1:
            return this->_sync_name;
        case 2:
            return "return 2";
    };
    return "default";
}


void Account::set_del(double del) {
    _del = del;
}

void Account::sort_acc() {
}

bool Account::operator>(const Account &acc) const {
    std::cout << acc._sync_name << std::endl;
    std::cout << this->_sync_name << std::endl;

    if (std::strcmp(acc._sync_name , this->_sync_name) > 0) {
        std::cout << "beijing " << std::endl;
        return true;
    } else if (std::strcmp(acc._sync_name, this->_sync_name) == 0) {
        std::cout << "tianjin" << std::endl;
        return false;
    } else {
        std::cout << "dapeng" << std::endl;
        return false;
    };
}

Account Account::operator+(const Account &t) const {
    Account a = Account(this->_name, this->_acc_id, this->_del + t._del, this->_sync_name);
    return a;
}

 Account operator+(double n, const Account &x)  {
      Account c(x._name, x._acc_id, x._del + n, "x name");
      return c;
}

Account::operator double() const {
    return this->_del;
}