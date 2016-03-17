# Music 

[![Build Status](https://travis-ci.org/go-music/music.svg?branch=master)](https://travis-ci.org/go-music/music) [![GoDoc](https://godoc.org/github.com/go-music/music?status.svg)](https://godoc.org/github.com/go-music/music) [![Go Report Card](https://goreportcard.com/badge/github.com/go-music/music)](https://goreportcard.com/report/github.com/go-music/music)

https://github.com/go-music/music

#### Opinionated models of certain atomic musical concepts

There's an example command-line utility `music.go` to demo the libraries:

    go run music.go -k "C minor 7"
    
This will output:
    
    root: C
    major: false
    minor: true
    tones:
      3: D#
      5: G
      7: A#

Author: [Charney Kaye](http://w.charney.io)

## Note

A Note is used to represent the relative duration and pitch of a sound.

[![GoDoc](https://godoc.org/github.com/go-music/music/note?status.svg)](https://godoc.org/github.com/go-music/music/note) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/note)

## Interval

An interval is the difference between two pitches.

[![GoDoc](https://godoc.org/github.com/go-music/music/interval?status.svg)](https://godoc.org/github.com/go-music/music/interval) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/interval)

## Key

The key of a piece is a group of pitches, or scale upon which a music composition is created in classical, Western art, and Western pop music.

[![GoDoc](https://godoc.org/github.com/go-music/music/key?status.svg)](https://godoc.org/github.com/go-music/music/key) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/key)

## Chord

A chord, in music, is any harmonic set of three or more notes that is heard as if sounding simultaneously.

[![GoDoc](https://godoc.org/github.com/go-music/music/chord?status.svg)](https://godoc.org/github.com/go-music/music/chord) [![Coverage](https://img.shields.io/badge/coverage-100%-brightgreen.svg?style=flat)](https://gocover.io/github.com/go-music/music/chord)
