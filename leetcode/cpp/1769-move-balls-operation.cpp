#include <iostream>
#include <vector>

using namespace std;


/*
    n boxes are given if boxes[i] = 0 it is empty if it is 1 it has one ball. 

    operation -> moving one ball from one box to another

    requirement is to find minimum number of operations to move all balls to each box.

    moving balls to specific box starts from initial state of boxes for each box.


    Brute force -> Calculating minimum number of operations for each box.

    1. Iterate through all boxes.
    2. Calculate Right operations to move all balls to i th ball.
    3. Calculate Left Operations to move all balls to i th  ball.

*/


class Solution {
public:
    vector<int> minOperations(string boxes) {
        
    }
};