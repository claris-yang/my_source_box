//
// Created by yangtao on 18-3-24.
//

#include "stock.h"
namespace NH{
    Stock::Stock(const std::string &company, long shares, double share_val ){
        _company = company;
        _shares = shares;
        _share_val = share_val;
    }

    NH::Stock::Stock() {
        _company = "baidu";
        _shares = 100;
        _share_val = 9.876543212456;
        std::cout << "shengcheng" << std::endl;
    }

    NH::Stock::~Stock() {
        std::cout << "bye bye" << std::endl;
    }

    NH::Stock::Stock(int n) {
        _share_val = n;
    }

    void NH::Stock::show() {
        std::cout << "secondary function()" << std::endl;
    };

    void NH::Stock::show() const{

        std::ios_base::fmtflags orig = std::cout.setf(std::ios_base::fixed, std::ios_base::floatfield);
        std::streamsize prec = std::cout.precision(3);

        std::cout << "company = " << _company << ": shares :" << _shares << " :share_val : ";
        std::cout.setf(std::ios_base::fixed);
        std::cout.precision(8);
        std::cout << _share_val << std::endl;

        std::cout.setf(orig);
        std::cout.precision(prec);
    }

    std::string Stock::longestCommomPrefix(std::vector<std::string> &strs){
        std::string result = "";
        if (strs.size() == 0)
        {
            return result;
        }
        for(int idx = 0 ; ; idx++){
            std::string one = "";
            for (int i = 0 ; i < strs.size(); i++ ){

                if (strs[i].length() <= idx)
                    goto over;

                if ( "" == one)
                    one = strs[i].substr(idx,1);

                if (one != strs[i].substr(idx, 1)) {


                    return result;
                }

            }
            result += one;
        }
        over:
        return result;
    }

    std::string Stock::addBinary(std::string a, std::string b) {
        int prior = 0 ;
        std::string str;
        for (int i = a.length() - 1, j = b.length() - 1; i >= 0 || j >= 0 ; ) {
            /*
            if (i < 0 ) {
                str += (std::stoi(b[j--] + "") + prior) % 2 + "";
                prior = (std::stoi(b[j--] + "") + prior) / 2 ;
                continue;
            }
            if (j < 0) {
                str += (std::stoi(a[i--] +"") + prior) % 2 + "";
                prior = (std::stoi(a[i--]+"") + prior ) / 2 ;
                continue;
            }

            str += (std::stoi(a[i--] + "") + std::stoi(b[j--] + "") + prior) % 2 + "";
            prior = (std::atoi(a[i--] + "") + std::stoi(b[j--] + "") + prior) / 2;
            */
        }

        if (prior == 1) {
            str += "1";
        }

        std::string result = "";
        for (auto s : str) {
            result += s;
        }

        return result;
    }


}

