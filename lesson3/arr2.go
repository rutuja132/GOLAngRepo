package main
import "fmt"
func elementExists(arr[5]int,target int)bool{
	for _, value :=range arr{
		if value == target{
			return true
		}
	}
	return false
}
func main(){
	arr :=[5]int{1,2,3,4,5}
	target:=3
	if elementExists(arr,target){
		fmt.Println(target,"exists in the array")
	}else{
		fmt.Println(target,"does not exist in the array")
	}
}