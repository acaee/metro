package main

import (
	"github.com/acaee/metro/service/parse"
	"github.com/acaee/metro/service/source"
	"github.com/acaee/metro/service/storage"
	_ "github.com/acaee/metro/service/storage"
	"net/http"
)

func main() {

	var city source.City
	source.Run(&city, source.Request{
		Method: http.MethodGet,
		URL:    "https://map.baidu.com/?qt=subwayscity",
	})
	cities := parse.ParseCity(&city)
	storage.Update(cities)
}

