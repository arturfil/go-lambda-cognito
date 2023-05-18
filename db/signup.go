package db

import (
	"fmt"

	"github.com/arturfil/go_lambdas/models"
	"github.com/arturfil/go_lambdas/tools"
)

func SignUp(sign models.SignUp) error {
	fmt.Println("starting registry")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	query := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUE ('" + sign.UserEmail + "' , '" + sign.UserUUID + "' '" + tools.MySqlDate() + "')"
    fmt.Println(query)

    _, err = Db.Exec(query)

    if err != nil {
        fmt.Println(err.Error())
        return err
    }

	return nil
}
