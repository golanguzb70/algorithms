#include <iostream>
#include <vector>

using namespace std;


class Solution {
public:
    unsigned short removeElement(vector<int>& nums, int val) {
        unsigned short int size = nums.size();
        unsigned short int index = 0;
        unsigned short i = 0;
        while(index < size) {
            if (nums[index] == val) {
                for(i = index; i < size-1; i++ )
                    nums[i] = nums[i+1];
                size--;
            }
        
            if (nums[index] != val) {
                index++;
            }
        }
        return  size;
    }
};


int main () {
    vector<int> nums = {0,1,2,2,3,0,4,2};
    Solution s;
    int res = s.removeElement(nums, 2);
    cout << res << endl;
    for(int i = 0; i < nums.size(); i++) {
        cout << nums[i] << ", ";
    }
    cout << endl;
    return 0;
}