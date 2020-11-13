package main

import (
<<<<<<< HEAD
        "encoding/json"
        "log"
        "net/http"
        "net/smtp"
)

type contactForm struct {
        Name string `json:"name"`
        Email string `json:"email"`
        Message string `json:"message"`
}
func sendEmail(rw http.ResponseWriter, req *http.Request) {
        c := &contactForm{}
        json.NewDecoder(req.Body).Decode(c)

        to := "shani.pathak98@gmail.com"
        subject := "Portfolio-Contact"
        body := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + "Name: " + c.Name + "\r\n\r\n" + "Email: " + c.Email + "\r\n\r\n" + "\r\n\r\n" + "Message: " + c.Message + "\r\n\r\n" + "OK to contact?: "
        auth := smtp.PlainAuth("", "***GMAIL USERNAME***", "***PASSWORD***", "smtp.gmail.com")
        err := smtp.SendMail("smtp.gmail.com:587", auth, "***GMAIL USERNAME***", []string{to},[]byte(body))
        if err != nil {
                log.Print("ERROR: attempting to send a mail ", err)
        }

}

func main() {
        //port := os.Getenv("PORT")
        //if port == "" {
        //        log.Fatal("$PORT must be set")
        //}
        //log.Fatal(http.ListenAndServe(":" + port, http.FileServer(http.Dir("."))))
        http.HandleFunc("/sendEmail", sendEmail)
        http.Handle( "/static/", http.StripPrefix( "/static", http.FileServer(http.Dir("./static")) ) )
        http.Handle( "/templates/", http.StripPrefix( "/templates", http.FileServer(http.Dir("./templates")) ) )
        http.Handle("/", http.FileServer(http.Dir(".")))
        log.Fatal(http.ListenAndServe(":8080", nil))

}

=======
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("."))))
}
>>>>>>> b60376713497a2e4d7fc21a5c23f92c8d9e561ce
