package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/joke", jokeHandler)
	http.HandleFunc("/madlib", madlibHandler)

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func jokeHandler(w http.ResponseWriter, r *http.Request) {
	jokes := []string{
		"Why did the scarecrow win an award? Because he was outstanding in his field!",
		"What do you call fake spaghetti? An impasta!",
		"Why don’t skeletons fight each other? They don’t have the guts.",
		"What’s orange and sounds like a parrot? A carrot!",
		"Why did the bicycle fall over? It was two-tired.",
	}

	joke := jokes[rand.Intn(len(jokes))]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: joke})
}

func madlibHandler(w http.ResponseWriter, r *http.Request) {
	names := []string{"Alex", "Jordan", "Taylor", "Chris", "Morgan"}
	ages := []int{25, 30, 35, 40, 45}
	occupations := []string{"engineer", "teacher", "artist", "developer", "scientist"}
	devices := []string{"smartphone", "tablet", "laptop", "desktop", "smartwatch"}
	bodyParts := []string{"wrist", "forehead", "neck", "chest", "ankle"}
	moods := []string{"stressed", "anxious", "happy", "relaxed", "angry"}
	actions := []string{"playing soothing sounds", "displaying calming images", "suggesting breathing exercises", "providing motivational quotes", "giving gentle reminders to take breaks"}

	madlib := fmt.Sprintf(
		"%s is a %d-year old %s who has been struggling with a lot of job-related stress. "+
			"He/she decided to try a new application to relieve stress, which runs on a/an %s to help improve his/her mood. "+
			"The application senses his/her mood through a device he/she wears on his/her %s. "+
			"When the device senses that he/she is %s, it responds by %s.",
		names[rand.Intn(len(names))],
		ages[rand.Intn(len(ages))],
		occupations[rand.Intn(len(occupations))],
		devices[rand.Intn(len(devices))],
		bodyParts[rand.Intn(len(bodyParts))],
		moods[rand.Intn(len(moods))],
		actions[rand.Intn(len(actions))],
	)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{Message: madlib})
}
