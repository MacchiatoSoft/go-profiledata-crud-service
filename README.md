# go-profiledata-crud-service
A simple service that can perform CRUD operations to store data on user profiles.

## Roadmap:

### Dockerised environment:

An environment in which this system is launched. It should have containers for both the application and database which should be able to talk to eachother.
- DONE Go application container
- DONE Postgres container

### User management:

Login system so users can only access their own data
- POST request to add a new user
- POST request to login
- DELETE request to remove a user

### Userdata management:

Figure out what data is being held

Upload and access userdata associated with a user
- POST request to store userdata
- GET request to retrieve userdata
- PATCH request to update userdata
- DELETE request to remove userdata

### Security

Add actual security features
- OAuth
- Encryption