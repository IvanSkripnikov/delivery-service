package helpers

import (
	"encoding/json"
	"net/http"

	"delivery-service/models"

	"github.com/IvanSkripnikov/go-gormdb"
	logger "github.com/IvanSkripnikov/go-logger"
)

func GetCouriersList(w http.ResponseWriter, _ *http.Request) {
	category := "/v1/couriers/list"
	var couriers []models.Courier

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Find(&couriers).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": couriers,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func GetDeliveresList(w http.ResponseWriter, r *http.Request) {
	category := "/v1/deliveres/list"
	var deliveres []models.Delivery

	db := gormdb.GetClient(models.ServiceDatabase)
	err := db.Find(&deliveres).Error
	if checkError(w, err, category) {
		return
	}

	data := ResponseData{
		"response": deliveres,
	}
	SendResponse(w, data, category, http.StatusOK)
}

func BookCourier(w http.ResponseWriter, r *http.Request) {
	category := "/v1/couriers/book"

	// получаем параметры
	var bookingParams models.BookingItem
	err := json.NewDecoder(r.Body).Decode(&bookingParams)

	if checkError(w, err, category) {
		return
	}

	// смотрим, есть ли нужное количество на складе
	var freeCouriers []models.Courier

	db := gormdb.GetClient(models.ServiceDatabase)
	err = db.Where("busy = ?", 0).Find(&freeCouriers).Error
	if checkError(w, err, category) {
		return
	}

	var result string
	if len(freeCouriers) > 0 {
		result = "success"

		// создаём новый заказ
		currentTimestamp := int(GetCurrentTimestamp())

		var newDelivery models.Delivery
		newDelivery.CourierID = freeCouriers[0].ID
		newDelivery.OrederID = bookingParams.OrderID
		newDelivery.Created = currentTimestamp
		newDelivery.Status = models.StatusNew
		err = db.Create(&newDelivery).Error
		if err != nil {
			result = "failure"
			logger.Errorf("Cant create new booking item: %v", err)
		} else {
			courier := freeCouriers[0]
			err = db.Model(&courier).Updates(models.Courier{Busy: 1, Updated: currentTimestamp}).Error
			if err != nil {
				result = "failure"
				logger.Errorf("Cant change courier status: %v", err)
			}
		}
	} else {
		result = "failure"
	}

	data := ResponseData{
		"response": result,
	}
	SendResponse(w, data, category, http.StatusOK)
}
