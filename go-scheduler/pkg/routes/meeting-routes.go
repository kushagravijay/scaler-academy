package routes

import (
	"github.com/gorilla/mux"
	"github.com/Kushagra Vijay/Desktop/scaler assignment/go-scheduler/pkg/controllers"
)

var MeetingRoutes = func(router *mux.Router) {
	router.HandleFunc("/meetings/", controllers.CreateMeeting).Method("POST")
	router.HandleFunc("/meetings/", controllers.GetMeeting).Method("GET")
	router.HandleFunc("/meeting/{meetingId}", controllers.GetMeetingById).Method("GET")
	router.HandleFunc("/meeting/{meetingId}", controllers.UpdateMeeting).Method("PUT")
	router.HandleFunc("/meeting/{meetingId}", controllers.DeleteMeeting).Method("DELETE")
}
