package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/web"
	"github.com/renaldid/hotel_booking_management.git/service"
	"net/http"
	"strconv"
)

type MeetingRoomPricingControllerImpl struct {
	MeetingRoomPricingService service.MeetingRoomPricingService
}

func NewMeetingRoomPricingController(meetingRoomPricingService service.MeetingRoomPricingService) MeetingRoomPricingController {
	return &MeetingRoomPricingControllerImpl{
		MeetingRoomPricingService: meetingRoomPricingService,
	}
}

func (controller *MeetingRoomPricingControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meetingRoomPricingCreateRequest := web.MeetingRoomPricingCreateRequest{}
	helper.ReadFromRequestBody(request, &meetingRoomPricingCreateRequest)

	meetingRoomPricingResponse := controller.MeetingRoomPricingService.Create(request.Context(), meetingRoomPricingCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meetingRoomPricingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomPricingControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meetingRoomPricingUpdateRequest := web.MeetingRoomPricingUpdateRequest{}
	helper.ReadFromRequestBody(request, &meetingRoomPricingUpdateRequest)

	meetingRoomPricingId := params.ByName("meeting_room_pricingId")
	id, err := strconv.Atoi(meetingRoomPricingId)
	helper.PanicIfError(err)

	meetingRoomPricingUpdateRequest.Id = id

	meetingRoomPricingResponse := controller.MeetingRoomPricingService.Update(request.Context(), meetingRoomPricingUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meetingRoomPricingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomPricingControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meetingRoomPricingId := params.ByName("meeting_room_pricingId")
	id, err := strconv.Atoi(meetingRoomPricingId)
	helper.PanicIfError(err)

	controller.MeetingRoomPricingService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomPricingControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meetingRoomPricingId := params.ByName("meeting_room_pricingId")
	id, err := strconv.Atoi(meetingRoomPricingId)
	helper.PanicIfError(err)

	meetingRoomPricingResponse := controller.MeetingRoomPricingService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meetingRoomPricingResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomPricingControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meetingRoomPricingResponses := controller.MeetingRoomPricingService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meetingRoomPricingResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
