package main

import (
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	c, err := smtp.Dial("postfix.test:25")

	if err != nil {
		log.Fatal(err.Error())
	}

	// Set the sender and recipient first
	if err := c.Mail("info@postfix.test"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("taro@recive-postfix.test"); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(wc, "Subject: Your Subject Here\r\n")
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(wc, "\r\n") // 空行を追加してヘッダーとボディを区切る
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
