/*
Package stacksblockchainapi

This file was automatically generated by APIMATIC v3.0 ( https://www.apimatic.io ).
*/
package stacksblockchainapi

import (
	"context"
	"net/http"

	"github.com/cbadawi/stacks-go-draft/controllers"

	"github.com/apimatic/go-core-runtime/https"
)

const UserAgent = "APIMATIC 3.0"

// Client is an interface representing the main client for accessing configuration and controllers.
type ClientInterface interface {
	Configuration() *Configuration
	AccountsController() *controllers.AccountsController
	BlocksController() *controllers.BlocksController
	BurnBlocksController() *controllers.BurnBlocksController
	FaucetsController() *controllers.FaucetsController
	FeesController() *controllers.FeesController
	InfoController() *controllers.InfoController
	MicroblocksController() *controllers.MicroblocksController
	NamesController() *controllers.NamesController
	NonFungibleTokensController() *controllers.NonFungibleTokensController
	RosettaController() *controllers.RosettaController
	SearchController() *controllers.SearchController
	SmartContractsController() *controllers.SmartContractsController
	StackingRewardsController() *controllers.StackingRewardsController
	TransactionsController() *controllers.TransactionsController
	MempoolController() *controllers.MempoolController
	StackingController() *controllers.StackingController
}

// client is an implementation of the Client interface.
type client struct {
	callBuilderFactory          https.CallBuilderFactory
	config                      Configuration
	accountsController          controllers.AccountsController
	blocksController            controllers.BlocksController
	burnBlocksController        controllers.BurnBlocksController
	faucetsController           controllers.FaucetsController
	feesController              controllers.FeesController
	infoController              controllers.InfoController
	microblocksController       controllers.MicroblocksController
	namesController             controllers.NamesController
	nonFungibleTokensController controllers.NonFungibleTokensController
	rosettaController           controllers.RosettaController
	searchController            controllers.SearchController
	smartContractsController    controllers.SmartContractsController
	stackingRewardsController   controllers.StackingRewardsController
	transactionsController      controllers.TransactionsController
	mempoolController           controllers.MempoolController
	stackingController          controllers.StackingController
}

// NewClient is the constructor for creating a new client instance.
// It takes a Configuration object as a parameter and returns the Client interface.
func NewClient(configuration Configuration) ClientInterface {

	client := &client{
		config: configuration,
	}

	client.callBuilderFactory = callBuilderHandler(
		func(server string) string {
			if server == "" {
				server = "default"
			}
			baseUri := getBaseUri(Server(server), client.config)
			return baseUri
		},
		nil,
		https.NewHttpClient(configuration.HttpConfiguration()),
		configuration.httpConfiguration.RetryConfiguration(),
		withUserAgent(UserAgent),
	)

	baseController := controllers.NewBaseController(controllers.CallBuilderLogger{Cb: client, Logger: client.config.logger})
	client.accountsController = *controllers.NewAccountsController(*baseController)
	client.blocksController = *controllers.NewBlocksController(*baseController)
	client.burnBlocksController = *controllers.NewBurnBlocksController(*baseController)
	client.faucetsController = *controllers.NewFaucetsController(*baseController)
	client.feesController = *controllers.NewFeesController(*baseController)
	client.infoController = *controllers.NewInfoController(*baseController)
	client.microblocksController = *controllers.NewMicroblocksController(*baseController)
	client.namesController = *controllers.NewNamesController(*baseController)
	client.nonFungibleTokensController = *controllers.NewNonFungibleTokensController(*baseController)
	client.rosettaController = *controllers.NewRosettaController(*baseController)
	client.searchController = *controllers.NewSearchController(*baseController)
	client.smartContractsController = *controllers.NewSmartContractsController(*baseController)
	client.stackingRewardsController = *controllers.NewStackingRewardsController(*baseController)
	client.transactionsController = *controllers.NewTransactionsController(*baseController)
	client.mempoolController = *controllers.NewMempoolController(*baseController)
	client.stackingController = *controllers.NewStackingController(*baseController)
	return client
}

// Configuration returns the configuration instance of the client.
func (c *client) Configuration() *Configuration {
	return &c.config
}

// AccountsController returns the accountsController instance of the client.
func (c *client) AccountsController() *controllers.AccountsController {
	return &c.accountsController
}

// BlocksController returns the blocksController instance of the client.
func (c *client) BlocksController() *controllers.BlocksController {
	return &c.blocksController
}

