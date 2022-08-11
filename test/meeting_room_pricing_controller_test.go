package test

import (
	"database/sql"
	"encoding/json"
	"github.com/agisnur24/booking_hotel_system.git/app/routers"
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/middleware"
	"github.com/agisnur24/booking_hotel_system.git/repository"
	"github.com/agisnur24/booking_hotel_system.git/service"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestMeetingRoomPricingDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_management_system")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupMeetingRoomPricingRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	meetingRoomPricingRepository := repository.NewMeetingRoomPricingRepository()
	meetingRoomPricingService := service.NewMeetingRoomPricingService(meetingRoomPricingRepository, db, validate)
	meetingRoomPricingController := controller.NewMeetingRoomPricingController(meetingRoomPricingService)
	router := routers.NewMeetingRoomPricingRouter(meetingRoomPricingController)

	return middleware.NewAuthMiddleware(router)
}

func truncateMeetingRoomPricing(db *sql.DB) {
	db.Exec("TRUNCATE meeting_room_pricings")
}

func TestCreateMeetingRoomPricingSuccess(t *testing.T) {
	db := setupTestMeetingRoomPricingDB()
	truncateMeetingRoomPricing(db)
	router := setupMeetingRoomPricingRouter(db)

	requestBody := strings.NewReader(`{"meeting_room_id" : 3, "price" : "300000", "price_type" : "hourly"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/meeting_room_pricings", requestBody)
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
	assert.Equal(t, 3, int(responseBody["data"].(map[string]interface{})["meeting_room_id"].(float64)))
	assert.Equal(t, "300000", responseBody["data"].(map[string]interface{})["price"])
	assert.Equal(t, "hourly", responseBody["data"].(map[string]interface{})["price_type"])
}

func TestCreateMeetingRoomPricingFailed(t *testing.T) {
	db := setupTestMeetingRoomPricingDB()
	truncateMeetingRoomPricing(db)
	router := setupMeetingRoomPricingRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/meeting_room_pricings", requestBody)
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

/*func TestUpdateMeetingRoomPricingSuccess(t *testing.T) {
	db := setupTestMeetingRoomPricingDB()
	truncateMeetingRoomPricing(db)

	tx, _ := db.Begin()
	meetingRoomPricingRepository := repository.NewMeetingRoomPricingRepository()
	meetingRoomPricing := meetingRoomPricingRepository.Create(context.Background(), tx, domain.MeetingRoomPricing{
		MeetingRoomId: 3,
		Price:         "300000",
		PriceType:     "hourly",
	})
	tx.Commit()

	router := setupMeetingRoomPricingRouter(db)

	requestBody := strings.NewReader(`{"meeting_room_id" : 3, "price" : "300000", "price_type" : "hourly"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/employees/"+strconv.Itoa(meetingRoomPricing.Id), requestBody)
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
	assert.Equal(t, meetingRoomPricing.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, 3, int(responseBody["data"].(map[string]interface{})["meeting_room_id"].(float64)))
	assert.Equal(t, "300000", responseBody["data"].(map[string]interface{})["price"])
	assert.Equal(t, "hourly", responseBody["data"].(map[string]interface{})["price_type"])
	assert.Equal(t, "eko", responseBody["data"].(map[string]interface{})["meeting_room_name"])
	assert.Equal(t, "5", responseBody["data"].(map[string]interface{})["meeting_room_capacity"])
}*/
