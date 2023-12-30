// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// NOTE: Requires an AuctionManagerV1.
// Needs to be called by someone at the end of the auction - will use the one time authorization token
// to fire up a bunch of printing tokens for use in participation redemptions.
type DeprecatedPopulateParticipationPrintingAccount struct {

	// [0] = [WRITE] safetyDepositTokenStorage
	// ··········· Safety deposit token store
	//
	// [1] = [WRITE] transientAccount
	// ··········· Transient account with mint of one time authorization account on master edition - you can delete after this txn
	//
	// [2] = [WRITE] printingToken
	// ··········· The printing token account on the participation state of the auction manager
	//
	// [3] = [WRITE] oneTimePrintingAuthorizationMint
	// ··········· One time printing authorization mint
	//
	// [4] = [WRITE] printingMint
	// ··········· Printing mint
	//
	// [5] = [WRITE] participationPrizeSafetyDeposit
	// ··········· Safety deposit of the participation prize
	//
	// [6] = [WRITE] vaultInfo
	// ··········· Vault info
	//
	// [7] = [] fractionMint
	// ··········· Fraction mint
	//
	// [8] = [] auctionInfo
	// ··········· Auction info
	//
	// [9] = [] auctionManagerInfo
	// ··········· Auction manager info
	//
	// [10] = [] tokenProgram
	// ··········· Token program
	//
	// [11] = [] tokenVaultProgram
	// ··········· Token vault program
	//
	// [12] = [] tokenMetadataProgram
	// ··········· Token metadata program
	//
	// [13] = [] auctionManagerStore
	// ··········· Auction manager store
	//
	// [14] = [] masterEdition
	// ··········· Master edition
	//
	// [15] = [] pdaBasedTransferAuthority
	// ··········· PDA-based Transfer authority to move the tokens from the store to the destination seed ['vault', program_id]
	// ··········· but please note that this is a PDA relative to the Token Vault program, with the 'vault' prefix
	//
	// [16] = [] payerWithRefund
	// ··········· Payer who wishes to receive refund for closing of one time transient account once we're done here
	//
	// [17] = [] rent
	// ··········· Rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDeprecatedPopulateParticipationPrintingAccountInstructionBuilder creates a new `DeprecatedPopulateParticipationPrintingAccount` instruction builder.
func NewDeprecatedPopulateParticipationPrintingAccountInstructionBuilder() *DeprecatedPopulateParticipationPrintingAccount {
	nd := &DeprecatedPopulateParticipationPrintingAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 18),
	}
	return nd
}

// SetSafetyDepositTokenStorageAccount sets the "safetyDepositTokenStorage" account.
// Safety deposit token store
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetSafetyDepositTokenStorageAccount(safetyDepositTokenStorage ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(safetyDepositTokenStorage).WRITE()
	return inst
}

// GetSafetyDepositTokenStorageAccount gets the "safetyDepositTokenStorage" account.
// Safety deposit token store
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetSafetyDepositTokenStorageAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetTransientAccount sets the "transientAccount" account.
// Transient account with mint of one time authorization account on master edition - you can delete after this txn
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetTransientAccount(transientAccount ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(transientAccount).WRITE()
	return inst
}

// GetTransientAccount gets the "transientAccount" account.
// Transient account with mint of one time authorization account on master edition - you can delete after this txn
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetTransientAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPrintingTokenAccount sets the "printingToken" account.
// The printing token account on the participation state of the auction manager
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetPrintingTokenAccount(printingToken ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(printingToken).WRITE()
	return inst
}

// GetPrintingTokenAccount gets the "printingToken" account.
// The printing token account on the participation state of the auction manager
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetPrintingTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOneTimePrintingAuthorizationMintAccount sets the "oneTimePrintingAuthorizationMint" account.
// One time printing authorization mint
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetOneTimePrintingAuthorizationMintAccount(oneTimePrintingAuthorizationMint ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(oneTimePrintingAuthorizationMint).WRITE()
	return inst
}

