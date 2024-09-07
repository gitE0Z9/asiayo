BUILD_DIR := dist/
APP_NAME := asiayo

build:
	swag init && go build -o $(BUILD_DIR)$(APP_NAME)

clean:
	rm -f $(APP_NAME)
