build-image:
	docker build -t luizhpferreira/finance .


run-app:
	docker-compose -f .devops/app.yml up -d

