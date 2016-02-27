# Pigeon App

Command `pigeon-app` is the beginning of the pigeon server, serving the image files to submit to the Google Cloud Vision API and return to client.


## Prerequisite

You need to export a service account json file to `GOOGLE_APPLICATION_CREDENTIALS` variable.

```
$ export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service_account.json
```


## Installation

Type the following line to install `pigeon`.

```shell
$ go get github.com/kaneshin/pigeon/tools/cmd/pigeon-app
```

Make sure that `pigeon` was installed correctly:

```shell
$ pigeon-app -h
```


## Run

```
$ pigeon-app -port=8080 -- -face -label -safe-search
```

## Example

![pigeon-app](https://raw.githubusercontent.com/kaneshin/pigeon/master/assets/pigeon-app.gif)


## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)


## Author

Shintaro Kaneko <kaneshin0120@gmail.com>

