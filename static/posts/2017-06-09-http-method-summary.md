---
layout: post
title:  "HTTP Method Summary"
date:   2017-06-09 16:48:17 -0500
categories: web-development
---

Here's a summary of the most common HTTP methods by use case.

The scenario: 

Imagine you're tasked to create a simple web api that hits a Car Dealership database with only two entities, Make and Model. It's a one-to-many relationship where a particular Make can have many Models. For example, Toyota produces the Camry, Tundra, 4Runner etc.

```
General uri format: 
api/{makes}/{makeId}/{models}/{modelId}
```

### Reading Resources

```
Request:
GET api/makes

Response:
200 Ok -> [{make}, {make}, ...]
404 Not Found
```

We send a GET request if we want a collection of resources. A successful response will return 200 Ok and said resource collection in the response body. We return 404 Not Found if the uri doesn't exist.

```
Request:
GET api/makes/{makeId}

Response:
200 Ok -> {make}
404 Not Found
```

If we want to read a specific resource, we add the id. A success will return that specific resource in the response body. A 404 is returned if the resource doesn't exist.

### Deleting Resources

```
Request:
DELETE api/makes/{makeId}

Response:
204 No Content 
404 Not Found
```

We return a 204 No Content after a successful DELETE request of a particular resource and return a 404 if the resource is not found. No Content means the server has successfully fulfilled the request and that there's nothing to be sent with the response body.

```
Request:
DELETE api/makes

Response:
204 No Content 
404 Not Found
```

Deleting a collection of resources follows the same guidelines as above but it's rarely done as it can be quite destructive. Generally, mass deletion is barred.

### Creating Resources

When creating resources, we have to differentiate between requests where the responsibility of creating the uri is at the server and where it's at the consumer of our api. 

#### Server
```
Request:
POST api/makes <- {make}

Response:
201 Created -> {make}
404 Not Found
```

To create a new car Make, we send a POST request to the Makes resource with the Make representation as the request body. If creation is successful then the status code is appropriately 201 Created and the response contains the created Make representation in the response body and location in the location header. If the Makes resource doesn't exist then we return 404 as usual.

```
Request:
POST api/makes/{makeId}

Response:
404 Not Found
409 Conflict
```

Issuing a POST request to a single resource uri can never be successful because either a 404 Not Found is returned if the Make doesn't exist or a 409 Conflict is returned because that particular Make already exists. 

```
Request:
POST api/makecollections <- {makeCollection}

Response:
201 Created -> {makeCollection}
404 Not Found
```  

Adding a collection of resources, if allowed, should have a new resource uri created for that collection. For example, we can call it MakeCollections and a POST request to that resource will contain an array of Makes. We return a 201 Created with the new collection in the response body for a successful response and return 404 if that resource collection doesn't exist.

#### Consumer

```
Request:
PUT api/makes/{makeId} <- {make}

Response:
201 Created -> {make}
``` 

When the consumer of our api is allowed to create resource uris, we can ["upsert"][Upsert] to create resources. For example, a PUT request is issued to update an existing resource, fine, but if the same PUT request is issued on a resource that doesn't exist then that will result in the resource being created. Hence, why a 201 Created is returned with the newly created Make in the response body. A 404 Not Found will never occur because a non-existant resource will be created. 

```
Request:
PATCH api/makes/{makeId} <- {JsonPatchDocument on make}

Response:
201 Created -> {make}
``` 

Upserting with a partial resource update with a PATCH request works identically to upserting with PUT except that a json patch document is included in the request.

### Updating Resources

```
Request:
PUT api/makes/{makeId} <- {make}

Response:
200 Ok -> {make}
204 No Content
404 Not Found
```  

Issuing a PUT request on a resource updates that resource. Since it's a PUT, we should pass in the request the full Make representation or else omitted fields will be set to their default values. We can either return 200 Ok with the successfully updated resource representation in the response body or just state 204 No Content. Non-existant resources return a 404 Not Found, unless we're [upserting][Upsert]. 

```
Request:
PUT api/makes <- [{make}, {make}, ...]

Response:
200 Ok -> [{make}, {make}, ...]
204 No Content
404 Not Found
```

Using a PUT request to mass update is usually never allowed but can be done. It follows the same guidelines as updating a single resource except the resource collection is pass in the request and the updated collection is returned in the successful response body. Non-existant resources return a 404 Not Found, unless we're [upserting][Upsert]. 

```
Request:
PATCH api/makes/{makeId} <- {JsonPatchDocument on make}

Response:
200 Ok -> {make}
204 No Content
404 Not Found
```

Partial updates with a PATCH request is similar to a PUT request. Except we do not have to pass with the request a full representation of the resource. We state the fields we want to update in a json patch document and omitted fields keep their present values. A successful response can return 200 Ok with the patched resource representation in the body or simply state 204 No Content. Non-existant resources return a 404 Not Found, unless we're [upserting][Upsert]. 

```
Request:
PATCH api/makes <- {JsonPatchDocument on make}

Response:
200 Ok -> [{make}, {make}, ...]
204 No Content
404 Not Found
```

Using PATCH to mass update resources, like most mass operations, are rarely allowed. If it is allowed then it follows the same guidelines as PATCHing a single resource. 

[Upsert]: https://en.wikipedia.org/wiki/Merge_(SQL) 