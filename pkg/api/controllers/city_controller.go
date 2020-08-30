package controllers

import (
	"github.com/yaien/clothes-store-api/pkg/api/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/api/models"
	"github.com/yaien/clothes-store-api/pkg/api/services"
	"net/http"
	"strconv"
)

type CityController struct {
	Cities services.CityService
}

func (cc *CityController) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	skip, _ := strconv.Atoi(query.Get("skip"))
	cities, err := cc.Cities.Search(services.SearchCityOptions{
		Name:     query.Get("name"),
		Province: query.Get("province"),
		Limit:    int64(limit),
		Skip:     int64(skip),
	})
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	if cities == nil {
		cities = make([]*models.City, 0)
	}
	response.Send(w, cities)
}
