FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o k8s-example-app .

FROM scratch

WORKDIR /app

COPY --from=builder /app/k8s-example-app /app

COPY --from=builder /app/app.yaml /app

EXPOSE 3000

ENTRYPOINT ["/app/k8s-example-app"]
