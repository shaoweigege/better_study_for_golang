/**
常用运算符
*/

package main

import "fmt"

func main() {
	suanshu()
	/*
	算数运算符的注意事项
		对于除号 "/"，它的整数除和小数除是有区别的:整数之间做除法时，只保留整数部分而舍弃小数部分。
		对一个数取模时，可以等价 a%b=a-a/b*b ， 这样我们可以看到 取模的一个本质运算。
		Golang的自增自减只能独立使用
		Golang 的++ 和 -- 只能写在变量的后面，不能写在变量的前面，即:只有 a++ a-- 没有 ++a --a
		Golang 的设计者去掉c/java中的自增自减的容易混淆的写法，让Golang更加简洁，统一。(强制性的)
	 */

	guanxi()
	/*
	关系运算符
		关系运算符的结果都是 bool 型，也就是要么是 true，要么是 false
		关系表达式经常用在if结构的条件中或循环结构的条件中
		比较运算符"=="不能误写成 "="
	*/

	luoji()
	/*
	逻辑运算符
		&&也叫短路与:如果第一个条件为false，则第二个条件不会判断，最终结果为false
		||也叫短路或:如果第一个条件为true，则第二个条件不会判断，最终结果为true
	 */
}

func suanshu(){
	//运算的数都是整数，最后值也是整数
	fmt.Println(5 / 4)
	var n1 float32 = 5 / 4
	fmt.Println(n1)
	//需要保留小数
	var n2 float32 = 5.0 / 4
	fmt.Println(n2)
	var n3 int = 10 % 3
	fmt.Println(n3)
	//返回值的正负取决于第一个操作数的正负!
	var n4 int = 10 % -3
	fmt.Println(n4)
	var n5 int = 5
	n5++
	fmt.Println(n5)
}

func guanxi()  {
	var flag bool
	flag = 5 > 4
	fmt.Println(flag);
}

func luoji(){
	var age int = 40
	if (age < 50 || age > 30) {
		fmt.Println("so young")
	}
}

/*
 总结
	1.

*/
