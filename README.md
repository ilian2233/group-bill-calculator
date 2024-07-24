# group-bill-calculator

## What is this
A simple calculator to give you the total amount of a multiple bills split between a group of people.

## How to use it
1. Clone the repository;
2. Enter bills in the `bills.json` file;
3. Commit and github actions will calculate the total amount of debt for each person in the group;

## How are the final amounts calculated
Each bill contains list of involved people and the amount each payed or should pay.
The amount for each person is calculated by the following formula: `amount = (How much I payed) - (How much do my items cost)`.
The sum of all amounts in a bill cannot be negative as this means the bill is not payed.
The amount can be positive meaning one or more people payed more than they should have. This positive amount will be split between all the people in the bill that have not entered amount explicitly.