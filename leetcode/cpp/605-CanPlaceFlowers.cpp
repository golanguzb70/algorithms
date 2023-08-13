#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    bool canPlaceFlowers(vector<int>& flowerbed, int n) {
        int length = flowerbed.size();
        if (length == 1) {
            return flowerbed[0]==0 || n == 0;
        } 
        for (int i = 0; i < length; i++) {
            if (i > 0 && i < length - 1) {
                if (flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i + 1] == 0) {
                    n--;
                    flowerbed[i] = 1;
                    i++;
                }
            } else if (i == 0) {
                if (flowerbed[i] == 0 && flowerbed[i + 1] == 0) {
                    n--;
                    flowerbed[i] = 1;
                    i++;
                }
            } else {
                if (flowerbed[i] == 0 && flowerbed[i-1] == 0) {
                    n--;
                    flowerbed[i] = 1;
                    i++;
                }
            }
            if (n  == 0) {
                return true;
            } 
        }
        return n <= 0;
    }
};

int main () 
{
    vector<int> flowerbed = {1,0,0,0,1};
    Solution s; 
    cout << s.canPlaceFlowers(flowerbed, 2) << endl;

    return 0;
}