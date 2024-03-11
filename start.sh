source /etc/profile
go build
cd fixtures
docker-compose down -v
docker-compose up -d
cd ..
./fabric-go-sdk