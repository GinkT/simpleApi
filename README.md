### Запуск  
Для запуска достаточно развернуть докер и ввести команды:

`docker build -t testapi .`

`docker run --publish 8081:8081 testapi`  
  
После билда можно запускать приложение через приведенную выше команду _docker run_.

У проекта нет особой надобновсти в докере, можно просто склонить проект и запустить через  

`go run main.go json_responses.go`

### Что делает

Имеет простую ручку, которая выдаёт json с системной информацией и прогнозом погоды в МСК.

`http://localhost:8081/api/v1/`

