package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/NicoNex/echotron/v3"
)

const maxBytes = 10 * 1024

func main() {

	tgBotToken := os.Getenv("TG_BOT_TOKEN")
	if tgBotToken == "" {
		log.Fatal("set TG_BOT_TOKEN env")
	}
	
	tgGroupID, err := strconv.ParseInt(os.Getenv("TG_GROUP_ID"), 10, 64) 
	if err != nil {
		log.Fatal("set TG_GROUP_ID env")
	}
	
 
	tgAPI := echotron.NewAPI(tgBotToken)

	http.HandleFunc("/submit", NewLeadHandler(tgAPI, tgGroupID))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func NewLeadHandler(tgAPI echotron.API, tgGroupID int64) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
		if err := r.ParseForm(); err != nil {
			if _, ok := errors.AsType[*http.MaxBytesError](err); ok {
				w.WriteHeader(http.StatusRequestEntityTooLarge)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		msg, err := formToMessage(r.Form)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		go tgAPI.SendMessage(msg, tgGroupID, nil)
	}
}


type field struct {
	emoji    string
	key      string
	required bool
}

var messageFields = []field{
	{"📖", "from", true},
	{"👤", "name", true},
	{"📞", "phone", true},
	{"🐣", "age", false},
	{"💬", "comment", false},
}

func formToMessage(form url.Values) (string, error) {
	var lines []string
	for _, f := range messageFields {
		if !form.Has(f.key) {
			if f.required {
				return "", fmt.Errorf("no requires field: %s", f.key)
			}
			continue
		}
		lines = append(lines, f.emoji+form.Get(f.key))
	}
	return strings.Join(lines, "\n"), nil
}
