## A simple todolist in GO

#### Description
The application uses SQLite as its persistence layer.\
It consists of 2 tables with PK-FK relationship: `todolist` & `todos`.\
`main.go` initializes the db and inserts dummy data into it (data insertion commented).

#### Usage

##### Create a todolist
Path: `/todolist` \
Method: `POST` \
Data: `{"listName":"<list_name>"}` \
Sample call: `curl -X POST http://localhost:1337/todolist -d '{"listName":"Sample todo list"}'`

##### *Get all todolist items*

##### Create a task under one of the todolist item

##### Update a todolist under one of the todolist item

##### Delete a task under one of the todolist item

##### Delete a todolist

#### TODO
1. Use HTTP status codes for API return values
2. Correct the PK-FK relationship in SQLite
3. Create universal DB object and pass it around instead of initialising everytime
4. Add a free-to-use educational license
5. Lowercase the folder `DL`
6. Dockerize the app
7. Host the app on private VPS for working demo
8. Password protect the APIs 