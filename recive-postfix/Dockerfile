FROM tozd/postfix:alpine-38

RUN mkdir -p /etc/skel/Maildir/{new,cur,tmp} && chmod -R 700 /etc/skel/Maildir/ && adduser taro -D -k /etc/skel