package main

import (
	"fmt"
	"math/rand"
	"time"
	"log"
	"net/http"
	// "encoding/json"

	"github.com/gorilla/mux"
	"github.com/bitly/go-simplejson"
)

type ship struct {
	name     string
	}

func returnShipName(w http.ResponseWriter, r *http.Request) {
	ships := [...]string {
		"So Much For Subtlety",
		"Little Rascal",
		"Unfortunate Conflict Of Evidence",
		"Just Read The Instructions",
		"Flexible Demeanour",
		"Of Course I Still Love You",
		"Limiting Factor",
		"Gunboat Diplomat",
		"Size Isn't Everything",
		"Congenital Optimist",
		"Sweet and Full of Grace",
		"Death And Gravity",
		"Ethics Gradient",
		"Honest Mistake",
		"Serious Callers Only",
		"Fate Amenable To Change",
		"Very Little Gravitas Indeed",
		"Problem Child",
		"It's Character Forming",
		"Killing Time",
		"Quietly Confident",
		"Experiencing A Significant Gravitas Shortfall",
		"Don't Try This At Home",
		"Now We Try It My Way",
		"Falling Outside the Normal Moral Constraints",
		"Mistake Not...",
		"Smile Tolerantly",
	}
	rand.Seed(time.Now().Unix())
	shipName := ships[rand.Intn(len(ships))]

	json := simplejson.New()
	json.Set("name", shipName)

	payload, err := json.MarshalJSON()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	}

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Nothing appearing on wide beam...")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/return-culture-ship-name", returnShipName).Methods("GET")
	http.Handle("/", myRouter)

    log.Fatal(http.ListenAndServe(":8089", nil))
}

func main() {

    handleRequests()
    // cmd.Execute()
	
}
