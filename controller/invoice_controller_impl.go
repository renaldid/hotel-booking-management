package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/web"
	"github.com/renaldid/hotel_booking_management.git/service"
	"net/http"
	"strconv"
)

type InvoiceControllerImpl struct {
	InvoiceService service.InvoiceService
}

func NewInvoiceController(InvoiceService service.InvoiceService) InvoiceController {
	return &InvoiceControllerImpl{
		InvoiceService: InvoiceService,
	}
}

func (controller *InvoiceControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceCreateRequest := web.InvoiceCreateRequest{}
	helper.ReadFromRequestBody(request, &invoiceCreateRequest)

	invoiceResponse := controller.InvoiceService.Create(request.Context(), invoiceCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *InvoiceControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceUpdateRequest := web.InvoiceUpdateRequest{}
	helper.ReadFromRequestBody(request, &invoiceUpdateRequest)

	invoiceId := params.ByName("invoiceId")
	id, err := strconv.Atoi(invoiceId)
	helper.PanicIfError(err)

	invoiceUpdateRequest.Id = id

	invoiceResponse := controller.InvoiceService.Update(request.Context(), invoiceUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *InvoiceControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceId := params.ByName("invoiceId")
	id, err := strconv.Atoi(invoiceId)
	helper.PanicIfError(err)

	controller.InvoiceService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *InvoiceControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceId := params.ByName("invoiceId")
	id, err := strconv.Atoi(invoiceId)
	helper.PanicIfError(err)

	invoiceResponse := controller.InvoiceService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *InvoiceControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	invoiceResponse := controller.InvoiceService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   invoiceResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
