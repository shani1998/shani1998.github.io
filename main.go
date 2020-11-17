package main

import (
        "fmt"
        "log"
        "net/http"
        "net/smtp"
)

type contactForm struct {
        Name string `json:"name"`
        Email string `json:"email"`
        Message string `json:"message"`
}
func sendEmail(w http.ResponseWriter, r *http.Request) {
        if r.Method=="POST"{
            name := r.FormValue("name")
            email := r.FormValue("email")
            message := r.FormValue("message")
            to := "shani.pathak98@gmail.com"
            subject := "Portfolio-Contact"
            body := "To: " + to + "\r\nSubject: " + subject + "\r\n\r\n" + "Name: " + name + "\r\n\r\n" + "Email: " + email + "\r\n\r\n" + "\r\n\r\n" + "Message: " + message
            auth := smtp.PlainAuth("", "skp80248@gmail.com", "vgmipaletguctnmv\n", "smtp.gmail.com")
            err := smtp.SendMail("smtp.gmail.com:587", auth,"skp80248@gmail.com" , []string{to},[]byte(body))
            if err != nil {
                    fmt.Fprintln(w,"ERROR: attempting to send a mail ", err)
            } else {
                    fmt.Fprintln(w,`<script type="text/javascript">alert("mail sent successfully!"); window.location = '/';</script>`)
            }

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



