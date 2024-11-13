"use strict";

const handler = require("./handler");
const config = require("./config");
const fastify = require("fastify");
const app = fastify();

app.get("/evaluate", async (req, reply) => {
    const model = req.query.model;
    if (!model) {
        reply.statusCode = 400;
        return;
    }
    const result = await handler(model);
    return result;
});

app.listen({
    host: config.server.host,
    port: config.server.port,
    listenTextResolver: (addr) => {
        console.log(`Server started: ${addr}`);
    },
});

process.on("SIGTERM", async () => {
    await app.close();
});
