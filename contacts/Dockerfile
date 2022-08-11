FROM golang

ENV PORT 50080

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE ${PORT}

ENTRYPOINT [ "go","run","main.go"]


