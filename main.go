package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/renaldid/hotel_booking_management.git/app"
	"github.com/renaldid/hotel_booking_management.git/controller"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/middleware"
	"github.com/renaldid/hotel_booking_management.git/repository"
	"github.com/renaldid/hotel_booking_management.git/service"
	"log"
	"net/http"
	"os"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	bookingRepository := repository.NewBookingRepository()
	bookingService := service.NewBookingService(bookingRepository, db, validate)
	bookingController := controller.NewBookingController(bookingService)

	roleRepository := repository.NewRoleRepository()
	roleService := service.NewRoleService(roleRepository, db, validate)
	roleController := controller.NewRoleController(roleService)

	hotelRepository := repository.NewHotelRepository()
	hotelService := service.NewHotelService(hotelRepository, db, validate)
	hotelController := controller.NewHotelController(hotelService)

	discountRepository := repository.NewDiscountRepository()
	discountService := service.NewDiscountService(discountRepository, db, validate)
	discountController := controller.NewDiscountController(discountService)

	meetingRoomRepository := repository.NewMeetingRoomRepository()
	meetingRoomService := service.NewMeetingRoomService(meetingRoomRepository, db, validate)
	meetingRoomController := controller.NewMeetingRoomController(meetingRoomService)

	invoiceRepository := repository.NewInvoiceRepository()
	invoiceService := service.NewInvoiceService(invoiceRepository, db, validate)
	invoiceController := controller.NewInvoiceController(invoiceService)

	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	mrpRepository := repository.NewMeetingRoomPricingRepository()
	mrpService := service.NewMeetingRoomPricingService(mrpRepository, db, validate)
	mrpController := controller.NewMeetingRoomPricingController(mrpService)

	facilityRepository := repository.NewFacilityRepository()
	facilityService := service.NewFacilityService(facilityRepository, db, validate)
	facilityController := controller.NewFacilityController(facilityService)

	floorRepository := repository.NewFloorRepository()
	floorService := service.NewFloorService(floorRepository, db, validate)
	floorController := controller.NewFloorController(floorService)

	employeeRepository := repository.NewEmployeeRepository()
	employeeService := service.NewEmployeeService(employeeRepository, db, validate)
	employeeController := controller.NewEmployeeController(employeeService)

	router := app.NewRouter(bookingController, discountController, employeeController, facilityController,
		floorController, hotelController, invoiceController, mrpController, meetingRoomController, roleController, userController)

	server := http.Server{
		//Addr:    "localhost:3000",
		Addr:    "https://booking-management-system-fp.herokuapp.com/" + os.Getenv("PORT") + "/",
		Handler: middleware.NewAuthMiddleware(router),
	}
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
