# Translation App in Golang
Command line app that can translate text

This was created initially to be submitted to Topcoder Development Practice Challenge.

## Installation

Clone the repository
```bash
git clone https://github.com/Rishav-12/translation-app.git
```

Of course, you need to have Go installed on your machine.

## Usage

Run
```bash
go build main.go
```
to create the executable

Use the following command
```bash
./main [source language] [target language] [query text]
```

For example
```bash
./main en de "good morning"
```

gives the output
```bash
[map[backend:1 orig:good morning trans:Guten Morgen] map[src_translit:ɡo͝od ˈmôrniNG]]
```

As we can see, we get the translation as "Guten Morgen".
