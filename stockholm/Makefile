NAME := stockholm

GOARCH=$(shell go env GOARCH)
GOOS=$(shell go env GOOS)

# Uncomment to build for a different OS
# GOOS=windows # (windows,linux,darwin,etc.)
# GOARCH=amd64 # (amd64,386,arm,etc.)

SRCS := stockholm.go

$(NAME):
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(NAME) $(SRCS)

clean:
	rm -f $(NAME)

re: clean $(NAME)

.PHONY: clean re