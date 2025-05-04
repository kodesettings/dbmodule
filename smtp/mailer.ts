import { EmailMessage } from "cloudflare:email";

const EMAIL_SENDER_ADDR = "<SENDER>@somedomain.domain";
const EMAIL_RECIPIENT_ADDR = "<SENDER>@somedomain.domain";

export default async (email_body: string) : Promise<boolean> => {
  var message = new EmailMessage(EMAIL_SENDER_ADDR /* sender */, EMAIL_RECIPIENT_ADDR /* recipient */, email_body);

  try {
    // TODO: this is cloudflare workers environment variable
    // await env.SEB.send(message);
  } catch (e: any) {
    console.log(e.message);
    return false;
  }

  return true;
}
