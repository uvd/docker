GIT_VERSION = $(shell git describe --always)
REPOSITORY_URL = uvdltd/terraform-helm

build:
	docker build -t ${REPOSITORY_URL}:${GIT_VERSION} -t ${REPOSITORY_URL}:0.11-2.7 .

publish:
	docker push ${REPOSITORY_URL}:${GIT_VERSION}
	docker push ${REPOSITORY_URL}:0.11-2.7
