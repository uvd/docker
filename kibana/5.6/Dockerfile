FROM docker.elastic.co/kibana/kibana:5.6.4

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT [ "/entrypoint.sh" ]
CMD []

ENV KIBANA_PLUGINS "https://github.com/sirensolutions/sentinl/releases/download/tag-5.6.2/sentinl-v5.6.4.zip;https://github.com/sivasamyk/logtrail/releases/download/v0.1.23/logtrail-5.6.4-0.1.23.zip"
