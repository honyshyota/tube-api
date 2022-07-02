<p align="center">
  <a href="" rel="noopener">
 <img width=400px height=400px src="https://github.com/honyshyota/tube-api/blob/master/images/super_gopher.webp" alt="Project logo"></a>
</p>


# Youtube API with Golang

## Task Description

- работа с АПИ стороннего сервиса (это может быть hh.ru, youtube, moex.com, любое)
- должно быть несколько связанных запросов (один запрос на получение чего-то, результат используется в другом запросе)

## How to

- Используйте make для запуска
 ```make docker```

*```http://localhost:8080/search``` для поиска каналов по ключу (key)
*```http://localhost:8080/video``` для поиска видео по id канала который подтягивается из БД (id)
*```http://localhost:8080/playlist``` для поиска плэйлистов по id канала который подтягивает из БД (id)

везде используется метод GET


## Look here

![alt text](https://github.com/honyshyota/tube-api/blob/master/images/docker_run.png)
![alt text](https://github.com/honyshyota/tube-api/blob/master/images/search_request.png)
![alt text](https://github.com/honyshyota/tube-api/blob/master/images/video_request.png)
![alt text](https://github.com/honyshyota/tube-api/blob/master/images/postgres_channel.png)
![alt text](https://github.com/honyshyota/tube-api/blob/master/images/postgres_video.png)
