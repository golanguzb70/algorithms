package leetcode

import "fmt"

/*
There is an m x n grid with a ball. The ball is initially at the position [startRow, startColumn].
You are allowed to move the ball to one of the four adjacent cells in the grid
(possibly out of the grid crossing the grid boundary). You can apply at most maxMove moves to the ball.

Given the five integers m, n, maxMove, startRow, startColumn, return the number of
paths to move the ball out of the grid boundary. Since the answer can be very large, return it modulo 10^9 + 7.

Hints:
1. Is traversing every path feasible? There are many possible paths for a small matrix. Try to optimize it.
  - We use memorization method to optimize and protect traversing every path.

2. Can we use some space to store the number of paths and update them after every move?
  - Yes...

3. One obvious thing: the ball will go out of the boundary only by crossing it.
Also, there is only one possible way the ball can go out of the boundary from the
boundary cell except for corner cells. From the corner cell, the ball can go out in two different ways.
Can you use this thing to solve the problem?

  - So let's use backtracking algorithm

    func a(m, n, maxMove, currentRow, currentColumn) {
    if maxMove == 0 {
    return 0
    }

    res := 0;
    if currentRow == 0 or currentRow == m-1 {
    res++
    }
    if currentColumn == 0 or currentColumn == n-1 {
    res++
    }

    if currentRow < m-1 {
    res += a(m, n, maxMove-1, currentRow+1, currentColumn)
    }

    if currentRow > 0 {
    res += a(m, n, maxMove-1, currentRow-1, currentColumn)
    }

    if currentColumn < n-1 {
    res += a(m, n, maxMove-1, currentRow, currentColumn+1)
    }

    if currentColumn > 0 {
    res += a(m, n, maxMove-1, currentRow, currentColumn-1)
    }

    return res
    }

With this solution we now achieved right answer
But Time Limit exceeded.

# Now lets use memorization algorithm to memorize number of paths

map[string]number{}
key = "m_n_maxMove"
value = number of paths the one can go through from the cell with maxMove number of movements
*/
const modulo = 1000000007
func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	return helperfindPaths(m, n, maxMove, startRow, startColumn, map[string]int{})
}

func helperfindPaths(m, n, maxMove, currentRow, currentColumn int, mp map[string]int) int {
	if maxMove == 0 {
		return 0
	}

	key := fmt.Sprintf("%d_%d_%d", currentRow, currentColumn, maxMove)

	v, ok := mp[key]
	if ok {
		return v % modulo
	}

	res := 0
	if currentRow == 0 {
		res++
	}
	if currentRow == m-1 {
		res++
	}

	if currentColumn == 0 {
		res++
	}
	if currentColumn == n-1 {
		res++
	}

	if currentRow < m-1 {
		res += helperfindPaths(m, n, maxMove-1, currentRow+1, currentColumn, mp)
	}

	if currentRow > 0 {
		res += helperfindPaths(m, n, maxMove-1, currentRow-1, currentColumn, mp)
	}

	if currentColumn < n-1 {
		res += helperfindPaths(m, n, maxMove-1, currentRow, currentColumn+1, mp)
	}

	if currentColumn > 0 {
		res += helperfindPaths(m, n, maxMove-1, currentRow, currentColumn-1, mp)
	}

	mp[key] = res % modulo

	return mp[key]
}
