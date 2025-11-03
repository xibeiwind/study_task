package main

import (
	// "time"
	"fmt"
	tasktwo "xibeiwind/task_two/task_two"
)

func main() { // 主函数
	fmt.Println("---Test add ten")
	a := 1
	tasktwo.AddTen(&a)
	fmt.Println(a)
	fmt.Println("---Test goroutine")
	tasktwo.GoroutinePrintNumber()

	// tasktwo.GoroutineTaskScheduler(
	// 	tasktwo.GoroutineTask(int[1,2,3,4,5,6,7,8,9])
	// 	)

	fmt.Println("---Test OOP")
	circle := tasktwo.Circle{Radius: 5}
	
	fmt.Println("Circle Area:", circle.Area(), "Perimeter:", circle.Perimeter())

	rectangle := tasktwo.Rectangle{Height: 5, Width: 10}
	fmt.Println("Rectangle Area:", rectangle.Area(), "Perimeter:", rectangle.Perimeter())

	employee := tasktwo.Employee{EmployeeID: 1, Person: tasktwo.Person{Age: 18, Name: "Person A"}}
	employee.PrintInfo()

	fmt.Println("---Test channel")
	tasktwo.ChannelMain()
	tasktwo.BufferChannelMain()
	fmt.Println("---Test lock")
	count := tasktwo.LockCounter()
	fmt.Println(count)
	count = tasktwo.AtomicCounter()
	fmt.Println(count)
}
