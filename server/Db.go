package server

import ("database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
	"log"
)

func ToInit(connect string,maxOpenConns int,maxIdleConns int)  {
	log.Println("Conecting...")
	db,_ = sql.Open("mysql",connect)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.Ping()
	log.Printf("Connecting suc. MaxOpenConns: %d , MaxIdleConns: %d \n",maxOpenConns,maxIdleConns)
}


func ToSelect(SQL string,db*sql.DB)(map[string]interface{}){
	rows,err := db.Query(SQL)
	if err == nil {
		columns, _ := rows.Columns()
		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		resp := make(map[string]interface{})
		var cou= 0
		for rows.Next() {
			rows.Scan(scanArgs...)
			record := make(map[string]string)
			for i, col := range values {
				if col != nil {
					record[columns[i]] = string(col.([]byte)[:])
				}
			}
			resp[strconv.Itoa(cou)] = record
			cou += 1
		}
		defer rows.Close()
		return ReturnResponse(resp,"response is suc.",200)
	}else {
		return ReturnResponse(make(map[string]interface{}),err,404)
	}
}

func ToExecute(SQL string,db*sql.DB) interface{}{
	ret,err :=db.Exec(SQL)
	if err == nil{
		CheckErr(err)
		rowsAffecteds,err := ret.RowsAffected()
		CheckErr(err)
		rowsAffectMap := make(map[string]interface{})
		rowsAffectMap["rowsAffect"] = rowsAffecteds
		return ReturnResponse(rowsAffectMap,"response is suc.",200)
	}else {
		fmt.Println(err)
		return ReturnResponse(make(map[string]interface{}),err,404)
	}


}

func ToBegin(SQL string,db*sql.DB) interface{}{
	Tx,err := db.Begin()
	CheckErr(err)
	ret,err := Tx.Exec(SQL)
	if err == nil{
		Tx.Commit()
		rowsAffecteds,err := ret.RowsAffected()
		CheckErr(err)
		rowsAffectMap := make(map[string]interface{})
		rowsAffectMap["rowsAffect"] = rowsAffecteds
		return ReturnResponse(rowsAffectMap,"response is suc.",200)
	}else {
		fmt.Println(err)
		Tx.Rollback()
		return ReturnResponse(make(map[string]interface{}),err,404)
	}

}


