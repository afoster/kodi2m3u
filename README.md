# Kodi to m3u

Suppose you're running [Kodi](https://kodi.tv/) and you've loaded a whole bunch of music videos into its library. Thanks to the magic of [theaudiodb.com](http://www.theaudiodb.com/) or similar scrapers, it's already organised your tracks into genres.

Now you want a playlist for each genre. This tool will export an [m3u playlist](https://en.wikipedia.org/wiki/M3U) for every genre Kodi has stored in its database.

## Install

* [Download the latest binary for your platform](https://github.com/afoster/kodi2m3u/releases/latest), or
* `go get github.com/afoster/kodi2m3u` to build from source

## Usage

First, find your [Kodi userdata folder](http://kodi.wiki/view/Userdata). That's where the Kodi database sits, and also where you'll need to copy the generated playlist files.

The database file should be in the `Database` sub-folder. On my system it's called `MyVideos107.db` (the `107` represents the Kodi database version).

Now run the script:

`kodi2m3u <path to your MyVideos107.db> <output folder>`

Now copy the m3u files into your Kodi `playlists` folder. Restart Kodi.

## Notes

This is my first try at [Golang](https://golang.org/), be kind.

