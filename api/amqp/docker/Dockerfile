FROM rabbitmq:3-management-alpine
COPY api/amqp/docker/rabbitmq.conf /etc/rabbitmq/
ENV RABBITMQ_NODENAME=rabbit@localhost
RUN chown rabbitmq:rabbitmq /etc/rabbitmq/rabbitmq.conf
USER rabbitmq:rabbitmq
