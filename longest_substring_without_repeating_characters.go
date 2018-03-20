package main

import "fmt"
import "strings"

func lengthOfLongestSubstring(s string) int {
    sz := len(s)
    if sz <= 1 {
        return sz
    }
    
    longest := 1
    begin := 0
    end := 1
    for end < sz {
        sub := s[begin:end]
        ch := s[end:end+1]
        i := strings.Index(sub,ch)
        if i==-1 {
            end++
            continue
        }

        length := end - begin
        if length > longest{
            longest = length
        }
        begin = begin + i + 1
        end++
        
    }
    length := end - begin
    if length > longest{
        longest = length
    }
    return longest
}

func main(){
    s := "abcdecfab"
    s = "aaaaabaaa"
    fmt.Printf("%d\n",lengthOfLongestSubstring(s))
}
