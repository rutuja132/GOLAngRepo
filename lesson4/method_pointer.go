package main
import "fmt"
type author struct{
	name string
	branch string
	particle int
}
//// Method with a receiver of author type 

func (a *author) show(abranch string){
	(*a).branch = abranch
}
func main(){
	res:= author{
		name: "ABC",
		branch:"CSE",
	}

	fmt.Println("Author's name:" , res.name)
	fmt.Println("Branch Name(Before):", res.branch)

//pouinter creastion
p:=&res


}