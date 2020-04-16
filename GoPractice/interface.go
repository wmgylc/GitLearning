//一个路径下只能存在一个package
package main


/*
	type interface_name interface {
		method_name1 (return type)
	}
*/

/*
	type struct_name struct {
	    variables
	}
 */

import "fmt"

type Phone interface {
	call()
}

//go没有类的概念，转而使用struct，所以struct就类似go里的类
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

	balance := [...]int{1, 2, 3, 4, 5, 6}

	for i := 0; i < len(balance); i++ {
		fmt.Println(balance[i])
	}
}

