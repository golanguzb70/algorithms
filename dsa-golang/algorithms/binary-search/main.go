package main

func main() {
	timeMap := Constructor()

	timeMap.Set("foo", "bar", 1)
	if res := timeMap.Get("foo", 1); res != "bar" {
		println("Expected 'bar', but got", res)
	} else {
		println("bar")
	}

	if res := timeMap.Get("foo", 3); res != "bar" {
		println("Expected 'bar', but got", res)
	} else {
		println("bar")
	}

	timeMap.Set("foo", "bar2", 4)
	if res := timeMap.Get("foo", 4); res != "bar2" {
		println("Expected 'bar2', but got", res)
	} else {
		println("bar2")
	}

	if res := timeMap.Get("foo", 5); res != "bar2" {
		println("Expected 'bar2', but got", res)
	} else {
		println("bar2")
	}
}
