#include <iostream> 
#include <vector>


using namespace std;

/*
    1. first write fiping algorithm
        in flipping k is given and arr[0...k-1] should be reversed.
    
    2. In each iteration scope of the array should be reduced. 
    for example; 
    3, 2, 4, 1; 
    in first iteration -> greatest number is found which four and search is O(n); 
    lets take index of n(greatest number)  flip the n -> 4, 2, 3, 1
    to lead n which 4 to the end of the array put reverse again. when 4 reaches to end our scope of array reduces to n - 1 
    in the next iteration


    1: 3, 2, 4, 1 -> O(n) to find n,  k = index of (4) + 1 => 3;  4, 2, 3, 1; whole flip. k = index of (4) + 1) => 1; 1, 3, 2, 4
    2: 1, 3, 2 |  4 ->
*/


class Solution {
public:
    vector<int> pancakeSort(vector<int>& arr) {
        vector<int> result;
        for (int i = arr.size()- 1; i > 0; --i) {
            for (int j = 1; j <= i; ++j) {
                if (arr[j] == i + 1) {
                    flip(arr, j);
                    result.push_back(j+1);
                    break;
                }
            }
            flip(arr, i);
            result.push_back(i+1);
        }

        return result;
    }
    void flip(vector<int>& arr, int k) {
        int temp;
        for (int i = 0; i <= k/2; ++i) {
            temp = arr[i];
            arr[i] =arr[k-i];
            arr[k-i] = temp;
        }
    }
};

int main() {

    Solution s;

    vector<int> arr = {3, 2, 4, 1};

    vector<int> result = s.pancakeSort(arr);

    for (int i = 0; i < result.size(); i++) {
        cout << result[i] << ' ';
    }
    cout << endl;
    return 0;
}