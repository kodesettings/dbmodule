package routes_apikey

var schema string = `
{
    id: {
        type: "object",
        required: ["id"],
        properties: {
            id: { type: "string" },
        }
    },
    create: {
        type: "object",
        required: ["key", "status", "permission"],
        properties: {
            key: { type: "string" },
            status: { type: "boolean" },
            permission: { type: "array" },
        }
    },
    update: {
        type: "object",
        required: ["key", "status", "permission"],
        properties: {
            key: { type: "string" },
            status: { type: "boolean" },
            permission: { type: "array" },
        }
    },
    findByKey: {
        type: "object",
        required: ["key"],
        properties: {
            key: { type: "string" }
        }
    }
}
`
