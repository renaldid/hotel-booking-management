package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/agisnur24/booking_hotel_system.git/app/routers"
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/middleware"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/repository"
	"github.com/agisnur24/booking_hotel_system.git/service"
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

func setupTestFacilityDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_management_system")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupFacilityRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	facilityRepository := repository.NewFacilityRepository()
	facilityService := service.NewFacilityService(facilityRepository, db, validate)
	facilityController := controller.NewFacilityController(facilityService)
	router := routers.NewFacilityRouter(facilityController)

	return middleware.NewAuthMiddleware(router)
}

func truncateFacility(db *sql.DB) {
	db.Exec("TRUNCATE facilities")
}

func TestCreateFacilitySuccess(t *testing.T) {
	db := setupTestFacilityDB()
	truncateFacility(db)
	router := setupFacilityRouter(db)

	requestBody := strings.NewReader(`{"name" : "swimming pool", "description" : "airnya bau"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/facilities", requestBody)
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
	assert.Equal(t, "swimming pool", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "airnya bau", responseBody["data"].(map[string]interface{})["description"])
}

func TestCreateFacilityFail(t *testing.T) {
	db := setupTestFacilityDB()
	truncateFacility(db)
	router := setupFacilityRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/facilities", requestBody)
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

func TestUpdateFacilitySuccess(t *testing.T) {
	db := setupTestFacilityDB()
	truncateFacility(db)

	tx, _ := db.Begin()
	facilityRepository := repository.NewFacilityRepository()
	facility := facilityRepository.Create(context.Background(), tx, domain.Facility{
		Name:        "swimming pool",
		Description: "airnya bau",
	})
	tx.Commit()

	router := setupFacilityRouter(db)

	requestBody := strings.NewReader(`{"name" : "swimming pool", "description" : "airnya bau"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/facilities/"+strconv.Itoa(facility.Id), requestBody)
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
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "swimming pool", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "airnya bau", responseBody["data"].(map[string]interface{})["description"])
}

func TestUpdateFacilityFailed(t *testing.T) {
	db := setupTestFacilityDB()
	truncateFacility(db)

	tx, _ := db.Begin()
	facilityRepository := repository.NewFacilityRepository()
	facility := facilityRepository.Create(context.Background(), tx, domain.Facility{
		Name:        "swimming pool",
		Description: "airnya bau",
	})
	tx.Commit()

	router := setupFacilityRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/facilities/"+strconv.Itoa(facility.Id), requestBody)
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

func TestGetFacilitySuccess(t *testing.T) {
	db := setupTestFacilityDB()
	truncateFacility(db)

	tx, _ := db.Begin()
	facilityRepository := repository.NewFacilityRepository()
	facility := facilityRepository.Create(context.Background(), tx, domain.Facility{
		Name:        "swimming pool",
		Description: "airnya bau",
	})
	tx.Commit()

	router := setupFacilityRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/facilities/"+strconv.Itoa(facility.Id), nil)
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
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "swimming pool", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "airnya bau", responseBody["data"].(map[string]interface{})["description"])

}

func TestGetFacilityFailed(t *testing.T) {
	db := setupTestFacilityDB()
	truncateFacility(db)
	router := setupFacilityRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/facilities/404", nil)
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

func TestDeleteFacilitySuccess(t *testing.T) {
	db := setupTestFacilityDB()
	truncateFacility(db)

	tx, _ := db.Begin()
	facilityRepository := repository.NewFacilityRepository()
	facility := facilityRepository.Create(context.Background(), tx, domain.Facility{
		Name:        "swimming pool",
		Description: "airnya bau",
	})
	tx.Commit()

	router := setupFacilityRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/facilities/"+strconv.Itoa(facility.Id), nil)
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

func TestDeleteFacilityFailed(t *testing.T) {
	db := setupTestFacilityDB()
	truncateFacility(db)
	router := setupFacilityRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/facilities/404", nil)
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

func TestListFacilitySuccess(t *testing.T) {
	db := setupTestFacilityDB()
	truncateFacility(db)

	tx, _ := db.Begin()
	facilityRepository := repository.NewFacilityRepository()
	facility1 := facilityRepository.Create(context.Background(), tx, domain.Facility{
		Name:        "swimming pool",
		Description: "airnya bau",
	})
	facility2 := facilityRepository.Create(context.Background(), tx, domain.Facility{
		Name:        "Apartemen",
		Description: "banyak cabe",
	})
	tx.Commit()

	router := setupFacilityRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/facilities", nil)
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

	var facilities = responseBody["data"].([]interface{})

	facilityResponse1 := facilities[0].(map[string]interface{})
	facilityResponse2 := facilities[1].(map[string]interface{})

	assert.Equal(t, facility1.Id, int(facilityResponse1["id"].(float64)))
	assert.Equal(t, facility1.Name, facilityResponse1["name"])
	assert.Equal(t, facility1.Description, facilityResponse1["description"])

	assert.Equal(t, facility2.Id, int(facilityResponse2["id"].(float64)))
	assert.Equal(t, facility2.Name, facilityResponse2["name"])
	assert.Equal(t, facility2.Description, facilityResponse2["description"])

}

func TestUnauthorized(t *testing.T) {
	db := setupTestFacilityDB()
	truncateFacility(db)
	router := setupFacilityRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/facilities", nil)
	request.Header.Add("X-API_Key", "SALAH")

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
