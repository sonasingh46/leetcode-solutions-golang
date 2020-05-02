package main

func twoSum(nums []int, target int) []int {
	hashMap:=make(map[int]int)
	arr:=make([]int,2)
	for i:=0;i<len(nums);i++{
		complement:=target-nums[i]
		if _,ok:=hashMap[complement];ok{
			arr[1]=i
			arr[0]=hashMap[complement]
			return arr
		}
		hashMap[nums[i]]=i
	}

	return arr
}