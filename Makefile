docker-build:
	docker build -t webhook:v0.0.1 .

docker-run:
	docker run -it -p 8080:8080 --env-file .env webhook:v0.0.1 

generate-secret:
	kubectl create secret generic webhook-secret --dry-run=client --from-env-file=.env -o yaml > k8s/secret.yaml