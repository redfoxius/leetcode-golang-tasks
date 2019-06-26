/**
*   https://leetcode.com/problems/two-sum/
* 	4 ms,	3.7 MB	
**/
func twoSum(nums []int, target int) []int {
    r := make(map[int]int)
        
    for i, x := range(nums) {        
        end := target - x
        if val, ok := r[end]; ok {
            return []int{val, i}
        }
        
        r[x] = i
    }
    
    return []int{}
}
