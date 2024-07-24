#!/usr/bin/env bats

setup() {
  go build -o bills_calculator .
}

teardown() {
  rm -f bills_calculator temp*.json
}

@test "Run with valid JSON file" {
  run ./bills_calculator -file testdata/valid-bill.json
  [ "$status" -eq 0 ]
  [[ "${lines[@]}" =~ "Final imbalances per person:" ]]
  [[ "${lines[@]}" =~ "John must receive: 10.00;" ]]
  [[ "${lines[@]}" =~ "Jane must give: 10.00;" ]]
}

@test "Run with equally split bill file" {
  run ./bills_calculator -file testdata/equally-split-bill.json
  [ "$status" -eq 0 ]
  [[ "${lines[@]}" =~ "Final imbalances per person:" ]]
  [[ "${lines[@]}" =~ "John must receive: 10.00;" ]]
  [[ "${lines[@]}" =~ "Jane must give: 5.00;" ]]
  [[ "${lines[@]}" =~ "Joe must give: 5.00;" ]]
}

@test "Run with unequally split bill file" {
  run ./bills_calculator -file testdata/unequally-split-three-way-bill.json
  [ "$status" -eq 0 ]
  [[ "${lines[@]}" =~ "Final imbalances per person:" ]]
  [[ "${lines[@]}" =~ "John must receive: 10.00;" ]]
  [[ "${lines[@]}" =~ "Jane must give: 3.00;" ]]
  [[ "${lines[@]}" =~ "Joe must give: 7.00;" ]]
}


@test "Run with invalid JSON file" {
  run ./bills_calculator -file temp_invalid.json
  [ "$status" -eq 1 ]
  [[ "${lines[0]}" =~ "Error reading bills" ]]
}

@test "Run with non-existent file" {
  run ./bills_calculator -file non_existent.json
  [ "$status" -eq 1 ]
  [[ "${lines[0]}" =~ "Error reading bills" ]]
}