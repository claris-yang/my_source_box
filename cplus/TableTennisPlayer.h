//
// Created by yangtao on 18-4-25.
//

#ifndef CPLUSPLUS_TABLETENNISPLAYER_H
#define CPLUSPLUS_TABLETENNISPLAYER_H

#include <string>
#include <iostream>
using std::string;

class TableTennisPlayer
{
private:
    string firstname;
    string lastname;
    bool hasTable;
    static int strings_num;
public:
    TableTennisPlayer(const string & fn= "none",
    const string& ln = "none", bool ht = false);
    TableTennisPlayer(TableTennisPlayer & t);
    void Name() const;
    bool HasTable() const {return hasTable;};
    void ResetTable(bool v) {hasTable =v ;};
    string & operator[](int n);
    virtual void ViewAcct() const;
    virtual ~TableTennisPlayer();
};



class RatedPlayer : public TableTennisPlayer, private std::string
{
private:
    unsigned int rating;
public:
    RatedPlayer(unsigned int r = 0, const string & fn = "none",
    const string &ln  = "none", bool ht = false) : std::string("nihao"){
        std::cout << "invoke RatedPlayer start()" << std::endl;

    }
    RatedPlayer(unsigned int r, const TableTennisPlayer &tp);
    unsigned int Rating() const { return rating;}
    void ResetRating(unsigned int r) { rating = r;}
    void ViewAcct() const;
    void showString() ;
    ~RatedPlayer();

};
#endif //CPLUSPLUS_TABLETENNISPLAYER_H
