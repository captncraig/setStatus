# setStatus
Command line utility for setting github commit status

Useful for ci environments like travis ci to set more specific commit statuses on your prs.

## Usage:

install:
`go get github.com/captncraig/setStatus`

setup a personal access token with repo:status scope, then:
`export SETSTATUS_TOKEN=9823748927349238947`

run:
`setStatus -o myGithubName -r myRepo -s success -c tests -d="unit Tests failed" -sha=$commit`
