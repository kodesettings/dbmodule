package auth_apikey

var schema string = `
{
    apiKey: {
        type: "object",
        required: ["x-api-key"],
        properties: {
            ["x-api-key"]: { type: "string" }
        }
    },
    auth: {
        type: "object",
        required: ["authorization"],
        properties: {
            deviceIdentifier: { type: "string" }
        }
    }
}
`
