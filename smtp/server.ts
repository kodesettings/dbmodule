import mail_helper from "./mailer"; // initialize mail sender

export default {
  async fetch(request: Request, env: Env, ctx: ExecutionContext): Promise<Response> {
    console.log("== request url: ", request.url);

    let isSent = await mail_helper(await request.json()["email_body"]);
    return new Response("email status: " + (isSent == true ? "success" : "failed"));
  }
}
