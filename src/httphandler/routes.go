package httphandler

import (
	"net/http"
	"regexp"

	"delivery-service/controllers"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	// system
	newRoute(http.MethodGet, "/health", controllers.HealthCheck),
	// notifications
	newRoute(http.MethodGet, "/v1/couriers/list", controllers.GetCouriersListV1),
	newRoute(http.MethodGet, "/v1/deliveres/list", controllers.GetDeliveresListV1),
	newRoute(http.MethodPost, "/v1/couriers/book", controllers.CourierBookV1),
}
