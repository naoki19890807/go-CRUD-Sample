package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	//init関数を呼ぶため
	fmt.Println(models.Db)

	//ユーザー作成処理
	/*
		u := &models.User{}
		u.Name = "test3"
		u.Email = "test3@test.com"
		u.PassWord = "test3test3"
		u.CreateUser()
	*/
	//ユーザー取得
	/*
		u, _ := models.GetUser(4)
		fmt.Println(u)
	*/
	//ユーザー更新
	/*
		u.Name = "test1"
		u.Email = "test1@test.com"
		fmt.Println(u)
		u.UpdateUser()
	*/

	//ユーザー削除
	//u.DeleteUser()

	//Todo作成
	/*
		u, _ := models.GetUser(5)
		u.CreateTodo("Third Todo")
	*/
	//Todo取得(単一)
	/*
		t, _ := models.GetTodo(1)
		fmt.Println(t)
	*/
	//Todo取得(複数)
	/*
		ts, _ := models.GetTodos()
		for _, v := range ts {
			fmt.Println(v)
		}
	*/
	//Todo取得(特定)
	/*
		u, _ := models.GetUser(2)
		u1, _ := u.GetTodosByUser()
		for _, v := range u1 {
			fmt.Println(v)
		}
	*/

	//Todo更新
	/*
		t, _ := models.GetTodo(1)
		t.Content = "Update Todo"
		t.UpdateTodo()
	*/

	//Todo削除
	t, _ := models.GetTodo(3)
	t.DeleteTodo()

}
