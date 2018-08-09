
    #include<iostream>
    #include<string>
    #include<vector>
    #include<cstring>

    using namespace std;
    int n;
    int line[100];
    string print_line(int cur) {
        string line;
        for(int i = 0; i < n; i++) {
            line += ".";
        }
        line[cur] = 'Q';
	return line;
    }
    bool search(vector<string> &queen, int cur) 
    {
        if(cur == n) {
		for(auto s : queen)
			std::cout << s << std::endl;
		std::cout << "北京" << std::endl;
			
		return true;
        } else {
            for(int i = 0; i < n ;i++) {
                int ok = 1;
                line[cur] = i;
                for(int j =0 ; j < cur; j++) {
                    if( (line[cur] == line[j]) || (cur - line[cur] == j - line[j]) || (cur + line[cur] == j + line[j] )) {
                        ok = 0;
                        break;
                    }
                }
                
                if (ok ) {
                    string star = print_line(cur);
                    queen.push_back(star);
                    search(queen, cur+1);
                }
            }
        }
    }
    vector<vector<string>> solveNQueens(int x) {
	  n = x;
          vector<vector<string>> result;
          for(int i = 0 ; i < n ;i++){
	      memset(line, 0, sizeof(line));
              vector<string> queen;
              if(search(queen, i)) {
                  result.push_back(queen);
              }
          }
	
          return result;
    }

    int main() {
	solveNQueens(4);	
	return 0;
    }
