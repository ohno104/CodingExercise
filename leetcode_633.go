import "math"

func judgeSquareSum(c int) bool {    
    a, b := 0, int(math.Sqrt(float64(c)))
    
    for a<=b {
        sum := a*a+b*b
        if sum == c {
            return true
        }
        
        if sum < c {
            a++
            
        } else {
            b--
        }
    }
    return false
}

//Time Complexity: O(Sqrt(N))
//Space Complexity: O(1)