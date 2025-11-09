package main

import "github.com/jmoiron/sqlx"

func ConnectSqlistDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=foo dbname=bar sslmode=disable")
	return db, err
}

// 题目1：使用SQL扩展库进行查询
// 假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
type Employee struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func AutoMigrateAndInitEmployee(db *sqlx.DB) { 
	
	db.MustExec(`CREATE TABLE IF NOT EXISTS employees (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		department VARCHAR(255) NOT NULL,
		salary DECIMAL(10, 2) NOT NULL
	)`)

	db.MustExec(`INSERT INTO employees (name, department, salary) VALUES
		("张三", "技术部", 5000.00),
		("李四", "销售部", 3000.00),
		("王五", "财务部", 4000.00),
		("赵六", "技术部", 6000.00),
		("孙七", "销售部", 2000.00)`)
}

// 要求 ：
// 编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。
func QueryEmployeesByDepartment(db *sqlx.DB) ([]Employee, error) { 
	var employees []Employee
	err := db.Select(&employees, "SELECT * FROM employees WHERE department = ?", "技术部")
	return employees, err
}
// 编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
func QueryHighestSalaryEmployee(db *sqlx.DB) (Employee, error) { 
	var employee Employee
	err := db.Get(&employee, "SELECT * FROM employees ORDER BY salary DESC LIMIT 1")
	return employee, err
}


func main() {
	db, err := ConnectSqlistDB()
	if err != nil {
		panic(err)
	}

	AutoMigrateAndInitEmployee(db)
	
	employees,_:= QueryEmployeesByDepartment(db)
	println("所有技术部员工信息：")
	for _, employee := range employees {
		println(employee.Name, employee.Department, employee.Salary)
	} 
	employee,_:= QueryHighestSalaryEmployee(db)
	println("工资最高的员工信息：")
	println(employee.Name, employee.Department, employee.Salary)

}
