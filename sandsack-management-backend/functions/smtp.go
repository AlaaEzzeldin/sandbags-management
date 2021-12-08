package functions

import (
	"bytes"
	mail "github.com/xhit/go-simple-mail/v2"
	"gorm.io/gorm"
	"html/template"
	"log"
)

type Template struct {
	Subject string
	Name string
	Reason string
	OTP string
}


func SendEmail(db *gorm.DB, email string, otp, reason string) error {
	templates := Template{
		Subject: "This is code to verify the email",
		Name: email,
		OTP: otp,
		Reason: "Verification of email",
	}

	if reason == "recovery" {
		templates.Subject = "This is code to recovery the password"
		templates.Reason = "Password recovery"
	}

	// create html file
	tmpl, err := template.New("tmpl").Parse(emailTemplate)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, templates)
	if err != nil {
		return err
	}

	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = "feuerwehr@passau.deee"
	server.Password = "KGsz2vV@"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Println("SMTP client connect error", err)
	}

	message := mail.NewMSG()
	message.SetFrom("feuerwehr@passau.deee")
	message.AddTo(email)
	message.SetSubject(templates.Subject)

	message.SetBody(mail.TextHTML, buf.String())

	err = message.Send(smtpClient)
	if err != nil {
		log.Println(err.Error())
		//todo: when we get real credentials uncomment error
		//return err
	}

	return nil

}

