// package handler

// import (
// 	"log"
// 	"net/http"
// 	"net/smtp"
// 	"os"

// 	"github.com/joho/godotenv"
// )

// func SendEmail(w http.ResponseWriter, r *http.Request) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Println("Error", err)
// 		return
// 	}

// 	targetEmail := ""

// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"
// 	senderEmail := os.Getenv("SENDEREMAIL")
// 	appPassword := os.Getenv("APP_PASS")

// 	auth := smtp.PlainAuth("", senderEmail, appPassword, smtpHost)

// 	to := []string{targetEmail}
// 	msg := []byte("To: " + targetEmail + "\r\n" +
// 		"Test\r\n" +
// 		"\r\n" +
// 		"Helo.\r\n")

// 	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, to, msg)

// 	if err != nil {
// 		log.Println("Error ", err)

// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
