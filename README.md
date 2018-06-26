### Kdakean
Kdakean is a web application written in Go(Golang). It is a task management application.

### Feature
...


### Installation
Make sure you already installed:
- [dep](https://golang.github.io/dep/docs/installation.html)
- [goose](https://bitbucket.org/liamstask/goose/src/master/)

Clone source code:
```
$ cd $GOPATH/src
$ mkdir -p github.com/kdakean
$ cd $_ && git clone https://your-forkrepo/kdakean.git
```
Create postgres user and database
```
CREATE ROLE kdakean WITH CREATEDB LOGIN ENCRYPTED PASSWORD 'YourPassword';
CREATE DATABASE kdakean;
```

Run an application:
```
$ dep ensure
$ goose up
$ ./watch.sh
```
Now you good to go...


### Development
...

### Test
...

### production
...

### Contributes
...
