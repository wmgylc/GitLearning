package main

import "fmt"

func main() {
	//a是int类型
	var a int = 10
	//p是int类指针
	var p *int
	//p指向a的地址，&是取地址符号
	p = &a

	j := a
	fmt.Println(&j, "地址不一样，可见是拷贝了值给新的地址")

	//&可以获得变量存储的地址值，*可以根据地址值取出对应的变量，所以*一定是针对的地址值

	fmt.Println("a的值是:", a)
	fmt.Println("p的值是:", p)
	fmt.Println("a的地址值是:", &a)
	fmt.Println("p的地址值是:", &p)
	//fmt.Println(*a) 不行，只有对指针才能进行*，*即指针取值
	fmt.Println("对指针p进行取值的值是: ", *p)

	//取出变量a的地址，再利用*取出a的地址对应的值，也就是a的值
	fmt.Println(*&a)

	//地址值不是字符串，而是指针类型，也就是指针变量，所以只有指针变量才能赋值地址

	b, c := 1, 2
	fmt.Println(b, c, &b, &c)
	swap(&b, &c)
	fmt.Println(b, c, &b, &c)
	fmt.Println("值交换了，地址没交换")

	fmt.Println(b, c, &b, &c)
	swap2(&b, &c)
	fmt.Println(b, c, &b, &c)

	swap3(a, b)
	fmt.Println(b, c, &b, &c)
}


func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap2(a, b *int) {
	a, b = b, a
}

func swap3(a, b int) {
	a, b = b, a
}


