FROM alpine:latest

RUN mkdir /app

COPY teamsApp /app

CMD [ "/app/teamsApp"]