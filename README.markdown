# haaska-go: Home Assistant Alexa Skill Adapter, but in Golang

This project does the same thing as [Haaska](https://github.com/mike-grant/haaska/), to set up you can follow the same [Wiki](https://github.com/mike-grant/haaska/wiki). The only difference is that it's written in [Golang](https://golang.org) instead of Python.

## What is the point of this?

There isn't one, really. I wrote it because I thought I could shave off a few milliseconds from my Alexa commands, because Haaska is written in Python and has some dependencies, and in theory a compiled language with all of the dependencies in the binary would be "faster".

That's not really the case, as it turns out [Python 3.7 is the fastest runtime on Lambda](https://medium.com/the-theam-journey/benchmarking-aws-lambda-runtimes-in-2019-part-i-b1ee459a293d) 🤷‍♂️. Maybe Go will get faster over time?

## Installation

Upload the zip file from Releases, or compile your own. Set `handler` to `main`, and specify the environment variables below. Publish a new version.

![Screenshot](./lambda-screenshot.png)

## Configuration

Haaska-go is configured with environment variables:

| Name                          | Description                                                                                                                                      | Default |
|-------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------|---------|
| HOME_ASSISTANT_URL            | The URL for your Home Assistant server, with no path. e.g. https://home-assistant.example.com                                                    | ""      |
| HOME_ASSISTANT_TOKEN          | A Home Assistant Long-Lived Access Token. Generate one from the UI by clicking your profile icon and scrolling to the bottom                     | ""      |
| SKIP_CERTIFICATE_VERIFICATION | Whether we should validate the TLS certificate given by home-assistant.example.com. If you're using a self-signed certificate set this to "true" | "false" |

## Compiling

`GOOS=linux go build ./main.go && zip haaska-go.zip main`
