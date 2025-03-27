package main
import "fmt"
type author struct{
	name string
	branch string
    particles int
	salary int
}
func main(){
	res:= author{
		name :"ABC",
		branch:"csd",
		particles: 203,
		salary:34000,
	}
	res.show()
}