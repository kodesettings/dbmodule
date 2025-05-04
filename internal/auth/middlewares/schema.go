package auth_middlewares

var schema string = `
{
    auth: {
        type: "object",
        required: ["authorization"],
        properties: {
            deviceIdentifier: { type: "string" }
        }
    }
}
`
