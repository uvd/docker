FROM alpine:latest

RUN apk --no-cache --update add ca-certificates

ADD dist/aws-ecr-cleaner /aws-ecr-cleaner

ENTRYPOINT ["/aws-ecr-cleaner"]

ENV AWS_REGION eu-west-1
ENV AWS_ACCESS_KEY_ID changeme
ENV AWS_SECRET_ACCESS_KEY changeme