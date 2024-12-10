FROM alpine:latest

RUN apk update && \
    apk upgrade && \
    apk add bash && \
    rm -rf /var/cache/apk/*

ADD https://github.com/pressly/goose/releases/download/v3.23.0/goose_linux_x86_64 /bin/goose
RUN chmod +x /bin/goose

WORKDIR /workspace

ADD migrations/*.sql migrations/
ADD scripts/migration.sh .
ADD .env* .

# Use appropriate .env file: if .env does not exist, use .env.default
RUN if [ -f .env ]; then echo ".env file exists, using .env"; \
    else cp .env.default .env && echo ".env file not found, using .env.default"; fi

RUN chmod +x migration.sh

ENTRYPOINT ["bash", "migration.sh"]