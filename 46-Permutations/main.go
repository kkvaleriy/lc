package main

//https://leetcode.com/problems/permutations/description/

func main() {
	input := []int{1, 2, 3}
	_ = permute(input)
}

func permute(nums []int) [][]int {
	output := make([][]int, 0)
	tmp := make([]int, 0)
	recursPermute(nums, tmp, &output)
	return output
}

func recursPermute(nums, tmp []int, out *[][]int) {

	if len(nums) == 0 {
		*out = append(*out, tmp)
		return
	}

	for k, v := range nums {
		recursPermute(append(append([]int{}, nums[:k]...), nums[k+1:]...), append(tmp, v), out)
	}
}
