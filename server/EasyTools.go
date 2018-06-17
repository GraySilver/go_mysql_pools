package server

import (_"os"
		 "log"
	"encoding/json"
)

func CheckErr(err error)  {
	if err != nil{
		log.Fatalln("Error is :", err)
		//os.Exit(-1)
	}
}

func MapToJson(m interface{}) string{
	data,err :=json.Marshal(m)
	CheckErr(err)
	return string(data)
}

func ReturnResponse(data map[string]interface{},msg interface{},status int) map[string]interface{}{
	resp := make(map[string]interface{})
	resp["data"] = data
	resp["msg"] = msg
	resp["status"] = status
	return resp
}