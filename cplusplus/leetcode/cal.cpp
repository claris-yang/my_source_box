
#include<iostream>
int calSum(int num) {
        int sum = 0;
        while(num > 0) {
            sum += num % 10;
            num /= 10;
        }
        return sum;
        
    }
    int addDigits(int num) {
        while(num > 10) {
            num = calSum(num);
        }
        return num;
    }

int main() {
	std::cout << addDigits(38) << std::endl;
}
