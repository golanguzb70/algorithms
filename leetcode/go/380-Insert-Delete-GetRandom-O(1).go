package leetcode

import "math/rand"

/*
Implement the RandomizedSet class:

    RandomizedSet() Initializes the RandomizedSet object.
    bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
    bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
    int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.

You must implement the functions of the class such that each function works in average O(1) time complexity.

*/

type RandomizedSet struct {
	values map[int]bool
}

func ConstructorRondomizedSet() RandomizedSet {
	return RandomizedSet{
		values: map[int]bool{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.values[val]
	if !ok {
		this.values[val] = true
	}
	return !ok
}

func (this *RandomizedSet) Remove(val int) bool {
	_, ok := this.values[val]
	if ok {
		delete(this.values, val)
	}
	return ok
}

func (this *RandomizedSet) GetRandom() int {
	rand := rand.Intn(len(this.values))
	i := 0

	for v := range this.values {
		if i == rand {
			return v
		}
		i++
	}
	return 0
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
