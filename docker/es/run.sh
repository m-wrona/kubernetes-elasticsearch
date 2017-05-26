export NODE_NAME=${NODE_NAME:-${HOSTNAME}}
export NODE_MASTER=${NODE_MASTER:-true}
export NODE_DATA=${NODE_DATA:-true}
export HTTP_PORT=${HTTP_PORT:-9200}
export TRANSPORT_PORT=${TRANSPORT_PORT:-9300}
export MINIMUM_MASTER_NODES=${MINIMUM_MASTER_NODES:-2}
export ES_JAVA_OPTS=${ES_JAVA_OPTS:-'-Xms4g -Xmx4g'}
export NAMESPACE=elk

/elasticsearch_logging_discovery >> /elasticsearch/config/elasticsearch.yml

chown -R elasticsearch:elasticsearch /data

/bin/su -c /elasticsearch/bin/elasticsearch elasticsearch
