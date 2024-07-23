# group-bill-calculator

## What is this
A simple calculator to give you the total amount of a multiple bills split between a group of people.

## How to use it
1. Clone the repository;
2. Enter bills in the `bills.json` file;
3. Commit and github actions will calculate the total amount of debt for each person in the group;

## How to add bills
1. Open the `bills.json` file;
2. The file should contain an empty array: 
```json 
 []
```
3. Add a new bill in the following format:
```json
{
  "name": "breakfast",
  "involved": [
    {
      "name": "John",
      "amount": 10
    },
    {
      "name": "Jane",
      "amount": -10
    }
  ]
}
```
4. A file with multiple bills should look like this:
```json
[
  {
    "name": "breakfast",
    "involved": [
      {
        "name": "John",
        "amount": 10
      },
      {
        "name": "Jane",
        "amount": -10
      }
    ]
  },
  {
    "name": "lunch",
    "involved": [
      {
        "name": "John",
        "amount": 20
      },
      {
        "name": "Jane",
        "amount": -20
      }
    ]
  }
]
```