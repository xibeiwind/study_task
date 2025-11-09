package main

import (
	"fmt"
	"gorm.io/gorm"
)

// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
// 	要求 ：
// 		编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
// 		编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
// 		编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
// 		编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

type Student struct {
	Id    int    `gorm:column:id;primary_key;auto_increment`
	Name  string `gorm:column:name;size:255;not null;index`
	Age   int    `gorm:column:age;not null`
	Grade string `gorm:column:grade;size:50;not null; index`
}

func AutoMigrateStudent(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	var students []Student
	db.Find(&students)
	db.Delete(&students)
}

// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
func SqlInsertStudent(db *gorm.DB) {
	var student Student
	student.Name = "张三"
	student.Age = 20
	student.Grade = "三年级"
	res := db.Create(&student)
	if res.Error != nil {
		fmt.Println("插入数据失败", res.Error)
	} else {
		fmt.Println("插入数据成功")
	}
}

// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
func SqlQueryStudent(db *gorm.DB) {
	var students []Student
	db.Where("age > ?", 18).Find(&students)
	for _, student := range students {
		fmt.Println(student.Name, student.Age, student.Grade)
	}
}

// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
func SqlUpdateStudent(db *gorm.DB) {
	var student Student
	db.Where("name = ?", "张三").First(&student)
	student.Grade = "四年级"
	db.Save(&student)
}

// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
func SqlDeleteStudent(db *gorm.DB) {
	var students []Student
	db.Where("age < ?", 15).Find(&students)
	for _, student := range students {
		db.Delete(&student)
	}
}
