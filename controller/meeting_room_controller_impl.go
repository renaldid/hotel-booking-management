package controller

import (
	"github.com/julienschmidt/httprouter"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"github.com/renaldid/hotel_booking_management.git/model/web"
	"github.com/renaldid/hotel_booking_management.git/service"
	"net/http"
	"strconv"
)

type MeetingRoomControllerImpl struct {
	MeetingRoomService service.MeetingRoomService
}

func NewMeetingRoomController(meetingRoomService service.MeetingRoomService) MeetingRoomController {
	return &MeetingRoomControllerImpl{
		MeetingRoomService: meetingRoomService,
	}
}

func (controller *MeetingRoomControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meetingRoomCreateRequest := web.MeetingRoomCreateRequest{}
	helper.ReadFromRequestBody(request, &meetingRoomCreateRequest)

	meetingRoomResponse := controller.MeetingRoomService.Create(request.Context(), meetingRoomCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meetingRoomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meetingRoomUpdateRequest := web.MeetingRoomUpdateRequest{}
	helper.ReadFromRequestBody(request, &meetingRoomUpdateRequest)

	meetingRoomId := params.ByName("meeting_roomId")
	id, err := strconv.Atoi(meetingRoomId)
	helper.PanicIfError(err)

	meetingRoomUpdateRequest.Id = id

	meetingRoomResponse := controller.MeetingRoomService.Update(request.Context(), meetingRoomUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meetingRoomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meetingRoomId := params.ByName("meeting_roomId")
	id, err := strconv.Atoi(meetingRoomId)
	helper.PanicIfError(err)

	controller.MeetingRoomService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meetingRoomId := params.ByName("meeting_roomId")
	id, err := strconv.Atoi(meetingRoomId)
	helper.PanicIfError(err)

	meetingRoomResponse := controller.MeetingRoomService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meetingRoomResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *MeetingRoomControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	meetingRoomResponses := controller.MeetingRoomService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   meetingRoomResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
