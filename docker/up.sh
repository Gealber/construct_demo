source deploy-config.sh 

docker-compose -p boukker -f docker-compose.yml up -d

#docker exec -i  boukker_v1Server_1 sh -c 'service memcached start'
