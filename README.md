# drucssGO
----------
drucssGO is an example of a REST API in Golang

## Design
```
┌─────────────────────────────────────┐
│ POST /v1/user                       │
├─────────────────────────────────────┤
│ GET /v1/users                       │
├─────────────────────────────────────┤
│ GET /v1/user                        │
├─────────────────────────────────────┤
│ PATCH /v1/user                      │
├─────────────────────────────────────┤
│ DELETE /v1/user                     │
└─────────────────────────────────────┘
```

The API has 5 routes for Creating, Reading, Updating and Deleting data from
a user. All users should be uniquely identified using a UUID in version 4.
The endpoints should work by the following rules:

**POST /v1/user**:  
    Add a user to the datasystem. The user name should be supplied using a
    json structure  
**GET /v1/users**:  
    Obtains all the users in the datasystem  
**GET /v1/user**:  
    Obtain a single user from the datasystem given an user_id  
**PATCH /v1/user**:  
    Updata a single user from the datasystem fiven an user_id. The user data
    should be supplied using a json structure  
**DELETE /v1/user**:  
    Delete a single user from the datasystem given an user_id  

## JSON Data structure
For input the JSON data structure used on this api has the following
format:

```
{
    "id": uuid
    "name": string
}
```

## LICENSE

See LICENSE file
