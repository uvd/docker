FROM fluent/fluentd-kubernetes-daemonset:v0.12-alpine-cloudwatch

RUN apk add --no-cache ruby-dev build-base && \
    fluent-gem install \
    fluent-plugin-slack \
    fluent-plugin-grepcounter \
    fluent-plugin-mail \
    && apk del build-base ruby-dev
