#include <iostream>
#include <vector>

/*
https://leetcode.com/problems/minimum-falling-path-sum/description/?envType=daily-question&envId=2024-01-19

Given an n x n array of integers matrix, return the minimum sum of any falling path through matrix.

A falling path starts at any element in the first row and chooses the element in the next row that is either directly below or diagonally left/right. Specifically, the next element from position (row, col) will be (row + 1, col - 1), (row + 1, col), or (row + 1, col + 1)



Algorithm: 
I am going to use recuresive algorithm do deep dive into the 2d array
Get minimum of the three paths while going out the stack

*/

using namespace std;


class Solution {
public:
    int minFallingPathSum(vector<vector<int>>& matrix) {
        vector<int> results;
        for (int i = 0; i < matrix[0].size(); i++) {
            results.push_back(helper(matrix, 0, i));
        }

        return min(results);
    }

    int helper(vector<vector<int>>& matrix, int i, int j) {
        if(i == matrix.size() - 1) {
            return matrix[i][j];
        }

        vector<int> temp;

        if (j > 0) {
            temp.push_back(helper(matrix, i+1, j-1));
        }
        if (j < matrix[i].size()) {
            temp.push_back(helper(matrix, i+1, j));
        }
        if (j+1 < matrix[i].size()) {
            temp.push_back(helper(matrix, i+1, j+1));
        }
        return matrix[i][j] + min(temp);
    }

    int min(vector<int> arg) {
        if (arg.size() == 0) {
            return 0;
        }
        for (int i = 1; i < arg.size(); i++) {
            if( arg[i] < arg[0]) {
                arg[0] = arg[i];
            }
        }

        return arg[0];
    }
};

int main() {
    
    return 0;
}