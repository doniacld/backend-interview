# backend-interview
This repository contains instructions and draft project for powder backend interview

## requirements

- go
- docker
- docker-compose
- make
- psql

## how to setup

- Setup local postgres db
```sh
> docker-compose up -d
```

- Create db schema + Populate with fake data
```sh
> make migrate
> make populate
```

- Build API binary and run locally
```sh
make api && ./bin/interview_api config/api/local.json
```

## instructions

You inherited this project and there are issues you need to fix.

Usecase: `user` has 1+ `account` and each `account` register its own `transactions`.

User must be able to insert new transactions and get his current total.

For this test, we don't care about security flaws or float64 approximations.

You can fork this project and add your code/answers into it.

## project architecture0

```
cmd_|_ # main
    |
    |__api # binary
    |
pkg_|_ # domain
    |
    |__user # domain name
      |
      |__sql # sql STORE implementation
      |__app # APP implementation, domain logic + store logic
      |__dto # DataTransferObject for external objects
```

### 0.0

- `GetUser` API route always returns a total = 0. Fix it to return approximated sum (due to float64). Hint: `account.FetchMany` SQL method already exists in draft code (`handler.account` store).

- Write a new API route `CreateTransaction` to add a new transaction. Hint: `account.InsertTransaction` SQL method already exists in draft code (`handler.account` store).

- Add some minimal tests on those routes to ensure at least 1 success path. (any kind of test is ok)

### 0.1

Questions (text response only):

- Following a technical refactoring, we decide to modify the previous REST API "/user" in grpc. Propose a protobuf contract using the principle of this API.

- This protobuf (and those that will follow) will be shared between the backend team, the android team and the ios team. Propose an efficient solution to document and synchronize the teams around this protobuf. 

- How would you proceed to make this synchronization and documentation automatic?

- Let's imagine that a new API with several microservices and asynchronous processing is designing, propose a notification solution to keep the clients (ios/android) up to date.

### 0.2 (bonus)

- (bonus) We want to get rid of account intermediary table and attach directly transactions to user. Write up a database migration plan (+ add some example queries).

- (bonus) Add a new rule where a transaction is not accepted if there is not enough money on account.

- (bonus) User now wants to know his largest expense (transaction) between 2 dates. Create a new route GetMaxTransaction which takes 2 timestamps in parameters.
