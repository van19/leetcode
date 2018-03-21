package main

import "fmt"

func median(nums []int) float64 {
	sz := len(nums)
	mid := sz / 2
	if sz%2 == 0 {
		return float64(nums[mid-1] + nums[mid])/2
	}
	return float64(nums[mid])
}

func find_one(arr []int, one int, k int) int {
	//fmt.Println(arr,one,k)
	if k==0 {
		if arr[0] > one {
			return one
		}
		return arr[0]
	}
	sz := len(arr)
	if sz == k {
		if arr[sz-1] > one {
			return arr[sz-1]
		}
		return one
	}
	
	if one >= arr[k] {
		return arr[k]
	}
	if one >= arr[k-1] {
		return one
	}
	return arr[k-1]
}
func find_k(nums1 []int, nums2 []int, k int) int {
	//fint kth in nums1 & nums2
	//fmt.Println(nums1,nums2,k)
	sz1 := len(nums1)
	sz2 := len(nums2)
	if sz1 == 0 {
		return nums2[k]
	}
	if sz2 == 0 {
		return nums1[k]
	}
	//处理一个数，主要是下面逻辑无法处理单数，所以预先处理了
	if sz1 == 1 {
		return find_one(nums2,nums1[0],k)
	}
	if sz2 == 1 {
		return find_one(nums1,nums2[0],k)
	}

	sn1 := nums1
	sn2 := nums2
	m1 := sz1/2
	m2 := sz2/2
	m  := (sz1+sz2)/2

	if k < m {
		if nums1[m1]< nums2[m2] {
			sn2 = nums2[0:m2]
		}else{
			sn1 = nums1[0:m1]
		}
	}else{
		if nums1[m1]<nums2[m2] {	
			sn1 = nums1[m1:]
			k = k - m1
		}else{
			sn2 = nums2[m2:]
			k = k - m2
		}
	}

	return find_k(sn1,sn2,k)
} 

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sz1 := len(nums1)
	sz2 := len(nums2)
	sz := sz1 + sz2
	k := sz /2
	if sz%2 != 0 {
		return float64(find_k(nums1,nums2,k))
	}else{
		right := find_k(nums1,nums2,k)
		left  := find_k(nums1,nums2,k-1)
		return float64(left+right)/2
	}

	
}

func main(){
    nums1 := []int{1,2,3,7}
	nums2 := []int{4,5,6,8}
    fmt.Println(findMedianSortedArrays(nums1,nums2))
}