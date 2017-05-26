FROM java:openjdk-8-jre


ENV DEBIAN_FRONTEND noninteractive
ENV ELASTICSEARCH_VERSION 2.4.1

RUN apt-get update \
    && apt-get install -y curl \
    && apt-get clean

RUN set -x \
    && cd / \
    && mkdir /elasticsearch \
    && curl -O https://download.elastic.co/elasticsearch/release/org/elasticsearch/distribution/tar/elasticsearch/$ELASTICSEARCH_VERSION/elasticsearch-$ELASTICSEARCH_VERSION.tar.gz \
    && tar xf elasticsearch-$ELASTICSEARCH_VERSION.tar.gz -C /elasticsearch --strip-components=1 \
    && rm elasticsearch-$ELASTICSEARCH_VERSION.tar.gz

RUN mkdir -p /elasticsearch/config/templates
COPY template-k8s-logstash.json /elasticsearch/config/templates/template-k8s-logstash.json
COPY template-k8s-kube-logstash.json /elasticsearch/config/templates/template-k8s-kube-logstash.json

COPY config /elasticsearch/config

COPY run.sh /
COPY elasticsearch_logging_discovery /

RUN useradd --no-create-home --user-group elasticsearch \
    && mkdir /data \
    && chown -R elasticsearch:elasticsearch /elasticsearch

RUN echo 'elasticsearch - nofile 65536' >> /etc/security/limits.conf
RUN echo 'elasticsearch - memlock unlimited' >> /etc/security/limits.conf

VOLUME ["/data"]
EXPOSE 9200 9300

CMD /run.sh
