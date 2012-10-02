# TentServers README

TentServers is a directory of [Tent.io](https://tent.io) servers

## Usage

To add your Tent.io server to the directory, run the following command
(or the following programmatic equivalent):

    curl -X POST -d '{"author": "My Name", "url": "https://mytent.mydomain.com"}' http://tentservers.appspot.com/tents

If it worked, the response will be a JSONified list of existing
Tent.io servers, including the one you just added.  Otherwise, you'll
see an error (in plain text, not JSON).

To see the current list of Tent apps, visit
<http://tentservers.appspot.com/tents> in your browser or run

    curl http://tentservers.appspot.com/tents

at the command line.


## Why did you create this directory?

To facilitate the programmatic discovery of new Tent.io servers.


## TODO

* Some kind of auth (right now anyone can post)

* Programmatic verification that a Tent.io server is actually at the
  given URL

  * Which of the Tent.io protocol features are supported, and which
    aren't?  (This could even be used as a sort of external unit
    testing suite.)
