#include <iostream>
#include <vector>
#include <map>
#include <string>
#include <algorithm>

using namespace std;


class Solution {
public:
    vector <vector <string> > groupAnagrams(vector<string>& strs) {
        unordered_map<string, vector<string> > mp;
        vector< vector <string> > response;

        for (int i = 0; i < strs.size(); i++) {
            string temp = strs[i];
            sort(temp.begin(), temp.end());
            auto value = mp.find(temp);

            if(value != mp.end()) {
                value->second.push_back(strs[i]);
                mp.insert(make_pair(temp, value->second));
            } else {
                vector<string> tempVec;
                tempVec.push_back(strs[i]);
                mp.insert(make_pair(temp, tempVec));
            }
        }

        for (auto it = mp.begin(); it != mp.end(); ++it) {
            response.push_back(it->second);
        }
        return response;
    }
};


int main() {
    Solution s;
    vector<string> strs ;
    strs.push_back("eat");
    strs.push_back("tea");
    strs.push_back("tan");
    strs.push_back("ate");
    strs.push_back("nat");
    strs.push_back("bat");
    vector< vector <string> > response = s.groupAnagrams(strs);
    for (int i = 0; i < response.size(); i++) {
        for (int j = 0; j < response[i].size(); j++) {
            cout << response[i][j] << " ";
        }
        cout << endl;
    }
    return 0;
}