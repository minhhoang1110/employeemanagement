package types

import (
	"encoding/json"
	"net/http"
)

type Respone struct {
	Success  bool        `json:"success"`
	HttpCode int         `json:"http_code"`
	Data     interface{} `json:"data"`
}

func (r *Respone) ResponseWithJson(writer http.ResponseWriter, status int, object interface{}, objectName string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	var res Respone
	res.Success = status == http.StatusOK
	res.HttpCode = status
	if objectName != "" {
		res.Data = map[string]interface{}{objectName: object}
	} else {
		res.Data = object
	}
	json.NewEncoder(writer).Encode(res)
}
