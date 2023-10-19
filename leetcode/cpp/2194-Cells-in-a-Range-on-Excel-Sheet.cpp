#include <iostream>
#include <vector>
#include <string>

using namespace std;

/*
    Two points are given (A, B) in excel in format A1:B3 
    Here A of A1 is a horizontal (x) axes and 1 is vertical(y) exes. 
    return list of all cell that between A1 and B3.

    Sorting first by columnes and then by rows:

    input: A1:B3

    | i |   A   |  B   |
    | 1 |   *   |      |
    | 2 |       |      |
    | 3 |       |   *  |

    response: [A1, A2, A3, B1, B2, B3]

    To sort convert row number to int and column letter to char then get its ASCI to sort easily.

    Time complexity -> O(x*y)
    Space Complexity -> O(1) -> consistent
*/
class Solution {
public:
    vector<string> cellsInRange(string s) {
        char x1 = s[0], x2 = s[3];
        char y1 = s[1], y2 = s[4];
        char temp = y1;
        vector<string> res;

        while(x1 <= x2) {
            temp = y1;
            while (temp <= y2) {
                res.push_back(string(1, x1) + string(1, temp));
                temp++;  
            }
            x1++;
        }
        return res;
    }
};


int main() {
    Solution s;
    vector<string> res = s.cellsInRange("A1:B3");
    for (int i = 0; i < res.size(); i++) {
        cout << res[i] << ", ";
    }
    cout << endl;
}