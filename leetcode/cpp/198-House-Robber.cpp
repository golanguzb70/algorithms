#include <iostream>
#include <vector>
#include <map>

using namespace std;


/*
You are a professional robber planning to rob houses along a street. 
Each house has a certain amount of money stashed, the only constraint stopping you from robbing 
each of them is that adjacent houses have security systems connected and it will automatically 
contact the police if two adjacent houses were broken into on the same night.

Given an integer array nums representing the amount of money of each house, return the maximum 
amount of money you can rob tonight without alerting the police.

Input: [1,2,3,1]
There are 2 initial houses to start robbing 1 and 2
Let's say we started robbing the house 1.
money = 1
Next I can rob house 3 or 4. 
    Let's say we robbed number 3
        money = 4
    Let's say we robbet house 4;
        money = 2
    Max money = 4
Let's say we started robbing house 2
    money=2
    then we can only rob house 4
        money = 3
    max money = 3

result = 4

Let's use recursive algirthm;
func(houses[], house_to_rob)
    if len(houses) <= house_to_rob {
        return 0;
    }

    return max(houses[house_to_rob]+func(houses, house_to_rob+2), houses[house_to_rob]+func(houses, house_to_rob+3));

With this solution we achieved right output bad bad performance: 
Time complexity O((n/2)!) which is really bad :(

So to optimize it I am going to use memorization technique by using hash map data structure
Time complexity (n);
*/

class Solution {    
    map<int, int> mp;
public:
    int rob(vector<int>& nums) {
        return max(helper(nums, 0), helper(nums, 1));
    }

    int helper(vector<int>& houses, int house_to_rob) {
        if (houses.size() <= house_to_rob) {
            return 0;
        }

        if (mp.find(house_to_rob) != mp.end()) {
            return mp[house_to_rob];
        }

        mp[house_to_rob] = max(houses[house_to_rob] + helper(houses, house_to_rob + 2),houses[house_to_rob] + helper(houses, house_to_rob + 3) );
        return mp[house_to_rob];
    }

    int max(int a, int b) {
        if (a > b) {
            return a;
        }

        return b;
    }
};

int main() {
    Solution s;
    vector<int> nums = {1,2,3,1};
    cout << s.rob(nums) << endl;
    return 0;
}