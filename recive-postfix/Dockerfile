FROM tozd/postfix:alpine-316

RUN mkdir -p /etc/skel/Maildir/{new,cur,tmp} && chmod -R 700 /etc/skel/Maildir/ && adduser taro -D -k /etc/skel
COPY ./main.cf /etc/postfix/main.cf