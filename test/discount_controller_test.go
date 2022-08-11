package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/renaldid/hotel_booking_management.git/app"
	"github.com/renaldid/hotel_booking_management.git/controller"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/middleware"
	"github.com/renaldid/hotel_booking_management.git/model/domain"
	"github.com/renaldid/hotel_booking_management.git/repository"
	"github.com/renaldid/hotel_booking_management.git/service"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupTestDiscountDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_management_system")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupDiscountRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	discountRepository := repository.NewDiscountRepository()
	discountService := service.NewDiscountService(discountRepository, db, validate)
	discountController := controller.NewDiscountController(discountService)
	router := app.NewRouter(discountController)

	return middleware.NewAuthMiddleware(router)
}

func truncateDiscount(db *sql.DB) {
	db.Exec("TRUNCATE discounts")
}

func TestCreateDiscountSuccess(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)
	router := setupDiscountRouter(db)

	requestBody := strings.NewReader(`{"employee_id" : 1, "hotel_id" : 1, "meeting_room_id" : 1, "rate" : "5", "status" : "waiting", "request_date" : "2022-05-08 00:00:00"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/discounts", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["employee_id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["hotel_id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["meeting_room_id"].(float64)))
	assert.Equal(t, "5", responseBody["data"].(map[string]interface{})["rate"])
	assert.Equal(t, "waiting", responseBody["data"].(map[string]interface{})["status"])
	assert.Equal(t, "2022-05-08 00:00:00", responseBody["data"].(map[string]interface{})["request_date"])
}

func TestCreateDiscountFail(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)
	router := setupDiscountRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/discounts", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateDiscountSuccess(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)

	tx, _ := db.Begin()
	discountRepository := repository.NewDiscountRepository()
	discount := discountRepository.Create(context.Background(), tx, domain.Discount{
		EmployeeId:    1,
		HotelId:       1,
		MeetingRoomId: 1,
		Rate:          "5",
		Status:        "waiting",
		RequestDate:   "2022-05-08 00:00:00",
	})
	tx.Commit()

	router := setupDiscountRouter(db)

	requestBody := strings.NewReader(`{"employee_id" : 1, "hotel_id" : 1, "meeting_room_id" : 1, "rate" : "5", "status" : "waiting", "request_date" : "2022-05-08 00:00:00"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/discounts/"+strconv.Itoa(discount.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, discount.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["employee_id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["hotel_id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["meeting_room_id"].(float64)))
	assert.Equal(t, "5", responseBody["data"].(map[string]interface{})["rate"])
	assert.Equal(t, "waiting", responseBody["data"].(map[string]interface{})["status"])
	assert.Equal(t, "2022-05-08 00:00:00", responseBody["data"].(map[string]interface{})["request_date"])

}

func TestUpdateDiscountFail(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)

	tx, _ := db.Begin()
	discountRepository := repository.NewDiscountRepository()
	discount := discountRepository.Create(context.Background(), tx, domain.Discount{
		Status: "gagal",
	})
	tx.Commit()

	router := setupDiscountRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/discounts/"+strconv.Itoa(discount.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetDiscountSuccess(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)

	tx, _ := db.Begin()
	discountRepository := repository.NewDiscountRepository()
	discount := discountRepository.Create(context.Background(), tx, domain.Discount{
		EmployeeId:    1,
		HotelId:       1,
		MeetingRoomId: 1,
		Rate:          "5",
		Status:        "rejected",
		RequestDate:   "2022-05-08 00:00:00",
	})
	tx.Commit()

	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/discounts/"+strconv.Itoa(discount.Id), nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, discount.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["employee_id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["hotel_id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["meeting_room_id"].(float64)))
	assert.Equal(t, "5", responseBody["data"].(map[string]interface{})["rate"])
	assert.Equal(t, "rejected", responseBody["data"].(map[string]interface{})["status"])
	assert.Equal(t, "2022-05-08 00:00:00", responseBody["data"].(map[string]interface{})["request_date"])
}

func TestGetDiscountFail(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)
	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/discounts/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteDiscountSuccess(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)

	tx, _ := db.Begin()
	discountRepository := repository.NewDiscountRepository()
	discount := discountRepository.Create(context.Background(), tx, domain.Discount{
		EmployeeId:    1,
		HotelId:       1,
		MeetingRoomId: 1,
		Rate:          "5",
		Status:        "waiting",
		RequestDate:   "2022-05-08 00:00:00",
	})
	tx.Commit()

	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/discounts/"+strconv.Itoa(discount.Id), nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteDiscountFail(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)
	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/discounts/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListDiscountSuccess(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)

	tx, _ := db.Begin()
	discountRepository := repository.NewDiscountRepository()
	discount1 := discountRepository.Create(context.Background(), tx, domain.Discount{
		EmployeeId:    1,
		HotelId:       1,
		MeetingRoomId: 1,
		Rate:          "5",
		Status:        "waiting",
		RequestDate:   "2022-05-08 00:00:00",
		EmployeeName:  "kecoa",
		HotelName:     "PRODEO",
		RoomName:      "kecoa",
	})
	discount2 := discountRepository.Create(context.Background(), tx, domain.Discount{
		EmployeeId:    1,
		HotelId:       1,
		MeetingRoomId: 1,
		Rate:          "10",
		Status:        "rejected",
		RequestDate:   "2022-05-09 00:00:00",
		EmployeeName:  "kecoa",
		HotelName:     "PRODEO",
		RoomName:      "kecoa",
	})
	tx.Commit()

	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/discounts", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var discounts = responseBody["data"].([]interface{})

	discountResponse1 := discounts[0].(map[string]interface{})
	discountResponse2 := discounts[1].(map[string]interface{})

	assert.Equal(t, discount1.Id, int(discountResponse1["id"].(float64)))
	assert.Equal(t, discount1.EmployeeId, int(discountResponse1["employee_id"].(float64)))
	assert.Equal(t, discount1.HotelId, int(discountResponse1["hotel_id"].(float64)))
	assert.Equal(t, discount1.MeetingRoomId, int(discountResponse1["meeting_room_id"].(float64)))
	assert.Equal(t, discount1.Rate, discountResponse1["rate"])
	assert.Equal(t, discount1.Status, discountResponse1["status"])
	assert.Equal(t, discount1.RequestDate, discountResponse1["request_date"])
	assert.Equal(t, discount1.EmployeeName, discountResponse1["employee_name"])
	assert.Equal(t, discount1.HotelName, discountResponse1["hotel_name"])
	assert.Equal(t, discount1.RoomName, discountResponse1["room_name"])

	assert.Equal(t, discount2.Id, int(discountResponse2["id"].(float64)))
	assert.Equal(t, discount2.EmployeeId, int(discountResponse2["employee_id"].(float64)))
	assert.Equal(t, discount2.HotelId, int(discountResponse2["hotel_id"].(float64)))
	assert.Equal(t, discount2.MeetingRoomId, int(discountResponse2["meeting_room_id"].(float64)))
	assert.Equal(t, discount2.Rate, discountResponse2["rate"])
	assert.Equal(t, discount2.Status, discountResponse2["status"])
	assert.Equal(t, discount2.RequestDate, discountResponse2["request_date"])
	assert.Equal(t, discount2.EmployeeName, discountResponse2["employee_name"])
	assert.Equal(t, discount2.HotelName, discountResponse2["hotel_name"])
	assert.Equal(t, discount2.RoomName, discountResponse2["room_name"])

}

func TestDiscountUnauthorized(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)
	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/discounts", nil)
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
