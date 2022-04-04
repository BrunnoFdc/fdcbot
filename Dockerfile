FROM ubuntu:20.04

ADD out/production/fdcbot_production /bot/fdcbot

CMD [ "/bot/fdcbot" ]
