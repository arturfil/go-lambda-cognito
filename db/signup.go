package db

import (
	"fmt"
    _ "github.com/go-sql-driver/mysql"
	"github.com/arturfil/go_lambdas/models"
	"github.com/arturfil/go_lambdas/tools"
)

func SignUp(sign models.SignUp) error {
	fmt.Println("starting registry")

	err := DbConnect()
	if err != nil {
        fmt.Println("ERROR -> ", err.Error())
		return err
	}

	defer Db.Close()

	query := fmt.Sprintf(`
        INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('%s' , '%s', '%s')`, 
        sign.UserEmail, 
        sign.UserUUID, 
        tools.MySqlDate(),
    )
	fmt.Println(query)

	_, err = Db.Exec(query)

	if err != nil {
        fmt.Println("Error with the query")
		fmt.Println(err.Error())
		return err
	}

	return nil
}
