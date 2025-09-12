package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
)

// Employee 结构体
type Employee struct {
    ID         int     `db:"id"`
    Name       string  `db:"name"`
    Department string  `db:"department"`
    Salary     float64 `db:"salary"`
}

// Book 结构体
type Book struct {
    ID     int     `db:"id"`
    Title  string  `db:"title"`
    Author string  `db:"author"`
    Price  float64 `db:"price"`
}

// 获取技术部员工信息
func getTechDepartmentEmployees(db *sqlx.DB) ([]Employee, error) {
    var employees []Employee
    query := "SELECT * FROM employees WHERE department = ?"
    err := db.Select(&employees, query, "技术部")
    if err != nil {
        return nil, err
    }
    return employees, nil
}

// 获取工资最高的员工信息
func getHighestSalaryEmployee(db *sqlx.DB) (Employee, error) {
    var employee Employee
    query := "SELECT * FROM employees ORDER BY salary DESC LIMIT 1"
    err := db.Get(&employee, query)
    if err != nil {
        return Employee{}, err
    }
    return employee, nil
}


// 获取价格大于 50 元的书籍
func getBooksAbovePrice(db *sqlx.DB, price float64) ([]Book, error) {
    var books []Book
    query := "SELECT * FROM books WHERE price > ?"
    err := db.Select(&books, query, price)
    if err != nil {
        return nil, err
    }
    return books, nil
}

func main() {
    // 数据库连接
    db, err := sqlx.Connect("mysql", "root:12345abc@tcp(127.0.0.1:3306)/sqlx")
    if err != nil {
        log.Fatalln(err)
    }

    // 查询技术部员工
    techEmployees, err := getTechDepartmentEmployees(db)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("技术部员工:", techEmployees)

    // 查询工资最高的员工
    highestSalaryEmployee, err := getHighestSalaryEmployee(db)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("工资最高的员工:", highestSalaryEmployee)

	 // 查询价格大于 50 元的书籍
	 books, err := getBooksAbovePrice(db, 50)
	 if err != nil {
		 log.Fatalln(err)
	 }
	 fmt.Println("价格大于 50 元的书籍:", books)
}
