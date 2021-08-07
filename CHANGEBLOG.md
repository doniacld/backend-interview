# Changeblog

## What I wanted to do

- Write more tests for GetUser
    - user invalid
    - no user
    
- Write more tests for CreationTransactions
    - Get the total amount to verify it has been udpated after the creation of a transaction
   
- Create integration tests with a script calling curls

- Use of the callback

- Answer better to the following questions! 

- Add the service in docker

## Answers

### 0.0.1

- Following a technical refactoring, we decide to modify the previous REST API "/user" in grpc. Propose a protobuf contract using the principle of this API.


    syntax = "proto3";
    package user;
    
    option go_package = "./model";
    
    message User {
      string ID = 1;
    }
    
    message GetUserResponse {
        User user = 1 ; 
        float total = 2 ;
    }
    
    service User {
      rpc GetUser(User) returns (GetUserResponse) {}
    }


- This protobuf (and those that will follow) will be shared between the backend team, the android team and the ios team. 
Propose an efficient solution to document and synchronize the teams around this protobuf. 

Share the API in a dedicated repository accessible by the teams.

- How would you proceed to make this synchronization and documentation automatic?

Create a CI that generates for all the languages the appropriate definition files triggered by a change. 
I could have a lot to say about this part.

- Let's imagine that a new API with several microservices and asynchronous processing is designing, propose a notification solution to keep the clients (ios/android) up to date.

I guess it is an extension of the given previous solution.

## Helpers

Helper to check what is in the db:

    docker exec -it postgres bash

    psql -h postgres -d postgres -U postgres

