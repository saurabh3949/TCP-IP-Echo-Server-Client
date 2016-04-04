docker network create -d bridge sg-network
echo "Created Network bridge"
cd data && docker build -t data-image . && cd ..
docker run -d --name datacontainer data-image /bin/bash
echo "Datacontainer has been setup! Now starting server and client..."
cd server-client && docker build -t server-client-image . && cd ..
docker run -d --net=sg-network --volumes-from datacontainer --name catserver server-client-image go run server.go /data/string.txt 8000
docker run -d --net=sg-network --volumes-from datacontainer --name catclient server-client-image go run client.go /data/string.txt 8000
echo "Now printing client logs.."
docker logs -f catclient
