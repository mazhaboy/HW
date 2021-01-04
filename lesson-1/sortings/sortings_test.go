package main

import (
	"testing"
)

func TestMergeSort(t *testing.T){
	tests:=[][]string{[]string{"cbcd","basd","abc"},[]string{"hello","world","!!!"},[]string{},[]string{"abc","ab", "a"}}
	for _,v:= range tests {
		MergeSort(&v)
		for i := 0; i<len(v)-1; i++{
		if v[i]>v[i+1]{
			t.Fail()
			break
		}
		}
	}
}
func TestBubbleSort(t *testing.T){
	tests:=[][]int{[]int{3,2,1},[]int{4,1000,64,0,1,1,34,5},[]int{},[]int{12,1,1,13,4,98,77,22}}
	for _,v:= range tests {
		BubbleSort(&v)
		for i := 0; i<len(v)-1; i++{
			if v[i]>v[i+1]{
				t.Fail()
				break
			}
		}
	}
}