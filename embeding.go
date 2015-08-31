package main
import "fmt"

type A struct{	}

func (a *A) M1() string {
	return "in A.M1"
}
type B struct{
	*A
}
func (b *B) M1() string {
	fmt.Printf("calling a.M1 %v\n", b.A.M1())
	return "in B.M1"
}

func main() {
	b := &B{}
	fmt.Printf("b.M1 %v\n", b.M1())
}