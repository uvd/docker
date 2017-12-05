FROM uvdltd/kubectl:1.8

RUN apk add --no-cache ca-certificates
ENV KOPS_VERSION 1.8.0
ADD https://github.com/kubernetes/kops/releases/download/1.8.0/kops-linux-amd64 /bin/kops
RUN chmod +x /bin/kops

VOLUME [ "/kube/.aws" ]

ENTRYPOINT ["/bin/kops"]
