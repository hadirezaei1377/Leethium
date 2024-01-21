we faced to problem and now we are aware of that,
we have this as a given :

func twoSum(nums []int, target int) []int {
    
}

change to it:
func TwoSum(nums []int, target int) []int {
	res := []int{}

	return res
}

so after that we write test function for getting inputs and return outputs based on examples:
Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]

write test file and test each testcase based on for loop.
run our test by go test

then write the files 

after that go test and go test -v
