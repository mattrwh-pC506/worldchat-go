FROM node:14.16.1 AS react-builder
ENV REACT_APP_NAME=chat-ui
WORKDIR /app
COPY $REACT_APP_NAME/package.json .
COPY $REACT_APP_NAME/package-lock.json .
RUN npm install
COPY $REACT_APP_NAME .
RUN npm run build

FROM golang:1.16 AS go-builder
ENV GO_APP_NAME=chat-server
WORKDIR /app
COPY $GO_APP_NAME .
RUN go build -o $GO_APP_NAME

FROM alpine:3.14
WORKDIR /app
COPY --from=react-builder /app/build /app/build
COPY --from=go-builder /app/$GO_APP_NAME /app/$GO_APP_NAME

CMD ["./chat-server"]
