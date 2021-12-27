package controller

import (
	"github.com/yaien/clothes-store-api/pkg/entity"
	"github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/service"
	"net/http"
	"strconv"
)

type CityController struct {
	Cities service.CityService
}

func (cc *CityController) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	skip, _ := strconv.Atoi(query.Get("skip"))
	cities, err := cc.Cities.Search(r.Context(), entity.SearchCityOptions{
		Name:     query.Get("name"),
		Province: query.Get("province"),
		Limit:    int64(limit),
		Skip:     int64(skip),
	})
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	response.Send(w, cities)
}
