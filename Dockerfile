FROM alpine:3.15

ADD out/production/fdcbot_production /bot/fdcbot

CMD [ "/bot/fdcbot" ]