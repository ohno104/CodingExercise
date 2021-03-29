func twoSum(numbers []int, target int) []int {
    indexLow := 0
    indexHigh := len(numbers) - 1
    
    for indexLow < indexHigh {
        sum := numbers[indexLow] + numbers[indexHigh]
        
        if sum == target {
            return []int {indexLow+1, indexHigh+1}            
        }
        
        if  sum > target {
            indexHigh--
            
        } else {
            indexLow++
        }
    }
    return nil
}

//Time Complexity: O(N)
//Space complexity: O(1)