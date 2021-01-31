package stuct

import "fmt"

type Person struct {
	Id int
	Name string
}

//实现string方法
func (person Person) String()string{
	return fmt.Sprintf("the personName:%v", person.Name)
}

type (
	Address struct {
		City string
	}

	Ass struct {
		Demo string
		Addr Address
	}

	Bss struct {
		Id int
		Address //组合
	}

	ErrorString struct {
		s string
	}
)

func (err *ErrorString) Error()string{
	return err.s
}

func (address Address) String() string{
	return fmt.Sprintf("the city is:%v",address.City)
}

//创建工厂函数，返回*Person
func NewPerson(name string)*Person{
	return &Person{Name:name}
}

//工厂函数，返回一个接口
func NewError(text string)error{
	return &ErrorString{text}
}




