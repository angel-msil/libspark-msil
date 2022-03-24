package database

import (
	"log"
	"context"
	//"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgconn"
)


func GetResultsExec(ctx context.Context, sQuery string, params ...string)(pgconn.CommandTag){

	dbPool,_ := GetDBConnection()

	log.Println("LENGTH",len(params))


	var Queryparams []interface{}

	i := 0

	for _,_ = range params{
		Queryparams = append(Queryparams, params[i])
		i = i + 1
	}


	rows,err_:= dbPool.Exec(ctx,sQuery ,Queryparams...)

	if err_ != nil {
		log.Println("error while executing query: ", err_)
	}

	return rows
}

func GetResultsExecCsv(ctx context.Context, sQuery string, Queryparams []interface{})(pgconn.CommandTag){

	dbPool,_ := GetDBConnection()

	log.Println("LENGTH",len(Queryparams))

	rows,err_:= dbPool.Exec(ctx,sQuery ,Queryparams...)

	if err_ != nil {
		log.Println("error while executing query: ", err_)
	}

	return rows
}

func GetResultsQuery(ctx context.Context, sQuery string, params ...string)(pgx.Rows){

	conn,_ := GetDBConnection()

	log.Println("LENGTH",len(params))


	var Queryparams []interface{}

	i := 0

	for _,_ = range params{
		Queryparams = append(Queryparams, params[i])
		i = i + 1
	}

	rows, err := conn.Query(ctx,sQuery ,Queryparams...)

	if err != nil {
		log.Println("error while executing query: ", err)
	}
	//defer rows.Close()

	return rows
}



