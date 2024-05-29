FROM alpine:20231219

# Set the appropriate timezone
ENV TZ="Europe/Lisbon"


# Install the required packages
RUN apk --no-cache add bash zsh go make \
  && rm -rf /tmp/*


# Copy the source files
COPY ../src /printer-src


# Add the entrypoint
USER $UNAME
COPY ./start.sh /start.sh
STOPSIGNAL SIGTERM
ENTRYPOINT [ "/start.sh" ]
