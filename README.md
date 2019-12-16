# Golang-working-with-RabbitMQ
#This is a small application will help you how to communicate with RabbitMQ. Please see the solution details below

### Project Description

**Scenario**

We have a `input` folder, which contains such data splitted in different files. Each of the files corresponds to the message, one service would provide. 

In order to present all findings, we need another service to restructure the provided data.

This second service (worker) will listen to the queue and restructure all messages into the required format.

**Example Input Structure**

```json
{
  "id": 1,
  "findings": [
    {
      "name": "finding 1",
      "severity": 5.0,
      "service": "A",
      "subfindings": [
        {
          "name": "subfinding 1",
          "severity": 6.0
        }
      ]
    },
    {
      "name": "finding 2",
      "severity": 1.0,
      "service": "F",
      "subfindings": [
        {
          "name": "subfinding 1",
          "severity": 0.0
        }
      ]
    },
    {
      "name": "finding 3",
      "severity": 7.6,
      "service": "B",
      "subfindings": [
        {
          "name": "subfinding 1",
          "severity": 3.9
        },
        {
          "name": "subfinding 2",
          "severity": 0.0
        }
      ]
    }
  ]
}
``` 

**Example Output Structure**

```json
{
  "categories": {
    "high": {
      "amount": 0,
      "max_severity": 0.0,
      "services": []
    },
    "low": {
      "amount": 1,
      "max_severity": 1.0,
      "services": [
        "F"
      ]
    },
    "medium": {
      "amount": 2,
      "max_severity": 7.6,
      "services": [
        "A",
        "B"
      ]
    }
  }
}
```

**Data Restrictions**

| Fieldname   | Restrictions                                    |
|-------------|-------------------------------------------------|
| Severity    | Float between 0.0 and 10.0                      |
| Service     | Letters from A to F                             |
| Category    | low (Severity: 0 -> 3.9), medium (Severity: 4 -> 7.9), high (Severity: 8-10)  |

### Solution

we will Build two Go Micro-Applications:

1. The first application (queue) contains a queue.
	* Implement a queue system
	* On startup of the application, read-in all files from the `inputs` folder, convert them to a struct and add the structs to the queue

2. The second application (worker) should be a worker which can read from the queue in the first application.
	* It should be able to handle 5 queue messages at the same time. 
	* Each message should be analysed and converted to a result struct.
	* Write the result struct to a file `result_ID.json` - where ID is the ID from the input - in the `outputs` subfolder.

Analysing and restructuring of a message does the following:

* Iterate over all findings and calculate the maximum severity value for each finding. The value is retrieved by comparing the findings severity with all its subfindings and returning the highest value.
* Cluster all findings by their `category` in the result struct.
* All possible categories should be listed in the output, even if there are no findings belonging to a specific category.


### Folder structure
- inputs/
- outputs/
- queue/
- worker/
- model/
- helper/
- categoryhandler/
- rabbitmqhandler/
