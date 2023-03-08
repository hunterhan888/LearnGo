package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var tom Employee

func main() {
	//字符串
	s := "hello, world"
	//字符串长度
	fmt.Println(len(s)) // "12"
	//字符串索引操作
	fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
	//截取字符串
	fmt.Println(s[0:5]) // "hello"
	fmt.Println(s[:5])  // "hello"
	fmt.Println(s[7:])  // "world"
	fmt.Println(s[:])   // "hello, world"
	//连接字符串
	fmt.Println("goodbye" + s[5:]) // "goodbye, world"
	//字符串处理常用bytes、strings、strconv和unicode包

	//字符串可以转byte切片
	b := []byte(s)
	fmt.Println(b)

	// 数组(数组的长度固定,Go语言中很少直接用到数组,一般用slice)
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // 0
	fmt.Println(a[len(a)-1]) // 0

	//初始化数组
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q[2]) // "3"
	fmt.Println(r[2]) // "0"

	//也可以这样,不写明长度
	p := [...]int{1, 2, 3}
	fmt.Printf("%T\n", p) // "[3]int"
	//修改数组元素的值
	p[2] = 4
	fmt.Println(p) //[1 2 4]
	//数组反转
	a1 := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Printf("a1 type=%T\n", a1)    // [6]int
	fmt.Printf("a1 type=%T\n", a1[:]) // "[]int" 转成slice了
	reverse(a1[:])                    //这样传入的是slice
	fmt.Printf("a1=%x\n", a1)         // "[5 4 3 2 1 0]"
	//如果这里是传slice的话
	a2 := []int{0, 1, 2, 3, 4, 5}
	reverse(a2)
	fmt.Printf("a2=%x\n", a2) // "[5 4 3 2 1 0]"

	//slice 切片
	//slice底层引用数组对象
	//slice由三部分构成:指针,长度,容量
	var sl []int = []int{1, 2, 3}
	fmt.Println(sl)
	//slice容量
	fmt.Println(cap(sl)) // 3
	//slice的切片操作s[i:j]，其中0 ≤ i≤ j≤ cap(s)
	fmt.Println(sl[1:2]) //[2]
	fmt.Println(sl[1:])  //[2 3] == sl[1:cap(sl)]
	fmt.Println(sl[:])   //[1 2 3] == sl[0:cap(sl)]
	//使用len(sl) == 0来判断slice是否为空，而不应该用sl == nil来判断
	//slice追加元素
	sl2 := append(sl, 4)
	fmt.Println(cap(sl))  // 3
	fmt.Println(cap(sl2)) // 6 容量扩充了一倍
	sl2 = append(sl2, 5, 6)
	fmt.Println(cap(sl2)) // 6 容量够,不需要扩充
	fmt.Println(sl2)      //[1 2 3 4 5 6]

	//模拟栈操作
	stack := []int{1, 2, 3, 4, 5}
	stack = append(stack, 6)        // push操作
	top := stack[len(stack)-1]      // 栈顶元素
	fmt.Printf("top=%x\n", top)     //top=6
	stack = stack[:len(stack)-1]    // pop操作
	fmt.Printf("stack=%x\n", stack) //stack=[1 2 3 4 5]
	//删除slice中的元素,并保持原有的顺序
	stack = remove(stack, 1)        //删除索引为1的元素
	fmt.Printf("stack=%x\n", stack) //stack=[1 3 4 5]

	//map
	//使用make初始化一个map
	ages := make(map[string]int)
	fmt.Println(ages) //map[]
	//或者直接带值
	ages2 := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages2) //map[alice:31 charlie:34]
	//等价于
	ages3 := make(map[string]int)
	ages3["alice"] = 31
	ages3["charlie"] = 34
	fmt.Println(ages3) //map[alice:31 charlie:34]

	fmt.Println(ages3["alice"]) // "31"
	//删除元素
	delete(ages3, "alice")
	fmt.Println(ages3) //map[charlie:34]
	//访问不存在的元素
	fmt.Println(ages3["bbb"]) //0

	_, ok := ages3["bbb"] //返回的第二个值的布尔值,可以判断元素是否存在
	if !ok {
		fmt.Println("ages中不存在bbb")
	}

	//所以也可以这样创建
	ages4 := map[string]int{}
	fmt.Println(ages4) //map[]

	//遍历map
	for name, age := range ages3 {
		fmt.Printf("%s\t%d\n", name, age) //charlie 34
	}

	//结构体
	tom.Name = "Tome"
	tom.Salary = 5000
	position := &tom.Position
	*position = "Senior " + *position
	fmt.Println(tom) //{0 Tome  0001-01-01 00:00:00 +0000 UTC Senior  5000 0}
	//计算奖金,结构体以指针形式传入
	fmt.Println(bonus(&tom, 80))                    //4000
	fmt.Println(bonus(&Employee{Salary: 5000}, 80)) //4000
	awardAnnualRaise(&tom)
	fmt.Println(tom) //{0 Tome  0001-01-01 00:00:00 +0000 UTC Senior  5250 0}
}

// 反转数组,函数的参数一般传slice,由于slice中包含指针,所以这里会直接修改传入的slice
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 删除slice中的元素
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

// 计算奖金,较大的结构体一般以指针的方式作为参数
func bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

// 使用指针直接修改结构体中的值
func awardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
