package utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x inteerface{}){
	if body, err := ioutil.ReadAll(r.Body); arr==nil{
		if err := json.Unmarshal([]byte(body), x); err!=nil{
			return
		}
	}
}