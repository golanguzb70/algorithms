#include <iostream>

using namespace std;

/*
What is hamming distance? 
 - Hamming distance is number of different bits between to same sized data.

 Algorithm to solve.
 - Two integers are give first they have to be converted into binaries
 - While converting into binaries compare the bit if they are not equal or one number have more bits increment count.

*/

class Solution {
public:
    int hammingDistance(int x, int y) {
        int count = 0;
        while(x > 0 || y > 0) {
            count += !((x % 2) == (y % 2));
            if (x > 0)
                x /= 2;
            if (y > 0) 
                y /= 2;
        }
        return count;
    }
};

int main() {
    Solution s;

    cout << s.hammingDistance(1, 4) << " == 2" << endl;
    cout << s.hammingDistance(3, 1) << " == 1" << endl;

    return 0;
}