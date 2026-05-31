# Full Project

- This project auto load sql "initdb.sql"

- ingress config "default.conf"

- "Dockerfile" for build custom API

How to run

```sh
git clone https://github.com/wachira90/docker-getstarted.git

cd Workshop-8-Full-Project

#connect-no-db
docker stack deploy -c project.yml app

#connect-db
docker stack deploy -c test.yml app

#deploy-app
docker stack deploy -c test.yml app

docker stack ps app

docker stack services app

#url
WEB-URL: http://app.127-0-0-1.nip.io       #for-web
API-URL: http://app.127-0-0-1.nip.io/api   #for-api
DB-URL: http://db.127-0-0-1.nip.io         #for-db-manage

#update-version
docker service update --image myphp:dev1.0.0 app_api
docker service update --image myphp:dev1.0.1 app_api

#destroy
docker stack rm app
```

## Appendix

```sh
#start-cmd
CMD ["php","-S","0.0.0.0:8080","-c","php.ini","-d","display_errors=1","-t","/public"]
CMD ["php","-S","0.0.0.0:8080","-c","-d","display_errors=1","-t","/public"]

#test-cmd
docker run -it --name myphp -p 8011:8080 docker.io/library/myphp:v1

#default-php.ini
/usr/local/etc/php/conf.d/php.ini

#example-ini
/usr/local/etc/php/conf.d/php.ini-development
/usr/local/etc/php/conf.d/php.ini-production

#additional-config
/usr/local/etc/php/conf.d/docker-fpm.ini

#build-cmd
docker buildx build -t myphp:v1 . --no-cache
docker run -it --name myphp -p 8011:8080 docker.io/library/myphp:v1
```