// BurnBlocksController returns the burnBlocksController instance of the client.
func (c *client) BurnBlocksController() *controllers.BurnBlocksController {
	return &c.burnBlocksController
}

// FaucetsController returns the faucetsController instance of the client.
func (c *client) FaucetsController() *controllers.FaucetsController {
	return &c.faucetsController
}

// FeesController returns the feesController instance of the client.
func (c *client) FeesController() *controllers.FeesController {
	return &c.feesController
}

// InfoController returns the infoController instance of the client.
func (c *client) InfoController() *controllers.InfoController {
	return &c.infoController
}

// MicroblocksController returns the microblocksController instance of the client.
func (c *client) MicroblocksController() *controllers.MicroblocksController {
	return &c.microblocksController
}

// NamesController returns the namesController instance of the client.
func (c *client) NamesController() *controllers.NamesController {
	return &c.namesController
}

// NonFungibleTokensController returns the nonFungibleTokensController instance of the client.
func (c *client) NonFungibleTokensController() *controllers.NonFungibleTokensController {
	return &c.nonFungibleTokensController
}

// RosettaController returns the rosettaController instance of the client.
func (c *client) RosettaController() *controllers.RosettaController {
	return &c.rosettaController
}

// SearchController returns the searchController instance of the client.
func (c *client) SearchController() *controllers.SearchController {
	return &c.searchController
}

// SmartContractsController returns the smartContractsController instance of the client.
func (c *client) SmartContractsController() *controllers.SmartContractsController {
	return &c.smartContractsController
}

// StackingRewardsController returns the stackingRewardsController instance of the client.
func (c *client) StackingRewardsController() *controllers.StackingRewardsController {
	return &c.stackingRewardsController
}

// TransactionsController returns the transactionsController instance of the client.
func (c *client) TransactionsController() *controllers.TransactionsController {
	return &c.transactionsController
}

// MempoolController returns the mempoolController instance of the client.
func (c *client) MempoolController() *controllers.MempoolController {
	return &c.mempoolController
}

// StackingController returns the stackingController instance of the client.
func (c *client) StackingController() *controllers.StackingController {
	return &c.stackingController
}

// GetCallBuilder returns the CallBuilderFactory used by the client.
func (c *client) GetCallBuilder() https.CallBuilderFactory {
	return c.callBuilderFactory
}

// TODO allow users to configure base uri
// getBaseUri returns the base URI based on the server and configuration.
func getBaseUri(
	server Server,
	configuration Configuration) string {
	if configuration.Environment() == Environment(PRODUCTION) {
		if server == Server(ENUMDEFAULT) {
			return "https://api.mainnet.hiro.so/"
		}
	}
	if configuration.Environment() == Environment(ENVIRONMENT2) {
		if server == Server(ENUMDEFAULT) {
			return "https://api.testnet.hiro.so/"
		}
	}
	if configuration.Environment() == Environment(ENVIRONMENT3) {
		if server == Server(ENUMDEFAULT) {
			return "http://localhost:3999/"
		}
	}
	return "TODO: Select a valid server."
}

// clientOptions is a function type representing options for the client.
type clientOptions func(cb https.CallBuilder)

// callBuilderHandler creates the call builder factory with various options.
func callBuilderHandler(
	baseUrlProvider func(server string) string,
	auth https.Authenticator,
	httpClient https.HttpClient,
	retryConfig RetryConfiguration,
	opts ...clientOptions) https.CallBuilderFactory {
	callBuilderFactory := https.CreateCallBuilderFactory(baseUrlProvider, auth, httpClient, retryConfig)
	return tap(callBuilderFactory, opts...)
}

// tap is a utility function to apply client options to the call builder factory.
func tap(
	callBuilderFactory https.CallBuilderFactory,
	opts ...clientOptions) https.CallBuilderFactory {
	return func(ctx context.Context, httpMethod, path string) https.CallBuilder {
		callBuilder := callBuilderFactory(ctx, httpMethod, path)
		for _, opt := range opts {
			opt(callBuilder)
		}
		return callBuilder
	}
}

// withUserAgent is an option to add a user agent header to the HTTP request.
func withUserAgent(userAgent string) clientOptions {
	f := func(request *http.Request) *http.Request {
		request.Header.Add("user-agent", userAgent)
		return request
	}
	return func(cb https.CallBuilder) {
		cb.InterceptRequest(f)
	}
}
