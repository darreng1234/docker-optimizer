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


RUN addgroup --gid "50000" "airflow" && \
    adduser --quiet "airflow" --uid "50000" \
        --gid "50000" \
        --home "/home/airflow"

ENV AIRFLOW_HOME=/home/airflow
ENV AIRFLOW_ENV=$env

RUN chown -R "airflow:root" "/home/airflow"

COPY requirements.txt /home/airflow/.

RUN pip install --no-cache-dir pip==21.3.1 wheel==0.37.1 && \
    pip install --no-cache-dir -r /home/airflow/requirements.txt