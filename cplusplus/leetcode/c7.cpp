
#include <vector>
#include <set>
#include <string>
#include <iostream>

using namespace std;

bool wordBreak(string s, vector<string>& wordDict) 
    {
 
        vector<int> pos(s.length() + 1, -1);
        set<string> mm(wordDict.begin(), wordDict.end());


	pos[0] = 1;
        
        if (s.length() == 0) return false;

        for( int i = 0; i < s.length(); i++) {
            
	    std::cout << pos[i+1] << std::endl;
            if(pos[i+1] != -1) continue;
            
	    for( int j = 0; j <= i ; j++) {
                
	    	std::cout << pos[j] << std::endl;
		if(  pos[j] != -1) {
                    string sub = s.substr(j , i - j + 1);
		    std::cout << sub << " = " << i << "j = " << j << std::endl;
                    set<string>::iterator it = mm.find(sub);
                    if ( it != mm.end()) {
                        pos[i+1] = 1;
			std::cout << "pos " << i << std::endl;
                        break;
                    }
		}
            }
        }

	for(int i = 0; i < pos.size(); i++)
	{
		for(int j = i+1; j < pos.size(); j++) {
			if(pos[j] != 1) {
				
			}
		}
	}
        
        return pos[s.length()] != -1;
    }

int main() {

	
//	vector<string> wordDict = {"leet", "code"};
//	std::cout << wordBreak("leetcode", wordDict) << std::endl;
	vector<string> wordDict = {"a"};
	std::cout << wordBreak("a", wordDict) << std::endl;
	
}
