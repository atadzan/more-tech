### Project for hackaton More Tech 5.0 organized by ВТБ
#### Service gives information about atms and offices with their detailed info.
<br><b>Used language and technologies:</b>
* GoLang 
* PostgreSQL as database

To run app you should check, if you have installed ``docker-compose`` and ``make``. Because third_party technology run in container.<br>
Run dependencies:<br> 

```make run-compose``` <br>

Run app:<br>

``make run-app`` <br>

If you don't have ``make`` tool installed, use commands below to run container and app <br>
Run container: <br>
``docker-compose -f ./deployment/docker-compose.yaml up -d`` <br>
Run app: <br>
``go run cmd/main.go``
