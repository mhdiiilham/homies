FROM golang:1.19-alpine AS builder
RUN apk update && apk add --no-cache git
RUN apk add tzdata
WORKDIR /app
COPY . /app/
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o homies cmd/rest/main.go

FROM scratch
COPY --from=builder /app/homies .
COPY --from=builder /app/config.yaml .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Jakarta
EXPOSE 8099
CMD [ "/homies" ]