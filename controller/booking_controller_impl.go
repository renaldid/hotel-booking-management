package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/web"
	"github.com/renaldid/hotel_booking_management.git/service"
	"net/http"
	"strconv"
)

type BookingControllerImpl struct {
	BookingService service.BookingService
}

func NewBookingController(BookingService service.BookingService) BookingController {
	return &BookingControllerImpl{
		BookingService: BookingService,
	}
}

func (controller *BookingControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingCreateRequest := web.BookingCreateRequest{}
	helper.ReadFromRequestBody(request, &bookingCreateRequest)

	bookingResponse := controller.BookingService.Create(request.Context(), bookingCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingUpdateRequest := web.BookingUpdateRequest{}
	helper.ReadFromRequestBody(request, &bookingUpdateRequest)

	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	bookingUpdateRequest.Id = id

	bookingResponse := controller.BookingService.Update(request.Context(), bookingUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	controller.BookingService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingId := params.ByName("bookingId")
	id, err := strconv.Atoi(bookingId)
	helper.PanicIfError(err)

	bookingResponse := controller.BookingService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *BookingControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	bookingResponse := controller.BookingService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   bookingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
