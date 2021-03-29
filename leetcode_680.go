func validPalindrome(s string) bool {
    left:=0
    right:=len(s)-1
    
    for left < right {
        if s[left] != s[right] {
            //return isPalindrome(s[left+1: right+1]) || isPalindrome(s[left:right])
            return isPalindrome(s, (left+1), right) || isPalindrome(s, left, (right-1))
        }
        
        left++
        right--
    }
    return true   
}

/*
func isPalindrome(s string) bool {
    left:=0
    right:=len(s) - 1
    for left < right  {
        if s[left] != s[right] {
            return false           
        }
        left++
        right--
    }
    return true
}
*/

func isPalindrome(s string, start int, end int) bool {
    for start < end {
        if s[start] != s[end] {
            return false
        }
        
        start++
        end--
    }
    return true
}

//Time Complexity: O(N)
//Space Complexity: O(1)