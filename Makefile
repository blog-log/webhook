docker-build:
	docker build -t webhook:v0.0.1 .

docker-run:
	docker run -it -p 8080:8080 --env-file .env webhook:v0.0.1 