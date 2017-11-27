GIT_VERSION = $(shell git describe --always)
REPOSITORY_URL = uvduvd/aws-ecr-cleaner

build:
	docker run --rm \
	--volume ${CURDIR}:/go/src/aws-ecr-cleaner \
	--workdir /go/src/aws-ecr-cleaner \
 	--env CGO_ENABLED=0 \
 	--env GOOS=linux \
	golang \
	/bin/sh -c "go get -u github.com/aws/aws-sdk-go && go build -a -installsuffix cgo -o dist/aws-ecr-cleaner"

	docker build -t ${REPOSITORY_URL}:${GIT_VERSION} -t ${REPOSITORY_URL}:latest .

publish:
	docker push ${REPOSITORY_URL}:${GIT_VERSION}
	docker push ${REPOSITORY_URL}:latest