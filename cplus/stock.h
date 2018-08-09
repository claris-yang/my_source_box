//
// Created by yangtao on 18-3-24.
//

#ifndef CPLUSPLUS_STOCK_H
#define CPLUSPLUS_STOCK_H

#include<string>
#include<iostream>
#include <vector>
#include <cstdlib>

namespace NH{
    struct ListNode {
        int val;
        ListNode *next;
    };
    class Stock {
    private:
        std::string _company;
        long _shares;
        double _share_val;
        enum {Months = 12};
        double cost[Months];
    public:
        Stock(const std::string &company, long shares = 0, double share_val = 1.234567890);
        Stock();
        Stock(int) ;
        ~Stock();
        void show() const;
        void show();
        std::string longestCommomPrefix(std::vector<std::string> &strs);
        ListNode* addTwoNumbers(ListNode *l1, ListNode *l2);
        std::string addBinary(std::string a, std::string b);
    };
}

#endif //CPLUSPLUS_STOCK_H
