package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type joked struct {
	joke1 string
	joke2 string
	joke3 string
	joke4 string
	joke5 string
}

var database = map[string]joked{
	"jokes": {
		joke1: "Why dont eggs tell jokes?: Because they might crack up.\n",
		joke2: "What does a lemon say when it answers the phone?:Yellow!\n",
		joke3: "Why couldnt the leopard play hide and seek?: Because he was always spotted.\n",
		joke4: "What do you call fake spaghetti?: An impasta.\n",
		joke5: "Why did the coffee file a police report?: It got mugged.\n"}}

func randlib(va []string) string {

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(va))

	return va[randomIndex]

}

func main() {
	http.HandleFunc("/madlib", madlibHandler)
	http.HandleFunc("/jokes", jokesHandler)

	values := []string{"Name: Charlie | Age: 78 Job: Construction | Sensing: I feel Charlie is having joint pains | Solution: He needs to retire ",
		"Name: Mary | Age: 65 Job: Banker | Sensing: I feel Mary is feeling stressed on a daily bases|  Solution: She needs to go on more vacation ",
		"Name: Sharon | Age: 81 Job: Baker | Sensing: I feel Sharon is always burnt out|  Solution: She needs move to a role ",
		"Name: Kunle | Age: 91 Job: Capenter | Sensing: I feel Kunle can't keep up|  Solution: Kunle needs to spend more time with his grandkids "}

	madlib := randlib(values)

	fmt.Println(madlib)

	fmt.Println("Server listening on port 4567")
	fmt.Println(database)

	port := ":4567"
	log.Fatal(http.ListenAndServe(port, nil))
}

func madlibHandler(w http.ResponseWriter, r *http.Request) {
	values := []string{"Name: Charlie | Age: 78 Job: Construction | Sensing: I feel Charlie is having joint pains | Solution: He needs to retire ",
		"Name: Mary | Age: 65 Job: Banker | Sensing: I feel Mary is feeling stressed on a daily bases|  Solution: She needs to go on more vacation ",
		"Name: Sharon | Age: 81 Job: Baker | Sensing: I feel Sharon is always burnt out|  Solution: She needs move to a role ",
		"Name: Kunle | Age: 91 Job: Capenter | Sensing: I feel Kunle can't keep up|  Solution: Kunle needs to spend more time with his grandkids "}

	madlib := randlib(values)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, madlib)
}

func jokesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, database)
}
