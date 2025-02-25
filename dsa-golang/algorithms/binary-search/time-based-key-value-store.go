package main

import (
	"slices"
)

/*Problem:
Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and
retrieve the key's value at a certain timestamp.

Implement the TimeMap class:

TimeMap() Initializes the object of the data structure.
void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
String get(String key, int timestamp) Returns a value such that set was called previously, with timestamp_prev <= timestamp.
If there are multiple such values, it returns the value associated with the largest timestamp_prev. If there are no values, it returns "".

*/

/*Solution:
Data structure is: map[string][[]struct{value, timestamp}]

Algorithm: Get by key and search for value using binary search.
*/

type Value struct {
	Val       string
	TimeStamp int
}

type TimeMap struct {
	Map map[string][]Value
}

func Constructor() TimeMap {
	return TimeMap{
		Map: map[string][]Value{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.Map[key] = append(this.Map[key], Value{Val: value, TimeStamp: timestamp})
	slices.SortFunc[[]Value](this.Map[key], func(a, b Value) int {
		if a.TimeStamp < b.TimeStamp {
			return -1
		}
		return 1
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	arr, ok := this.Map[key]
	if !ok {
		return ""
	}

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i].TimeStamp <= timestamp {
			return arr[i].Val
		}
	}

	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
