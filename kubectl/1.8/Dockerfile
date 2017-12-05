FROM alpine:3.7

ENV VERSION "v1.8.0"

ADD https://storage.googleapis.com/kubernetes-release/release/${VERSION}/bin/linux/amd64/kubectl /bin/kubectl
RUN chmod +x /bin/kubectl

ENV HOME "/home/kube"
VOLUME ["/home/kube/.kube"]

ENTRYPOINT ["/bin/kubectl"]
