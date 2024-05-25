package main

import (
	"fmt"
	"go-todo/lesson_18/app/models"
)

// lesson_18まで移動してから実行すること(でないとconfig.iniが読み込めない)
func main() {
	// これが実行される前にconfig.iniが読み込まれる
	// fmt.Println(config.Config.Port)

	// log.Println("test")

	// 意味はないが、initを実行するためにmodelsパッケージをインポートしている
	fmt.Println(models.Db)

	// create
	// u := &models.User{
	// 	Name:     "test",
	// 	Email:    "test@test.com",
	// 	Password: "test",
	// }
	// u.CreateUser()

	// get
	u, _ := models.GetUser(1)
	fmt.Println(u)

	// update
	u.Name = "test2"
	u.Email = "test2@test2.com"
	u.UpdateUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

	// delete
	u.DeleteUser()
	u, _ = models.GetUser(1)
	fmt.Println(u)

}
