package source

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	Method string
	URL    string
	Body   io.Reader
}

type Data interface {
	Parse(bye []byte)
}

type Result interface {

}

func Run(data Data, request Request) {
	req, _ := http.NewRequest(request.Method, request.URL, request.Body)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	byt, err := io.ReadAll(res.Body)
	data.Parse(byt)


}

type City struct {
	Result struct {
		Type          string `json:"type"`
		Error         string `json:"error"`
		SubwayVersion string `json:"subwayVersion"`
	} `json:"result"`
	SubwaysCity struct {
		Cities []struct {
			CnName string `json:"cn_name"`
			Cename string `json:"cename"`
			Code   int    `json:"code"`
			Cpre   string `json:"cpre"`
			CxfDis int    `json:"cxfDis,omitempty"`
		} `json:"cities"`
	} `json:"subways_city"`
}

func (city *City) Parse(byt []byte) {
	json.Unmarshal(byt, city)
}