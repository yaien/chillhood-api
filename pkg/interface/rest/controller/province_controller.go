package controller

import (
	"github.com/yaien/clothes-store-api/pkg/entity"
	response "github.com/yaien/clothes-store-api/pkg/interface/rest/helpers/response"
	"github.com/yaien/clothes-store-api/pkg/service"
	"net/http"
	"strconv"
)

type ProvinceController struct {
	Provinces service.ProvinceService
}

func (pc *ProvinceController) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	skip, _ := strconv.Atoi(query.Get("skip"))
	provinces, err := pc.Provinces.Search(r.Context(), entity.SearchProvinceOptions{
		Name:  query.Get("name"),
		Limit: int64(limit),
		Skip:  int64(skip),
	})
	if err != nil {
		response.Error(w, err, http.StatusInternalServerError)
		return
	}
	if provinces == nil {
		provinces = make([]*entity.Province, 0)
	}
	response.Send(w, provinces)
}
