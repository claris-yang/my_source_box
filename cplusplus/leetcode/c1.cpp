
#include<iostream>
#include<vector>

using namespace std;

int main() {
	
	int target = 4;

	vector<int> nums = {4,5,6,7,0,1,2};
        
        int repeat = 0;
        bool find = false;
        for(int i = 0 ; i < nums.size(); i++) {
            if(nums[i] < target) {
                repeat++;
                
            } else if (nums[i] == target){
                find = true;
            }
        }

	std::cout << repeat << std::endl;
        
        if(find)
            return repeat;
        return -1;


}
