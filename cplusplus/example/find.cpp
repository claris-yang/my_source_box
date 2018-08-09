
#include<cassert>
#include<iostream>
#include<algorithm>
#include<cstring>

using namespace std;

int main() {
	char s[] = "C++ is a better C";
	int len = strlen(s);
	const char * where = find(&s[0], &s[len], 'e');
	
	assert(*where == 'e' && *(where+1) == 't');
	where += 10;
	assert(where[1000]);
	cout << " ---- Ok " << endl;
}
