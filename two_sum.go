package main

import "fmt"

func twoSum(nums []int, target int) []int {
    sz := len(nums)
    res := make([]int,2)
    for i,_ := range nums{
        for j:=i+1;j<sz;j++ {
            t := nums[i] + nums[j]
            if t == target {
                res[0]=i
                res[1]=j
                return res
            }
        }
        //fmt.Println(nums[i])
    }
    return res
}

func main(){
    
    nums := []int{2, 7, 11, 15}
    target := 9
    fmt.Println(nums)
    ret := twoSum(nums,target)
    fmt.Println(ret)
}
