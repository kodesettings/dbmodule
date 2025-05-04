import { EmailMessage } from "cloudflare:email";
import mimecontent from "./mimecontent";

const EMAIL_SENDER_NAME = "NAME of the SENDER";
const EMAIL_SENDER_ADDR = "<SENDER>@somedomain.domain";

export async function sendEmailHelperAction(email: string, fullname: string, subject: string, body: string): Promise<boolean> {
  const email_body = mimecontent(email, fullname, EMAIL_SENDER_NAME, EMAIL_SENDER_ADDR, subject, body);
  var message = new EmailMessage(EMAIL_SENDER_ADDR /* sender */, email /* recipient */, email_body);

  try {
    // TODO: this is cloudflare workers environment variable
    // await env.SEB.send(message);
  } catch (e: any) {
    console.log(e.message);
    return false;
  }

  return true;
}
