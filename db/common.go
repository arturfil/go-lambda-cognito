package db

import (
	"database/sql"
	"fmt"
	"os"

    _ "github.com/go-sql-driver/mysql"
	"github.com/arturfil/go_lambdas/models"
	secretmanager "github.com/arturfil/go_lambdas/secret_manager"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
    SecretModel, err = secretmanager.GetSecret(os.Getenv("SecretName"))
    return err
}

func DbConnect() error {
    Db, err = sql.Open("mysql", ConnStr(SecretModel))
    if err != nil {
        fmt.Println(err.Error())
        return err
    }

    err = Db.Ping()
    if err != nil {
        fmt.Println(err.Error())
        return err
    }

    fmt.Println("*** Successfull connection to db ***")
    return nil
}

func ConnStr(keys models.SecretRDSJson) string {
   var dbUser, authToken, dbEndpoint, dbName string
   dbUser = keys.Username
   authToken = keys.Password
   dbEndpoint = keys.Host
   dbName = "goecommerce"
   dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
   fmt.Println(dsn)
   return dsn
} 
