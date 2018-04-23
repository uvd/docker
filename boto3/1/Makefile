GIT_VERSION = $(shell git describe --always)
REPOSITORY_URL = uvdltd/boto3

build:
	docker build -t ${REPOSITORY_URL}:${GIT_VERSION} -t ${REPOSITORY_URL}:1 .

publish:
	docker push ${REPOSITORY_URL}:${GIT_VERSION}
	docker push ${REPOSITORY_URL}:1
