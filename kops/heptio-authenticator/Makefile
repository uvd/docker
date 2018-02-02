GIT_VERSION = $(shell git describe --always)
REPOSITORY_URL = uvdltd/kops-heptio-authenticator

build:
	docker build -t ${REPOSITORY_URL}:${GIT_VERSION} -t ${REPOSITORY_URL}:latest .

publish:
	docker push ${REPOSITORY_URL}:${GIT_VERSION}
	docker push ${REPOSITORY_URL}:latest
