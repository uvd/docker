FROM uvdltd/kubectl:1.8

ENV HELM_VERSION "v2.7.2"
ADD https://storage.googleapis.com/kubernetes-helm/helm-${HELM_VERSION}-linux-amd64.tar.gz /helm.tar.gz
RUN tar -xzvf /helm.tar.gz && \
    mv linux-amd64/helm /bin/helm && \
    chmod +x /bin/helm && \
    rm -rf /linux-amd64 /helm.tar.gz

ENTRYPOINT ["/bin/helm"]
