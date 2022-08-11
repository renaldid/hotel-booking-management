package test

import (
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
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func setupTestEmployeeDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_management_system")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupEmployeeRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	employeeRepository := repository.NewEmployeeRepository()
	employeeService := service.NewEmployeeService(employeeRepository, db, validate)
	employeeController := controller.NewEmployeeController(employeeService)
	router := routers.NewEmployeeRouter(employeeController)

	return middleware.NewAuthMiddleware(router)
}

func truncateEmployee(db *sql.DB) {
	db.Exec("TRUNCATE employees")
}

func TestCreateEmployeeSuccess(t *testing.T) {
	db := setupTestEmployeeDB()
	truncateEmployee(db)
	router := setupEmployeeRouter(db)

	requestBody := strings.NewReader(`{"role_id" : 2, "hotel_id" : 1, "name" : "Ranger Biru", "gender" : "female", "address" : "Jl. Tamiya", "email" : "email12@email.com", "phone_number" : "+20 81230666999"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/employees", requestBody)
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
	assert.Equal(t, 2, int(responseBody["data"].(map[string]interface{})["role_id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["hotel_id"].(float64)))
	assert.Equal(t, "Ranger Biru", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "female", responseBody["data"].(map[string]interface{})["gender"])
	assert.Equal(t, "Jl. Tamiya", responseBody["data"].(map[string]interface{})["address"])
	assert.Equal(t, "email12@email.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "+20 81230666999", responseBody["data"].(map[string]interface{})["phone_number"])
}

func TestCreateEmployeeFail(t *testing.T) {
	db := setupTestEmployeeDB()
	truncateEmployee(db)
	router := setupEmployeeRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/employees", requestBody)
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

func TestUpdateEmployeeSuccess(t *testing.T) {
	db := setupTestEmployeeDB()
	truncateEmployee(db)

	tx, _ := db.Begin()
	employeeRepository := repository.NewEmployeeRepository()
	employee := employeeRepository.Create(context.Background(), tx, domain.Employee{
		RoleId:      2,
		HotelId:     1,
		Name:        "ironman",
		Gender:      "female",
		Address:     "Jl. Kebenaran",
		Email:       "rorojok24@rocketmail.com",
		PhoneNumber: "+62 89000111222",
	})
	tx.Commit()

	router := setupEmployeeRouter(db)

	requestBody := strings.NewReader(`{"role_id" : 2, "hotel_id" : 1, "name" : "Ronaldo", "gender" : "male", "address" : "Jl. Jalan yuk", "email" : "ronaldo07@gmail.com", "phone_number" : "+62 87321456789"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/employees/"+strconv.Itoa(employee.Id), requestBody)
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
	assert.Equal(t, employee.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, 2, int(responseBody["data"].(map[string]interface{})["role_id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["hotel_id"].(float64)))
	assert.Equal(t, "Ronaldo", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "male", responseBody["data"].(map[string]interface{})["gender"])
	assert.Equal(t, "Jl. Jalan yuk", responseBody["data"].(map[string]interface{})["address"])
	assert.Equal(t, "ronaldo07@gmail.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "+62 87321456789", responseBody["data"].(map[string]interface{})["phone_number"])
	assert.Equal(t, "employee", responseBody["data"].(map[string]interface{})["role"])
	assert.Equal(t, "Ikan Hijau Hotel", responseBody["data"].(map[string]interface{})["hotel"])
}

func TestUpdateEmployeeFail(t *testing.T) {
	db := setupTestEmployeeDB()
	truncateEmployee(db)

	tx, _ := db.Begin()
	employeeRepository := repository.NewEmployeeRepository()
	employee := employeeRepository.Create(context.Background(), tx, domain.Employee{
		Name: "Badang",
	})
	tx.Commit()

	router := setupEmployeeRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/employees/"+strconv.Itoa(employee.Id), requestBody)
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

func TestGetEmployeeSuccess(t *testing.T) {
	db := setupTestEmployeeDB()
	truncateEmployee(db)

	tx, _ := db.Begin()
	employeeRepository := repository.NewEmployeeRepository()
	employee := employeeRepository.Create(context.Background(), tx, domain.Employee{
		RoleId:      2,
		HotelId:     1,
		Name:        "ironman",
		Gender:      "female",
		Address:     "Jl. Kebenaran",
		Email:       "rorojok24@rocketmail.com",
		PhoneNumber: "+62 89000111222",
	})
	tx.Commit()

	router := setupEmployeeRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/employees/"+strconv.Itoa(employee.Id), nil)
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
	assert.Equal(t, employee.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, 2, int(responseBody["data"].(map[string]interface{})["role_id"].(float64)))
	assert.Equal(t, 1, int(responseBody["data"].(map[string]interface{})["hotel_id"].(float64)))
	assert.Equal(t, "ironman", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "female", responseBody["data"].(map[string]interface{})["gender"])
	assert.Equal(t, "Jl. Kebenaran", responseBody["data"].(map[string]interface{})["address"])
	assert.Equal(t, "rorojok24@rocketmail.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "+62 89000111222", responseBody["data"].(map[string]interface{})["phone_number"])
	assert.Equal(t, "employee", responseBody["data"].(map[string]interface{})["role"])
	assert.Equal(t, "Ikan Hijau Hotel", responseBody["data"].(map[string]interface{})["hotel"])
}

func TestGetEmployeeFail(t *testing.T) {
	db := setupTestEmployeeDB()
	truncateEmployee(db)
	router := setupEmployeeRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/employees/404", nil)
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

func TestDeleteEmployeeSuccess(t *testing.T) {
	db := setupTestEmployeeDB()
	truncateEmployee(db)

	tx, _ := db.Begin()
	employeeRepository := repository.NewEmployeeRepository()
	employee := employeeRepository.Create(context.Background(), tx, domain.Employee{
		RoleId:      2,
		HotelId:     1,
		Name:        "ironman",
		Gender:      "female",
		Address:     "Jl. Kebenaran",
		Email:       "rorojok24@rocketmail.com",
		PhoneNumber: "+62 89000111222",
	})
	tx.Commit()

	router := setupEmployeeRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/employees/"+strconv.Itoa(employee.Id), nil)
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

func TestDeleteEmployeeFail(t *testing.T) {
	db := setupTestEmployeeDB()
	truncateEmployee(db)
	router := setupEmployeeRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/employees/404", nil)
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

func TestListEmployeesSuccess(t *testing.T) {
	db := setupTestEmployeeDB()
	truncateEmployee(db)

	tx, _ := db.Begin()
	employeeRepository := repository.NewEmployeeRepository()
	employee1 := employeeRepository.Create(context.Background(), tx, domain.Employee{
		RoleId:      2,
		HotelId:     1,
		Name:        "ironman",
		Gender:      "male",
		Address:     "Jl. Kebenaran",
		Email:       "rorojok24@rocketmail.com",
		PhoneNumber: "+62 89000111222",
		Role:        "employee",
		Hotel:       "Ikan Hijau Hotel",
	})
	employee2 := employeeRepository.Create(context.Background(), tx, domain.Employee{
		RoleId:      2,
		HotelId:     1,
		Name:        "ironwoman",
		Gender:      "female",
		Address:     "Jl. Kebenaran",
		Email:       "rorok24@rocketmail.com",
		PhoneNumber: "+62 8900019867222",
		Role:        "employee",
		Hotel:       "Ikan Hijau Hotel",
	})
	tx.Commit()

	router := setupEmployeeRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/employees", nil)
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

	var employees = responseBody["data"].([]interface{})

	employeeResponse1 := employees[0].(map[string]interface{})
	employeeResponse2 := employees[1].(map[string]interface{})

	assert.Equal(t, employee1.Id, int(employeeResponse1["id"].(float64)))
	assert.Equal(t, employee1.RoleId, int(employeeResponse1["role_id"].(float64)))
	assert.Equal(t, employee1.HotelId, int(employeeResponse1["hotel_id"].(float64)))
	assert.Equal(t, employee1.Name, employeeResponse1["name"])
	assert.Equal(t, employee1.Gender, employeeResponse1["gender"])
	assert.Equal(t, employee1.Address, employeeResponse1["address"])
	assert.Equal(t, employee1.Email, employeeResponse1["email"])
	assert.Equal(t, employee1.PhoneNumber, employeeResponse1["phone_number"])
	assert.Equal(t, employee1.Role, employeeResponse1["role"])
	assert.Equal(t, employee1.Hotel, employeeResponse1["hotel"])

	assert.Equal(t, employee2.Id, int(employeeResponse2["id"].(float64)))
	assert.Equal(t, employee2.RoleId, int(employeeResponse2["role_id"].(float64)))
	assert.Equal(t, employee2.HotelId, int(employeeResponse2["hotel_id"].(float64)))
	assert.Equal(t, employee2.Name, employeeResponse2["name"])
	assert.Equal(t, employee2.Gender, employeeResponse2["gender"])
	assert.Equal(t, employee2.Address, employeeResponse2["address"])
	assert.Equal(t, employee2.Email, employeeResponse2["email"])
	assert.Equal(t, employee2.PhoneNumber, employeeResponse2["phone_number"])
	assert.Equal(t, employee2.Role, employeeResponse2["role"])
	assert.Equal(t, employee2.Hotel, employeeResponse2["hotel"])
}

func TestEmployeeUnauthorized(t *testing.T) {
	db := setupTestEmployeeDB()
	truncateEmployee(db)
	router := setupEmployeeRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/employees", nil)
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
