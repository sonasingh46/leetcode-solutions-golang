package main

const (
	INFPositive = 999999
	INFNegative = -999999
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	firstArrayLength := len(nums1)
	secondArrayLength := len(nums2)

	low := 0
	high := firstArrayLength

	for low <= high {
		firstArrayPartition := (low + high) / 2
		secondArrayPartition := (firstArrayLength+secondArrayLength+1)/2 - firstArrayPartition
		var maxLeftFirst, minRightFirst int

		if firstArrayPartition == 0 {
			maxLeftFirst = INFNegative
		} else {
			maxLeftFirst = nums1[firstArrayPartition-1]
		}

		if firstArrayPartition == firstArrayLength {
			minRightFirst = INFPositive
		} else {
			minRightFirst = nums1[firstArrayPartition]
		}

		var maxLeftSecond, minRightSecond int

		if secondArrayPartition == 0 {
			maxLeftSecond = INFNegative
		} else {
			maxLeftSecond = nums2[secondArrayPartition-1]
		}

		if secondArrayPartition == secondArrayLength {
			minRightSecond = INFPositive
		} else {
			minRightSecond = nums2[secondArrayPartition]
		}

		if maxLeftFirst <= minRightSecond && maxLeftSecond <= minRightFirst {

			if (firstArrayLength+secondArrayLength)%2 == 0 {
				return float64(max(maxLeftFirst, maxLeftSecond) +
					min(minRightFirst, minRightSecond)) / 2.0
			} else {
				return float64(max(maxLeftFirst, maxLeftSecond))
			}

		} else if maxLeftFirst > minRightSecond {
			high = firstArrayPartition - 1
		} else {
			low = firstArrayPartition + 1
		}
	}
	panic("Illegal Input")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
