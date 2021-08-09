.PHONE: all

all: docker go

go:
	@printf  "RUN SERVER \n\n"
	go run main.go

docker:
	@printf  "RUN DOCKER \n\n"
	cd resources/docker && docker-compose up -d

