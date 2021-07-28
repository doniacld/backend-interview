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

Questions :

we have an app - Powder - which is a video creation and sharing app. There is a feed of Posts (one post = video + comments + likes).
At the moment, the feed is the same for all users.

We want to customise it. To do this, we need to:
- Retrieve relevant information from the user regarding the viewing of posts
- store this information in a cloud
- use this information to customise the feeds for each user
Please describe all the steps needed to obtain these feeds, on the mobile applications and the backend.
Be as specific as possible about the tools you will be using (especially in the cloud) - please describe as much as possible about the settings of the tools, and make a diagram explaining all the steps of your solution. You can add your schema/diagram directly in this repository.
