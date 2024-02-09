FROM golang:1.22.0-alpine3.19 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN  go build -o /rinha .

#Segunda etapa
# Use uma imagem mínima do alpine como base para a imagem final
FROM alpine:latest


WORKDIR /

COPY --from=builder /rinha /rinha

EXPOSE 8082

CMD ["/rinha"]
