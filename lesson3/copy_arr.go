package main
import "fmt"
func main(){
	arr1:=[3]int{1,2,3}
	var arr2[3]intarr2=arr1
	arr2[0]=100
	fmt.Println("arr1",arr1)
	fmt.Println("arr2",arr2)

}