package main
import (
	"fmt"
)
type Std struct{
   name string
   age int
   email string
   address address
}
type address struct{
	address1 string
	address2 string
	pin int
}
func main(){
	//r:=Std{"rohith",19,"rohith.knaidu015@gmail.com",address{"hyfrtfuy","rctfrt",500090}}
	//fmt.Printf("%+v\n",r)
	r:=Std{
		name:"rohith",
		age:19,
		email:"rohiitrvjy",
		address: address{
			address1: "hyderabad",
			address2: "anantapur",
			pin:500090,
		},
	}

	fmt.Printf("%+v\n",r.address)
}