ARG IMAGE=golang:1.23-alpine
ARG EXPOSED_PORT=3000

# Build stage

FROM ${IMAGE} AS build-stage

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

RUN ls -la

RUN CGO_ENABLED=0 GOOS=linux go build -o /llm-size-service ./cmd/llm-size-service/main.go

# Run stage

FROM ${IMAGE} AS run-stage

RUN apk add git git-lfs

COPY --from=build-stage /llm-size-service /llm-size-service

EXPOSE ${EXPOSED_PORT}

CMD [ "/llm-size-service" ]
