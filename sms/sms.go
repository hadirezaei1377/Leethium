package sms

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sfreiberg/gotwilio"
)

var twilio *gotwilio.Twilio

func sendSMS(w http.ResponseWriter, r *http.Request) {
	to := "+1234567890"
	msg := "Hello, this is a test message from your Go application."

	_, _, err := twilio.SendSMS("+1987654321", to, msg, "", "")
	if err != nil {
		log.Println("Error sending SMS:", err)
		http.Error(w, "Error sending SMS", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "SMS sent successfully to %s", to)
}

func main() {
	twilio = gotwilio.NewTwilioClient("YOUR_TWILIO_SID", "YOUR_TWILIO_AUTH_TOKEN")

	r := mux.NewRouter()
	r.HandleFunc("/send-sms", sendSMS).Methods("POST")

	http.Handle("/", r)

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
