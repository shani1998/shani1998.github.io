package main

import (
	"fmt"
	"net"
	"net/http"
	"net/smtp"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/shani1998/shani1998.github.io/pkg/logger"
)

type contactForm struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func init() {
	pflag.Parse()
	viper.AutomaticEnv()

	logger.InitializeLogger(
		viper.GetString("log.format"),
		viper.GetString("log.level"),
	)
}

func main() {
	address := net.JoinHostPort(viper.GetString("server.address"), viper.GetString("server.port"))
	log.Infof("Starting portfolio server on address: %s", address)

	http.HandleFunc("/sendEmail", sendEmail)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.Handle("/templates/", http.StripPrefix("/templates", http.FileServer(http.Dir("./templates"))))
	http.Handle("/", http.FileServer(http.Dir(".")))

	// start http server
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Errorf("failed to start portfolio server, Reason %v", err)
	}
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
