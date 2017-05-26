FROM gcr.io/google_containers/ubuntu-slim:0.4

# Ensure there are enough file descriptors for running Fluentd.
RUN ulimit -n 65536

# Disable prompts from apt.
ENV DEBIAN_FRONTEND noninteractive

# Copy the Fluentd configuration file.
COPY td-agent.conf /etc/td-agent/td-agent.conf

COPY build.sh /tmp/build.sh
RUN /tmp/build.sh

ENV LD_PRELOAD /opt/td-agent/embedded/lib/libjemalloc.so

# Run the Fluentd service.
ENTRYPOINT ["td-agent", "-p=/etc/fluent/plugin"]
