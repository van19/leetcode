
package main

import "fmt"
type ListNode struct {
	Val int
	Next *ListNode
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := ListNode{}
	res := &head
	p1 := l1
	p2 := l2
	add := 0
	for p1!=nil && p2 !=nil {
		v1 := p1.Val
		v2 := p2.Val
		v := v1 + v2 + add
		x := v % 10
		add = v / 10
		node := ListNode{Val:x,Next:nil}
		res.Next = &node

		p1=p1.Next
		p2=p2.Next
		res = res.Next
	}
	p := p1
	if p2 != nil {
		p = p2
	}
	for p!=nil {
		v := p.Val + add
		x := v % 10
		add = v / 10
		node := ListNode{Val:x,Next:nil}
		res.Next = &node

		p=p.Next
		res=res.Next
	}
	if(add!=0){
		node := ListNode{Val:add,Next:nil}
		res.Next = &node
	}
	return head.Next
	
}
func make(x int) *ListNode {
	node := ListNode{}
	head := &node
	a := x / 10
	b := x % 10
	head.Next=nil
	head.Val = b
	res := head
	for a!=0 {
		b = a%10
		node := ListNode{}
		node.Val =b
		node.Next = nil
		res.Next = &node 
		res = res.Next
		a = a/10
	}
	return head
}
func (node ListNode) output() {
	p := &node
	for p!=nil {
		fmt.Printf("%d,",p.Val)
		p=p.Next
	}
	fmt.Printf("\n")
}
func main(){
	a := 5
	b := 5
	l1 := make(a)
	l1.output()
	l2 := make(b)
	l2.output()
	
	res := addTwoNumbers(l1,l2)
	res.output()
}