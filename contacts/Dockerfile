FROM golang

ENV PORT 50091

WORKDIR /app

COPY . .

RUN go mod tidy

CMD ["go" ,"run", "main.go"]

