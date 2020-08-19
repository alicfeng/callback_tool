FROM alpine:3.7

COPY release/callback_tool_unix /usr/local/sbin/callback_tool

#server default port
ENV PORT 80
# server default host
ENV HOST "0.0.0.0"
# server default route
ENV ROUTE "/api/debugging/callback"
# server default output
ENV OUTPUT "/var/log/callback_tool"

EXPOSE ${PORT}

RUN mkdir /var/log/callback_tool

WORKDIR /var/log/callback_tool

CMD ["/usr/local/sbin/callback_tool","-h","${HOST}","-p","${PORT}","-r","${ROUTE}","-o","${OUTPUT}"]