package main

import "fmt"

func main(){
	arr:=&[]string{"Nur-Sultan","Almaty","Qaragandy"}
	arr1:=&[]int{2,1,9}
	BubbleSort(arr1)
	fmt.Println(*arr1)
	fmt.Println(arr)
	MergeSort(arr)
	fmt.Println(arr)
}

func MergeSort(arr *[]string){
	i:=len(*arr)
	if i<2{
		return
	}
	left:=make([]string,i-i/2)
	right:=make([]string,i/2)

	x:=0
	for j:=0;j<len(left);j++{
		left[j]=(*arr)[j]
		x=j+1
	}

	for k:=0;k<len(right);k++{
		right[k]=(*arr)[x]
		x++
	}

	MergeSort(&left)
	MergeSort(&right)
	Merge(&left,&right,arr)

}

func Merge(left, right,arr *[]string){
	l:=len(*left)
	r:=len(*right)
	i:=0
	j:=0
	k:=0
	for i<l && j<r {
		if (*left)[i]<=(*right)[j]{
			(*arr)[k]=(*left)[i]
			k++
			i++
		} else {
			(*arr)[k]=(*right)[j]
			k++
			j++
		}
	}
	for i<l{
			(*arr)[k]=(*left)[i]
			k++
			i++
	}
	for j<r{
		(*arr)[k]=(*right)[j]
		j++
		k++
	}
}
func BubbleSort(arr *[]int){
	for i:=0; i<len(*arr)-1; {
		first:=(*arr)[i]
		second:=(*arr)[i+1]
		if (*arr)[i]>(*arr)[i+1]{
			(*arr)[i]=second
			(*arr)[i+1]=first
			i=0
		} else {
			i++
		}
	}
}