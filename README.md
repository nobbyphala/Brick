# Brick (Money Transfer Test Assigment)

## Code Structure

### adapter
This interface adapter layer. This layer responsibility is handle all API request coming from the outside and build the
data to be consumed by the business logic. This layer define all the request handler (Rest API, GRPC, etc).

### domain
Domain contains all of business entities. Domain responsibility to separate the business related entity from the rest of
application. This layer will contain all structure related to the business rule. For example in this project is Disbursement struct.

### external
External contains the framework and drivers needed by the application. This layer will wrap all dependencies such as Database,
Http Request, and validator library. By wrapping the external dependencies will make us easier to migrate to new framework or
driver in the future since we only need to re-wrap the new library following the external layer interface and no need to touch
the business logic

### usecase
Usecase contains application business logic. Inside usecase there are repository and api, both are the interface to communicate
with the outside service such as Database or API request to other service. In summary usecase hold the business logic and
all the interface needed by the business logic

### config
Contains additional config for the application such as Database credentials. Config also the place hold all the environment
variable needed by the application

## How To Run

1. Spin up the postgresql database using docker compose

   ```
   docker compose up
   ```
2. Run main.go
    ```
   go run main.go
   ```
3. This application use Mockoon to mock the third party bank service API. Please go to this link https://mockoon.com/download/#download-section
to read the details how to install. After Mockoon installed import mockoon.json file and run the mock api server. 
If you plan to mock using another service please change the base url config here https://github.com/nobbyphala/Brick/blob/337ba33b16cad8902c20e4b6b08bc833480aed19/config/bank.go#L4.

Note: I use Mockoon instead of mockapi.io because it's free and open source

## Improvement
This section explain a bit about what can be improved from this project

1. Wrap all the sqlx interface. As you can see not all the sqlx api is wrapped in external and one of the repository (util repository)
use the sqlx directly. Next improvement need to wrap all the needed sqlx api and refactor util repository. ref: https://github.com/nobbyphala/Brick/blob/9bcc1ca0893f0a44f08f6e0df20d681d3eee631e/usecase/repository/utils.go#L9
2. Add unit test for external package