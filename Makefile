# Define variables
REACT_APP_NAME := chat-ui
GO_APP_NAME := chat-server
GO_APP_BINARY := $(GO_APP_NAME).exe

# Build the React app
.PHONY: react-build
react-build:
	cd $(REACT_APP_NAME) && npm install && npm run build

# Build the Go app
.PHONY: go-build
go-build:
	cd $(GO_APP_NAME) && go build -o $(GO_APP_BINARY)

# Run the Go server
.PHONY: go-run
go-run:
	cd $(GO_APP_NAME) && $(GO_APP_BINARY)

# Build and run the app
.PHONY: build
build: react-build go-build go-run