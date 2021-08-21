package parse

import (
	"github.com/acaee/metro/service/source"
	"github.com/acaee/metro/service/storage"
)

func ParseCity(sourceCity *source.City) []storage.City{
	var cities []storage.City
	for _, v := range sourceCity.SubwaysCity.Cities {
		var city storage.City
		city.Code = v.Code
		city.CnName = v.CnName
		city.CeName = v.Cename
		city.Cpre = v.Cpre
		cities = append(cities, city)
	}
	return cities
}
