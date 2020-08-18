FROM alpine:3.7

COPY release/callback_tool_unix /usr/local/sbin/callback_tool

EXPOSE 80

RUN mkdir /var/log/callback_tool

WORKDIR /var/log/callback_tool

CMD ["/usr/local/sbin/callback_tool","-h","0.0.0.0","-p","80","-r","/api/debugging/callback","-o","/var/log/callback_tool"]