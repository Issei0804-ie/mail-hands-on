FROM ubuntu

RUN sed -i.bak -e "s%http://[^ ]\+%http://mirror.yuki.net.uk/ubuntu-ports/ubuntu-ports/%g" /etc/apt/sources.list

RUN apt-get update -y && apt-get install golang -y