BUILD_DIR := dist/
APP_NAME := asiayo

.PHONY: doc build clean

all: doc build

doc:
	swag init

build:
	 go build -o $(BUILD_DIR)$(APP_NAME).exe

clean:
	rm -f $(APP_NAME)
