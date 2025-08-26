package twosum

func run(nums []int, target int) []int {
	// Define result for mem efficiency
	var res []int;

	//Create HashMap for fast lookup of O(1) after looking 
	numMap := make(map[int]int)

	// Loops over all numbers in nums
	for i, v := range nums {

		// Determines the remainder of the number if the remainder is -100 it means that it should know a value that is 100 to reach the target
		b := target - v

		// Checks if the remainder value exists in the map, if it exists it can return the position of the value
		if j, exists := numMap[b]; exists {
			res = []int{i, j}
			break;
		}

		// If not store value in map
		numMap[v] = i
	}

	return res
}
