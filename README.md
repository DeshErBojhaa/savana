# Savana

Savana is the controller for a parking lot. It finds the optimal parking spot for vehicles in a 
multy entry multi storey parking space.

## Usage guide
1. Run "setup" script in bin folder.
```
./bin/setup
```
2. This will place executable file in ./target folder which can be run like:
```
./bin/parking_lot
or
./bin/parking_lot "INPUT_FILE_FULL_NAME"
```

Run test from `root` directory:
```
go test ./... -v
```

## WARNING!!
Running `run_functional_test` first time results to missmatch.
Subsequent runs are workig fine.

TODO: Fix it.
