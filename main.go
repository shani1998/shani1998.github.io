package main

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"

	log "github.com/sirupsen/logrus"
)

type contactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Errorf("Invalid request type for sending a email, %s", r.Method)
	}
	// get credentials set in env
	from := os.Getenv("FROMMAIL")
	password := os.Getenv("PASSWD")
	to := os.Getenv("TOMAIL")

	// read form data
	name := r.FormValue("name")
	email := r.FormValue("email")
	message := r.FormValue("message")

	// send mail with above info
	subject := "Portfolio-Contact"
	body := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + "Name: " + name + "\r\n\r\n" + "Email: " + email + "\r\n\r\n" + "\r\n\r\n" + "Message: " + message
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(body))
	if err != nil {
		log.Errorf("Error occurred while sending an email, %v", err)
		fmt.Fprintln(w, `<script type="text/javascript">alert("ERROR: attempting to send a mail !"); window.location = '/';</script>`)
	} else {
		log.Infof("Successfully sent email with  name:msg, %s:%s", name, message)
		fmt.Fprintln(w, `<script type="text/javascript">alert("mail sent successfully!"); window.location = '/';</script>`)
	}
}

func main() {
	// retrieve port variable set by heroku
	// https://devcenter.heroku.com/articles/getting-started-with-go#define-config-vars
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	logger := &log.Logger{
		Out:   os.Stderr,
		Level: log.DebugLevel,
		Formatter: &log.TextFormatter{
			DisableColors:   false,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		},
	}

	logger.Printf("Starting portfolio server....")
	http.HandleFunc("/sendEmail", sendEmail)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("./templates"))))
	http.Handle("/", http.FileServer(http.Dir(".")))
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Errorf("failed to start portfolio server, Reason %v", err)
	}
}
