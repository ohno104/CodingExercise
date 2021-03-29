/*
func reverseVowels(s string) string {       
    vowels:="aeiouAEIOU"
    leftIndex := 0
    rightIndex := len(s)-1
    
    rs := []rune(s)    
    for leftIndex < rightIndex {
        isLeftVoewl := strings.ContainsRune(vowels, rs[leftIndex])
        isRightVoewl := strings.ContainsRune(vowels, rs[rightIndex])               
        
        if isLeftVoewl && isRightVoewl {
            rs[leftIndex] , rs[rightIndex] = rs[rightIndex], rs[leftIndex]
            
            leftIndex++
            rightIndex--
            continue
        }
        
        if !isLeftVoewl {
            leftIndex++
        }
        
        if !isRightVoewl {
            rightIndex--
        }
    }
    
    return string(rs)
}
*/


func reverseVowels(s string) string {
    rs:= []byte(s)
    left:=0
    right:=len(s)-1
    
    for left<right {
        
        if isVowels(s[left]) && isVowels(s[right]) {
            rs[left] , rs[right] = rs[right], rs[left]
            left++
            right--
            continue
        }
        
        if !isVowels(s[left]) {
            left++
        }
        
        if !isVowels(s[right]) {
            right--
        }                
    }
    
    return string(rs)
}

func isVowels(ascii byte) bool {
    if ascii < 'a' {
        ascii += 32
    }
    
    return ascii == 'a' || ascii == 'e' || ascii == 'i' || ascii == 'o' || ascii == 'u'
}

//Time Complexity: O(N)
//Space Complexity: O(1)