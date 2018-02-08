FROM fluent/fluentd-kubernetes-daemonset:v0.12-alpine-cloudwatch

RUN apk add --no-cache ruby-dev build-base git && \
    git clone https://github.com/MerlinDMC/fluent-plugin-cloudwatch-logs.git && \
    cd fluent-plugin-cloudwatch-logs/ && git checkout master && \
    gem build fluent-plugin-cloudwatch-logs.gemspec && \
    gem uninstall fluent-plugin-cloudwatch-logs && \
    gem install --no-document --no-ri *.gem && \
    gem install --no-document --no-ri \
    fluent-plugin-slack \
    fluent-plugin-grepcounter \
    fluent-plugin-mail \
    fluent-plugin-rewrite-tag-filter \
    && apk del build-base ruby-dev curl git
