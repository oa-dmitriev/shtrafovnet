all: run

run: .
	docker-compose build
	docker-compose up
