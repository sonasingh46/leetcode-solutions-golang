package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) <= 0 {
		return false
	}
	if len(matrix) > 0 {
		if len(matrix[0]) <= 0 {
			return false
		}
	}
	return searchColumn(matrix, target, 0, len(matrix)-1)
}

func searchColumn(matrix [][]int, key, low, high int) bool {
	if high < 0 {
		return false
	}
	if low <= high {
		mid := (low + high) / 2
		if key == matrix[mid][0] {
			return true
		}

		if key < matrix[mid][0] {
			return searchColumn(matrix, key, low, mid-1)
		}

		return searchColumn(matrix, key, mid+1, high)
	}
	return searchRow(matrix[high], key, 0, len(matrix[high])-1)
}

func searchRow(row []int, key, low, high int) bool {

	if low <= high {
		mid := (low + high) / 2
		if key == row[mid] {
			return true
		}

		if key < row[mid] {
			return searchRow(row, key, low, mid-1)
		}

		return searchRow(row, key, mid+1, high)
	}
	return false
}
