# FROM alpine:3.20
FROM node:lts-alpine

RUN apk add git git-lfs

WORKDIR /app

COPY . .

EXPOSE 3000

CMD [ "node", "src/index.js" ]