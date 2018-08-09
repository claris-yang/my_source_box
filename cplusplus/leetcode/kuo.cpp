
#include<iostream>

using namespage std;

int longestValidParentheses(string s) {
	int resp;
	vector<int> maxlen(s.length(), 0);
	for(int i = 1; i < s.length(); i++) 
	{
		if( s[i] == ')')
		{
			if (s[i-1] == '(') {
				
			}
		}
	}        
}

int main() {
}
