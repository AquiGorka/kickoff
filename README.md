# Kickoff [![Build Status](https://travis-ci.org/AquiGorka/kickoff.svg?branch=master)](https://travis-ci.org/AquiGorka/kickoff) [![Coverage Status](https://coveralls.io/repos/github/AquiGorka/kickoff/badge.svg?branch=master)](https://coveralls.io/github/AquiGorka/kickoff?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/AquiGorka/kickoff/game)](https://goreportcard.com/report/github.com/AquiGorka/kickoff/game)

A multiplayer real-time turn-based soccer HTML5 game for virtual reality.


## Todo

* Analysis
  * Setup flows for an automated game (this will show the basic micro-services and their interactions)
   * game flow
   * match flow
   * player flow (bot)
   * viewer flow
* Project setup
  * Define milestones to reach the basic automated game
* Development
  * Tests
* MVP
  * Automated game


## Flow

* execute game
* execute match
* execute N players for match
* execute viewer


### Game Simulation (Soccer-like game)

A micro-service for a soccer game-like simulation.

This service will wait for connections to setup and initiate an instance of a soccer-like game where a physics simulation representing a field (with defined size & boundaries defined by parameters sent to the service), a number of players (defined by parameters sent to the service) and a ball will take place.

This service will receive game interactions from the connected peer that will modify the game state (in a soccer-like fashion: giving the ball direction & force to move it).

This service will publish to the connected peer the current state of the physics simulation in a defined interval (defined via parameters sent to the service).

The soccer-like game simulation will take responsibility for soccer rules (what to do when the ball goes out of the defined boundaries, how to proceed after a goal has been scored, duration of the match, player keeping the ball for too long, etc).

```
execute
  wait for connections
end
```

```
setup
  init simulation
    world
    stadium
      field (@param size)
    players (@param number of players)
    ball
  init game rules
    duration (@param duration)
    start datetime (@param datetime) (this is when this service will start publishing updates)
    listens to game specific events
      ball out of bounds
      score
      ball stuck
```

```
start match
  start simulation loop
  listen to player interactions
    update game state
  publish game state
end match
```


### Match

A micro-service for players and viewers to connect to.

This service works as the gateway to connect the outside world and the game simulation.

This service encapsulates the auth logic, match settings for the game, player interactions with the game and lets viewers subscribe to watch the game.

```
execute
  setup automated match
    setup players
    setup field
    setup duration + start datetime
  wait for connections
    player
    viewer
end
```

```
connection
  player
    Send to player game state and user interactions
    Receive from player game interactions
  viewer
    Send to viewer game state and user interactions
```


### Player

A micro-service to automate playing the game.

This service will connect to the game and send interactions randomly when needed to.

This service will receive the game state everytime the match service publishes such state.

```
execute
  connect
    wait/join game loop

  game loop
    check the ball, if it is inside a certain radius the player will send an interaction to the match micro-service
  end game loop
end
```


### Viewer

A micro-service to watch the game.

This service will receive the game state everytime the match service publishes such state and render a visualization of the game.

```
execute
  connect
    wait/join game loop

  game loop
    render game state
  end game loop
end
```

