# worldchat-go

This is an implementation of a single room chatroom using Go + Gorilla/Websockets for its server, and Typescript + React for its ui.

### Demo
https://user-images.githubusercontent.com/6784933/234136980-96829fe0-aefd-4eed-8942-b6060fb857af.mp4

## Setup

### Makefile
If `make` is supported on your machine (as part of gcc/gnu compiler), you can do the following:
```
# Requirements Golang=1.20.3 + Node=18.16.0
WCG_PASSWORD=YOUR_PASSWORD make build
```

### Notes/Quirks/Things I would do if I had all the time in the world :)
- This app has a naive approach to password management, in that when the server is started, it uses whatever value you've set for WCG_PASSWORD. It then generates a secret token from that password value to authorize logins and auth checks. For a real application, we would handle this with an actual secret key, stored in a secured key vault, and we would also have the concept of true users, with their own usernames and passwords.
- Unfortunately, two unit tests (one in the react app, and one in the go webserver), both of which attempt to test the sending of messages via our websocket, had to be commented out due to the inherent trickiness of testing websocket behavior, and of course, time constraints.
- Since there is no concept of users in this application, only active "clients", we assume that all incoming messages are from external users. We use this to distinguish current user messages from external user messages so that we can style them differently in the chat thread. As a side effect of that, when you refresh the page, all messages, now read from the server, are treated as external user messages (even the current user's messages), which means they all get pushed to the right and given the external user ux treatment. If I had more time, I would have encoded these messages with user information so that they could be distinguished beyond one session.
- At some point I'd like to refactor this to use gosf instead of Gorilla/Websockets, which is no longer actively supported.
