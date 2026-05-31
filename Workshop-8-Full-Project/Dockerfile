FROM  php:8.3.31-fpm

RUN apt update
RUN apt install -y libpq-dev unixodbc-dev gnupg zip unzip git
RUN docker-php-ext-install mysqli
RUN docker-php-ext-install pdo pdo_mysql
RUN docker-php-ext-enable pdo_mysql

WORKDIR /public

COPY index.php .

ENV TZ="Asia/Bangkok"

EXPOSE 8080

CMD ["php","-S","0.0.0.0:8080","-d","display_errors=1","-t","/public"]

