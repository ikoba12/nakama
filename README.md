# Nakama Configuration RPC
## Overview
The purpose of the project is to implement a simple RPC function that accepts a request, 
reads the data from the json, stores it in the DB, checks the file hash and returns the response.  

The data stored in the json files for testing purposes is the information about the user's server.

```
{
  "maxNumPlayers" : 200,
  "server" : "europe",
  "bestAvenger" : "Captain america"
}
```
It provides the client with the information about maximum number of players, the location of the server   
and who is considered the best avenger on that server :)

The following content is stored in the database to demonstrate how data is stored in Nakama's internal storage as a json blob   
It can be accessed via the nakama client or the using Nakama's JSONB queries.

For calculating the hash of the file SHA256 is used. The client should provide the same file content hash if that data needs to be returned.  

The codebase is covered with unit tests. Unfortunately due to lack of time could not find a reliable framework for Nakama go to perform integration tests(There was a third party but went with Unit testing for now)
Heroiclabs official website provided a framework for testing Typescript server runtime but they don't have one for go.

I tried to decouple the project as much as possible to make the code more readable and easily maintainable.

Given more time I would try out the framework for testing the RPC function end to end, try to implement a better error handling mechanism to have more detailed error messages for the client. 
Also, would take some time to figure out what are the best practices for go and maybe arrange the code better.  
A bit more context would allow to figure out what the data should be used for and how it should be stored properly in the storage for optimal performance in the future.




## Getting Started

### Prerequisites

- Go 1.21.6+
- Docker (for running Nakama and PostgreSQL)

### Quickstart

1. Clone the repository:

```sh
git clone https://github.com/ikoba12/nakama.git
cd nakama
```
2. Install the dependencies
```sh
go get github.com/heroiclabs/nakama-common/runtime
go mod vendor 
```
3. Running the project
```sh
docker compose up --build
```
4. Login 
```
Go to http://localhost:7351/

Use nakama default user and password to login

user : admin  
password : password

find the configrationinforpc in the Api Explorer and test it out
```

