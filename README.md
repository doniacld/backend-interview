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

## project architecture

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

- `GetUser` API route always returns a total = 0. Fix it to return approximated sum (due to float64). Hint: `account.FetchManyAccount` SQL method already exists in draft code (`handler.account` store).

- Write a new API route `CreateTransaction` to add a new transaction. Hint: `account.InsertTransaction` SQL method already exists in draft code (`handler.account` store).

- Add some minimal tests on those routes to ensure at least 1 success path. (any kind of test is ok)

### 0.1

Questions (text only):

- You are running this service in production under ~100 req/s, what are your main concerns about scaling and stability ?

>

- We want to get rid of `account` intermediary table and attach directly transactions to `user`. Write up a database migration plan (+ add some example queries).

>

### 0.2 (bonus)

- (bonus) Add a new rule where a transaction is not accepted if there is not enough money on account.

- (bonus) User now wants to know his largest expense (`transaction`) between 2 dates. Create a new route `GetMaxTransaction` which takes 2 timestamps in parameters.

- (bonus) Create and implement a mock on a `store` or `app` (of your choice). Using this mock, write up a benchmark comparing a route (of your choice) with and without mocks.
Mock implementation will be check in the code part.
For the benchmark part you need to provide a code part (no verification on clean/maintenance for this one, it's benchmark code) + benchmark results (in easy to read format plz)

