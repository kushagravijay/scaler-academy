package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/kushagravijay/scaler assignment/go-scheduler/pkg/utils"
	"github.com/kushagravijay/scaler assignment/go-scheduler/pkg/models"
)

var NewMeet= models.Meeting

func GetAllMeetings(w http.ResponseWriter, r *http.Request){
	newMeets:=models.GetAllMeetings()
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.write(res)
}

func GetMeetingById(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	meetId := vars["meetId"]
	ID, err := strconv.ParseInt(meetId,0,0)
	if err !=nil {
		fmt.Println("error while parsing")
	}
	meetDetails, _:= models.GetMeetingById(ID)
	res, _ := json.Marshal(meetDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.writeHeader(http.StatusOK)
	w.Write(res)
}


func CreateMeeting(w http.ResponseWriter, r *http.Request){
	createMeet := &model.Meeting{}
	utils.parseBody(r, createMeet)
	m:= createMeet.CreateMeeting()
	res, _ := json.Marshal(m)
	w.WriteHeader(http.StatusOK)
	w.write(res)
}

func DeleteMeeting(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	meetId := vars["meetId"]
	ID, err := strconv.ParseInt(meetId,0,0)
	if err !=nil {
		fmt.Println("error while parsing")
	}
	meet := models.DeleteMeeting(ID)
	res,_ := json.Marshal(meet)
	w.Header().Set("Content-Type","pkglication/json")
	w.writeHeader(http.StatusOK)
	w.Write(res)
}


func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateMeet = &models.Meeting{}
	utils.ParseBody(r, updateMeet)
	vars:=mux.Vars(r)
	meetId := vars["meetId"]
	ID, err := strconv.ParseInt(meetId,0,0)
	if err !=nil {
		fmt.Println("error while parsing")
	}
	meetDetails, db := models.GetMeetingById(ID)
	if updateMeet.Name != ""{
		meetDetails.Name = updateMeet.Name
	}
	if updateMeet.Start_time != ""{
		meetDetails.Start_time = updateMeet.Start_time
	}
	if updateMeet.End_time != ""{
		meetDetails.End_time = updateMeet.End_time
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.writeHeader(http.StatusOK)
	w.Write(res)
}

