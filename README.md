## Receipt Processor Project

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Dependencies](#dependencies)


## Overview

A webservice built with Go that fulfils two APIs - /receipts/process & /receipts/{id}/points. The major code is in the file `main.go`

## Installation

After installing Go environment, go to the project directory(`fetch` folder), and run commands

`go run main.go`

It will automatically install dependencies and run the service.

## Usage

If project successfully runs, there will be two APIs running on the localhost:3000, which is localhost:3000/receipts/process and localhost:3000/receipts/{id}/points. The data is stored in memory, so it does not persist after application stops.

* For Post API - localhost:3000/receipts/process: 

the request body is like:

```JSON
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    }
  ],
  "total": "35.35"
}
```
the response is like:

```JSON
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

*  For Get API - localhost:3000/receipts/{id}/points

construct the url using the id in the response from the post API, like this: `localhost:3000/receipts/7fb1377b-b223-49d9-a31a-5a02701dd310/points`

the response is like:

```JSON
{ "points": 28 }
```

## API Endpoints

The detailed API description can be found in the given `api.yml` file. Additionally, there are some important notes regarding edge case handling in the implementation.

- POST API: `localhost:3000/receipts/process`

  - If the field names in the request body differ from those mentioned in the description file, no errors will be returned. However, only the fields with correct names will be matched and saved. During point calculation, only these correct fields will be considered.

  - The following error conditions will result in a reponse of error message:

    - If the request body cannot be read, possibly due to a network or similar issue.

    - If the request body is not valid JSON and cannot be parsed as a JSON object.

    - If the provided JSON cannot be unmarshaled into the specified struct due to structural mismatches or incorrect field types.

- GET API: `localhost:3000/receipts/{id}/points`

  - If the specified ID does not exist, the API will return an error message: "Receipt not found. ID does not exist."


## Dependencies

The external libraries used in project are: "github.com/gin-gonic/gin" (for routing HTTP requests) and "github.com/google/uuid" (for generating random ID)



