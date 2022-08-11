package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldid/hotel_booking_management.git/controller"
	"github.com/renaldid/hotel_booking_management.git/exception"
)

func NewRouter(booking controller.BookingController, discount controller.DiscountController, employee controller.EmployeeController,
	facility controller.FacilityController, floor controller.FloorController, hotel controller.HotelController,
	invoice controller.InvoiceController, meetingRoomPricing controller.MeetingRoomPricingController,
	meetingRoom controller.MeetingRoomController, role controller.RoleController, user controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/bookings", booking.FindAll)
	router.GET("/api/bookings/:bookingId", booking.FindById)
	router.POST("/api/bookings", booking.Create)
	router.PUT("/api/bookings/:bookingId", booking.Update)
	router.DELETE("/api/booking/:bookingId", booking.Delete)

	router.GET("/api/discounts", discount.FindAll)
	router.GET("/api/discounts/:discountId", discount.FindById)
	router.POST("/api/discounts", discount.Create)
	router.PUT("/api/discounts/:discountId", discount.Update)
	router.DELETE("/api/discounts/:discountId", discount.Delete)

	router.GET("/api/employees", employee.FindAll)
	router.GET("/api/employees/:employeeId", employee.FindById)
	router.POST("/api/employees", employee.Create)
	router.PUT("/api/employees/:employeeId", employee.Update)
	router.DELETE("/api/employees/:employeeId", employee.Delete)

	router.GET("/api/facilities", facility.FindAll)
	router.GET("/api/facilities/:facilityId", facility.FindById)
	router.POST("/api/facilities", facility.Create)
	router.PUT("/api/facilities/:facilityId", facility.Update)
	router.DELETE("/api/facilities/:facilityId", facility.Delete)

	router.GET("/api/floors", floor.FindAll)
	router.GET("/api/floors/:floorId", floor.FindById)
	router.POST("/api/floors", floor.Create)
	router.PUT("/api/floors/:floorId", floor.Update)
	router.DELETE("/api/floors/:floorId", floor.Delete)

	router.GET("/api/hotels", hotel.FindAll)
	router.GET("/api/hotels/:hotelId", hotel.FindById)
	router.POST("/api/hotels", hotel.Create)
	router.PUT("/api/hotels/:hotelId", hotel.Update)
	router.DELETE("/api/hotels/:hotelId", hotel.Delete)

	router.GET("/api/invoices", invoice.FindAll)
	router.GET("/api/invoices/:invoiceId", invoice.FindById)
	router.POST("/api/invoices", invoice.Create)
	router.PUT("/api/invoices/:invoiceId", invoice.Update)
	router.DELETE("/api/invoices/:invoiceId", invoice.Delete)

	router.GET("/api/meeting_room_pricings", meetingRoomPricing.FindAll)
	router.GET("/api/meeting_room_pricings/:meeting_room_pricingId", meetingRoomPricing.FindById)
	router.POST("/api/meeting_room_pricings", meetingRoomPricing.Create)
	router.PUT("/api/meeting_room_pricings/:meeting_room_pricingId", meetingRoomPricing.Update)
	router.DELETE("/api/meeting_room_pricings/:meeting_room_pricingId", meetingRoomPricing.Delete)

	router.GET("/api/meeting_rooms", meetingRoom.FindAll)
	router.GET("/api/meeting_rooms/:meeting_roomId", meetingRoom.FindById)
	router.POST("/api/meeting_rooms", meetingRoom.Create)
	router.PUT("/api/meeting_rooms/:meeting_roomId", meetingRoom.Update)
	router.DELETE("/api/meeting_rooms/:meeting_roomId", meetingRoom.Delete)

	router.GET("/api/roles", role.FindAll)
	router.GET("/api/roles/:roleId", role.FindById)
	router.POST("/api/roles", role.Create)
	router.PUT("/api/roles/:roleId", role.Update)
	router.DELETE("/api/roles/:roleId", role.Delete)

	router.GET("/api/users", user.FindAll)
	router.GET("/api/users/:userId", user.FindById)
	router.POST("/api/users", user.Create)
	router.PUT("/api/users/:userId", user.Update)
	router.DELETE("/api/users/:userId", user.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
