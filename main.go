package main

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
)

type quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}

var quotes = []quote{
	{"Programs must be written for people to read, and only incidentally for machines to execute.", "Harold Abelson"},
	{"Simplicity is prerequisite for reliability.", "Edsger W. Dijkstra"},
	{"Premature optimization is the root of all evil.", "Donald Knuth"},
	{"The best way to predict the future is to invent it.", "Alan Kay"},
	{"Any sufficiently advanced technology is indistinguishable from magic.", "Arthur C. Clarke"},
	{"We can only see a short distance ahead, but we can see plenty there that needs to be done.", "Alan Turing"},
	{"The question of whether a computer can think is no more interesting than the question of whether a submarine can swim.", "Edsger W. Dijkstra"},
	{"Talk is cheap. Show me the code.", "Linus Torvalds"},
	{"Controlling complexity is the essence of computer programming.", "Brian Kernighan"},
	{"Debugging is twice as hard as writing the code in the first place.", "Brian Kernighan"},
	{"Clear is better than clever.", "Rob Pike"},
	{"A little copying is better than a little dependency.", "Rob Pike"},
	{"The Web as I envisaged it, we have not seen it yet. The future is still so much bigger than the past.", "Tim Berners-Lee"},
	{"The Internet is becoming the town square for the global village of tomorrow.", "Bill Gates"},
	{"Software is eating the world.", "Marc Andreessen"},
	{"AI is the new electricity.", "Andrew Ng"},
	{"Machine intelligence is the last invention that humanity will ever need to make.", "Nick Bostrom"},
	{"The computer was born to solve problems that did not exist before.", "Bill Gates"},
	{"There are two ways of constructing a software design: make it so simple that there are obviously no deficiencies, or so complicated that there are no obvious deficiencies.", "C. A. R. Hoare"},
	{"Adding manpower to a late software project makes it later.", "Fred Brooks"},
	{"Walking on water and developing software from a specification are easy if both are frozen.", "Edward V. Berard"},
	{"First, solve the problem. Then, write the code.", "John Johnson"},
	{"Make it work, make it right, make it fast.", "Kent Beck"},
	{"Any fool can write code that a computer can understand. Good programmers write code that humans can understand.", "Martin Fowler"},
	{"The most dangerous phrase in the language is: we've always done it this way.", "Grace Hopper"},
	{"The real problem is not whether machines think but whether men do.", "B. F. Skinner"},
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	q := quotes[rand.IntN(len(quotes))]
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(q); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", handleRoot)

	log.Printf("listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
