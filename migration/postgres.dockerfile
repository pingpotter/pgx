FROM postgres:11.2-alpine

COPY migration/postgres/00_create_db.sh /docker-entrypoint-initdb.d/00_init.sh
COPY migration/postgres/*.sql /docker-entrypoint-initdb.d/

# Set host time zone 
ENV TZ=Asia/Bangkok
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone