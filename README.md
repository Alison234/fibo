# fibo
### Установка memcashe 
sudo apt install memcached libmemcached-tools
###  Билд сервера 
make build
###  запуск http серва
./apiserver --serverType http
###  запуск grpc серва
./apiserver --serverType grpc
### для запуска grpc клиента 
 перейти в fibo/grpc/proto/pkg/client выполить go build && ./client
 ### Примеры отправки curl запроса 
 curl -X POST -H "Content-Type: application/json"     -d '{"StartIndex": 0, "EndIndex": 10}'     http://localhost:8080/fromTo   
 curl -X POST -H "Content-Type: application/json"     -d '{"EndIndex": 10}'     http://localhost:8080/to
