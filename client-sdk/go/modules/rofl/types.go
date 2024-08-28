package rofl

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/x25519"

	beacon "github.com/oasisprotocol/oasis-core/go/beacon/api"
	"github.com/oasisprotocol/oasis-core/go/common/crypto/signature"
	"github.com/oasisprotocol/oasis-core/go/common/node"

	"github.com/oasisprotocol/oasis-sdk/client-sdk/go/types"
)

// Create new ROFL application call.
type Create struct {
	// Policy is the application authentication policy.
	Policy AppAuthPolicy `json:"policy"`
	// Scheme is the identifier generation scheme.
	Scheme IdentifierScheme `json:"scheme"`
}

// IdentifierScheme is a ROFL application identifier generation scheme.
type IdentifierScheme uint8

const (
	CreatorRoundIndex IdentifierScheme = 0
	CreatorNonce      IdentifierScheme = 1
)

// Update an existing ROFL application call.
type Update struct {
	// ID is the application identifier.
	ID AppID `json:"id"`
	// Policy is the application authentication policy.
	Policy AppAuthPolicy `json:"policy"`
	// Admin is the application administrator address.
	Admin *types.Address `json:"admin"`
}

// Remove an existing ROFL application call.
type Remove struct {
	// ID is the application identifier.
	ID AppID `json:"id"`
}

// Register a ROFL application instance call.
type Register struct {
	// App is the application identifier of the app the caller is registering for.
	App AppID `json:"app"`
	// EndorsedCapability is the endorsed TEE capability.
	EndorsedCapability node.EndorsedCapabilityTEE `json:"ect"` //nolint: misspell
	// Expiration is the epoch when the ROFL registration expires if not renewed.
	Expiration beacon.EpochTime `json:"expiration"`
	// ExtraKeys are the extra public keys to endorse (e.g. secp256k1 keys).
	//
	// All of these keys need to co-sign the registration transaction to prove ownership.
	ExtraKeys []types.PublicKey `json:"extra_keys"`
}

// AppQuery is an application-related query.
type AppQuery struct {
	// ID is the application identifier.
	ID AppID `json:"id"`
}

// AppConfig is a ROFL application configuration.
type AppConfig struct {
	// ID is the application identifier.
	ID AppID `json:"id"`
	// Policy is the application authentication policy.
	Policy AppAuthPolicy `json:"policy"`
	// Admin is the application administrator address.
	Admin *types.Address `json:"admin"`
	// Stake is the staked amount.
	Stake types.BaseUnits `json:"stake"`
}

// Registration is a ROFL enclave registration descriptor.
type Registration struct {
	// App is the application this enclave is registered for.
	App AppID `json:"app"`
	// NodeID is the identifier of the endorsing node.
	NodeID signature.PublicKey `json:"node_id"`
	// RAK is the Runtime Attestation Key.
	RAK signature.PublicKey `json:"rak"`
	// REK is the Runtime Encryption Key.
	REK x25519.PublicKey `json:"rek"`
	// Expiration is the epoch when the ROFL registration expires if not renewed.
	Expiration beacon.EpochTime `json:"expiration"`
	// ExtraKeys are the extra public keys to endorse (e.g. secp256k1 keys).
	ExtraKeys []types.PublicKey `json:"extra_keys"`
}

// Parameters are the parameters for the rofl module.
type Parameters struct{}

// ModuleName is the rofl module name.
const ModuleName = "rofl"

const (
	// AppCreatedEventCode is the event code for the application created event.
	AppCreatedEventCode = 1
	// AppUpdatedEventCode is the event code for the application updated event.
	AppUpdatedEventCode = 2
	// AppRemovedEventCode is the event code for the application removed event.
	AppRemovedEventCode = 3
)

// AppCreatedEvent is an application created event.
type AppCreatedEvent struct {
	ID AppID `json:"id"`
}

// AppUpdatedEvent is an application updated event.
type AppUpdatedEvent struct {
	ID AppID `json:"id"`
}

// AppRemovedEvent is an application removed event.
type AppRemovedEvent struct {
	ID AppID `json:"id"`
}

// Event is a rofl module event.
type Event struct {
	AppCreated *AppCreatedEvent
	AppUpdated *AppUpdatedEvent
	AppRemoved *AppRemovedEvent
}
