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

func setupTestRoleDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_management_system")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRoleRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)
	router := routers.NewRoleRouter(roleController)

	return middleware.NewAuthMiddleware(router)
}

func truncateRole(db *sql.DB) {
	db.Exec("TRUNCATE roles")
}

func TestCreateRoleSuccess(t *testing.T) {
	db := setupTestRoleDB()
	truncateRole(db)
	router := setupRoleRouter(db)

	requestBody := strings.NewReader(`{"role_name" : "Owner"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/roles", requestBody)
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
	assert.Equal(t, "Owner", responseBody["data"].(map[string]interface{})["role_name"])
}

func TestCreateRoleFail(t *testing.T) {
	db := setupTestRoleDB()
	truncateRole(db)
	router := setupRoleRouter(db)

	requestBody := strings.NewReader(`{"role_name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/roles", requestBody)
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

func TestUpdateRoleSuccess(t *testing.T) {
	db := setupTestRoleDB()
	truncateRole(db)

	tx, _ := db.Begin()
	roleRepository := repository.NewRoleRepository()
	role := roleRepository.Create(context.Background(), tx, domain.Role{
		RoleName: "Employee",
	})
	tx.Commit()

	router := setupRoleRouter(db)

	requestBody := strings.NewReader(`{"role_name" : "owner"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/roles/"+strconv.Itoa(role.Id), requestBody)
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
	assert.Equal(t, role.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "owner", responseBody["data"].(map[string]interface{})["role_name"])
}

func TestUpdateRoleFail(t *testing.T) {
	db := setupTestRoleDB()
	truncateRole(db)

	tx, _ := db.Begin()
	roleRepository := repository.NewRoleRepository()
	role := roleRepository.Create(context.Background(), tx, domain.Role{
		RoleName: "Employee",
	})
	tx.Commit()

	router := setupRoleRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/roles/"+strconv.Itoa(role.Id), requestBody)
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

func TestGetRoleSuccess(t *testing.T) {
	db := setupTestRoleDB()
	truncateRole(db)

	tx, _ := db.Begin()
	roleRepository := repository.NewRoleRepository()
	role := roleRepository.Create(context.Background(), tx, domain.Role{
		RoleName: "Owner",
	})
	tx.Commit()

	router := setupRoleRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/roles/"+strconv.Itoa(role.Id), nil)
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
	assert.Equal(t, role.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, role.RoleName, responseBody["data"].(map[string]interface{})["role_name"])
}

func TestGetRoleFail(t *testing.T) {
	db := setupTestRoleDB()
	truncateRole(db)
	router := setupRoleRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/roles/404", nil)
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

func TestDeleteRoleSuccess(t *testing.T) {
	db := setupTestRoleDB()
	truncateRole(db)

	tx, _ := db.Begin()
	roleRepository := repository.NewRoleRepository()
	role := roleRepository.Create(context.Background(), tx, domain.Role{
		RoleName: "Owner",
	})
	tx.Commit()

	router := setupRoleRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/roles/"+strconv.Itoa(role.Id), nil)
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

func TestDeleteRoleFail(t *testing.T) {
	db := setupTestRoleDB()
	truncateRole(db)
	router := setupRoleRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/roles/404", nil)
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

func TestListRoleSuccess(t *testing.T) {
	db := setupTestRoleDB()
	truncateRole(db)

	tx, _ := db.Begin()
	roleRepository := repository.NewRoleRepository()
	role1 := roleRepository.Create(context.Background(), tx, domain.Role{
		RoleName: "Owner",
	})
	role2 := roleRepository.Create(context.Background(), tx, domain.Role{
		RoleName: "Employee",
	})
	tx.Commit()

	router := setupRoleRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/roles", nil)
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

	assert.Equal(t, role1.Id, int(userResponse1["id"].(float64)))
	assert.Equal(t, role1.RoleName, userResponse1["role_name"])

	assert.Equal(t, role2.Id, int(userResponse2["id"].(float64)))
	assert.Equal(t, role2.RoleName, userResponse2["role_name"])
}

func TestRoleUnauthorized(t *testing.T) {
	db := setupTestRoleDB()
	truncateRole(db)
	router := setupRoleRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/roles", nil)
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
