package database

import (
        "github.com/aws/aws-sdk-go/aws/session"
        "github.com/aws/aws-sdk-go/service/dynamodb"
        "fmt"
)

var Conn *dynamodb.DynamoDB
func InitDatabase() {


        sess := session.Must(session.NewSessionWithOptions(session.Options{
                SharedConfigState: session.SharedConfigEnable,
        }))

        Conn = dynamodb.New(sess)

        fmt.Println("Db Connected")
}

func GetDBConnection() (*dynamodb.DynamoDB) {
    return Conn
}

