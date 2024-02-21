FROM golang:1.22
#LABEL authors="vladyslav.koshulko"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

ENV SERVER_IP="localhost"
ENV SERVER_PORT="8080"


CMD ["/docker-gs-ping"]



ENTRYPOINT ["top", "-b"]