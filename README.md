# Kickoff

A multiplayer real-time turn-based soccer html5 game for virtual reality.

## TODO

* Define world setup
* Define game flow
* Define game limits (time & players)

## File structure
/
  src/
    game/
    auth/
    gateway/
    player/
    viewer/

## Architecture & Microservices

* Game
* Auth
* Gateway
* Client
** Player
** Viewer

## Ideas

* Go as server?
** What about p2p streaming?
* Use cannon from Go?
* Bot for dev-flow steps
** Tells you where to continue and how (might even fill out required info)
** Ask for status, milestones, etc <- maybe integrate with git?
*** https://hubot.github.com/

## Flow

### Management

teams
  add
  edit
match
  play?
    setup?
    join?

### Player

* play
start
  connect
     which match? (if there is only one, guess where you're gonna get redirected to)
  loop
end

* Substitution

### Match

start
  setup (player ids?, size, duration, start date+time, weather?)
    setup players
    setup field
    setup duration + start datetime
    setup conditions
end

### Game

setup (size, time, ..extra conditions)
  setup world
  prepare physics engine
start
  init loop
  loop
end

### Connection

#### Notes

* Whitelist for players.
** Only selected players can join.
** If one disconnects then he will be allowed to come back when reconnected.

