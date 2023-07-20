package models

import (
	"log"
	"time"
)

// 構造体生成
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

// ユーザー作成----------------------------
func (u *User) CreateUser() (err error) {
	cmd := `insert into users (
		uuid,
		name,
		email,
		password,
		created_at) values (?, ?, ?, ?, ?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ユーザー取得-----------------------
func GetUser(id int) (u User, err error) {
	//取得したデータを格納する構造体
	u = User{}
	cmd := `select * from users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&u.ID,
		&u.UUID,
		&u.Name,
		&u.Email,
		&u.PassWord,
		&u.CreatedAt)

	if err != nil {
		log.Fatalln(err)
	}
	return u, err
}

// ユーザー更新
func (u *User) UpdateUser() (err error) {
	cmd := `update users set name = ?, email = ? where id = ?`
	_, err = Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

// ユーザー削除
func (u *User) DeleteUser() (err error) {
	cmd := `delete from users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
