
#include<iostream>
#include<vector>

using namespace std;

int max(int i, int j)
{
	return i>j ? i:j;
}

int uniquePathsWithObstacles(vector<vector<int>>& obstacleGrid) {

        if ( obstacleGrid.size() == 0)
            return 0;
        int path[obstacleGrid.size() + 1][obstacleGrid[0].size() + 1]; 
	for(int i = 0 ; i <= obstacleGrid.size(); i++)
		for(int j = 0 ; j <= obstacleGrid[0].size(); j++)
			path[i][j] = 0;
        path[1][1] = 1;
        for(int i = 1; i <= obstacleGrid.size(); i++) {
            for(int j= 1; j <= obstacleGrid[i-1].size(); j++) {
                if(obstacleGrid[i-1][j-1] != 1 ) 
		{
                    path[i][j] = max(path[i][j], path[i-1][j] + path[i][j-1]);
		    std::cout << path[i][j] << std::endl;
		}
            }
        }
        return path[obstacleGrid.size()][obstacleGrid[0].size()];
}

int main() {
	
	vector<int> n1 = {0,0,0};
	vector<int> n2 = {0,1,0};
	vector<int> n3 = {0,0,0};
	vector<vector<int>> nums;
	nums.push_back(n1);
	nums.push_back(n2);
	nums.push_back(n3);
std::cout<< 	uniquePathsWithObstacles(nums) << std::endl;
}
