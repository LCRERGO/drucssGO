# drucssGO
----------
drucssGO is an example of a REST API in Golang

## Design
```
┌─────────────────────────────────────┐
│ GET /v1/users                       │
├─────────────────────────────────────┤
│ GET /v1/users/{user_id}             │
├─────────────────────────────────────┤
│ POST /v1/users                      │
├─────────────────────────────────────┤
│ DELETE /v1/users/{user_id}          │
├─────────────────────────────────────┤
│ PATCH /v1/users/{user_id}           │
└─────────────────────────────────────┘
```

The API has 5 routes for Creating, Reading, Updating and Deleting data from
a user. All users should be uniquely identified using a UUID in version 4.
The endpoints should work by the following rules:

**GET /v1/users**:  
    Obtains all the users in the datasystem  
**GET /v1/users/{user_id}**:  
    Obtain a single user from the datasystem given an user_id  
**POST /v1/users**:  
    Add a user to the datasystem. The user data should be supplied using a
    json structure  
**DELETE /v1/users/{user_id}**:  
    Delete a single user from the datasystem given an user_id  
**PATCH /v1/users/{user_id}**:  
    Updata a single user from the datasystem fiven an user_id. The user data
    should be supplied using a json structure  

## LICENSE

See LICENSE file
