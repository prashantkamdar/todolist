## A simple todolist in GO

#### Description
The application uses SQLite as its persistence layer.\
It consists of 2 tables with PK-FK relationship: `todolist` & `todos`.\
`main.go` initializes the db and inserts dummy data into it (data insertion commented).

#### Usage

##### *Create a todolist*
Path: `/todolist` \
Method: `POST` \
Data: `{"listName":"<list_name>"}` \
Sample call: `curl -X POST http://localhost:1337/todolist -d '{"listName":"Sample todo list"}'`

##### *Get all todolist items*
Path: `/todolist` \
Method: `GET` \
Sample call: `curl -X GET http://localhost:1337/todolist`

##### *Delete a todolist*
Path: `/todolist/{todolistid}` \
Method: `DELETE` \
URL params: `<todolistid>` \
Sample call: `curl -X DELETE http://localhost:1337/todolist/7`

##### *Create a task under one of the todolist item*
Path: `/todos/{todolistid}` \
Method: `POST` \
URL params: `<todolistid>` \
Data: `{"taskName": "<task_name>"}` \
Sample call: `curl -X POST http://localhost:1337/todos/1 -d '{"taskName": "Sample task name"}'`

##### *Update a todolist under one of the todolist item*
Path: `/todos` \
Method: `PUT` \
Data: `{"id": <todo_id>, "todoListId":<new_todo_list_id>, "taskName":"<new_task_name>"}` \
Sample call: `curl -X PUT http://localhost:1337/todos -d '{"id": 1, "todoListId":1, "taskName":"just updated this"}'`

##### *Delete a task under one of the todolist item*
Path: `/todos/{todolistid}` \
Method: `DELETE` \
URL params: `<todolistid>` \
Sample call: `curl -X DELETE http://localhost:1337/todos/7`

##### *Get all todos*
Path: `/todos` \
Method: `GET` \
Sample call: `curl -X GET http://localhost:1337/todos`

#### TODO
1. Use HTTP status codes for API return values
2. Correct the PK-FK relationship in SQLite
3. Create universal DB object and pass it around instead of initialising everytime
4. Add a free-to-use educational license
5. Lowercase the folder `DL`
6. Dockerize the app
7. Host the app on private VPS for working demo
8. Password protect the APIs
9. Add `bool` to track whether todo is completed or no