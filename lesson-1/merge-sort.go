package main

import "fmt"

func main(){
	arr:=[]int{2,4,1,6,8,5,3,7}
	fmt.Println(MergeSort(arr))
	//left:=[]int{1,2,4,6}
	//right:=[]int{3,5,7,8}
	//fmt.Println(Merge(left,right))
}
func MergeSort(arr []int)[]int{
	i:=len(arr)
	if i<2{
		return arr
	}
	left:=make([]int,i-i/2)
	right:=make([]int,i/2)

	x:=0
	for j:=0;j<len(left);j++{
		left[j]=arr[j]
		x=j+1
	}


	for k:=0;k<len(right);k++{
		right[k]=arr[x]
		x++
	}

	MergeSort(left)
	MergeSort(right)


	fmt.Println(Merge(left,right))
	return Merge(left,right)
}

func Merge(left, right []int)[]int{
	l:=len(left)
	r:=len(right)
	arr:=make([]int,l+r)
	i:=0
	j:=0
	k:=0
	for i<l && j<r {
		if left[i]<=right[j]{
			arr[k]=left[i]
			k++
			i++
		} else {
			arr[k]=right[j]
			k++
			j++
		}
	}
	for i<l{
			arr[k]=left[i]
			k++
			i++
	}
	for j<r{
		arr[k]=right[j]
		j++
		k++
	}
	return arr
}