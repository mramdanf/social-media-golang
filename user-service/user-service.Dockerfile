FROM alpine:latest

RUN mkdir /app
RUN mkdir /app/config
RUN mkdir /app/scripts

COPY --from=ghcr.io/ufoscout/docker-compose-wait:latest /wait /wait

COPY userApp /app

COPY scripts/user.sql /app/scripts

COPY app/config/appConfigDev.yaml /app/config/

CMD [ "/app/userApp" ]