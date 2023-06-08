# Project Name

Receipt Processor

## Table of Contents

- [Overview](#overview)
- [Installation](#installation)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Dependencies](#dependencies)
- [Contributing](#contributing)
- [License](#license)

## Overview

A webservice built with Go that fulfils the two APIs - /receipts/process and /receipts/{id}/points. The major code file is main.go

## Installation

After installing Go environment, go to the project directory(`fetch` folder), and run `go run main.go`

## Usage

If project successfully runs, there will be two apis running on the localhost:3000, which is localhost:3000/receipts/process and localhost:3000/receipts/{id}/points. The data is stored in memory, so it does not persist after application stops.

## API Endpoints

Detailed API descreption is already stated clearly in the given api.yml file. There are some notes on edge case handling in implementation.

For Post API: localhost:3000/receipts/process

If the filed names in the request body are different from the struct field names(as given in the description file), it will not return errors. An object with ID is still generated. However, only the field with correct names will be matched and saved. When later doing point calculation, only those correct fileds will be considered.

It will return an error under following conditions:

* If the request body cannot be read: This might occur if there's a network error or a similar problem.

* If the request body is not valid JSON: If the request body cannot be parsed as a JSON object, you'll get an error.

* If the JSON cannot be unmarshaled into the provided struct: If the JSON structure doesn't match the structure of your struct (for example, if it has fields of the wrong type), you'll get an error.


## Dependencies

The external libraries used in project are: 	"github.com/gin-gonic/gin" "github.com/google/uuid"



