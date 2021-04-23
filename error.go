package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// User 用户
type User struct {
	ID       int64
	Username string
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hongtu?charset=utf8&loc=Local")
	if err != nil {
		fmt.Printf("%#v", err)
		return
	}
	defer db.Close()

	user, err := getUserByID(0)
	// 直接判定user是否有内容，此处查询应不允许降级，如果没有则退出，并打印error
	if user.ID == 0 {
		fmt.Printf("Not found：%+v\n", err)
		return
	}
	fmt.Printf("Query User：%#v\n", user)
}

// getUserByID 根据用户ID查询单个用户信息
func getUserByID(id int64) (user *User, err error) {
	user = new(User)
	if db == nil {
		err = errors.New("DB is nil")
		return
	}

	err = db.QueryRow(`SELECT id,username from mt_user WHERE id = ?`, id).Scan(&user.ID, &user.Username)
	// 个人认为此处没有必要额外对sql.ErrNoRows进行处理，只需要在上层对user进行判定即可
	return
}
