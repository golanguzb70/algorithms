package main

import "fmt"

/*Problem:
Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
*/

/*
Solution:

Validness rules & algorithms to detect:
1. Unique digits in each row -> key: "{row_index}{digit}", if exists return invalid
2. Unique digits in each column -> key: "{column_index}{digit}", if exists return invalid
3. Unique digits in each 3x3 sub-boxes -> key: "{box_number}{digit}", if exists return invalid

Box number calculation: (row/3)*3 + (column / 3)

Algorithm:
1. Loop through the matrix
2. if matrix[i][j] is not valid return false
3. at last return true

Time complexity: O(n^2)
Space complexity: O(n^2)
*/

func find3x3Box(column, row int) int {
	return (row/3)*3 + (column / 3)
}

func IsValidSudoku(board [][]byte) bool {
	rows := map[string]struct{}{}
	columns := map[string]struct{}{}
	box3x3 := map[string]struct{}{}
	var (
		rowKey, columnKey, box3x3Key string
	)

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '.' {
				continue
			}

			rowKey = fmt.Sprintf("%d%d", i, board[i][j])
			columnKey = fmt.Sprintf("%d%d", j, board[i][j])
			box3x3Key = fmt.Sprintf("%d%d", find3x3Box(j, i), board[i][j])
			if _, ok := rows[rowKey]; ok {
				return false
			} else {
				rows[rowKey] = struct{}{}
			}

			if _, ok := columns[columnKey]; ok {
				return false
			} else {
				columns[columnKey] = struct{}{}
			}

			if _, ok := box3x3[box3x3Key]; ok {
				return false
			} else {
				box3x3[box3x3Key] = struct{}{}
			}
		}
	}

	return true
}
