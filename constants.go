package fordefi

// API Paths

const (
	CreateTransactionPath  = "/api/v1/transactions"
	ApproveTransactionPath = "/api/v1/transactions/%s/approve"
	AbortTransactionPath   = "/api/v1/transactions/%s/abort"
	ReleaseTransactionPath = "/api/v1/transactions/%s/release"
	ListChainsPath         = "/api/v1/blockchains"
	GetTransactionPath     = "/api/v1/transactions/%s"
	ListTransactionPath    = "/api/v1/transactions"
	ListVaultsPath         = "/api/v1/vaults"
)

// Chain Names

const (
	ArbitrumMainnet = "evm_arbitrum_mainnet"
	AvalancheCChain = "evm_avalanche_chain"
	BaseMainnet     = "evm_base_mainnet"
	BitcoinMainnet  = "bitcoin_mainnet"
	BitcoinTestnet  = "bitcoin_testnet"
	BscMainnet      = "evm_bsc_mainnet"
	CantoMainnet    = "evm_canto_mainnet"
	ConfluxESpace   = "evm_conflux_mainnet"
	CosmosMainnet   = "cosmos_cosmoshub-4"
	CosmosDydxNet   = "cosmos_dydx-mainnet-1"
	CosmosDydxTest  = "cosmos_dydx-testnet-4"
	SepoliaTestnet  = "evm_ethereum_sepolia"
	EthereumMainnet = "ethereum_mainnet"

	SolanaMainnet = "solana_mainnet"
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

const (
	StatusWaitingForApproval = "waiting_for_approval"
	StatusApproved           = "approved"
	StatusPushedToBlockchain = "pushed_to_blockchain"
	StatusCompleted          = "completed"
	StatusReverted           = "reverted"
	StatusStuck              = "stuck"
	StatusAborted            = "aborted"
)

var OnChainStatuses = map[string]bool{
	StatusPushedToBlockchain: true,
	StatusCompleted:          true,
	StatusReverted:           true,
}

// Misc Values

const MaxRepr = -1
