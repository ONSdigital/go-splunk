package auth

import (
	"net/http"

	"github.com/ONSdigital/go-ns/log"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

//Client is a generic method for authenticating with google and return a
//client for the given google area.
func Client(scope string, create func(*http.Client) (interface{}, error)) interface{} {
	ctx := context.Background()

	ts, err := google.DefaultTokenSource(ctx, scope)
	if err != nil {
		log.ErrorC("authentication", err, nil)
	}
	client := oauth2.NewClient(ctx, ts)
	c, err := create(client)
	if err != nil {
		log.ErrorC("authentication", err, nil)
	}
	return c
}
