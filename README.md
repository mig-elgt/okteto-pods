# Okteto Pods

Okteto Pods is an [Offline Exercise](https://gist.github.com/jmacelroy/0de9eb394fd4f0cc49869c25aef70b66) that provides an API REST to get Pods Data.

## Features

* Endpoint to get Total Pods.
* Endpoint to get a Sorted List of Pods using the next posible options: name, restarts and age.

## API Usage & Example

URL Base: https://pods-mig-elgt.assessment.jdm.okteto.net

### Endpoint to get the number of pods.

```
Endpoint: /pods/total

Example

Request:
GET https://pods-mig-elgt.assessment.jdm.okteto.net/pods/total

Responses:

Status: 200 - Ok
Body Respose
{
   "total": 100
}

Status: 500 - Internal Server Error
Body Respose
{
   "error": {
      "status": 500,
      "error": "INTERNAL",
      "description": "Something went wrong...",
   }
}
```

### Endpoint to a Sort List of Pods

```

Endpoint: /pods?sort=field:order

Query Parameters
 - sort
 - values
    // field posible values
    - field: name, restarts, age
    // order posible values
    - order: asc, desc

Example

Request:
GET https://pods-mig-elgt.assessment.jdm.okteto.net/pods?sort=age:desc,name:desc

Responses:

Status: 200 - Ok
Body Respose
{
  "status": 200,
  "pods": [
    {
      "name": "mongodb-0",
      "restarts": 0,
      "status": "Running",
      "age": "18h2m37s"
    },
    {
      "name": "pods-79bf8477c8-wc64q",
      "restarts": 0,
      "status": "Running",
      "age": "43m28s"
    },
    {
      "name": "pods-79bf8477c8-st7wx",
      "restarts": 0,
      "status": "Running",
      "age": "43m28s"
    },
    {
      "name": "pods-79bf8477c8-mj7v6",
      "restarts": 0,
      "status": "Running",
      "age": "43m27s"
    },
    {
      "name": "pods-79bf8477c8-4bfvh",
      "restarts": 0,
      "status": "Running",
      "age": "43m27s"
    }
  ]
}

Status 400 - Bad Request
{
  "error": {
    "status": 400,
    "error": "INVALID_ARGUMENT",
    "description": "One or more fields raised validation errors.",
    "fields": {
      "sort": "Incorrect field order value, should be asc or desc."
    }
  }
}


Status: 500 - Internal Sever Error
Body Respose
{
   "error": {
      "status": 500,
      "error": "INTERNAL",
      "description": "Something went wrong...",
   }
}
```
