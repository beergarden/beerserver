# beerserver

An API server for home brewery.

## Prerequisite

- Go
- [Go workspace](https://golang.org/doc/code.html)
- [Godep](https://github.com/tools/godep) `go get github.com/tools/godep`
- Heroku Toolbelt (for deployment to Heroku)
- MongoDB

## Usage

Install dependencies:

```
go get github.com/tools/godep
godep restore
```

Create `.env`:

```
MONGOLAB_URI=localhost:27017/beerserver
```

Run server:

```
go install && beerserver
```

and open http://localhost:3000

## API

### Channels

```
GET /channels
GET /channels/{id} {"name": "Beer Temperture"}
POST /channels
```

## Datapoints

```
GET /channels/{channelId}/datapoints
POST /channels/{channelId}/datapoints {"at": "2012-04-23T18:25:43.511Z", value: 123.456}
```

## Test

```
go test
```

## Dependencies

Dependencies are managed with [Godep](https://github.com/tools/godep). Whenever you add new dependency, make sure to:

```
godep save
```

and **check the content of `Godep` directory into Git**.

## Deployment to Heroku

```
git push heroku master
heroku open
```
