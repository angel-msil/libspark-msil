package database


import (
	"log"
	"context"
	"strconv"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/angel-msil/libspark-msil/configs"
	"github.com/angel-msil/libspark-msil/constants"
)

var dbInstance *pgxpool.Pool


func InitPsqlDBProvider() (error){
	
	DataConfig:= configs.Get((constants.ApplicationConfig))
	fmt.Println("after comfig get : ",DataConfig.GetString(constants.Port))
	port,_ := strconv.Atoi(DataConfig.GetString(constants.Port))
	dsn:=fmt.Sprintf(" user =%s password=%s host=%s port=%d database=%s ",
	DataConfig.GetString(constants.User) ,DataConfig.GetString(constants.Password), DataConfig.GetString(constants.Host),port,DataConfig.GetString(constants.Database))

	var err error
	dbInstance, err = pgxpool.Connect(context.Background(),dsn)
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
		return err
	}

	return nil

}

func GetDBConnection() (*pgxpool.Pool,error) {
	if dbInstance == nil{
		err := InitPsqlDBProvider()
		log.Println("DB connection null")
		if(err != nil){
			log.Println("Failed to open a DB connection: ", err)
			return nil, err
		}
	}
	return dbInstance,nil
}

