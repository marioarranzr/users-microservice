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

# Solution

## REST API

This API must comply with the following contract:
### GET /health

Responses:

* **200 OK** When the application is running.

### GET /?parameters

Return the user that matches the requirements.

Sample:

`?firstName=Mario&lastName=Arranz`

Responses:

* **200 OK** When there are one or more users that match.
```json
[
    {
        "first_name": "Mario",
        "last_name": "Arranz",
        "nickname": "marioarranzr",
        "password": "*****",
        "email": "mario@omg.lol",
        "country": "Spain"
    }
]

```
* **400 Bad Request** When there is a failure in the request format expected.
* **404 Not Found** When there is no user matching the parameters requested.

### POST /

**Body** _required_ User to load in the system. Nickname mandatory.

Sample:

```json
{
  "first_name": "Mario",
  "last_name": "Arranz",
  "nickname": "marioarranzr",
  "email": "mario@omg.lol",
  "country": "Spain"
}
```

Responses:

* **201 Created** When the user is registered correctly.
* **400 Bad Request** When there is a failure in the request format expected.
* **409 Conflict** When there is a user with the same nickname already in the system.

### Put /

**Body** _required_ User to modify. It will search by nickname and apply the changes to that user.

Sample:

```json
{
  "first_name": "Mario",
  "last_name": "Arranz",
  "nickname": "marioarranzr",
  "email": "mario@omg.lol",
  "password": "LolLolLol",
  "country": "Spain"
}
```

Responses:

* **200 OK** When the user is modified correctly.
* **400 Bad Request** When there is a failure in the request format expected.

### Delete /

**Body** _required_ User to delete. It will search by all the fields and detele all the users that match the requirements.

Sample:

```json
{
  "first_name": "Mario",
  "last_name": "Arranz",
  "nickname": "marioarranzr"
}
```

Responses:

* **200 OK** When the users are deleted correctly.
* **400 Bad Request** When there is a failure in the request format expected.

## Run

Run locally in docker:
```
make run-locally
```
Run locally (not using docker):
```
make run
```
Run tests with coverage:
```
make test
```

The app will run in port `9091`


### Assumptions:

- The in-memory database is simply a list of users.
- When a user is inserted, the system checks that the nickname is not already in the database.
- Nickname is mandatory when insterting and not allowed to modify.
- When updating a user, it will search by nickname and update the other fields in the request.
- The comunication among microservices will be done in the API Gateway that would make the requests to each microservice. I.e. To add a payment method in the system, from the Gateway it will be necessary at least 2 requests:
  
  - user microservice
  - if the response is **201 Created**, the next request would be:
  - payment microservice
  - if the response is **201 Created**, we know that new user is in the system and has payment details configured.

- ###### (Async communication) Another way of comunication among microservices, without passing through the Gateway, would be using messages queues and posting messages for other microservices to read.
