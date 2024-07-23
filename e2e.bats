#!/usr/bin/env bats

setup() {
  go build -o bills_calculator .
}

teardown() {
  rm -f bills_calculator temp*.json
}

@test "Run with valid JSON file" {
  run ./bills_calculator -file testdata/test-valid-bill.json
  [ "$status" -eq 0 ]
  [ "${lines[0]}" = "Final imbalances per person:" ]
  [ "${lines[1]}" = "John must receive: 10;" ]
  [ "${lines[2]}" = "Jane must give: 10;" ]
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