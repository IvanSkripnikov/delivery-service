package controllers

import (
	"net/http"

	"delivery-service/helpers"
)

func GetCouriersListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetCouriersList(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/couriers/list")
	}
}

func GetDeliveresListV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		helpers.GetDeliveresList(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/deliveres/get")
	}
}

func CourierBookV1(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		helpers.BookCourier(w, r)
	default:
		helpers.FormatResponse(w, http.StatusMethodNotAllowed, "/v1/couriers/book")
	}
}
