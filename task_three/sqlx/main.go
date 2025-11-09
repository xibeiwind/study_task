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

// 题目2：实现类型安全映射
// 假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// 要求 ：
// 定义一个 Book 结构体，包含与 books 表对应的字段。
type Book struct {
	Id     int     `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func AutoMigrateAndInitBook(db *sqlx.DB) {
	db.MustExec(`CREATE TABLE IF NOT EXISTS books (
		id INT AUTO_INCREMENT PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		price DECIMAL(10, 2) NOT NULL
	)`)

	db.MustExec(`INSERT INTO books (title, author, price) VALUES
		("《Go 语言基础》", "小王子", 25.00),
		("《Go 语言进阶》", "小王子", 40.00),
		("《Go 语言实战》", "小王子", 50.00),
		("《Go 语言微服务》", "小王子", 70.00),
		("《Go 语言分布式》", "小王子", 100.00)`)

}

// 编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
func QueryBooksByPrice(db *sqlx.DB) ([]Book, error) {
	var books []Book
	err := db.Select(&books, "SELECT * FROM books WHERE price > ?", 50.00)
	return books, err
}

func main() {
	{
		db, err := ConnectSqlistDB()
		if err != nil {
			panic(err)
		}

		AutoMigrateAndInitEmployee(db)

		employees, _ := QueryEmployeesByDepartment(db)
		println("所有技术部员工信息：")
		for _, employee := range employees {
			println(employee.Name, employee.Department, employee.Salary)
		}
		employee, _ := QueryHighestSalaryEmployee(db)
		println("工资最高的员工信息：")
		println(employee.Name, employee.Department, employee.Salary)
	}
	{
		db, err := ConnectSqlistDB()
		if err != nil {
			panic(err)
		}
		AutoMigrateAndInitBook(db)
		books, _ := QueryBooksByPrice(db)
		println("价格大于 50 元的书籍信息：")
		for _, book := range books {
			println(book.Title, book.Author, book.Price)
		}
	}
}
