FROM golang:1.14-alpine as builder
WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o igor .

FROM golang:1.14-alpine
WORKDIR /app

COPY --from=builder /app/igor /app

ENTRYPOINT [ '/app/igor' ]
