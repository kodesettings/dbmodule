#ifndef HELPERS_MIMECONTENT_H
#define HELPERS_MIMECONTENT_H

#include "context.hpp"

std::string mimeTextBuilder(std::string email_r, std::string fullname_r, std::string email_s, std::string fullname_s, std::string subject, std::string body) {
  std::string email_sender,email_recipient, email_subject, email_body;

  email_sender = std::string("From: ").append(fullname_s).append(" <").append(email_s).append(">\n");
  email_recipient = std::string("To: ").append(fullname_r).append(" <").append(email_r).append(">\n");
  email_subject = std::string("Subject: ").append(subject).append("\n");
  email_body = std::string(body).append("\n");

  std::string content_type = "Content-Type: text/plain;charset=utf-8";

  char mime_content_text[2000];
  sprintf(mime_content_text, "%s\n%s%s%s\n%s", content_type.c_str(), email_sender.c_str(), email_recipient.c_str(), email_subject.c_str(), email_body.c_str());

  return std::string(mime_content_text);
}

#endif // HELPERS_MIMECONTENT_H
