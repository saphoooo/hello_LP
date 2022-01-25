FROM ubuntu:20.04

LABEL maintainer="stephane.beuret@gmail.com"

RUN apt-get update -y && \
    apt-get install -y python3-pip python3.8-venv

RUN python3 -m venv venv
RUN . venv/bin/activate
RUN pip install Flask

COPY . .

ENV FLASK_APP=hello
ENTRYPOINT [ "flask", "run" ]