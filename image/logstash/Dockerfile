FROM docker.elastic.co/logstash/logstash:6.3.1
RUN rm -f /usr/share/logstash/pipeline/logstash.conf
ADD pipeline/ /etc/pipeline
ADD config/ /usr/share/logstash/config/
