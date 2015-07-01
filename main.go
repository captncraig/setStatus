package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/google/go-github/github"
)

type myRoundTripper struct {
	accessToken string
}

func (rt myRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Authorization", fmt.Sprintf("token %s", rt.accessToken))
	return http.DefaultTransport.RoundTrip(r)
}

var (
	owner = flag.String("o", "", "repository owner")
	repo  = flag.String("r", "", "repository name")
	ctx   = flag.String("c", "", "context of status")
	state = flag.String("s", "", "state of commit")
	desc  = flag.String("d", "", "text description")
	url   = flag.String("u", "", "target url")
	sha   = flag.String("sha", "", "sha of commit to update")
)

func main() {
	var token string
	if token = os.Getenv("SETSTATUS_TOKEN"); token == "" {
		log.Fatal("SETSTATUS_TOKEN env var required.")
	}
	flag.Parse()
	if *owner == "" {
		flag.PrintDefaults()
		log.Fatal("-o required")
	}
	if *repo == "" {
		flag.PrintDefaults()
		log.Fatal("-r required")
	}
	if *owner == "" {
		flag.PrintDefaults()
		log.Fatal("-o required")
	}
	if *sha == "" {
		flag.PrintDefaults()
		log.Fatal("-sha required")
	}
	http.DefaultClient.Transport = myRoundTripper{token}
	client := github.NewClient(http.DefaultClient)
	st := &github.RepoStatus{}
	st.State = state
	if *url != "" {
		st.TargetURL = url
	}
	if *desc != "" {
		st.Description = desc
	}
	if *ctx != "" {
		st.Context = ctx
	}
	st, _, err := client.Repositories.CreateStatus(*owner, *repo, *sha, st)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created status", *st.ID)
}
