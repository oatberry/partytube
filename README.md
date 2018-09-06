# partytube
An Internet Video Jukebox

## Description
Partytube is a command-line program that lets people send internet video URLs to be queued up and played, kinda like a jukebox for internet video! URLs are sent as simple plain text over a TCP connection, using a tool like netcat.

## Features
* fully concurrent utilizing Go's "goroutines"
* video fetching + playback done by `mpv` and `youtube-dl`

## Dependencies
* a working [Go](https://golang.org) installation (for building)
* [mpv](https://github.com/mpv-player/mpv)
* [youtube-dl](https://github.com/rg3/youtube-dl)

## Installation
To fetch, build, and install the binary into your `$GOPATH`, run:
```shell
$ go get github.com/oatberry/partytube
```

## Usage
1. Run the `partytube` binary
1. Connect to `partytube` using a program like netcat, for example:
    ```shell
    $ nc <ip-address> 2338
    ```
1. Send a video link!
    ``` shell
    partytube > https://youtu.be/dQw4w9WgXcQ
    ```

## Todo
* enable configuration of things like the port being listened on
* add the ability for users to vote to stop the current video
* web interface for video submission

## License - GPL3
Copyright © 2018 Thomas Berryhill <oats@oatberry.me>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
