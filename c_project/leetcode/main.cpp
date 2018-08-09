
#include<vector>
#include<iostream>

int main() {
	int num = 13;

	std::vector<int> result ;
	std::vector<int> resp;
	
	while(num ) {
		if (num & 0x01) 
			result.push_back(1);
		else 
			result.push_back(0);
		num = num >> 1;
	}
		
	for(int i = result.size() -1 ; i >= 0 ; i--)
		resp.push_back(result[i]);

	for ( auto s: resp)
		std::cout << s << " " << std::endl;
	
	
}
