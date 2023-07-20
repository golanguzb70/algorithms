#include <iostream>
#include <vector>
#include <deque>

using namespace std;

class Solution {
public:
    vector<int> asteroidCollision(vector<int>& asteroids) {
        deque<int> aliveAstroids;
        int len = asteroids.size(), first, second;

        for(int i = 0; i < len; i++) {
            if (i == 0) {
                aliveAstroids.push_back(asteroids[i]);
            } else {
                second=asteroids[i];
                while(true) {
                    if (aliveAstroids.empty()) {
                        aliveAstroids.push_back(second);
                        break;
                    }
                    first = aliveAstroids.back();
                    if ((first > 0 && second > 0) || (first < 0 && second < 0) || (first < 0 && second > 0)) {
                        aliveAstroids.push_back(second);
                        break;
                    } else {
                        if (abs(first) < abs(second)) {
                            if (!aliveAstroids.empty()) {
                                aliveAstroids.pop_back();
                            }
                            if (aliveAstroids.empty()) {
                                aliveAstroids.push_back(second);
                                break;
                            }
                            first = aliveAstroids.back();
                        } else if (abs(first) == abs(second)) {
                            if (!aliveAstroids.empty()) {
                                aliveAstroids.pop_back();
                            }
                            break;
                        } else {
                            break;
                        }
                    }

                }
            }
        }
        vector<int> temp = {};
        while(!aliveAstroids.empty()) {
            temp.push_back(aliveAstroids.front());
            aliveAstroids.pop_front();
        }
        
        return temp;
    }
};

int main() {
    vector<int> astroids = {1,-2,-2,-2};
    
    astroids.clear();
    astroids.push_back(1);
    for (int i = 0; i < astroids.size(); i++) {
        cout << i + 1 << astroids[i] << endl;
    }
    
    return 0;
}