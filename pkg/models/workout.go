package models

import (
	"github.com/usamah/go-workout/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Workout struct {
	gorm.Model
	Name   string `json:"name"`
	Muscle string `json:"muscle"`
	Reps   string `json:"reps"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Workout{})
}

func (w *Workout) CreateWorkout() *Workout {
	// db.NewRecord(w)
	db.Create(&w)
	return w
}

func GetAllWorkouts() []Workout {
	var Workouts []Workout
	db.Find(&Workouts)
	return Workouts
}

func GetWorkoutById(Id int) (*Workout, *gorm.DB) {
	var getWorkout Workout
	db := db.Where("ID=?", Id).Find(&getWorkout)
	return &getWorkout, db
}

func DeleteWorkout(Id int) Workout {
	var workout Workout
	db.Where("ID=?", Id).Delete(workout)
	return workout
}
