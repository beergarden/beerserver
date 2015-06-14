# beerserver

An API server for home brewery.

## Prerequisite

- Go
- [Go workspace](https://golang.org/doc/code.html)
- [Godep](https://github.com/tools/godep) `go get github.com/tools/godep`
- Node, NPM
- Heroku Toolbelt (for deployment to Heroku)
- MongoDB
- [flow](http://flowtype.org/) (optional)

## Usage

Install Go dependencies:

```
go get github.com/tools/godep
godep restore
```

Install Node dependencies and build assets:

```
npm install
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

### Datapoints

```
GET /channels/{channelId}/datapoints
POST /channels/{channelId}/datapoints {"at": "2012-04-23T18:25:43.511Z", value: 123.456}
```

## Front-end development

Asset files in `assets` directory are copied/compiled into `static` directory. Edit only files in `assets`.

Copy static files and compile js once:

```
npm run build
```

Copy static files once, watch js files and compile them as needed:

```
npm run watch
```

or

```
npm start
```

Statically check common errors:

```
npm run eslint
```

Statically check types:

```
flow
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

Use heroku-buildpack-multi:

```
heroku buildpacks:set https://github.com/ddollar/heroku-buildpack-multi.git
```

Install also `devDependencies` of `package.json`:

```
heroku config:set NPM_CONFIG_PRODUCTION=false
```

Deploy and open browser:

```
git push heroku master
heroku open
```
