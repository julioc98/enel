package api

import (
	"context"
	"errors"

	firebase "firebase.google.com/go"
	joker "github.com/julioc98/enel/firestore/joker"
	"google.golang.org/api/option"
)

// Add ...
func Add(add joker.Joker) (joker.Joker, error) {
	ctx := context.Background()
	opt := option.WithCredentialsFile("./key.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, errors.New("Foi Aqui 1")
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, errors.New("Foi Aqui 2")
	}
	defer client.Close()

	cli, _, err := client.Collection("users").Add(ctx, add)
	if err != nil {
		return nil, errors.New("Foi Aqui 3")
	}

	dsnap, err := cli.Get(ctx)
	if err != nil {
		return nil, errors.New("Foi Aqui 4")
	}
	m := dsnap.Data()
	// fmt.Printf("Document data: %#v\n", m)

	return m, nil
}
