package main

import(

	"fmt"
)

func returnSliced (arr []int) []int{

	return arr[2:4];
}

func main(){

	fmt.Println("Hello this is my first GO program");
	a :=[5]int{1,2,3,4,5};
	b := a[2:4]
	var c int;
	c = 4;
	var d = 5;
	fmt.Println(b,c,d)
	fmt.Println(returnSliced(a[:]));
}




