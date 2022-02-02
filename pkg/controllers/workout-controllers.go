package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/usamah/go-workout/pkg/models"
	"github.com/usamah/go-workout/pkg/utils"
)

// var newWorkout models.Workout

func GetWorkout(w http.ResponseWriter, r *http.Request) {
	newWorkouts := models.GetAllWorkouts()
	// convert to Json
	res, _ := json.Marshal(newWorkouts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// send json version
	w.Write(res)
}

func GetWorkoutById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workoutId := vars["workoutId"]
	ID, _ := strconv.Atoi(workoutId)
	workoutDetails, _ := models.GetWorkoutById(ID)
	res, _ := json.Marshal(workoutDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateWorkout(w http.ResponseWriter, r *http.Request) {
	CreateWorkout := &models.Workout{}
	utils.ParseBody(r, CreateWorkout)
	workout := CreateWorkout.CreateWorkout()
	res, _ := json.Marshal(workout)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteWorkout(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	workoutId := vars["workoutId"]
	ID, _ := strconv.Atoi(workoutId)
	workout := models.DeleteWorkout(ID)
	res, _ := json.Marshal(workout)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateWorkout(w http.ResponseWriter, r *http.Request) {
	var UpdateWorkout = &models.Workout{}
	utils.ParseBody(r, UpdateWorkout)
	vars := mux.Vars(r)
	workoutId := vars["workoutId"]
	ID, _ := strconv.Atoi(workoutId)
	workoutDetails, db := models.GetWorkoutById(ID)
	if workoutDetails.Name != "" {
		workoutDetails.Name = UpdateWorkout.Name
	}
	if workoutDetails.Muscle != "" {
		workoutDetails.Muscle = UpdateWorkout.Muscle
	}
	if workoutDetails.Reps != "" {
		workoutDetails.Reps = UpdateWorkout.Reps
	}
	db.Save(&workoutDetails)
	res, _ := json.Marshal(workoutDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
