# mars-rover

A microservice application

Code Organization

    - cmd - Main & Binaries
        - bin
            - mars-rover
    - config - config data from env
    - internal
        - world - world data from input
        - pkg
            - parser - to parse the individual robot data and command to execute
            - robot - robot model
    go.mod
    go.sum
    Makefile
    Readme.md
    .sample.env
    sample-input.txt

## Build

```bash
  make build
```

## Run Locally

```bash
  cp sample-input.txt input.txt
  cp .sample.env .env
  cmd/bin/mars-rover
```

## Running Tests

To run tests, run the following command

```bash
  make test
```

## Working

    Application loads the config and reads the world data from input
    
    ----------------------------------------------------------------------------------------------------------------------------------------------------------
    Sample input
    ----------------------------------------------------------------------------------------------------------------------------------------------------------
    4 8 //grid size
    (2, 3, E) LFRFF // Robot initial position, orientation (x, y, orientation) and command to execute
    ----------------------------------------------------------------------------------------------------------------------------------------------------------

    Move the Robots to the target position by running instructions in the command

    ----------------------------------------------------------------------------------------------------------------------------------------------------------
    Details of the robot movement based on the command
    ----------------------------------------------------------------------------------------------------------------------------------------------------------
    Each robot has a position (x, y), and an orientation (N, E, S, W)
    Each robot can move forward one space (F), rotate left by 90 degrees (L), or rotate right by 90 degrees (R)
    If a robot moves off the grid, it is marked as ‘lost’ and its last valid grid position and orientation is recorded
    Going from x -> x + 1 is in the easterly direction, and y -> y + 1 is in the northerly direction. i.e. (0, 0) represents the south-west corner of the grid    
