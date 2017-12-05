FROM uvdltd/kubectl:1.8

USER root
ENV KOPS_VERSION 1.8.0
ADD https://github.com/kubernetes/kops/releases/download/1.8.0/kops-linux-amd64 /bin/kops
RUN chmod +x /bin/kops
USER kube

VOLUME [ "/kube/.aws" ]

ENTRYPOINT ["/bin/kops"]
