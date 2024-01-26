https://www.youtube.com/watch?v=p08c0-99SyU&t=52s
1. create a docker file `touch Dockerfile`
2. nvim Dockerfile
``` Dockerfile
	FROM golang:1.19.0

	WORKDIR /usr/src/app
```
3. create a docker-compose.yml `touch docker-compose.yml`
4. nvim docker-compose.yml
``` docker-compose.yml
version: '3.8'

services:
  web:
    build: .
    ports:
      - "3000:3000"
    volumes:
      - .:/usr/src/app
```
5. in this part you can build this `sudo docker-compose up --build`
6. and with this part working ok. you can use bash of docker for work with go 1.19 !!!
