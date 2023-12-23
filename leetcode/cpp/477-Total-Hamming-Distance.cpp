#include <iostream>
#include <vector>
#include <bitset> 
#include <string> 


using namespace std;

/*
What is hamming distance? 
 - Hamming distance is number of different bits between to same sized data.

Requirement: 
- Array of integers are given 
- Return SUM of hamming distance between all pair numbers.

Algorithm to solve: -> (this idea gave time limit exceeded)
- Initialize count variable with 0 inital value and return it in the end
- Get pairs by iterating through two loops, one is inner loop. 
- Get hamming distance of pairs and add to count

*/

class Solution {
public:
    int totalHammingDistance(vector<int>& nums) {
        int size = nums.size();
        if(size < 2) return 0;

        int ans = 0;
        int *zeroOne = new int[2];

        while(true)
        {
            int zeroCount = 0;
            zeroOne[0] = 0;
            zeroOne[1] = 0;
            for(int i = 0; i < nums.size(); i++)
            {
                if(nums[i] == 0) zeroCount++;
                zeroOne[nums[i] % 2]++;
                nums[i] = nums[i] >> 1;
            }
            ans += zeroOne[0] * zeroOne[1];
            if(zeroCount == nums.size()) return ans;
        }
    }
};

int main() {
    Solution s;
    vector<int> input = {4, 14, 2};

    cout << s.totalHammingDistance(input) << " == 6" << endl;

    return 0;
}