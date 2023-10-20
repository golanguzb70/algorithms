#include <iostream>
#include <vector>



using namespace std;

/*
https://leetcode.com/problems/two-furthest-houses-with-different-colors/description/

I am going use 2 pointer algorithm. 

max_distance=0
p1                                  P2
1  -  1  -  1  -  6  -  1  -  1  -  1  : max_distance = (abs(p1-p2) > max_distance ? abs(p1-p2) : max_distance)
if distance between p2 and p1 is already smaller than max_distance there is no meaning of iterating through other parts.
in innner loop p2 is decreases and in outur loop p1 increases.
*/ 

class Solution {
public:
    int maxDistance(vector<int>& colors) {
        uint8_t max = 0, p1 = 0, p2 = colors.size() - 1;
        while(abs(p1-p2) > max) {
            for (int i = p2; i > p1; i--) {
                if (colors[p1] != colors[i]) {
                    max = (max < abs(p1-i) ? abs(p1-i)  : max);
                    break;
                }
            }
            p1++;
        }

        return max;
    }
};

int main() {
    Solution s;
    vector<int>  req = {1,1,1,6,1,1,1};
    int res =  s.maxDistance(req);

    cout << "Response is: " << res << endl;
    return 0;
}