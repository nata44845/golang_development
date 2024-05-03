# 03-01-migration

В этом задании нам нужно поработать с миграциями. Для начала установите следующие базы данных Postgres, MongoDB и Redis и запустите их.
Можно воспользоваться ссылками из лекции. Нам неважно как вы запустить эти DB через системы пакетов или docker. Я
исхожу из того, что вам знакомы основы docker.

Я прилагаю вариант через docker, но можно воспользоваться любым удобным способом. Для начала установить docker любым
удобным способом для вашей [ОС](https://www.docker.com/get-started/)

Обратите внимание, в команде запуска сделана привязка портов. Нам нужно будет обращаться на порт 5434 в случае
работы с postgres и 6381 для работы с redis, 27018 для mongodb Это сделано, чтобы при запуске не было конфликтов с
вашими прошлыми экспериментами.

Поднять необходимые DB нам нужно для выполнения ДЗ текущего и последующих.

```shell
docker pull postgres
docker run --name gb-workshop-pg -e POSTGRES_DB=university -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d 
-p 5434:5432 postgres

docker pull redis
docker run --name gb-workshop-redis -d -p 6381:6379 redis

docker pull mongo
docker run --name gb-workshop-mongodb -d -p 27018:27017 mongo 
```

Теперь, чтобы получить доступ к интерактивной консоли контейнера можно также запустить cli

Postgres
```shell
docker exec -it gb-workshop-pg psql -U postgres -d university
```
В интерфейсе psql можно выполнить \d чтобы посмотреть схему данных

Redis
```shell
docker exec -it db-workshop-redis redis-cli
```


Mongo

```shell
docker exec -it gb-workshop-mongodb mongosh "mongodb://localhost:27017"
```

Установить приложение go migrate одним из способов

[go migrate cli](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
Также рекомендую сразу установить к ней драйверы для работы с БД
```shell
go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

Теперь попробуем написать миграцию для учебной схемы данных

```sql
CREATE TABLE Students (
    StudentID INT PRIMARY KEY,
    FirstName VARCHAR(50),
    LastName VARCHAR(50),
    EnrollmentDate DATE
);

CREATE TABLE Professors (
    ProfessorID INT PRIMARY KEY,
    FirstName VARCHAR(50),
    LastName VARCHAR(50),
    Department VARCHAR(50)
);

CREATE TABLE Courses (
    CourseID INT PRIMARY KEY,
    CourseName VARCHAR(100),
    Department VARCHAR(50),
    Credits INT,
    ProfessorID INT,
    FOREIGN KEY (ProfessorID) REFERENCES Professors(ProfessorID)
);

CREATE TABLE Grades (
    GradeID INT PRIMARY KEY,
    StudentID INT,
    CourseID INT,
    Grade VARCHAR(2),
    FOREIGN KEY (StudentID) REFERENCES Students(StudentID),
    FOREIGN KEY (CourseID) REFERENCES Courses(CourseID)
);

```

Вам потребуется команда  `migrate create` и `migrate up`

После того, как сделаете ваш первую миграцию, давайте сделаем еще 2.
Далее добавим индекс в таблице Courses по professorID и составной индекс StudentID, CourseID для Grades. Каждая
операция в отдельной миграции.

Все должно отрабатывать без ошибок, если вы сделаете migrate down, migrate up все таблицы должны раскатится
автоматически.