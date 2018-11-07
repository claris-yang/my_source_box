
 #include<iostream>
 #include<algorithm>
 #include<vector>

	using namespace std;

struct Interval {
    int start;
    int end;
    Interval() : start(0), end(0) {}
    Interval(int s, int e) : start(s), end(e) {}
};

bool mycompare(Interval &a, Interval &b) {
    return (a.start < b.start);
}


 vector<Interval> merge(vector<Interval>& intervals) {
        
        vector<Interval> nums;
        sort(intervals.begin(), intervals.end(), mycompare);
        if(intervals.size() > 0) {
            nums.push_back(intervals[0]);
        }

	for(int i = 0 ; i < intervals.size(); i++ ) {
		std::cout << intervals[i].start << " " << intervals[i].end << std::endl;
	}
            
        for(int i = 1; i < intervals.size(); i++) {
            int l = nums.size() -1;
            if(l >=0 && intervals[i].start == nums[l].start ) {
                Interval elem(nums[i].start, nums[l].end > intervals[i].end ? nums[l].end:intervals[i].end);
		std::cout << "intsert data " << elem.start << " " << elem.end << " " << nums[i].start << " " << intervals[i].start << std::endl;
                nums.pop_back();
                nums.push_back(elem);
            } else if(l >= 0 && intervals[i].start <= nums[l].end) {
                Interval elem(nums[l].start, nums[l].end > intervals[i].end ? nums[l].end: intervals[i].end);
                nums.pop_back();
                nums.push_back(elem);
            } else {
                nums.push_back(intervals[i]);
            }
        }
        return nums;
    }

    int main() {
	
	vector<Interval> nums = {{1,4},{1,4}};
	vector<Interval> result = merge(nums);
	for(int i = 0 ; i < result.size(); i++ ) {
		std::cout << result[i].start << " " << result[i].end << std::endl;
	}

    }