// GetOneTimePrintingAuthorizationMintAccount gets the "oneTimePrintingAuthorizationMint" account.
// One time printing authorization mint
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetOneTimePrintingAuthorizationMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPrintingMintAccount sets the "printingMint" account.
// Printing mint
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetPrintingMintAccount(printingMint ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(printingMint).WRITE()
	return inst
}

// GetPrintingMintAccount gets the "printingMint" account.
// Printing mint
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetPrintingMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetParticipationPrizeSafetyDepositAccount sets the "participationPrizeSafetyDeposit" account.
// Safety deposit of the participation prize
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetParticipationPrizeSafetyDepositAccount(participationPrizeSafetyDeposit ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(participationPrizeSafetyDeposit).WRITE()
	return inst
}

// GetParticipationPrizeSafetyDepositAccount gets the "participationPrizeSafetyDeposit" account.
// Safety deposit of the participation prize
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetParticipationPrizeSafetyDepositAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetVaultInfoAccount sets the "vaultInfo" account.
// Vault info
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetVaultInfoAccount(vaultInfo ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(vaultInfo).WRITE()
	return inst
}

// GetVaultInfoAccount gets the "vaultInfo" account.
// Vault info
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetVaultInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetFractionMintAccount sets the "fractionMint" account.
// Fraction mint
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetFractionMintAccount(fractionMint ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(fractionMint)
	return inst
}

// GetFractionMintAccount gets the "fractionMint" account.
// Fraction mint
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetFractionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAuctionInfoAccount sets the "auctionInfo" account.
// Auction info
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetAuctionInfoAccount(auctionInfo ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(auctionInfo)
	return inst
}

// GetAuctionInfoAccount gets the "auctionInfo" account.
// Auction info
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetAuctionInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAuctionManagerInfoAccount sets the "auctionManagerInfo" account.
// Auction manager info
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetAuctionManagerInfoAccount(auctionManagerInfo ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(auctionManagerInfo)
	return inst
}

// GetAuctionManagerInfoAccount gets the "auctionManagerInfo" account.
// Auction manager info
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetAuctionManagerInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenVaultProgramAccount sets the "tokenVaultProgram" account.
// Token vault program
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetTokenVaultProgramAccount(tokenVaultProgram ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenVaultProgram)
	return inst
}

// GetTokenVaultProgramAccount gets the "tokenVaultProgram" account.
// Token vault program
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetTokenVaultProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
// Token metadata program
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
// Token metadata program
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetAuctionManagerStoreAccount sets the "auctionManagerStore" account.
// Auction manager store
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetAuctionManagerStoreAccount(auctionManagerStore ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(auctionManagerStore)
	return inst
}

// GetAuctionManagerStoreAccount gets the "auctionManagerStore" account.
// Auction manager store
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetAuctionManagerStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetMasterEditionAccount sets the "masterEdition" account.
// Master edition
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(masterEdition)
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
// Master edition
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetPdaBasedTransferAuthorityAccount sets the "pdaBasedTransferAuthority" account.
// PDA-based Transfer authority to move the tokens from the store to the destination seed ['vault', program_id]
// but please note that this is a PDA relative to the Token Vault program, with the 'vault' prefix
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetPdaBasedTransferAuthorityAccount(pdaBasedTransferAuthority ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(pdaBasedTransferAuthority)
	return inst
}

// GetPdaBasedTransferAuthorityAccount gets the "pdaBasedTransferAuthority" account.
// PDA-based Transfer authority to move the tokens from the store to the destination seed ['vault', program_id]
// but please note that this is a PDA relative to the Token Vault program, with the 'vault' prefix
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetPdaBasedTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetPayerWithRefundAccount sets the "payerWithRefund" account.
// Payer who wishes to receive refund for closing of one time transient account once we're done here
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetPayerWithRefundAccount(payerWithRefund ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(payerWithRefund)
	return inst
}

// GetPayerWithRefundAccount gets the "payerWithRefund" account.
// Payer who wishes to receive refund for closing of one time transient account once we're done here
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetPayerWithRefundAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

