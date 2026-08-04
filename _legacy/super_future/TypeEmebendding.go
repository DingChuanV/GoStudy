package main

import (
	"fmt"
)

/*
	如果包含结构体具有嵌入字段同名的属性(即字段)或方法，则需要使用嵌入字段的类型来引用被遮蔽的字段或方法
*/
type Inner struct{
	X int
}

type Outer struct{
	Inner
	X int
}

type Employee struct{
	Name string
	Age int
}

func (e Employee) Description() string{
	return fmt.Sprintf("Name: %s, Age: %d", e.Name, e.Age)
}

type Manager struct{
	Employee  // Manager 包含一个Employee 的字段，但该字段没有分配名称。这使得Employee 成为一个嵌入字段
	Reports []Employee
}

func (m Manager) FindNewEmployee() []Employee{
	return m.Reports
}

func main(){
	o := Outer{
		Inner: Inner{
			X: 100,
		},
		X: 200,
	}
	fmt.Println(o.X)

	m := Manager{
		Employee: Employee{
			Name: "Dingchuan",
			Age: 30,
		},
		Reports: []Employee{
			Employee{
				Name: "Dingchuan",
				Age: 30,
			},
		},
	}

	fmt.Println(m.Age)
	fmt.Println(m.Description())
}

/*
	嵌入不是继承
		

*/
