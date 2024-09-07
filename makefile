APP_NAME := asiayo

build:
	swag init && go build -o $(APP_NAME)

clean:
	rm -f $(APP_NAME)
