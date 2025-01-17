package wasmtest

import (
	"os"
	"testing"

	"github.com/suborbital/appspec/capabilities"
	"github.com/suborbital/e2core/scheduler"

	"github.com/suborbital/sat/api"
	"github.com/suborbital/sat/engine"
)

func TestDisabledHTTP(t *testing.T) {
	config := capabilities.DefaultCapabilityConfig()
	config.HTTP = &capabilities.HTTPConfig{Enabled: false}

	api, _ := api.NewWithConfig(config)

	e := engine.NewWithAPI(api)

	// test a WASM module that is loaded directly instead of through the bundle
	doWasm, _ := e.RegisterFromFile("wasm", "../testdata/fetch/fetch.wasm")

	_, err := doWasm("https://1password.com").Then()
	if err != nil {
		if err.Error() != `{"code":1,"message":"capability is not enabled"}` {
			t.Error("received incorrect error", err.Error())
		}
	} else {
		t.Error("runnable should have failed")
	}
}

func TestDisabledGraphQL(t *testing.T) {
	// bail out if GitHub auth is not set up (i.e. in Travis)
	// we want the Runnable to fail because graphql is disabled,
	// not because auth isn't set up correctly
	if _, ok := os.LookupEnv("GITHUB_TOKEN"); !ok {
		return
	}

	config := capabilities.DefaultCapabilityConfig()
	config.GraphQL = &capabilities.GraphQLConfig{Enabled: false}
	config.Auth = &capabilities.AuthConfig{
		Enabled: true,
		Headers: map[string]capabilities.AuthHeader{
			"api.github.com": {
				HeaderType: "bearer",
				Value:      "env(GITHUB_TOKEN)",
			},
		},
	}

	api, _ := api.NewWithConfig(config)

	e := engine.NewWithAPI(api)

	e.RegisterFromFile("rs-graphql", "../testdata/rs-graphql/rs-graphql.wasm")

	_, err := e.Do(scheduler.NewJob("rs-graphql", nil)).Then()
	if err != nil {
		if err.Error() != `{"code":1,"message":"capability is not enabled"}` {
			t.Error("received incorrect error ", err.Error())
		}
	} else {
		t.Error("runnable should have produced an error")
	}
}
