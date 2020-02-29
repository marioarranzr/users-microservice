# Users microservice

Write a small Users microservice.

Available languages are Java (preferred), GO or Node.js. The user entity consists of:
- First name
- Last name
- Nickname
- Password
- email
- country

The service must allow to:
- add a new user
- modify an existing user
- remove a user
- return the list of the users satisfying certain criteria (e.g. for country)

The microservice will be part of a more complex architecture, so consider for example that the Search microservice will need to be notified when a new user is added, or that the Competition microservice will need to be notified when the user changes his nickname. Think at how to implement a system that is scalable.

The application must be a “good citizen”:
- Meaningful logs
- Self-documented end points
- Health checks
- ...

You can “mock” the database if you like (e.g. saving the data in memory).
Please provide the instructions to start the application on localhost.
We expect to be able to run unit tests and to be able to add/modify/delete/list the users by calling the end points using http calls after starting the application.
Please explain what are the criteria and assumptions you used to take decisions. A clear and correct explanation is part of the test.


## REST API

This API must comply with the following contract:
### GET /health

Responses:

* **200 OK** When the application is running.

### POST /users

**Body** _required_ User to load in the system.

Sample:

```json
{
  "first_name": "Mario",
  "last_name": "Arranz",
  "email": "mario@omg.lol",
}
```

Responses:

* **202 Accepted** When the user is registered correctly.
* **400 Bad Request** When there is a failure in the request format, expected

### GET /users

Return the user that fulfils the requirements.

Responses:

* **200 OK** When the user is registered correctly.
```json
{
  "first_name": "Mario",
  "last_name": "Arranz",
  "nickname": "marioarranzr",
  "password": "******",
  "email": "mario@omg.lol",
  "country": "Spain"
}

```
* **400 Bad Request** When there is a failure in the request format.
* **404 Not Found** When there is no user with the parameters requested.

# Solution

Run locally in docker:
```
make run-locally
```
Run locally (not using docker):
```
make debug
```
Run tests with coverage:
```
make test
```

The app will run in port `9091`


### Briefing about the solution:

- The in-memory database is simply a list of users.
- When a user is inserted, the system check it is not already in the database

- Unit tests for repository methods and one extra that tests the complete integration. 
- Unit tests for service
