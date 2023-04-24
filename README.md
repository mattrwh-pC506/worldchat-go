# worldchat-go

This is an implementation of a single room chatroom using Go + Gorilla/Websockets for its server, and Typescript + React for its ui.

## Setup
To run this on your local machine, you can do one of the following:

### Dockerfile
This method is preferred:
```
# Build single image, for both server and ui
docker build -t chat .

# Run chat image in a container
docker run chat
```

### Makefile
If `make` is supported on your machine (as part of gcc/gnu compiler), you can do the following:
```
# Requirements Golang=1.20.3 + Node=18.16.0
make build
```

### Demo
https://user-images.githubusercontent.com/6784933/234136980-96829fe0-aefd-4eed-8942-b6060fb857af.mp4

