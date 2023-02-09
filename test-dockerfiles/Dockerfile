FROM asia.gcr.io/co-169315/base/python:3.6-slim

ARG env

RUN apt-get update -y && apt-get install -y \
    libczmq-dev \
    libssl-dev \
    inetutils-telnet \
    bind9utils \
    gcc \
    nano \
    curl \
    htop \
    libpq-dev \
    default-libmysqlclient-dev \
    curl \
    && apt-get clean


