# Okteto Pods

Okteto Pods is an [Offline Exercise](https://gist.github.com/jmacelroy/0de9eb394fd4f0cc49869c25aef70b66) that provides an API REST to get Pods Data.

## Features

* Endpoint to get Total Pods.
* Endpoint to get a Sorted List of Pods using the next posible options: name, restarts and age.

## API Usage & Example

URL Base: https://pods-mig-elgt.assessment.jdm.okteto.net

### Endpoint to get the number of pods.

```
Request

GET /pods/total

Posible Responses

Status: 200
Body Respose
{
   "total": 100,
}

Status: 500
Body Respose
{
   "error": {
      "status": 500,
      "error": "INTERNAL",
      "description": "Something went wrong...",
   }
}
```
