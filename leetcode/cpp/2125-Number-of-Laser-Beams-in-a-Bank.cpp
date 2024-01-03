#include <iostream>
#include <vector>

using namespace std;

/*
Anti-theft security devices are activated inside a bank. You are given a 0-indexed binary string array bank representing the floor plan of the bank, which is an m x n 2D matrix. bank[i] represents the ith row, consisting of '0's and '1's. '0' means the cell is empty, while'1' means the cell has a security device.

There is one laser beam between any two security devices if both conditions are met:

    The two devices are located on two different rows: r1 and r2, where r1 < r2.
    For each row i where r1 < i < r2, there are no security devices in the ith row.

Laser beams are independent, i.e., one beam does not interfere nor join with another.

Return the total number of laser beams in the bank.

Algorithm to solve: 
1. Lets assume 0 string as number 0
2. Get sum of i'th row 
3. Multiply i'th rows laser devices count with i+1'th rows count if there are at least 1 device in the row.
4. Gain all the additions to one variable.
*/

class Solution {
public:
    int numberOfBeams(vector<string>& bank) {
        int result = 0, previus_row = 0, current_row = 0;
        
        for (int i = 0; i < bank.size(); i++) {
            if(i == 0) {
                for (int j = 0; j < bank[i].size(); j++) {
                    previus_row += bank[i][j] == '1';
                }
                continue;
            } 
            
            for (int j = 0; j < bank[i].size(); j++) {
                current_row += bank[i][j] == '1';
            }

            if (current_row == 0) {
                continue;
            }

            result += current_row  * previus_row;
            previus_row = current_row;
            current_row = 0;
        }
        
        return result;
    }
};

int main() {
    Solution s;
    vector<string> in = {"011001","000000","010100","001000"};

    cout << s.numberOfBeams(in) << endl;
    
    return 0;
}