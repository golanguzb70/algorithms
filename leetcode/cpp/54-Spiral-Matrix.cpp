#include <iostream>
#include <vector>

using namespace std;


class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        vector<int> result;
        int left = 0, right = matrix[0].size(), up = 0, down = matrix.size(), direction = 0;

        while (left < right && up < down) {
            switch (direction)
            {
            case 0:
                for (int i = left; i < right; i++) 
                    

                direction = 1;
                break;
            }
        }

    }
};

int main() {



    return 0;
}