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

func setupTestMeetingRoomDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_management_system")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupMeetingRoomRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	meetingRoomRepository := repository.NewMeetingRoomRepository()
	meetingRoomService := service.NewMeetingRoomService(meetingRoomRepository, db, validate)
	meetingRoomController := controller.NewMeetingRoomController(meetingRoomService)
	router := routers.NewMeetingRoomRouter(meetingRoomController)

	return middleware.NewAuthMiddleware(router)
}

func truncateMeetingRoom(db *sql.DB) {
	db.Exec("TRUNCATE meeting_rooms")
}

func TestCreateMeetingRoomSuccess(t *testing.T) {
	db := setupTestMeetingRoomDB()
	truncateMeetingRoom(db)
	router := setupMeetingRoomRouter(db)

	requestBody := strings.NewReader(`{"floor_id" : 1, "name" : "eko", "capacity" : "5", "facility_id" : 5}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/meeting_rooms", requestBody)
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
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["floor_id"].(float64)))
	assert.Equal(t, "eko", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "5", responseBody["data"].(map[string]interface{})["capacity"])
	assert.Equal(t, 5, int(responseBody["data"].(map[string]interface{})["facility_id"].(float64)))
}

func TestCreateMeetingRoomFailed(t *testing.T) {
	db := setupTestMeetingRoomDB()
	truncateMeetingRoom(db)
	router := setupMeetingRoomRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/meeting_rooms", requestBody)
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

func TestUpdateMeetingRoomSuccess(t *testing.T) {
	db := setupTestMeetingRoomDB()
	truncateMeetingRoom(db)

	tx, _ := db.Begin()
	meetingRoomRepository := repository.NewMeetingRoomRepository()
	meetingRoom := meetingRoomRepository.Create(context.Background(), tx, domain.MeetingRoom{
		FloorId:    1,
		Name:       "eko",
		Capacity:   "5",
		FacilityId: 5,
	})
	tx.Commit()

	router := setupMeetingRoomRouter(db)

	requestBody := strings.NewReader(`{"floor_id" : 1, "name" : "eko", "capacity" : "5", "facility_id" : 5}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/meeting_rooms/"+strconv.Itoa(meetingRoom.Id), requestBody)
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
	assert.Equal(t, meetingRoom.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["floor_id"].(float64)))
	assert.Equal(t, "eko", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "5", responseBody["data"].(map[string]interface{})["capacity"])
	assert.Equal(t, 5, int(responseBody["data"].(map[string]interface{})["facility_id"].(float64)))
	assert.Equal(t, "mawar", responseBody["data"].(map[string]interface{})["floor_name"])
	assert.Equal(t, "lumayan", responseBody["data"].(map[string]interface{})["floor_description"])
	assert.Equal(t, "swimming pool", responseBody["data"].(map[string]interface{})["facility_name"])
	assert.Equal(t, "airnya bau", responseBody["data"].(map[string]interface{})["facility_description"])

}

/*func TestUpdateMeetingRoomFailed(t *testing.T) {
	db := setupTestMeetingRoomDB()
	truncateMeetingRoom(db)

	tx, _ := db.Begin()
	meetingRoomRepository := repository.NewMeetingRoomRepository()
	meetingRoom := meetingRoomRepository.Create(context.Background(), tx, domain.MeetingRoom{
		Name: "eko",
	})
	tx.Commit()

	router := setupMeetingRoomRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/meeting_rooms/"+strconv.Itoa(meetingRoom.Id), requestBody)
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
}*/

/*func TestGetMeetingRoomSuccess(t *testing.T) {
	db := setupTestMeetingRoomDB()
	truncateMeetingRoom(db)

	tx, _ := db.Begin()
	meetingRoomRepository := repository.NewMeetingRoomRepository()
	meetingRoom := meetingRoomRepository.Create(context.Background(), tx, domain.MeetingRoom{
		FloorId:    1,
		Name:       "eko",
		Capacity:   "5",
		FacilityId: 5,
	})
	tx.Commit()

	router := setupMeetingRoomRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/meeting_rooms/"+strconv.Itoa(meetingRoom.Id), nil)
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
	assert.Equal(t, meetingRoom.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["floor_id"].(float64)))
	assert.Equal(t, "eko", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "5", responseBody["data"].(map[string]interface{})["capacity"])
	assert.Equal(t, 5, int(responseBody["data"].(map[string]interface{})["facility_id"].(float64)))

}*/

