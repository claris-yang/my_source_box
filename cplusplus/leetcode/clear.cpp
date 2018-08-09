
#include<iostream>
#include<vector>

using namespace std;

void clear(vector<vector<int>> &matrix, int m, int n, int sz, int yz) {
        for(int i = 0 ; i < yz; i++) 
            matrix[m][i] = 0;
        for(int j=0; j < sz; j++)
            matrix[j][n] = 0;
    }
    void setZeroes(vector<vector<int>>& matrix) {
        for(int i =0; i < matrix.size();i++) {
            for(int j=0; j < matrix[j].size(); j++) {
                if (matrix[i][j] == 0) {
                    clear(matrix, i, j, matrix.size(), matrix[j].size());
		     break;
                }
            }
        }
    }

int main() {
	vector<vector<int>> matrix = {{1,0}};
	setZeroes( matrix);
	for(int i = 0; i < matrix.size(); i++) {
		for (int j = 0; j < matrix[i].size(); j++)
			std::cout << matrix[i][j] << " ";
		std::cout << std::endl;
	}
	return 0;
}
