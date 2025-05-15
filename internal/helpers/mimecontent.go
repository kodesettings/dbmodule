package helpers

import (
	"fmt"
	"strings"
)

func MimeTextBuilder(email_r string, fullname_r string, email_s string, fullname_s string, subject string, body string) string {
	var email_sender,email_recipient, email_subject, email_body string = "", "", "", "";

	email_sender = strings.Join([]string{"From: ", fullname_s, " <", email_s, ">"}, "\n");
	email_recipient = strings.Join([]string{"To: ", fullname_r, " <", email_r, ">"}, "\n");
	email_subject = strings.Join([]string{"Subject: ", subject}, "\n");
	email_body = strings.Join([]string{body}, "\n");

	content_type := "Content-Type: text/plain;charset=utf-8";
	mime_content_text := fmt.Sprintf("%s\n%s%s%s\n%s", content_type, email_sender, email_recipient, email_subject, email_body);

	return mime_content_text;
}
