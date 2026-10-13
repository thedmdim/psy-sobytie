package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/SevereCloud/vksdk/v3/api"
	"github.com/rs/zerolog/log"
)

const maxBytes = 10 * 1024

func main() {

	vkToken := os.Getenv("VK_TOKEN")
	if vkToken == "" {
		log.Fatal().Msg("set VK_TOKEN env")
	}
	
    vkChatID, err := strconv.ParseInt(os.Getenv("VK_CHAT_ID"), 10, 64) 
	if err != nil {
		log.Fatal().Msg("set VK_CHAT_ID env")
	}
 
	vkAPI := api.NewVK(vkToken)

	http.HandleFunc("/submit", NewLeadHandler(vkAPI, vkChatID))
	log.Fatal().Err(http.ListenAndServe(":8080", nil)).Msg("stopped listening")
}

func NewLeadHandler(vkAPI *api.VK, vkChatID int64) func(w http.ResponseWriter, r *http.Request) {
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

		go func(){
			_, err := vkAPI.MessagesSend(api.Params{
				"peer_id":  vkChatID,
				"random_id": 0,
				"message":  msg,
			})
			if err != nil {
				log.Error().Err(err).Int64("peer_id", vkChatID).Msg("cannot send message")
			}
		}()

		http.Redirect(w, r, "/success/", http.StatusSeeOther)
	}
}


type field struct {
	emoji    string
	key      string
	required bool
	checker    func(string)bool  //must ruturn false to pass
}

var messageFields = []field{
	{"📖", "from", true, nil},
	{"👤", "name", true, nil},
	{"📞", "phone", true, nil},
	{"🐣", "age", false, nil},
	{"💬", "comment", false, commentChecker},
}

var linkRE = regexp.MustCompile(`https?://[^\s]+`)

func commentChecker(s string) bool {
	if linkRE.MatchString(s) {
		return true
	}

	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		if unicode.Is(unicode.Cyrillic, r) {
			return false
		}
	}

	return true
}

func formToMessage(form url.Values) (string, error) {
	var lines []string
	for _, field := range messageFields {
		if form.Has(field.key) {
			if field.checker != nil && field.checker(form.Get(field.key)) {
				return "", fmt.Errorf("field hasn't passed its checker: %s", field.key)
			}
		} else {
			if field.required {
				return "", fmt.Errorf("no requires field: %s", field.key)
			}
			continue
		}

		lines = append(lines, field.emoji+form.Get(field.key))
	}
	return strings.Join(lines, "\n"), nil
}
