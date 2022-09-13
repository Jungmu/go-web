NAME=go-web
TAG=$1

rm $NAME
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $NAME -a main.go

if [ "$(docker image ls -q ''$NAME'')" != "" ]
then
    docker rmi $(docker image ls -q ''$NAME'')
fi

docker build \
--no-cache \
--file ./dockerfile \
--tag ''$NAME':'$TAG'' \
./

if [ "$(docker ps --filter 'name='$NAME'' -q)" != "" ]
then
    docker rm -f $(docker ps --filter 'name='$NAME'' -q)
fi

docker run \
-d \
--restart=always \
--name ''$NAME'_'$TAG'' \
--net host \
''$NAME':'$TAG''

if [ "$(docker images -f "dangling=true" -q)" != "" ]
then
    docker rmi $(docker images -f "dangling=true" -q)
fi

