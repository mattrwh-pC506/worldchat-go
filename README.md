# worldchat-go

This is an implementation of a single room chatroom using Go + Gorilla/Websockets for its server, and Typescript + React for its ui.

### Demo
https://user-images.githubusercontent.com/6784933/234136980-96829fe0-aefd-4eed-8942-b6060fb857af.mp4

## Setup

### Dockerfile
This can be run with docker using the following commands:
```
export WCG_PASSWORD=<YOUR_PASSWORD>
docker build -t worldchat-go --build-arg pwd=$WCG_PASSWORD .
docker run -p 8080:8080 --name=worldchat-go-container worldchat-go
```

### Makefile
If `make` is supported on your machine (as part of gcc/gnu compiler), you can do the following:
```
# Requirements Golang=1.20.3 + Node=18.16.0
WCG_PASSWORD=<YOUR_PASSWORD> make build
```

### Using the application
The password will be whatever you set in <YOUR_PASSWORD> in the setup steps above. Once you are in that chat, you can type and submit messages, either by clicking Send or hitting the Enter key. As a new user, logging in, you should see a list of existing messages for that server session. A logout button was added for convenience, since our JWT token is stored in localstorage and we'd otherwise need to clear that or wait till the JWT expires to test subsequent logins.

### Requirements

- [x] The application should support only one chat room.
- [x] Upon entering the chat room, the [user must provide a password](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-ui/src/routes/Login.test.tsx) that is validated on the [server side](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-server/handlers/login_handler.go). If the password is incorrect, the user can't access the chat.
    - NOTE: before determining if the user should go to the login screen, I am doing an auth check [here](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-ui/src/routes/App.tsx#L6), [here](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-ui/src/routes/router.tsx#L27), [here](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-ui/src/auth.ts#L3) and [here](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-server/handlers/auth_handler.go#L1)
- [x] When the user enters the message, it is immediately shown to other users in the chat.
- [x] New users should have a history of messages available to them. There is no requirement to persist the history in case of a server restart. ([here](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-server/handlers/room.go#L63))
- [x] In case of errors, the user should receive the message. The error should also be logged on the server side. (e.g BE [here](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-server/handlers/chat_handler.go#L58), FE [here](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-ui/src/routes/Chat.tsx#L70) and [here](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-ui/src/components/MessageList.tsx#L128))
- [x] Backend - Go, Frontend - Typescript/Javascript

### Notes/Quirks/Things I would do if I had all the time in the world :)
- This app has a naive approach to password management, in that when the server is started, it uses whatever value you've set for WCG_PASSWORD. It then generates a secret token from that password value to authorize logins and auth checks. For a real application, we would handle this with an actual secret key, stored in a secured key vault, and we would also have the concept of true users, with their own usernames and passwords.
- Unfortunately, two unit tests (one in the [react app](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-ui/src/routes/Chat.test.tsx#L41), and one in the [go webserver](https://github.com/mattrwh-pC506/worldchat-go/blob/main/chat-server/tests/chat_handler_test.go#L32)), both of which attempt to test the sending of messages via our websocket, had to be commented out due to the inherent trickiness of testing websocket behavior, and of course, time constraints.
- Since there is no concept of users in this application, only active "clients", we assume that all incoming messages are from external users. We use this to distinguish current user messages from external user messages so that we can style them differently in the chat thread. As a side effect of that, when you refresh the page, all messages, now read from the server, are treated as external user messages (even the current user's messages), which means they all get pushed to the right and given the external user ux treatment. If I had more time, I would have encoded these messages with user information so that they could be distinguished beyond one session.
- At some point I'd like to refactor this to use gosf instead of Gorilla/Websockets, since the latter is no longer actively supported and gosf seems to have more idiomatic patterns.
