# todo note commands
docker swarm init 
docker swarm leave --force
docker node ls
docker stack rm sbd
docker volume rm sbd_order_pg_vol sbd_minio_vol
docker stack deploy -c docker-compose.yml sbd
docker service ls
docker service logs sbd_postgres
docker service ps sbd_postgres --no-trunc