// SetRentAccount sets the "rent" account.
// Rent
func (inst *DeprecatedPopulateParticipationPrintingAccount) SetRentAccount(rent ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent
func (inst *DeprecatedPopulateParticipationPrintingAccount) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(17)
}

func (inst DeprecatedPopulateParticipationPrintingAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_DeprecatedPopulateParticipationPrintingAccount),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeprecatedPopulateParticipationPrintingAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeprecatedPopulateParticipationPrintingAccount) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.SafetyDepositTokenStorage is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.TransientAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PrintingToken is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.OneTimePrintingAuthorizationMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PrintingMint is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.ParticipationPrizeSafetyDeposit is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.VaultInfo is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.FractionMint is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.AuctionInfo is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.AuctionManagerInfo is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenVaultProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.AuctionManagerStore is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.PdaBasedTransferAuthority is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.PayerWithRefund is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *DeprecatedPopulateParticipationPrintingAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeprecatedPopulateParticipationPrintingAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=18]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("       safetyDepositTokenStorage", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                       transient", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                   printingToken", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("oneTimePrintingAuthorizationMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                    printingMint", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta(" participationPrizeSafetyDeposit", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                       vaultInfo", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                    fractionMint", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                     auctionInfo", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("              auctionManagerInfo", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                    tokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("               tokenVaultProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("            tokenMetadataProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("             auctionManagerStore", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("                   masterEdition", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("       pdaBasedTransferAuthority", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("                 payerWithRefund", inst.AccountMetaSlice.Get(16)))
						accountsBranch.Child(ag_format.Meta("                            rent", inst.AccountMetaSlice.Get(17)))
					})
				})
		})
}

func (obj DeprecatedPopulateParticipationPrintingAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *DeprecatedPopulateParticipationPrintingAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewDeprecatedPopulateParticipationPrintingAccountInstruction declares a new DeprecatedPopulateParticipationPrintingAccount instruction with the provided parameters and accounts.
func NewDeprecatedPopulateParticipationPrintingAccountInstruction(
	// Accounts:
	safetyDepositTokenStorage ag_solanago.PublicKey,
	transientAccount ag_solanago.PublicKey,
	printingToken ag_solanago.PublicKey,
	oneTimePrintingAuthorizationMint ag_solanago.PublicKey,
	printingMint ag_solanago.PublicKey,
	participationPrizeSafetyDeposit ag_solanago.PublicKey,
	vaultInfo ag_solanago.PublicKey,
	fractionMint ag_solanago.PublicKey,
	auctionInfo ag_solanago.PublicKey,
	auctionManagerInfo ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	tokenVaultProgram ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	auctionManagerStore ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	pdaBasedTransferAuthority ag_solanago.PublicKey,
	payerWithRefund ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *DeprecatedPopulateParticipationPrintingAccount {
	return NewDeprecatedPopulateParticipationPrintingAccountInstructionBuilder().
		SetSafetyDepositTokenStorageAccount(safetyDepositTokenStorage).
		SetTransientAccount(transientAccount).
		SetPrintingTokenAccount(printingToken).
		SetOneTimePrintingAuthorizationMintAccount(oneTimePrintingAuthorizationMint).
		SetPrintingMintAccount(printingMint).
		SetParticipationPrizeSafetyDepositAccount(participationPrizeSafetyDeposit).
		SetVaultInfoAccount(vaultInfo).
		SetFractionMintAccount(fractionMint).
		SetAuctionInfoAccount(auctionInfo).
		SetAuctionManagerInfoAccount(auctionManagerInfo).
		SetTokenProgramAccount(tokenProgram).
		SetTokenVaultProgramAccount(tokenVaultProgram).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetAuctionManagerStoreAccount(auctionManagerStore).
		SetMasterEditionAccount(masterEdition).
		SetPdaBasedTransferAuthorityAccount(pdaBasedTransferAuthority).
		SetPayerWithRefundAccount(payerWithRefund).
		SetRentAccount(rent)
}
