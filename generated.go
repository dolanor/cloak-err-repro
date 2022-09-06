package main

import (
	"context"
	"encoding/json"

	"github.com/dagger/cloak/sdk/go/dagger"
)

func (r *query) hugo(ctx context.Context) (*Hugo, error) {

	return new(Hugo), nil

}

type query struct{}
type hugo struct{}

func main() {
	dagger.Serve(context.Background(), map[string]func(context.Context, dagger.ArgsInput) (interface{}, error){
		"Query.hugo": func(ctx context.Context, fc dagger.ArgsInput) (interface{}, error) {
			var bytes []byte
			_ = bytes
			var err error
			_ = err

			return (&query{}).hugo(ctx)
		},
		"hugo.generate": func(ctx context.Context, fc dagger.ArgsInput) (interface{}, error) {
			var bytes []byte
			_ = bytes
			var err error
			_ = err

			var src dagger.FSID

			bytes, err = json.Marshal(fc.Args["src"])
			if err != nil {
				return nil, err
			}
			if err := json.Unmarshal(bytes, &src); err != nil {
				return nil, err
			}

			return (&hugo{}).generate(ctx,

				src,
			)
		},
	})
}
