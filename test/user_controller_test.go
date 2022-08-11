package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/agisnur24/booking_hotel_system.git/app/routers"
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/middleware"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"github.com/agisnur24/booking_hotel_system.git/repository"
	"github.com/agisnur24/booking_hotel_system.git/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupTestUserDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_management_system")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupUserRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := routers.NewUserRouter(userController)

	return middleware.NewAuthMiddleware(router)
}

func truncateUser(db *sql.DB) {
	db.Exec("TRUNCATE users")
}

func TestCreateUserSuccess(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)
	router := setupUserRouter(db)

	requestBody := strings.NewReader(`{"name" : "Renaldi", "email" : "renaldi099@rocketmail.com", "password" : "yaha5432j", "role_id" : 2}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", requestBody)
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
	assert.Equal(t, "Renaldi", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "renaldi099@rocketmail.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "yaha5432j", responseBody["data"].(map[string]interface{})["password"])
	assert.Equal(t, 2, int(responseBody["data"].(map[string]interface{})["role_id"].(float64)))
}

func TestCreateUserFail(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)
	router := setupUserRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/users", requestBody)
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

func TestUpdateUserSuccess(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Create(context.Background(), tx, domain.User{
		Name:     "ironman",
		Email:    "rorojok24@rocketmail.com",
		Password: "lalalal123",
		RoleId:   2,
	})
	tx.Commit()

	router := setupUserRouter(db)

	requestBody := strings.NewReader(`{"name" : "ironwoman", "email" : "ironwoman69@gmail.com", "password" : "lawalwal999"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), requestBody)
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
	assert.Equal(t, user.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "ironwoman", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "ironwoman69@gmail.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "lawalwal999", responseBody["data"].(map[string]interface{})["password"])
	assert.Equal(t, 2, int(responseBody["data"].(map[string]interface{})["role_id"].(float64)))
	assert.Equal(t, "employee", responseBody["data"].(map[string]interface{})["role_name"])
}

func TestUpdateUserFail(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Create(context.Background(), tx, domain.User{
		Name: "Bolang",
	})
	tx.Commit()

	router := setupUserRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), requestBody)
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

func TestGetUserSuccess(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Create(context.Background(), tx, domain.User{
		Name:     "Rama",
		Email:    "ramram69@yehee.com",
		Password: "aplwpalwp",
		RoleId:   2,
		RoleName: "employee",
	})
	tx.Commit()

	router := setupUserRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), nil)
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
	assert.Equal(t, user.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, user.Name, responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, user.Email, responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, user.Password, responseBody["data"].(map[string]interface{})["password"])
	assert.Equal(t, user.RoleId, int(responseBody["data"].(map[string]interface{})["role_id"].(float64)))
	assert.Equal(t, user.RoleName, responseBody["data"].(map[string]interface{})["role_name"])
}

func TestGetUserFail(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)
	router := setupUserRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users/404", nil)
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

func TestDeleteUserSuccess(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user := userRepository.Create(context.Background(), tx, domain.User{
		Name:     "Fadil",
		Email:    "fadiljon@gmail.com",
		Password: "mcnjdsk223",
		RoleId:   2,
	})
	tx.Commit()

	router := setupUserRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/users/"+strconv.Itoa(user.Id), nil)
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

func TestDeleteUserFail(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)
	router := setupUserRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/users/404", nil)
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

func TestListUsersSuccess(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)

	tx, _ := db.Begin()
	userRepository := repository.NewUserRepository()
	user1 := userRepository.Create(context.Background(), tx, domain.User{
		Name:     "Andika",
		Email:    "emailku33@yuhoo.com",
		Password: "aokwoakwoa",
		RoleId:   2,
	})
	user2 := userRepository.Create(context.Background(), tx, domain.User{
		Name:     "Pradipta",
		Email:    "pradipta222@weerte.com",
		Password: "oakwoakwoka",
		RoleId:   2,
	})
	tx.Commit()

	router := setupUserRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users", nil)
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

	fmt.Println(responseBody)

	var users = responseBody["data"].([]interface{})

	userResponse1 := users[0].(map[string]interface{})
	userResponse2 := users[1].(map[string]interface{})

	assert.Equal(t, user1.Id, int(userResponse1["id"].(float64)))
	assert.Equal(t, user1.Name, userResponse1["name"])
	assert.Equal(t, user1.Email, userResponse1["email"])
	assert.Equal(t, user1.Password, userResponse1["password"])
	assert.Equal(t, user1.RoleId, int(userResponse1["role_id"].(float64)))

	assert.Equal(t, user2.Id, int(userResponse2["id"].(float64)))
	assert.Equal(t, user2.Name, userResponse2["name"])
	assert.Equal(t, user2.Email, userResponse2["email"])
	assert.Equal(t, user2.Password, userResponse2["password"])
	assert.Equal(t, user2.RoleId, int(userResponse1["role_id"].(float64)))
}

func TestUserUnauthorized(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)
	router := setupUserRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/users", nil)
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
