FROM golang:1.19.4-bullseye AS build
WORKDIR /app/src/api-go-gin
ENV GOPATH=/app
COPY . .
RUN go mod download
RUN go build -o FoobarGoApp

FROM gcr.io/distroless/base-debian11 AS deploy
WORKDIR /
COPY --from=build /app/src/api-go-gin/FoobarGoApp ./
EXPOSE 9000
USER nonroot:nonroot
ENTRYPOINT ["/FoobarGoApp"]

FROM deploy AS final