/*func TestGetMeetingRoomFailed(t *testing.T) {
	db := setupTestMeetingRoomDB()
	truncateMeetingRoom(db)
	router := setupMeetingRoomRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/meeting_rooms/404", nil)
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
}*/

/*func TestDeleteMeetingRoomSuccess(t *testing.T) {
	db := setupTestMeetingRoomDB()
	truncateMeetingRoom(db)

	tx, _ := db.Begin()
	meetingRoomRepository := repository.NewMeetingRoomRepository()
	meetingRoom := meetingRoomRepository.Create(context.Background(), tx, domain.MeetingRoom{
		FloorId:    1,
		Name:       "eko",
		Capacity:   "5",
		FacilityId: 5,
	})
	tx.Commit()

	router := setupMeetingRoomRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/meeting_rooms/"+strconv.Itoa(meetingRoom.Id), nil)
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
}*/

/*func TestDeleteMeetingRoomFail(t *testing.T) {
	db := setupTestMeetingRoomDB()
	truncateMeetingRoom(db)
	router := setupMeetingRoomRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/meeting_rooms/404", nil)
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
}*/

/*func TestListMeetingRoomSuccess(t *testing.T) {
	db := setupTestMeetingRoomDB()
	truncateMeetingRoom(db)

	tx, _ := db.Begin()
	meetingRoomRepository := repository.NewMeetingRoomRepository()
	meetingRoom1 := meetingRoomRepository.Create(context.Background(), tx, domain.MeetingRoom{
		FloorId:             1,
		Name:                "eko",
		Capacity:            "5",
		FacilityId:          5,
		FloorName:           "mawar",
		FloorDescription:    "lumayan",
		FacilityName:        "swimming pool",
		FacilityDescription: "airnya bau",
	})
	meetingRoom2 := meetingRoomRepository.Create(context.Background(), tx, domain.MeetingRoom{
		FloorId:             2,
		Name:                "edi",
		Capacity:            "6",
		FacilityId:          5,
		FloorName:           "melati",
		FloorDescription:    "bagus",
		FacilityName:        "gym",
		FacilityDescription: "berat",
	})
	tx.Commit()

	router := setupMeetingRoomRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/meeting_rooms", nil)
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

	var meeting_rooms = responseBody["data"].([]interface{})

	meetingRoomResponse1 := meeting_rooms[0].(map[string]interface{})
	meetingRoomResponse2 := meeting_rooms[1].(map[string]interface{})

	assert.Equal(t, meetingRoom1.Id, int(meetingRoomResponse1["id"].(float64)))
	assert.Equal(t, meetingRoom1.FloorId, int(meetingRoomResponse1["floor_id"].(float64)))
	assert.Equal(t, meetingRoom1.Name, meetingRoomResponse1["name"])
	assert.Equal(t, meetingRoom1.Capacity, meetingRoomResponse1["capacity"])
	assert.Equal(t, meetingRoom1.FacilityId, int(meetingRoomResponse1["facility_id"].(float64)))
	assert.Equal(t, meetingRoom1.FloorName, meetingRoomResponse1["floor_name"])
	assert.Equal(t, meetingRoom1.FloorDescription, meetingRoomResponse1["floor_description"])
	assert.Equal(t, meetingRoom1.FacilityName, meetingRoomResponse1["facility_name"])
	assert.Equal(t, meetingRoom1.FacilityDescription, meetingRoomResponse1["facility_description"])

	assert.Equal(t, meetingRoom2.Id, int(meetingRoomResponse2["id"].(float64)))
	assert.Equal(t, meetingRoom2.FloorId, int(meetingRoomResponse2["floor_id"].(float64)))
	assert.Equal(t, meetingRoom2.Name, meetingRoomResponse2["name"])
	assert.Equal(t, meetingRoom2.Capacity, meetingRoomResponse2["capacity"])
	assert.Equal(t, meetingRoom2.FacilityId, int(meetingRoomResponse2["facility_id"].(float64)))
	assert.Equal(t, meetingRoom2.FloorName, meetingRoomResponse2["floor_name"])
	assert.Equal(t, meetingRoom2.FloorDescription, meetingRoomResponse2["floor_description"])
	assert.Equal(t, meetingRoom2.FacilityName, meetingRoomResponse2["facility_name"])
	assert.Equal(t, meetingRoom2.FacilityDescription, meetingRoomResponse2["facility_description"])
}*/

/*func TestMeetingRoomUnauthorized(t *testing.T) {
	db := setupTestMeetingRoomDB()
	truncateMeetingRoom(db)
	router := setupMeetingRoomRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/meeting_rooms", nil)
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
}*/
