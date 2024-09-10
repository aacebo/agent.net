# Based on GithubGist https://gist.github.com/jgould22/3280fc0f531485f4fe19a2ef1ef67361#file-dockerfile
FROM postgres:alpine

ENV PG_PARTMAN_VERSION v5.1.0

# Install pg_partman
RUN set -ex \
    && apk update \
    && apk upgrade \
    # Get the depends required to build pg_jobmon and name this set of depends build-deps so we can delete them later
    && apk add --no-cache --virtual .build-deps \
        autoconf \
        automake \
        g++ \
        clang15 \
        llvm \
        libtool \
        libxml2-dev \
        make \
        perl \
    && make --version \
    # Download pg_partman
    && wget -O pg_partman.tar.gz "https://github.com/pgpartman/pg_partman/archive/$PG_PARTMAN_VERSION.tar.gz" \
    # Create a folder to put the src files in
    && mkdir -p /usr/src/pg_partman \
    # Extract the src files
    && tar \
        --extract \
        --file pg_partman.tar.gz \
        --directory /usr/src/pg_partman \
        --strip-components 1 \
    # Delete src file tar
    && rm pg_partman.tar.gz \
    # Move to src file folder
    && cd /usr/src/pg_partman \
    # Install the extension
    && make NO_BGW=1 install \
    # Delete the src files for pg_partman
    && rm -rf /usr/src/pg_partman \
