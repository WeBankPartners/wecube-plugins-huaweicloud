FROM webankpartners/alpine-ca-certificate
LABEL maintainer = "Webank CTB Team"

ENV APP_HOME=/home/app/wecube-plugins-huawecloud
ENV LOG_PATH=$APP_HOME/logs

RUN mkdir -p $APP_HOME  $LOG_PATH

COPY build/start.sh $APP_HOME/ 
COPY build/stop.sh $APP_HOME/ 
RUN chmod +x $APP_HOME/*.*

COPY wecube-plugins-huaweicloud $APP_HOME/

WORKDIR $APP_HOME

ENTRYPOINT ["/bin/sh", "start.sh"]
