
#include<iostream>
#include<vector>

using namespace std;

void sortColors(vector<int>& nums) {
        int red = 0, white = 0, blue = 0;
        for(int i = 0 ; i < nums.size() ; i++) {
            if (nums[i] == 0) {
                red++;
                
            }
            if(nums[i] == 2) {
                blue++;
            }
            nums[i] = 1;
        }
        
        for(int i = 0 ; i < red; i++) {
            nums[i] = 0;
        }
        
	int side = nums.size() - blue;
        for(int i = nums.size() - 1; i >= side; i--) {
            nums[i] = 2;
        } 
        
        
}

int main() {
	vector<int> nums = {2};
	sortColors(nums);
	return 0;
}
