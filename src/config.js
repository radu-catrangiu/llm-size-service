"use strict";

module.exports = {
    server: {
        host: process.env["SERVER_HOST"] ?? "localhost",
        port: process.env["SERVER_PORT"] ?? 3000,
    },
    huggingface: {
        password: process.env["HF_API_KEY"],
    }
};