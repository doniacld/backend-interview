# backend-interview
This repository contains instructions and draft project for powder backend interview

### requirements

- go
- docker
- docker-compose
- make
- psql

### how to setup

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

### 0.0

- Write a new API route `CreateTransaction` to add a new transaction. Hint: `account.InsertTransaction` already exists.

- `GetUser` API route always returns a total = 0. Fix it to return approximated sum (due to float64). Hint: `account.FetchManyAccount` already exists.

### 0.1

With above routes, write some api tests to ensure app basic behaviors.
You can write tests with any tool/language.
(focus on defining different test flows to cover most cases, "technical" dimension is not what we're looking for here)

### 0.2

(bonus) Add a new rule where a transaction is not accepted if there is not enough money on account.
Add a test for this case.

(bonus) User now wants to know his largest expense (`transaction`) between 2 dates. Create a new route `GetMaxTransaction` which takes 2 timestamps in parameters.
Add a test for this route.

### 0.3

Questions (text only):

- What metrics could be interesting to observe to ensure app integrity and stability ?

>

- What kind of alerts based on those metrics could we use here ? What critical conditions should we look at ?

>

