build:
	docker build -t another-cloud-project-api .

docker-push:
	docker tag another-cloud-project-api:latest  another-cloud-project-api:v1
	docker push krzysztofla/ziqq-docker-repo:v1
	