# build the Go API
FROM golang:latest AS builder
ADD . /app
WORKDIR /app/chat-server
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# build the React App
FROM node:alpine AS node_builder
COPY --from=builder /app/chat-ui .
RUN npm install
RUN npm run build

# copy the build assets to a minimal
# alpine image
FROM alpine:latest
ARG pwd
ENV WCG_PASSWORD=$pwd
COPY --from=builder /main .
COPY --from=node_builder /build ./chat-ui/build
RUN chmod +x ./main
EXPOSE 8080
CMD ./main
