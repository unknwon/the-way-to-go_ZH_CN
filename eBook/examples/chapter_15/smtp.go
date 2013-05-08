// smtp.go
package main

import (
        "bytes"
        "log"
        "net/smtp"
)

func main() {
        // Connect to the remote SMTP server.
        client, err := smtp.Dial("mail.example.com:25")
        if err != nil {
                log.Fatal(err)
        }
        // Set the sender and recipient.
        client.Mail("sender@example.org")
        client.Rcpt("recipient@example.net")
        // Send the email body.
        wc, err := client.Data()
        if err != nil {
                log.Fatal(err)
        }
        defer wc.Close()
        buf := bytes.NewBufferString("This is the email body.")
        if _, err = buf.WriteTo(wc); err != nil {
                log.Fatal(err)
        }
}

