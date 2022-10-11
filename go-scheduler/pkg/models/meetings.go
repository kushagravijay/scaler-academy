package models

import(
	"time"
	"github.com/jinzhu/gorm"
	"github.com/Kushagra Vijay/Desktop/scaler assignment/go-scheduler/pkg/connection"
)


var db *gorm.DB

type Meeting struct{
	gorm.model
	Subject string `gorm:""json:"name"`
	Start_time time.time `"json:start_time"`
	Duration time.Duration `json:"duration"`
	End_time time.time `"json:start_time"`
	Number_of_participants uint `"json:number_of_participants"`
} 

func init(){
	db = connection.connectDB();
	db.AutoMigrate(&Meeting{})
}

func (meeting *Meeting) CreateMeeting() *Meeting{
	db.NewRecord(meeting)
	db.Create(&meeting)
	return meeting
}

func GetAllMeetings() []Meeting{
	var meetings []Meeting
	db.Find(&meetings)
	return meetings
}

func GetMeetingById(Id int64) (*Meeting, *gorm.DB){
	var getMeet Meeting
	db :=db.Where("ID=?",Id).Find(&getMeet)
	return &getMeet, db
}

Func DeleteBook(Id int64) Meeting{
	var meet Meeting
	db:=db.Where("ID=?", Id).Delete(book)
	return book
}






