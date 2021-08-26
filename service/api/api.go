package api

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"
)



type Service struct {
	logger *zap.Logger
	Config
}

func (service *Service) WithLogger(log *zap.Logger)  {
	service.logger = log.With(zap.String("service","api"))
}
func (service *Service) Open() error {
	http.HandleFunc("/query", Query)
	return http.ListenAndServe(service.Config.Host + ":" + strconv.Itoa(service.Config.Post), http.DefaultServeMux)
}

func Query(w http.ResponseWriter, r *http.Request) {
	//TODO 解析这个查询语句
	w.Write([]byte("Hello, This is a query result. "))

}

func (service *Service) Close() error {

	return nil
}

