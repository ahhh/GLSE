package main

import (
  "bytes"
  "fmt"
  "log"
  "net/smtp"
  "strings"
  gomail "gopkg.in/gomail.v2"
)

var (
  sendmail = "/usr/sbin/sendmail"
  userList = []string{
    "test11@mailinator.com",
    "test12@mailinator.com",
    "test13@mailinator.com",
  }
  body = `
    Hello Friend,
      Thank you for subscribing.
      Have some <a href="https://www.google.com">linkz</a>
    Bye!
  `
)

type Mail struct {
  Sender  string
  To      []string
  Cc      []string
  Bcc     []string
  Subject string
  Body    string
}

func main(){
  fmt.Println("Sending spoofed email now!")
  sendSampleEmail()
  fmt.Println("Succesfully sent emails!")
}

func (mail *Mail) BuildMessage() string {
  header := ""
  header += fmt.Sprintf("From: %s\r\n", mail.Sender)
  if len(mail.To) > 0 {
    header += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
  }
  if len(mail.Cc) > 0 {
    header += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
  }
  header += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
  header += "\r\n" + mail.Body
  return header
}

func sendSampleEmail() {
  // Connect to the SMTP server.
  c, err := smtp.Dial("mail.mailinator.com:25")
  if err != nil {
    log.Fatal(err)
  }
  defer c.Close()
  m := gomail.NewMessage()
  m.SetHeader("From", "obama@whitehouse.gov", "Barack Obama")
  m.SetHeader("To", "test11@mailinator.com")
  //m.SetHeader("Bcc", userList...)
  m.SetHeader("Subject", "Thanks for Subscribing!")
  m.SetBody("text/html", body)
  buf := new(bytes.Buffer)
  m.WriteTo(buf)
  c.Mail("obama@whitehouse.gov")
  c.Rcpt("test11@mailinator.com")
  wc, err := c.Data()
  if err != nil {
    log.Fatal(err)
  }
  defer wc.Close()
  if _, err = buf.WriteTo(wc); err != nil {
    log.Fatal(err)
  }
}
