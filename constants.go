package fordefi

// API Paths

const (
	CreateTransactionPath  = "/api/v1/transactions"
	ApproveTransactionPath = "/api/v1/transactions/%s/approve"
	AbortTransactionPath   = "/api/v1/transactions/%s/abort"
	ReleaseTransactionPath = "/api/v1/transactions/%s/release"
)

// Chain Names

const (
	EthereumMainnet = "ethereum_mainnet"
	SolanaMainnet   = "solana_mainnet"
)

// Send Tx Payload Modes

const (
	ModeByFunctionName = "full_details"
	ModeByHexData      = "hex"
	ModeByBase64Data   = "base64"
)

const (
	SignPersonal = "personal_message_type"
	SignTyped    = "typed_message_type"
)

const (
	ReleaseCancel     = "cancel"
	ReleaseAccelerate = "accelerate"
)

// Misc Values

const MaxRepr = -1
