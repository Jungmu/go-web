NAME=go-web
TAG=$1

rm $NAME

echo go build start
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o $NAME -a main.go
echo go build end

if [ "$(docker image ls -q ''$NAME'')" != "" ]
then
    docker rmi $(docker image ls -q ''$NAME'')
fi

echo docker build start
docker build \
--no-cache \
--file ./dockerfile \
--tag ''$NAME':'$TAG'' \
./
echo docker build end

if [ "$(docker ps --filter 'name='$NAME'' -q)" != "" ]
then
    docker rm -f $(docker ps --filter 'name='$NAME'' -q)
fi

echo docker run start
docker run \
-d \
--restart=always \
--name ''$NAME'_'$TAG'' \
--net host \
''$NAME':'$TAG''
echo docker run end

if [ "$(docker images -f "dangling=true" -q)" != "" ]
then
    docker rmi $(docker images -f "dangling=true" -q)
fi

