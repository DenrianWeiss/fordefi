package fordefi

import "time"

type EvmTransferReq struct {
	SignerType string `json:"signer_type"`
	Type       string `json:"type"`
	Details    struct {
		Type            string    `json:"type"`
		Gas             EvmGasOps `json:"gas"`
		To              string    `json:"to"`
		AssetIdentifier struct {
			Type    string                 `json:"type"`
			Details map[string]interface{} `json:"details"`
		} `json:"asset_identifier"`
		Value struct {
			Type  string `json:"type"`
			Value string `json:"value,omitempty"`
		} `json:"value"`
	} `json:"details"`
	Note    string `json:"note"`
	VaultId string `json:"vault_id"`
}

type EvmTransactResp struct {
	Id                     string    `json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	ModifiedAt             time.Time `json:"modified_at"`
	ManagedTransactionData struct {
		CreatedBy struct {
			Id       string `json:"id"`
			UserType string `json:"user_type"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			State    string `json:"state"`
			Role     string `json:"role"`
		} `json:"created_by"`
		AbortedBy struct {
			Id       string `json:"id"`
			UserType string `json:"user_type"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			State    string `json:"state"`
			Role     string `json:"role"`
		} `json:"aborted_by"`
		DeviceSigningRequest struct {
			Signers []struct {
				User struct {
					Id       string `json:"id"`
					UserType string `json:"user_type"`
					Name     string `json:"name"`
					Email    string `json:"email"`
					State    string `json:"state"`
					Role     string `json:"role"`
				} `json:"user"`
				ModifiedAt time.Time `json:"modified_at"`
				HasSigned  bool      `json:"has_signed"`
			} `json:"signers"`
		} `json:"device_signing_request"`
		ApprovalRequest struct {
			State      string `json:"state"`
			QuorumSize int    `json:"quorum_size"`
			Approvers  []struct {
				User struct {
					Id       string `json:"id"`
					UserType string `json:"user_type"`
					Name     string `json:"name"`
					Email    string `json:"email"`
					State    string `json:"state"`
					Role     string `json:"role"`
				} `json:"user"`
				ModifiedAt time.Time `json:"modified_at"`
				Decision   string    `json:"decision"`
			} `json:"approvers"`
			ApprovalGroups []struct {
				QuorumSize int `json:"quorum_size"`
				Approvers  []struct {
					User struct {
						Id       string `json:"id"`
						UserType string `json:"user_type"`
						Name     string `json:"name"`
						Email    string `json:"email"`
						State    string `json:"state"`
						Role     string `json:"role"`
					} `json:"user"`
					ModifiedAt time.Time `json:"modified_at"`
					Decision   string    `json:"decision"`
				} `json:"approvers"`
			} `json:"approval_groups"`
		} `json:"approval_request"`
		PolicyMatch struct {
			IsDefault  bool   `json:"is_default"`
			RuleId     string `json:"rule_id"`
			RuleName   string `json:"rule_name"`
			ActionType string `json:"action_type"`
		} `json:"policy_match"`
		SignedCreateRequest struct {
			RawData            string `json:"raw_data"`
			TimestampSignature struct {
				Signature string `json:"signature"`
				Timestamp int    `json:"timestamp"`
			} `json:"timestamp_signature"`
			UrlPath string `json:"url_path"`
		} `json:"signed_create_request"`
		SignerType string `json:"signer_type"`
		Risks      []struct {
			Type        string `json:"type"`
			Severity    string `json:"severity"`
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"risks"`
		ErrorPushingToBlockchainMessage         string `json:"error_pushing_to_blockchain_message"`
		OriginalErrorPushingToBlockchainMessage string `json:"original_error_pushing_to_blockchain_message"`
		Vault                                   struct {
			Id      string `json:"id"`
			Name    string `json:"name"`
			Address string `json:"address"`
			EndUser struct {
				Id         string `json:"id"`
				UserType   string `json:"user_type"`
				ExternalId string `json:"external_id"`
				State      string `json:"state"`
			} `json:"end_user"`
		} `json:"vault"`
		IdempotenceId string `json:"idempotence_id"`
	} `json:"managed_transaction_data"`
	Signatures []struct {
		Data     string `json:"data"`
		SignedBy struct {
			Id       string `json:"id"`
			UserType string `json:"user_type"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			State    string `json:"state"`
			Role     string `json:"role"`
		} `json:"signed_by"`
	} `json:"signatures"`
	Note                      string `json:"note"`
	SpamState                 string `json:"spam_state"`
	Type                      string `json:"type"`
	EvmTransactionTypeDetails struct {
		Type      string `json:"type"`
		Direction string `json:"direction"`
	} `json:"evm_transaction_type_details"`
	Chain struct {
		ChainType            string `json:"chain_type"`
		NamedChainId         string `json:"named_chain_id"`
		ChainId              int    `json:"chain_id"`
		UniqueId             string `json:"unique_id"`
		Name                 string `json:"name"`
		NativeCurrencySymbol string `json:"native_currency_symbol"`
		NativeCurrencyName   string `json:"native_currency_name"`
		BlockchainExplorer   struct {
			TransactionUrl       string `json:"transaction_url"`
			AddressUrl           string `json:"address_url"`
			RootUrl              string `json:"root_url"`
			TransactionFormatUrl string `json:"transaction_format_url"`
			AddressFormatUrl     string `json:"address_format_url"`
		} `json:"blockchain_explorer"`
		GasType            string `json:"gas_type"`
		SupportsSecureNode bool   `json:"supports_secure_node"`
	} `json:"chain"`
	From struct {
		Address string `json:"address"`
		Vault   struct {
			Id      string `json:"id"`
			Name    string `json:"name"`
			Address string `json:"address"`
			EndUser struct {
				Id         string `json:"id"`
				UserType   string `json:"user_type"`
				ExternalId string `json:"external_id"`
				State      string `json:"state"`
			} `json:"end_user"`
		} `json:"vault"`
		Contact struct {
			Id         string `json:"id"`
			Name       string `json:"name"`
			AddressRef struct {
				ChainType string `json:"chain_type"`
				Address   string `json:"address"`
				Chains    []struct {
					ChainType    string `json:"chain_type"`
					NamedChainId string `json:"named_chain_id"`
					ChainId      int    `json:"chain_id"`
					UniqueId     string `json:"unique_id"`
				} `json:"chains"`
			} `json:"address_ref"`
		} `json:"contact"`
		ExplorerUrl string `json:"explorer_url"`
		Contract    struct {
			Name string `json:"name"`
			Dapp struct {
				Id      string `json:"id"`
				Name    string `json:"name"`
				Url     string `json:"url"`
				LogoUrl string `json:"logo_url"`
			} `json:"dapp"`
			IsVerified bool `json:"is_verified"`
			Token      struct {
				Address struct {
					Chain struct {
						ChainType    string `json:"chain_type"`
						NamedChainId string `json:"named_chain_id"`
						ChainId      int    `json:"chain_id"`
						UniqueId     string `json:"unique_id"`
					} `json:"chain"`
					HexRepr string `json:"hex_repr"`
				} `json:"address"`
				Name     string `json:"name"`
				Symbol   string `json:"symbol"`
				Type     string `json:"type"`
				Decimals int    `json:"decimals"`
				LogoUrl  string `json:"logo_url"`
			} `json:"token"`
		} `json:"contract"`
	} `json:"from"`
	To struct {
		Address string `json:"address"`
		Vault   struct {
			Id      string `json:"id"`
			Name    string `json:"name"`
			Address string `json:"address"`
			EndUser struct {
				Id         string `json:"id"`
				UserType   string `json:"user_type"`
				ExternalId string `json:"external_id"`
				State      string `json:"state"`
			} `json:"end_user"`
		} `json:"vault"`
		Contact struct {
			Id         string `json:"id"`
			Name       string `json:"name"`
			AddressRef struct {
				ChainType string `json:"chain_type"`
				Address   string `json:"address"`
				Chains    []struct {
					ChainType    string `json:"chain_type"`
					NamedChainId string `json:"named_chain_id"`
					ChainId      int    `json:"chain_id"`
					UniqueId     string `json:"unique_id"`
				} `json:"chains"`
			} `json:"address_ref"`
		} `json:"contact"`
		ExplorerUrl string `json:"explorer_url"`
		Contract    struct {
			Name string `json:"name"`
			Dapp struct {
				Id      string `json:"id"`
				Name    string `json:"name"`
				Url     string `json:"url"`
				LogoUrl string `json:"logo_url"`
			} `json:"dapp"`
			IsVerified bool `json:"is_verified"`
			Token      struct {
				Address struct {
					Chain struct {
						ChainType    string `json:"chain_type"`
						NamedChainId string `json:"named_chain_id"`
						ChainId      int    `json:"chain_id"`
						UniqueId     string `json:"unique_id"`
					} `json:"chain"`
					HexRepr string `json:"hex_repr"`
				} `json:"address"`
				Name     string `json:"name"`
				Symbol   string `json:"symbol"`
				Type     string `json:"type"`
				Decimals int    `json:"decimals"`
				LogoUrl  string `json:"logo_url"`
			} `json:"token"`
		} `json:"contract"`
	} `json:"to"`
	Value      string `json:"value"`
	Data       string `json:"data"`
	ParsedData struct {
		Method          string `json:"method"`
		MethodArguments []struct {
			Name            string `json:"name"`
			Type            string `json:"type"`
			Value           string `json:"value"`
			EnrichedAddress struct {
				Address string `json:"address"`
				Vault   struct {
					Id      string `json:"id"`
					Name    string `json:"name"`
					Address string `json:"address"`
					EndUser struct {
						Id         string `json:"id"`
						UserType   string `json:"user_type"`
						ExternalId string `json:"external_id"`
						State      string `json:"state"`
					} `json:"end_user"`
				} `json:"vault"`
				Contact struct {
					Id         string `json:"id"`
					Name       string `json:"name"`
					AddressRef struct {
						ChainType string `json:"chain_type"`
						Address   string `json:"address"`
						Chains    []struct {
							ChainType    string `json:"chain_type"`
							NamedChainId string `json:"named_chain_id"`
							ChainId      int    `json:"chain_id"`
							UniqueId     string `json:"unique_id"`
						} `json:"chains"`
					} `json:"address_ref"`
				} `json:"contact"`
				ExplorerUrl string `json:"explorer_url"`
				Contract    struct {
					Name string `json:"name"`
					Dapp struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Url     string `json:"url"`
						LogoUrl string `json:"logo_url"`
					} `json:"dapp"`
					IsVerified bool `json:"is_verified"`
					Token      struct {
						Address struct {
							Chain struct {
								ChainType    string `json:"chain_type"`
								NamedChainId string `json:"named_chain_id"`
								ChainId      int    `json:"chain_id"`
								UniqueId     string `json:"unique_id"`
							} `json:"chain"`
							HexRepr string `json:"hex_repr"`
						} `json:"address"`
						Name     string `json:"name"`
						Symbol   string `json:"symbol"`
						Type     string `json:"type"`
						Decimals int    `json:"decimals"`
						LogoUrl  string `json:"logo_url"`
					} `json:"token"`
				} `json:"contract"`
			} `json:"enriched_address"`
		} `json:"method_arguments"`
	} `json:"parsed_data"`
	Hash  string `json:"hash"`
	Nonce int    `json:"nonce"`
	Block struct {
		Number  int       `json:"number"`
		Hash    string    `json:"hash"`
		MinedAt time.Time `json:"mined_at"`
	} `json:"block"`
	ExpectedResult struct {
		Reversion struct {
			State  string `json:"state"`
			Reason string `json:"reason"`
		} `json:"reversion"`
		GasDebit struct {
			GasUsed  string `json:"gas_used"`
			GasPrice string `json:"gas_price"`
			TotalFee string `json:"total_fee"`
		} `json:"gas_debit"`
		Effects struct {
			BalanceChanges []struct {
				Type    string `json:"type"`
				Address struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					Contact struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					ExplorerUrl string `json:"explorer_url"`
					Contract    struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Address struct {
								Chain struct {
								} `json:"chain"`
								HexRepr string `json:"hex_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Type     string `json:"type"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"address"`
				Diff  string `json:"diff"`
				Price struct {
					Price        string `json:"price"`
					FiatCurrency struct {
						CurrencySymbol string `json:"currency_symbol"`
						Decimals       int    `json:"decimals"`
					} `json:"fiat_currency"`
				} `json:"price"`
			} `json:"balance_changes"`
			Transfers []struct {
				Type string `json:"type"`
				From struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					Contact struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					ExplorerUrl string `json:"explorer_url"`
					Contract    struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Address struct {
								Chain struct {
								} `json:"chain"`
								HexRepr string `json:"hex_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Type     string `json:"type"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"from"`
				To struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					Contact struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					ExplorerUrl string `json:"explorer_url"`
					Contract    struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Address struct {
								Chain struct {
								} `json:"chain"`
								HexRepr string `json:"hex_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Type     string `json:"type"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"to"`
				Amount string `json:"amount"`
				Price  struct {
					Price        string `json:"price"`
					FiatCurrency struct {
						CurrencySymbol string `json:"currency_symbol"`
						Decimals       int    `json:"decimals"`
					} `json:"fiat_currency"`
				} `json:"price"`
			} `json:"transfers"`
			Allowances []struct {
				Type  string `json:"type"`
				Owner struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					Contact struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					ExplorerUrl string `json:"explorer_url"`
					Contract    struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Address struct {
								Chain struct {
								} `json:"chain"`
								HexRepr string `json:"hex_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Type     string `json:"type"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"owner"`
				Spender struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					Contact struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					ExplorerUrl string `json:"explorer_url"`
					Contract    struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Address struct {
								Chain struct {
								} `json:"chain"`
								HexRepr string `json:"hex_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Type     string `json:"type"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"spender"`
				Amount        string `json:"amount"`
				TokenContract struct {
					Name string `json:"name"`
					Dapp struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Url     string `json:"url"`
						LogoUrl string `json:"logo_url"`
					} `json:"dapp"`
					IsVerified bool `json:"is_verified"`
					Token      struct {
						Address struct {
							Chain struct {
								ChainType    string `json:"chain_type"`
								NamedChainId string `json:"named_chain_id"`
								ChainId      int    `json:"chain_id"`
								UniqueId     string `json:"unique_id"`
							} `json:"chain"`
							HexRepr string `json:"hex_repr"`
						} `json:"address"`
						Name     string `json:"name"`
						Symbol   string `json:"symbol"`
						Type     string `json:"type"`
						Decimals int    `json:"decimals"`
						LogoUrl  string `json:"logo_url"`
					} `json:"token"`
				} `json:"token_contract"`
				Price struct {
					Price        string `json:"price"`
					FiatCurrency struct {
						CurrencySymbol string `json:"currency_symbol"`
						Decimals       int    `json:"decimals"`
					} `json:"fiat_currency"`
				} `json:"price"`
			} `json:"allowances"`
			Bridge []struct {
				BridgeName string `json:"bridge_name"`
				Source     struct {
					Type          string `json:"type"`
					SourceAddress struct {
						Address string `json:"address"`
						Vault   struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Address string `json:"address"`
							EndUser struct {
								Id         string `json:"id"`
								UserType   string `json:"user_type"`
								ExternalId string `json:"external_id"`
								State      string `json:"state"`
							} `json:"end_user"`
						} `json:"vault"`
						Contact struct {
							Id         string `json:"id"`
							Name       string `json:"name"`
							AddressRef struct {
								ChainType string        `json:"chain_type"`
								Address   string        `json:"address"`
								Chains    []interface{} `json:"chains"`
							} `json:"address_ref"`
						} `json:"contact"`
						ExplorerUrl string `json:"explorer_url"`
						Contract    struct {
							Name string `json:"name"`
							Dapp struct {
								Id      string `json:"id"`
								Name    string `json:"name"`
								Url     string `json:"url"`
								LogoUrl string `json:"logo_url"`
							} `json:"dapp"`
							IsVerified bool `json:"is_verified"`
							Token      struct {
								Address struct {
								} `json:"address"`
								Name     string `json:"name"`
								Symbol   string `json:"symbol"`
								Type     string `json:"type"`
								Decimals int    `json:"decimals"`
								LogoUrl  string `json:"logo_url"`
							} `json:"token"`
						} `json:"contract"`
					} `json:"source_address"`
					Price struct {
						Price        string `json:"price"`
						FiatCurrency struct {
							CurrencySymbol string `json:"currency_symbol"`
							Decimals       int    `json:"decimals"`
						} `json:"fiat_currency"`
					} `json:"price"`
				} `json:"source"`
				Destination struct {
					Type        string `json:"type"`
					DestAddress struct {
						Address string `json:"address"`
						Vault   struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Address string `json:"address"`
							EndUser struct {
								Id         string `json:"id"`
								UserType   string `json:"user_type"`
								ExternalId string `json:"external_id"`
								State      string `json:"state"`
							} `json:"end_user"`
						} `json:"vault"`
						Contact struct {
							Id         string `json:"id"`
							Name       string `json:"name"`
							AddressRef struct {
								ChainType string        `json:"chain_type"`
								Address   string        `json:"address"`
								Chains    []interface{} `json:"chains"`
							} `json:"address_ref"`
						} `json:"contact"`
						ExplorerUrl string `json:"explorer_url"`
						Contract    struct {
							Name string `json:"name"`
							Dapp struct {
								Id      string `json:"id"`
								Name    string `json:"name"`
								Url     string `json:"url"`
								LogoUrl string `json:"logo_url"`
							} `json:"dapp"`
							IsVerified bool `json:"is_verified"`
							Token      struct {
								Address struct {
								} `json:"address"`
								Name     string `json:"name"`
								Symbol   string `json:"symbol"`
								Type     string `json:"type"`
								Decimals int    `json:"decimals"`
								LogoUrl  string `json:"logo_url"`
							} `json:"token"`
						} `json:"contract"`
					} `json:"dest_address"`
					DestChain struct {
						ChainType            string `json:"chain_type"`
						NamedChainId         string `json:"named_chain_id"`
						ChainId              int    `json:"chain_id"`
						UniqueId             string `json:"unique_id"`
						Name                 string `json:"name"`
						NativeCurrencySymbol string `json:"native_currency_symbol"`
						NativeCurrencyName   string `json:"native_currency_name"`
						BlockchainExplorer   struct {
							TransactionUrl       string `json:"transaction_url"`
							AddressUrl           string `json:"address_url"`
							RootUrl              string `json:"root_url"`
							TransactionFormatUrl string `json:"transaction_format_url"`
							AddressFormatUrl     string `json:"address_format_url"`
						} `json:"blockchain_explorer"`
						GasType            string `json:"gas_type"`
						SupportsSecureNode bool   `json:"supports_secure_node"`
					} `json:"dest_chain"`
				} `json:"destination"`
				Amount string `json:"amount"`
			} `json:"bridge"`
		} `json:"effects"`
	} `json:"expected_result"`
	SimulationStatusResult struct {
		SimulationStatus string `json:"simulation_status"`
		Details          string `json:"details"`
	} `json:"simulation_status_result"`
	MinedResult struct {
		Reversion struct {
			State  string `json:"state"`
			Reason string `json:"reason"`
		} `json:"reversion"`
		GasDebit struct {
			GasUsed  string `json:"gas_used"`
			GasPrice string `json:"gas_price"`
			TotalFee string `json:"total_fee"`
		} `json:"gas_debit"`
		Effects struct {
			BalanceChanges []struct {
				Type    string `json:"type"`
				Address struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					Contact struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					ExplorerUrl string `json:"explorer_url"`
					Contract    struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Address struct {
								Chain struct {
								} `json:"chain"`
								HexRepr string `json:"hex_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Type     string `json:"type"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"address"`
				Diff  string `json:"diff"`
				Price struct {
					Price        string `json:"price"`
					FiatCurrency struct {
						CurrencySymbol string `json:"currency_symbol"`
						Decimals       int    `json:"decimals"`
					} `json:"fiat_currency"`
				} `json:"price"`
			} `json:"balance_changes"`
			Transfers []struct {
				Type string `json:"type"`
				From struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					Contact struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					ExplorerUrl string `json:"explorer_url"`
					Contract    struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Address struct {
								Chain struct {
								} `json:"chain"`
								HexRepr string `json:"hex_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Type     string `json:"type"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"from"`
				To struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					Contact struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					ExplorerUrl string `json:"explorer_url"`
					Contract    struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Address struct {
								Chain struct {
								} `json:"chain"`
								HexRepr string `json:"hex_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Type     string `json:"type"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"to"`
				Amount string `json:"amount"`
				Price  struct {
					Price        string `json:"price"`
					FiatCurrency struct {
						CurrencySymbol string `json:"currency_symbol"`
						Decimals       int    `json:"decimals"`
					} `json:"fiat_currency"`
				} `json:"price"`
			} `json:"transfers"`
			Allowances []struct {
				Type  string `json:"type"`
				Owner struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					Contact struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					ExplorerUrl string `json:"explorer_url"`
					Contract    struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Address struct {
								Chain struct {
								} `json:"chain"`
								HexRepr string `json:"hex_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Type     string `json:"type"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"owner"`
				Spender struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					Contact struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					ExplorerUrl string `json:"explorer_url"`
					Contract    struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Address struct {
								Chain struct {
								} `json:"chain"`
								HexRepr string `json:"hex_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Type     string `json:"type"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"spender"`
				Amount        string `json:"amount"`
				TokenContract struct {
					Name string `json:"name"`
					Dapp struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Url     string `json:"url"`
						LogoUrl string `json:"logo_url"`
					} `json:"dapp"`
					IsVerified bool `json:"is_verified"`
					Token      struct {
						Address struct {
							Chain struct {
								ChainType    string `json:"chain_type"`
								NamedChainId string `json:"named_chain_id"`
								ChainId      int    `json:"chain_id"`
								UniqueId     string `json:"unique_id"`
							} `json:"chain"`
							HexRepr string `json:"hex_repr"`
						} `json:"address"`
						Name     string `json:"name"`
						Symbol   string `json:"symbol"`
						Type     string `json:"type"`
						Decimals int    `json:"decimals"`
						LogoUrl  string `json:"logo_url"`
					} `json:"token"`
				} `json:"token_contract"`
				Price struct {
					Price        string `json:"price"`
					FiatCurrency struct {
						CurrencySymbol string `json:"currency_symbol"`
						Decimals       int    `json:"decimals"`
					} `json:"fiat_currency"`
				} `json:"price"`
			} `json:"allowances"`
			Bridge []struct {
				BridgeName string `json:"bridge_name"`
				Source     struct {
					Type          string `json:"type"`
					SourceAddress struct {
						Address string `json:"address"`
						Vault   struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Address string `json:"address"`
							EndUser struct {
								Id         string `json:"id"`
								UserType   string `json:"user_type"`
								ExternalId string `json:"external_id"`
								State      string `json:"state"`
							} `json:"end_user"`
						} `json:"vault"`
						Contact struct {
							Id         string `json:"id"`
							Name       string `json:"name"`
							AddressRef struct {
								ChainType string        `json:"chain_type"`
								Address   string        `json:"address"`
								Chains    []interface{} `json:"chains"`
							} `json:"address_ref"`
						} `json:"contact"`
						ExplorerUrl string `json:"explorer_url"`
						Contract    struct {
							Name string `json:"name"`
							Dapp struct {
								Id      string `json:"id"`
								Name    string `json:"name"`
								Url     string `json:"url"`
								LogoUrl string `json:"logo_url"`
							} `json:"dapp"`
							IsVerified bool `json:"is_verified"`
							Token      struct {
								Address struct {
								} `json:"address"`
								Name     string `json:"name"`
								Symbol   string `json:"symbol"`
								Type     string `json:"type"`
								Decimals int    `json:"decimals"`
								LogoUrl  string `json:"logo_url"`
							} `json:"token"`
						} `json:"contract"`
					} `json:"source_address"`
					Price struct {
						Price        string `json:"price"`
						FiatCurrency struct {
							CurrencySymbol string `json:"currency_symbol"`
							Decimals       int    `json:"decimals"`
						} `json:"fiat_currency"`
					} `json:"price"`
				} `json:"source"`
				Destination struct {
					Type        string `json:"type"`
					DestAddress struct {
						Address string `json:"address"`
						Vault   struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Address string `json:"address"`
							EndUser struct {
								Id         string `json:"id"`
								UserType   string `json:"user_type"`
								ExternalId string `json:"external_id"`
								State      string `json:"state"`
							} `json:"end_user"`
						} `json:"vault"`
						Contact struct {
							Id         string `json:"id"`
							Name       string `json:"name"`
							AddressRef struct {
								ChainType string        `json:"chain_type"`
								Address   string        `json:"address"`
								Chains    []interface{} `json:"chains"`
							} `json:"address_ref"`
						} `json:"contact"`
						ExplorerUrl string `json:"explorer_url"`
						Contract    struct {
							Name string `json:"name"`
							Dapp struct {
								Id      string `json:"id"`
								Name    string `json:"name"`
								Url     string `json:"url"`
								LogoUrl string `json:"logo_url"`
							} `json:"dapp"`
							IsVerified bool `json:"is_verified"`
							Token      struct {
								Address struct {
								} `json:"address"`
								Name     string `json:"name"`
								Symbol   string `json:"symbol"`
								Type     string `json:"type"`
								Decimals int    `json:"decimals"`
								LogoUrl  string `json:"logo_url"`
							} `json:"token"`
						} `json:"contract"`
					} `json:"dest_address"`
					DestChain struct {
						ChainType            string `json:"chain_type"`
						NamedChainId         string `json:"named_chain_id"`
						ChainId              int    `json:"chain_id"`
						UniqueId             string `json:"unique_id"`
						Name                 string `json:"name"`
						NativeCurrencySymbol string `json:"native_currency_symbol"`
						NativeCurrencyName   string `json:"native_currency_name"`
						BlockchainExplorer   struct {
							TransactionUrl       string `json:"transaction_url"`
							AddressUrl           string `json:"address_url"`
							RootUrl              string `json:"root_url"`
							TransactionFormatUrl string `json:"transaction_format_url"`
							AddressFormatUrl     string `json:"address_format_url"`
						} `json:"blockchain_explorer"`
						GasType            string `json:"gas_type"`
						SupportsSecureNode bool   `json:"supports_secure_node"`
					} `json:"dest_chain"`
				} `json:"destination"`
				Amount string `json:"amount"`
			} `json:"bridge"`
		} `json:"effects"`
	} `json:"mined_result"`
	GasSubmitted struct {
		Price    string `json:"price"`
		Priority string `json:"priority"`
		Limit    string `json:"limit"`
		Type     string `json:"type"`
	} `json:"gas_submitted"`
	State        string `json:"state"`
	StateChanges []struct {
		ChangedAt time.Time `json:"changed_at"`
		Prices    struct {
			NativeCurrencyPrice struct {
				Price        string `json:"price"`
				FiatCurrency struct {
					CurrencySymbol string `json:"currency_symbol"`
					Decimals       int    `json:"decimals"`
				} `json:"fiat_currency"`
			} `json:"native_currency_price"`
			TokenPrices []struct {
				Price        string `json:"price"`
				FiatCurrency struct {
					CurrencySymbol string `json:"currency_symbol"`
					Decimals       int    `json:"decimals"`
				} `json:"fiat_currency"`
				Token struct {
					Address struct {
						Chain struct {
							ChainType    string `json:"chain_type"`
							NamedChainId string `json:"named_chain_id"`
							ChainId      int    `json:"chain_id"`
							UniqueId     string `json:"unique_id"`
						} `json:"chain"`
						HexRepr string `json:"hex_repr"`
					} `json:"address"`
					Name     string `json:"name"`
					Symbol   string `json:"symbol"`
					Type     string `json:"type"`
					Decimals int    `json:"decimals"`
					LogoUrl  string `json:"logo_url"`
				} `json:"token"`
			} `json:"token_prices"`
		} `json:"prices"`
		AssetPrices []struct {
			AssetIdentifier struct {
				Type    string `json:"type"`
				Details struct {
					Chain struct {
						ChainType            string `json:"chain_type"`
						UniqueId             string `json:"unique_id"`
						Name                 string `json:"name"`
						NativeCurrencySymbol string `json:"native_currency_symbol"`
						NativeCurrencyName   string `json:"native_currency_name"`
						BlockchainExplorer   struct {
							TransactionUrl       string `json:"transaction_url"`
							AddressUrl           string `json:"address_url"`
							RootUrl              string `json:"root_url"`
							TransactionFormatUrl string `json:"transaction_format_url"`
							AddressFormatUrl     string `json:"address_format_url"`
						} `json:"blockchain_explorer"`
					} `json:"chain"`
					Type string `json:"type"`
					Coin struct {
						Type        string `json:"type"`
						Name        string `json:"name"`
						Symbol      string `json:"symbol"`
						Display     string `json:"display"`
						Description string `json:"description"`
						Denom       string `json:"denom"`
						Decimals    int    `json:"decimals"`
						LogoUrl     string `json:"logo_url"`
						ExplorerUrl string `json:"explorer_url"`
					} `json:"coin"`
				} `json:"details"`
			} `json:"asset_identifier"`
			Price struct {
				Price        string `json:"price"`
				FiatCurrency struct {
					CurrencySymbol string `json:"currency_symbol"`
					Decimals       int    `json:"decimals"`
				} `json:"fiat_currency"`
			} `json:"price"`
		} `json:"asset_prices"`
		PreviousState string `json:"previous_state"`
		NewState      string `json:"new_state"`
	} `json:"state_changes"`
	ParentTransactionId                             string `json:"parent_transaction_id"`
	ChildTransactionId                              string `json:"child_transaction_id"`
	CurrentPrecedingPushedToBlockchainTransactionId string `json:"current_preceding_pushed_to_blockchain_transaction_id"`
	IsCancelation                                   bool   `json:"is_cancelation"`
	IsAcceleration                                  bool   `json:"is_acceleration"`
	UseSecureNode                                   bool   `json:"use_secure_node"`
	ExplorerUrl                                     string `json:"explorer_url"`
}

type EvmSignResp struct {
	Id                     string    `json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	ModifiedAt             time.Time `json:"modified_at"`
	ManagedTransactionData struct {
		CreatedBy struct {
			Id       string `json:"id"`
			UserType string `json:"user_type"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			State    string `json:"state"`
			Role     string `json:"role"`
		} `json:"created_by"`
		AbortedBy struct {
			Id       string `json:"id"`
			UserType string `json:"user_type"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			State    string `json:"state"`
			Role     string `json:"role"`
		} `json:"aborted_by"`
		DeviceSigningRequest struct {
			Signers []struct {
				User struct {
					Id       string `json:"id"`
					UserType string `json:"user_type"`
					Name     string `json:"name"`
					Email    string `json:"email"`
					State    string `json:"state"`
					Role     string `json:"role"`
				} `json:"user"`
				ModifiedAt time.Time `json:"modified_at"`
				HasSigned  bool      `json:"has_signed"`
			} `json:"signers"`
		} `json:"device_signing_request"`
		ApprovalRequest struct {
			State      string `json:"state"`
			QuorumSize int    `json:"quorum_size"`
			Approvers  []struct {
				User struct {
					Id       string `json:"id"`
					UserType string `json:"user_type"`
					Name     string `json:"name"`
					Email    string `json:"email"`
					State    string `json:"state"`
					Role     string `json:"role"`
				} `json:"user"`
				ModifiedAt time.Time `json:"modified_at"`
				Decision   string    `json:"decision"`
			} `json:"approvers"`
			ApprovalGroups []struct {
				QuorumSize int `json:"quorum_size"`
				Approvers  []struct {
					User struct {
						Id       string `json:"id"`
						UserType string `json:"user_type"`
						Name     string `json:"name"`
						Email    string `json:"email"`
						State    string `json:"state"`
						Role     string `json:"role"`
					} `json:"user"`
					ModifiedAt time.Time `json:"modified_at"`
					Decision   string    `json:"decision"`
				} `json:"approvers"`
			} `json:"approval_groups"`
		} `json:"approval_request"`
		PolicyMatch struct {
			IsDefault  bool   `json:"is_default"`
			RuleId     string `json:"rule_id"`
			RuleName   string `json:"rule_name"`
			ActionType string `json:"action_type"`
		} `json:"policy_match"`
		SignedCreateRequest struct {
			RawData            string `json:"raw_data"`
			TimestampSignature struct {
				Signature string `json:"signature"`
				Timestamp int    `json:"timestamp"`
			} `json:"timestamp_signature"`
			UrlPath string `json:"url_path"`
		} `json:"signed_create_request"`
		SignerType string `json:"signer_type"`
		Risks      []struct {
			Type        string `json:"type"`
			Severity    string `json:"severity"`
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"risks"`
		ErrorPushingToBlockchainMessage         string `json:"error_pushing_to_blockchain_message"`
		OriginalErrorPushingToBlockchainMessage string `json:"original_error_pushing_to_blockchain_message"`
		Vault                                   struct {
			Id      string `json:"id"`
			Name    string `json:"name"`
			Address string `json:"address"`
			EndUser struct {
				Id         string `json:"id"`
				UserType   string `json:"user_type"`
				ExternalId string `json:"external_id"`
				State      string `json:"state"`
			} `json:"end_user"`
		} `json:"vault"`
		IdempotenceId string `json:"idempotence_id"`
	} `json:"managed_transaction_data"`
	Signatures []struct {
		Data     string `json:"data"`
		SignedBy struct {
			Id       string `json:"id"`
			UserType string `json:"user_type"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			State    string `json:"state"`
			Role     string `json:"role"`
		} `json:"signed_by"`
	} `json:"signatures"`
	Note           string `json:"note"`
	SpamState      string `json:"spam_state"`
	Type           string `json:"type"`
	EvmMessageType string `json:"evm_message_type"`
	State          string `json:"state"`
	StateChanges   []struct {
		ChangedAt time.Time `json:"changed_at"`
		Prices    struct {
			NativeCurrencyPrice struct {
				Price        string `json:"price"`
				FiatCurrency struct {
					CurrencySymbol string `json:"currency_symbol"`
					Decimals       int    `json:"decimals"`
				} `json:"fiat_currency"`
			} `json:"native_currency_price"`
			TokenPrices []struct {
				Price        string `json:"price"`
				FiatCurrency struct {
					CurrencySymbol string `json:"currency_symbol"`
					Decimals       int    `json:"decimals"`
				} `json:"fiat_currency"`
				Token struct {
					Address struct {
						Chain struct {
							ChainType    string `json:"chain_type"`
							NamedChainId string `json:"named_chain_id"`
							ChainId      int    `json:"chain_id"`
							UniqueId     string `json:"unique_id"`
						} `json:"chain"`
						HexRepr string `json:"hex_repr"`
					} `json:"address"`
					Name     string `json:"name"`
					Symbol   string `json:"symbol"`
					Type     string `json:"type"`
					Decimals int    `json:"decimals"`
					LogoUrl  string `json:"logo_url"`
				} `json:"token"`
			} `json:"token_prices"`
		} `json:"prices"`
		AssetPrices []struct {
			AssetIdentifier struct {
				Type    string `json:"type"`
				Details struct {
					Chain struct {
						ChainType            string `json:"chain_type"`
						UniqueId             string `json:"unique_id"`
						Name                 string `json:"name"`
						NativeCurrencySymbol string `json:"native_currency_symbol"`
						NativeCurrencyName   string `json:"native_currency_name"`
						BlockchainExplorer   struct {
							TransactionUrl       string `json:"transaction_url"`
							AddressUrl           string `json:"address_url"`
							RootUrl              string `json:"root_url"`
							TransactionFormatUrl string `json:"transaction_format_url"`
							AddressFormatUrl     string `json:"address_format_url"`
						} `json:"blockchain_explorer"`
					} `json:"chain"`
					Type string `json:"type"`
					Coin struct {
						Type        string `json:"type"`
						Name        string `json:"name"`
						Symbol      string `json:"symbol"`
						Display     string `json:"display"`
						Description string `json:"description"`
						Denom       string `json:"denom"`
						Decimals    int    `json:"decimals"`
						LogoUrl     string `json:"logo_url"`
						ExplorerUrl string `json:"explorer_url"`
					} `json:"coin"`
				} `json:"details"`
			} `json:"asset_identifier"`
			Price struct {
				Price        string `json:"price"`
				FiatCurrency struct {
					CurrencySymbol string `json:"currency_symbol"`
					Decimals       int    `json:"decimals"`
				} `json:"fiat_currency"`
			} `json:"price"`
		} `json:"asset_prices"`
		PreviousState string `json:"previous_state"`
		NewState      string `json:"new_state"`
	} `json:"state_changes"`
	RawData   string `json:"raw_data"`
	TypedData struct {
		Domain struct {
			Name              string `json:"name"`
			Version           string `json:"version"`
			ChainId           int    `json:"chain_id"`
			VerifyingContract struct {
				Address string `json:"address"`
				Vault   struct {
					Id      string `json:"id"`
					Name    string `json:"name"`
					Address string `json:"address"`
					EndUser struct {
						Id         string `json:"id"`
						UserType   string `json:"user_type"`
						ExternalId string `json:"external_id"`
						State      string `json:"state"`
					} `json:"end_user"`
				} `json:"vault"`
				Contact struct {
					Id         string `json:"id"`
					Name       string `json:"name"`
					AddressRef struct {
						ChainType string `json:"chain_type"`
						Address   string `json:"address"`
						Chains    []struct {
							ChainType    string `json:"chain_type"`
							NamedChainId string `json:"named_chain_id"`
							ChainId      int    `json:"chain_id"`
							UniqueId     string `json:"unique_id"`
						} `json:"chains"`
					} `json:"address_ref"`
				} `json:"contact"`
				ExplorerUrl string `json:"explorer_url"`
				Contract    struct {
					Name string `json:"name"`
					Dapp struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Url     string `json:"url"`
						LogoUrl string `json:"logo_url"`
					} `json:"dapp"`
					IsVerified bool `json:"is_verified"`
					Token      struct {
						Address struct {
							Chain struct {
								ChainType    string `json:"chain_type"`
								NamedChainId string `json:"named_chain_id"`
								ChainId      int    `json:"chain_id"`
								UniqueId     string `json:"unique_id"`
							} `json:"chain"`
							HexRepr string `json:"hex_repr"`
						} `json:"address"`
						Name     string `json:"name"`
						Symbol   string `json:"symbol"`
						Type     string `json:"type"`
						Decimals int    `json:"decimals"`
						LogoUrl  string `json:"logo_url"`
					} `json:"token"`
				} `json:"contract"`
			} `json:"verifying_contract"`
			Salt string `json:"salt"`
		} `json:"domain"`
		Type string `json:"type"`
	} `json:"typed_data"`
	Vault struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Address string `json:"address"`
		EndUser struct {
			Id         string `json:"id"`
			UserType   string `json:"user_type"`
			ExternalId string `json:"external_id"`
			State      string `json:"state"`
		} `json:"end_user"`
	} `json:"vault"`
	Chain struct {
		ChainType            string `json:"chain_type"`
		NamedChainId         string `json:"named_chain_id"`
		ChainId              int    `json:"chain_id"`
		UniqueId             string `json:"unique_id"`
		Name                 string `json:"name"`
		NativeCurrencySymbol string `json:"native_currency_symbol"`
		NativeCurrencyName   string `json:"native_currency_name"`
		BlockchainExplorer   struct {
			TransactionUrl       string `json:"transaction_url"`
			AddressUrl           string `json:"address_url"`
			RootUrl              string `json:"root_url"`
			TransactionFormatUrl string `json:"transaction_format_url"`
			AddressFormatUrl     string `json:"address_format_url"`
		} `json:"blockchain_explorer"`
		GasType            string `json:"gas_type"`
		SupportsSecureNode bool   `json:"supports_secure_node"`
	} `json:"chain"`
}

type EvmGasOps struct {
	GasLimit string `json:"gas_limit"`
	Type     string `json:"type"`
	Details  struct {
		Type                 string `json:"type"`
		MaxPriorityFeePerGas string `json:"max_priority_fee_per_gas"`
		MaxFeePerGas         string `json:"max_fee_per_gas"`
	} `json:"details"`
}

type SolanaTransferReq struct {
	SignerType string `json:"signer_type"`
	Type       string `json:"type"`
	Details    struct {
		Type  string `json:"type"`
		To    string `json:"to"`
		Value struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"value"`
		AssetIdentifier struct {
			Type    string                 `json:"type"`
			Details map[string]interface{} `json:"details"`
		} `json:"asset_identifier"`
	} `json:"details"`
	Note    string `json:"note"`
	VaultId string `json:"vault_id"`
}

type SolanaTransactResp struct {
	Id                     string    `json:"id"`
	CreatedAt              time.Time `json:"created_at"`
	ModifiedAt             time.Time `json:"modified_at"`
	ManagedTransactionData struct {
		CreatedBy struct {
			Id       string `json:"id"`
			UserType string `json:"user_type"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			State    string `json:"state"`
			Role     string `json:"role"`
		} `json:"created_by"`
		AbortedBy struct {
			Id       string `json:"id"`
			UserType string `json:"user_type"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			State    string `json:"state"`
			Role     string `json:"role"`
		} `json:"aborted_by"`
		DeviceSigningRequest struct {
			Signers []struct {
				User struct {
					Id       string `json:"id"`
					UserType string `json:"user_type"`
					Name     string `json:"name"`
					Email    string `json:"email"`
					State    string `json:"state"`
					Role     string `json:"role"`
				} `json:"user"`
				ModifiedAt time.Time `json:"modified_at"`
				HasSigned  bool      `json:"has_signed"`
			} `json:"signers"`
		} `json:"device_signing_request"`
		ApprovalRequest struct {
			State      string `json:"state"`
			QuorumSize int    `json:"quorum_size"`
			Approvers  []struct {
				User struct {
					Id       string `json:"id"`
					UserType string `json:"user_type"`
					Name     string `json:"name"`
					Email    string `json:"email"`
					State    string `json:"state"`
					Role     string `json:"role"`
				} `json:"user"`
				ModifiedAt time.Time `json:"modified_at"`
				Decision   string    `json:"decision"`
			} `json:"approvers"`
			ApprovalGroups []struct {
				QuorumSize int `json:"quorum_size"`
				Approvers  []struct {
					User struct {
						Id       string `json:"id"`
						UserType string `json:"user_type"`
						Name     string `json:"name"`
						Email    string `json:"email"`
						State    string `json:"state"`
						Role     string `json:"role"`
					} `json:"user"`
					ModifiedAt time.Time `json:"modified_at"`
					Decision   string    `json:"decision"`
				} `json:"approvers"`
			} `json:"approval_groups"`
		} `json:"approval_request"`
		PolicyMatch struct {
			IsDefault  bool   `json:"is_default"`
			RuleId     string `json:"rule_id"`
			RuleName   string `json:"rule_name"`
			ActionType string `json:"action_type"`
		} `json:"policy_match"`
		SignedCreateRequest struct {
			RawData            string `json:"raw_data"`
			TimestampSignature struct {
				Signature string `json:"signature"`
				Timestamp int    `json:"timestamp"`
			} `json:"timestamp_signature"`
			UrlPath string `json:"url_path"`
		} `json:"signed_create_request"`
		SignerType string `json:"signer_type"`
		Risks      []struct {
			Type        string `json:"type"`
			Severity    string `json:"severity"`
			Title       string `json:"title"`
			Description string `json:"description"`
		} `json:"risks"`
		ErrorPushingToBlockchainMessage         string `json:"error_pushing_to_blockchain_message"`
		OriginalErrorPushingToBlockchainMessage string `json:"original_error_pushing_to_blockchain_message"`
		Vault                                   struct {
			Id      string `json:"id"`
			Name    string `json:"name"`
			Address string `json:"address"`
			EndUser struct {
				Id         string `json:"id"`
				UserType   string `json:"user_type"`
				ExternalId string `json:"external_id"`
				State      string `json:"state"`
			} `json:"end_user"`
		} `json:"vault"`
		IdempotenceId string `json:"idempotence_id"`
	} `json:"managed_transaction_data"`
	Signatures []struct {
		Data     string `json:"data"`
		SignedBy struct {
			Id       string `json:"id"`
			UserType string `json:"user_type"`
			Name     string `json:"name"`
			Email    string `json:"email"`
			State    string `json:"state"`
			Role     string `json:"role"`
		} `json:"signed_by"`
	} `json:"signatures"`
	Note                         string `json:"note"`
	SpamState                    string `json:"spam_state"`
	Type                         string `json:"type"`
	SolanaTransactionTypeDetails struct {
		Type      string `json:"type"`
		Direction string `json:"direction"`
		Sender    struct {
			Address string `json:"address"`
			Vault   struct {
				Id      string `json:"id"`
				Name    string `json:"name"`
				Address string `json:"address"`
				EndUser struct {
					Id         string `json:"id"`
					UserType   string `json:"user_type"`
					ExternalId string `json:"external_id"`
					State      string `json:"state"`
				} `json:"end_user"`
			} `json:"vault"`
			ExplorerUrl string `json:"explorer_url"`
			Contact     struct {
				Id         string `json:"id"`
				Name       string `json:"name"`
				AddressRef struct {
					ChainType string `json:"chain_type"`
					Address   string `json:"address"`
					Chains    []struct {
						ChainType    string `json:"chain_type"`
						NamedChainId string `json:"named_chain_id"`
						ChainId      int    `json:"chain_id"`
						UniqueId     string `json:"unique_id"`
					} `json:"chains"`
				} `json:"address_ref"`
			} `json:"contact"`
			Contract struct {
				Name string `json:"name"`
				Dapp struct {
					Id      string `json:"id"`
					Name    string `json:"name"`
					Url     string `json:"url"`
					LogoUrl string `json:"logo_url"`
				} `json:"dapp"`
				IsVerified bool `json:"is_verified"`
				Token      struct {
					Type    string `json:"type"`
					Address struct {
						Chain struct {
							ChainType string `json:"chain_type"`
							UniqueId  string `json:"unique_id"`
						} `json:"chain"`
						Base58Repr string `json:"base58_repr"`
					} `json:"address"`
					Name     string `json:"name"`
					Symbol   string `json:"symbol"`
					Decimals int    `json:"decimals"`
					LogoUrl  string `json:"logo_url"`
				} `json:"token"`
			} `json:"contract"`
		} `json:"sender"`
		Recipient struct {
			Address string `json:"address"`
			Vault   struct {
				Id      string `json:"id"`
				Name    string `json:"name"`
				Address string `json:"address"`
				EndUser struct {
					Id         string `json:"id"`
					UserType   string `json:"user_type"`
					ExternalId string `json:"external_id"`
					State      string `json:"state"`
				} `json:"end_user"`
			} `json:"vault"`
			ExplorerUrl string `json:"explorer_url"`
			Contact     struct {
				Id         string `json:"id"`
				Name       string `json:"name"`
				AddressRef struct {
					ChainType string `json:"chain_type"`
					Address   string `json:"address"`
					Chains    []struct {
						ChainType    string `json:"chain_type"`
						NamedChainId string `json:"named_chain_id"`
						ChainId      int    `json:"chain_id"`
						UniqueId     string `json:"unique_id"`
					} `json:"chains"`
				} `json:"address_ref"`
			} `json:"contact"`
			Contract struct {
				Name string `json:"name"`
				Dapp struct {
					Id      string `json:"id"`
					Name    string `json:"name"`
					Url     string `json:"url"`
					LogoUrl string `json:"logo_url"`
				} `json:"dapp"`
				IsVerified bool `json:"is_verified"`
				Token      struct {
					Type    string `json:"type"`
					Address struct {
						Chain struct {
							ChainType string `json:"chain_type"`
							UniqueId  string `json:"unique_id"`
						} `json:"chain"`
						Base58Repr string `json:"base58_repr"`
					} `json:"address"`
					Name     string `json:"name"`
					Symbol   string `json:"symbol"`
					Decimals int    `json:"decimals"`
					LogoUrl  string `json:"logo_url"`
				} `json:"token"`
			} `json:"contract"`
		} `json:"recipient"`
	} `json:"solana_transaction_type_details"`
	Chain struct {
		ChainType            string `json:"chain_type"`
		UniqueId             string `json:"unique_id"`
		Name                 string `json:"name"`
		NativeCurrencySymbol string `json:"native_currency_symbol"`
		NativeCurrencyName   string `json:"native_currency_name"`
		BlockchainExplorer   struct {
			TransactionUrl       string `json:"transaction_url"`
			AddressUrl           string `json:"address_url"`
			RootUrl              string `json:"root_url"`
			TransactionFormatUrl string `json:"transaction_format_url"`
			AddressFormatUrl     string `json:"address_format_url"`
		} `json:"blockchain_explorer"`
	} `json:"chain"`
	State        string `json:"state"`
	StateChanges []struct {
		ChangedAt   time.Time `json:"changed_at"`
		AssetPrices []struct {
			AssetIdentifier struct {
				Type    string `json:"type"`
				Details struct {
					Chain struct {
						ChainType            string `json:"chain_type"`
						UniqueId             string `json:"unique_id"`
						Name                 string `json:"name"`
						NativeCurrencySymbol string `json:"native_currency_symbol"`
						NativeCurrencyName   string `json:"native_currency_name"`
						BlockchainExplorer   struct {
							TransactionUrl       string `json:"transaction_url"`
							AddressUrl           string `json:"address_url"`
							RootUrl              string `json:"root_url"`
							TransactionFormatUrl string `json:"transaction_format_url"`
							AddressFormatUrl     string `json:"address_format_url"`
						} `json:"blockchain_explorer"`
					} `json:"chain"`
					Type string `json:"type"`
					Coin struct {
						Type        string `json:"type"`
						Name        string `json:"name"`
						Symbol      string `json:"symbol"`
						Display     string `json:"display"`
						Description string `json:"description"`
						Denom       string `json:"denom"`
						Decimals    int    `json:"decimals"`
						LogoUrl     string `json:"logo_url"`
						ExplorerUrl string `json:"explorer_url"`
					} `json:"coin"`
				} `json:"details"`
			} `json:"asset_identifier"`
			Price struct {
				Price        string `json:"price"`
				FiatCurrency struct {
					CurrencySymbol string `json:"currency_symbol"`
					Decimals       int    `json:"decimals"`
				} `json:"fiat_currency"`
			} `json:"price"`
		} `json:"asset_prices"`
		PreviousState string `json:"previous_state"`
		NewState      string `json:"new_state"`
	} `json:"state_changes"`
	Version      string `json:"version"`
	Instructions []struct {
		ProgramIndex   int    `json:"program_index"`
		Data           string `json:"data"`
		AccountIndexes []int  `json:"account_indexes"`
	} `json:"instructions"`
	Accounts []struct {
		Address struct {
			Address string `json:"address"`
			Vault   struct {
				Id      string `json:"id"`
				Name    string `json:"name"`
				Address string `json:"address"`
				EndUser struct {
					Id         string `json:"id"`
					UserType   string `json:"user_type"`
					ExternalId string `json:"external_id"`
					State      string `json:"state"`
				} `json:"end_user"`
			} `json:"vault"`
			ExplorerUrl string `json:"explorer_url"`
			Contact     struct {
				Id         string `json:"id"`
				Name       string `json:"name"`
				AddressRef struct {
					ChainType string `json:"chain_type"`
					Address   string `json:"address"`
					Chains    []struct {
						ChainType    string `json:"chain_type"`
						NamedChainId string `json:"named_chain_id"`
						ChainId      int    `json:"chain_id"`
						UniqueId     string `json:"unique_id"`
					} `json:"chains"`
				} `json:"address_ref"`
			} `json:"contact"`
			Contract struct {
				Name string `json:"name"`
				Dapp struct {
					Id      string `json:"id"`
					Name    string `json:"name"`
					Url     string `json:"url"`
					LogoUrl string `json:"logo_url"`
				} `json:"dapp"`
				IsVerified bool `json:"is_verified"`
				Token      struct {
					Type    string `json:"type"`
					Address struct {
						Chain struct {
							ChainType string `json:"chain_type"`
							UniqueId  string `json:"unique_id"`
						} `json:"chain"`
						Base58Repr string `json:"base58_repr"`
					} `json:"address"`
					Name     string `json:"name"`
					Symbol   string `json:"symbol"`
					Decimals int    `json:"decimals"`
					LogoUrl  string `json:"logo_url"`
				} `json:"token"`
			} `json:"contract"`
		} `json:"address"`
		Writable bool `json:"writable"`
		Signer   bool `json:"signer"`
	} `json:"accounts"`
	AddressTableLookups []struct {
		AccountKey struct {
			Chain struct {
				ChainType string `json:"chain_type"`
				UniqueId  string `json:"unique_id"`
			} `json:"chain"`
			Base58Repr string `json:"base58_repr"`
		} `json:"account_key"`
		WritableIndexes []int `json:"writable_indexes"`
		ReadonlyIndexes []int `json:"readonly_indexes"`
	} `json:"address_table_lookups"`
	Sender struct {
		Address string `json:"address"`
		Vault   struct {
			Id      string `json:"id"`
			Name    string `json:"name"`
			Address string `json:"address"`
			EndUser struct {
				Id         string `json:"id"`
				UserType   string `json:"user_type"`
				ExternalId string `json:"external_id"`
				State      string `json:"state"`
			} `json:"end_user"`
		} `json:"vault"`
		ExplorerUrl string `json:"explorer_url"`
		Contact     struct {
			Id         string `json:"id"`
			Name       string `json:"name"`
			AddressRef struct {
				ChainType string `json:"chain_type"`
				Address   string `json:"address"`
				Chains    []struct {
					ChainType    string `json:"chain_type"`
					NamedChainId string `json:"named_chain_id"`
					ChainId      int    `json:"chain_id"`
					UniqueId     string `json:"unique_id"`
				} `json:"chains"`
			} `json:"address_ref"`
		} `json:"contact"`
		Contract struct {
			Name string `json:"name"`
			Dapp struct {
				Id      string `json:"id"`
				Name    string `json:"name"`
				Url     string `json:"url"`
				LogoUrl string `json:"logo_url"`
			} `json:"dapp"`
			IsVerified bool `json:"is_verified"`
			Token      struct {
				Type    string `json:"type"`
				Address struct {
					Chain struct {
						ChainType string `json:"chain_type"`
						UniqueId  string `json:"unique_id"`
					} `json:"chain"`
					Base58Repr string `json:"base58_repr"`
				} `json:"address"`
				Name     string `json:"name"`
				Symbol   string `json:"symbol"`
				Decimals int    `json:"decimals"`
				LogoUrl  string `json:"logo_url"`
			} `json:"token"`
		} `json:"contract"`
	} `json:"sender"`
	Hash            string `json:"hash"`
	RecentBlockhash string `json:"recent_blockhash"`
	Block           struct {
		Slot    int       `json:"slot"`
		Hash    string    `json:"hash"`
		MinedAt time.Time `json:"mined_at"`
	} `json:"block"`
	ExpectedResult struct {
		TransactionError string `json:"transaction_error"`
		Fee              string `json:"fee"`
		Effects          struct {
			BalanceChanges []struct {
				Type    string `json:"type"`
				Address struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					ExplorerUrl string `json:"explorer_url"`
					Contact     struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					Contract struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Type    string `json:"type"`
							Address struct {
								Chain struct {
								} `json:"chain"`
								Base58Repr string `json:"base58_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"address"`
				Diff  string `json:"diff"`
				Price struct {
					Price        string `json:"price"`
					FiatCurrency struct {
						CurrencySymbol string `json:"currency_symbol"`
						Decimals       int    `json:"decimals"`
					} `json:"fiat_currency"`
				} `json:"price"`
			} `json:"balance_changes"`
			Transfers []struct {
				Type string `json:"type"`
				From struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					ExplorerUrl string `json:"explorer_url"`
					Contact     struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					Contract struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Type    string `json:"type"`
							Address struct {
								Chain struct {
								} `json:"chain"`
								Base58Repr string `json:"base58_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"from"`
				To struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					ExplorerUrl string `json:"explorer_url"`
					Contact     struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					Contract struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Type    string `json:"type"`
							Address struct {
								Chain struct {
								} `json:"chain"`
								Base58Repr string `json:"base58_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"to"`
				Amount string `json:"amount"`
				Price  struct {
					Price        string `json:"price"`
					FiatCurrency struct {
						CurrencySymbol string `json:"currency_symbol"`
						Decimals       int    `json:"decimals"`
					} `json:"fiat_currency"`
				} `json:"price"`
			} `json:"transfers"`
		} `json:"effects"`
		InstructionError struct {
			ErrorType        string `json:"error_type"`
			InstructionIndex int    `json:"instruction_index"`
			ErrorDescription string `json:"error_description"`
		} `json:"instruction_error"`
	} `json:"expected_result"`
	MinedResult struct {
		TransactionError string `json:"transaction_error"`
		Fee              string `json:"fee"`
		Effects          struct {
			BalanceChanges []struct {
				Type    string `json:"type"`
				Address struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					ExplorerUrl string `json:"explorer_url"`
					Contact     struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					Contract struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Type    string `json:"type"`
							Address struct {
								Chain struct {
								} `json:"chain"`
								Base58Repr string `json:"base58_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"address"`
				Diff  string `json:"diff"`
				Price struct {
					Price        string `json:"price"`
					FiatCurrency struct {
						CurrencySymbol string `json:"currency_symbol"`
						Decimals       int    `json:"decimals"`
					} `json:"fiat_currency"`
				} `json:"price"`
			} `json:"balance_changes"`
			Transfers []struct {
				Type string `json:"type"`
				From struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					ExplorerUrl string `json:"explorer_url"`
					Contact     struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					Contract struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Type    string `json:"type"`
							Address struct {
								Chain struct {
								} `json:"chain"`
								Base58Repr string `json:"base58_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"from"`
				To struct {
					Address string `json:"address"`
					Vault   struct {
						Id      string `json:"id"`
						Name    string `json:"name"`
						Address string `json:"address"`
						EndUser struct {
							Id         string `json:"id"`
							UserType   string `json:"user_type"`
							ExternalId string `json:"external_id"`
							State      string `json:"state"`
						} `json:"end_user"`
					} `json:"vault"`
					ExplorerUrl string `json:"explorer_url"`
					Contact     struct {
						Id         string `json:"id"`
						Name       string `json:"name"`
						AddressRef struct {
							ChainType string `json:"chain_type"`
							Address   string `json:"address"`
							Chains    []struct {
							} `json:"chains"`
						} `json:"address_ref"`
					} `json:"contact"`
					Contract struct {
						Name string `json:"name"`
						Dapp struct {
							Id      string `json:"id"`
							Name    string `json:"name"`
							Url     string `json:"url"`
							LogoUrl string `json:"logo_url"`
						} `json:"dapp"`
						IsVerified bool `json:"is_verified"`
						Token      struct {
							Type    string `json:"type"`
							Address struct {
								Chain struct {
								} `json:"chain"`
								Base58Repr string `json:"base58_repr"`
							} `json:"address"`
							Name     string `json:"name"`
							Symbol   string `json:"symbol"`
							Decimals int    `json:"decimals"`
							LogoUrl  string `json:"logo_url"`
						} `json:"token"`
					} `json:"contract"`
				} `json:"to"`
				Amount string `json:"amount"`
				Price  struct {
					Price        string `json:"price"`
					FiatCurrency struct {
						CurrencySymbol string `json:"currency_symbol"`
						Decimals       int    `json:"decimals"`
					} `json:"fiat_currency"`
				} `json:"price"`
			} `json:"transfers"`
		} `json:"effects"`
		InstructionError struct {
			ErrorType        string `json:"error_type"`
			InstructionIndex int    `json:"instruction_index"`
			ErrorDescription string `json:"error_description"`
		} `json:"instruction_error"`
	} `json:"mined_result"`
	ExplorerUrl string `json:"explorer_url"`
}
