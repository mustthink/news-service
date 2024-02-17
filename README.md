# news-service
simple application with CRUD operations for news

## Installation
Clone the repository via `git clone`
> If you have Go installed
> install dependencies via `go mod download`

## Running 

### Configuration
You can configure the application via environment variables. 
- `CONFIG_PATH` - path to the configuration file. 

You can configure the application via flags on the start:
- `--config` - path to the configuration file.

Config file structure:
```yaml
env: local # environment (local/test/prod)
server: # server configuration
  host: localhost # host for the server 
  port: 8080 # port for the server
db: # database configuration
  user: user # user for the database
  pass: pass # password for the database
  host: localhost # host for the database
  port: 5432 # port for the database
  db: default # database name
```


### 1) Run in Docker
1) Run PostgreSQL locally or via Docker. 
>You can use `task docker-postgres` or `docker-compose up postgres` for running.
2) Run application via command `task docker` which runs migrations and starts the application. 
3) Now you can access the application via `http://localhost:8080`  

### 2) Run via Go
1) Run PostgreSQL locally or via Docker. (check details in the previous section)
2) Run migrations via `task migrate`
3) Run application via `task run`
4) Now you can access the application via `http://localhost:8080`

## Testing
1) In Docker

run test container via `task docker-test`

2) Locally
run tests via `task test`
> tips:
> 1) Run service locally before testing for integration tests
> 2) After running tests, coverage will be shown in the browser


## API
Application has the following endpoints:
- `GET /news` - get all news. 
  - Parameters: none. 
  - Returns: array of news in JSON format.
- `GET /news/{id}` - get news by id (from URL)
  - Parameters: `id` — id of the news. 
  - Returns: `news` in JSON format.
- `POST /news` - create news (from request body)
  - Parameters: 
    - `title` — title of the news. It should be unique per author and be less than 255 characters. 
    - `content` — content of the news. 
    - `author_id` — id of the author.
    - `topic_id` — id of the topic.
  - Returns:
    - `id` of created news.
- `PUT /news/{id}` - update news by id (from URL)
  - Parameters: 
    - `id` — id of the news. 
    - `title` — title of the news. It should be unique per author and be less than 255 characters. 
    - `content` — content of the news. 
    - `author_id` — id of the author.
    - `topic_id` — id of the topic.
  - Returns:
    - `success` `true/false` if the news was updated.
- `DELETE /news/{id}` - delete news by id (from URL)
  - Parameters: 
    - `id` — id of the news. 
  - Returns:
    - `success` `true/false` if the news was deleted.
