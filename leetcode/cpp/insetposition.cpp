#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        int low = 0, high = nums.size() - 1;

        int mid = (low + high) / 2;

        if (nums[0] > target) {
            return 0;
        } else if (nums[nums.size() - 1] < target) {
            return nums.size();
        }
        while (low <= high)
        {
            mid = (low + high) / 2;
            cout << "mid: " << mid << endl;
            if (nums[mid] == target) {
                return mid;
            } else if (nums[mid] < target) {
                low = mid + 1;
            } else {
                high = mid - 1; 
            }

            if (mid != 0) {
                if (nums[mid] > target && nums[mid - 1] < target) {
                     return mid;
                }
            }
        }
        return nums.size();
    }
};

int main() {
    Solution sol;
    vector<int> nums = {1, 2, 3, 4};
    int res = sol.searchInsert(nums, -5);
    cout << "Response: " << res << endl;
    return 0;
}