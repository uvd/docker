FROM golang:1.9

ENV CGO_ENABLED=0
RUN mkdir -p $GOPATH/src/github.com/mcuadros && \
    git clone https://github.com/mcuadros/terraform-provider-helm.git $GOPATH/src/github.com/mcuadros/terraform-provider-helm && \
    cd $GOPATH/src/github.com/mcuadros/terraform-provider-helm && \
    GOOS=linux go build -a -installsuffix cgo .

FROM hashicorp/terraform:0.11.3

RUN mkdir -p /root/.terraform.d/plugins/
COPY --from=0 /go/src/github.com/mcuadros/terraform-provider-helm /root/.terraform.d/plugins/terraform-provider-helm

ENV HELM_VERSION "v2.7.2"
ADD https://storage.googleapis.com/kubernetes-helm/helm-${HELM_VERSION}-linux-amd64.tar.gz /helm.tar.gz
RUN tar -xzvf /helm.tar.gz && \
    mv linux-amd64/helm /bin/helm && \
    chmod +x /bin/helm && \
    rm -rf /linux-amd64 /helm.tar.gz

ENV VERSION "v1.8.0"

ADD https://storage.googleapis.com/kubernetes-release/release/${VERSION}/bin/linux/amd64/kubectl /bin/kubectl
RUN chmod +x /bin/kubectl
