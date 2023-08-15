FROM alpine:latest

RUN mkdir /app
RUN mkdir /app/config

COPY userApp /app

COPY app/config/appConfigDev.yaml /app/config/

CMD [ "/app/userApp" ]