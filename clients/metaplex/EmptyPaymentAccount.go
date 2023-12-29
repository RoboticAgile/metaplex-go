// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// At any time, the auction manager authority may empty whatever funds are in the accept payment account
// on the auction manager. Funds come here from fixed price payments for partipation nfts, and from draining bid payments
// from the auction.
//
// This action specifically takes a given safety deposit box, winning config, and creator on a metadata for the token inside that safety deposit box
// and pumps the requisite monies out to that creator as required by the royalties formula.
//
// It's up to the UI to iterate through all winning configs, all safety deposit boxes in a given winning config tier, and all creators for
// each metadata attached to each safety deposit box, to get all the money. Note that one safety deposit box can be used in multiple different winning configs,
// but this shouldn't make any difference to this function.
//
// We designed this function to be called in this loop-like manner because there is a limit to the number of accounts that can
// be passed up at once (32) and there may be many more than that easily in a given auction, so it's easier for the implementer to just
// loop through and call it, and there is an incentive for them to do so (to get paid.) It's permissionless as well as it
// will empty into any destination account owned by the creator that has the proper mint, so anybody can call it.
//
// For the participation NFT, there is no winning config, but the total is figured by summing the winning bids and subtracting
// from the total escrow amount present.
type EmptyPaymentAccount struct {
	Args *EmptyPaymentAccountArgs

	// [0] = [WRITE] acceptPayment
	// ··········· The accept payment account on the auction manager
	//
	// [1] = [WRITE] destinationAccount
	// ··········· The destination account of same mint type as the accept payment account. Must be an Associated Token Account.
	//
	// [2] = [WRITE] auctionManager
	// ··········· Auction manager
	//
	// [3] = [WRITE] payoutTicketInfo
	// ··········· Payout ticket info to keep track of this artist or auctioneer's payment, pda of [metaplex, auction manager, winning config index OR 'participation', safety deposit key]
	//
	// [4] = [SIGNER] payer
	// ··········· payer
	//
	// [5] = [] metadata
	// ··········· The metadata
	//
	// [6] = [] masterEdition
	// ··········· The master edition of the metadata (optional if exists)
	// ··········· (pda of ['metadata', program id, metadata mint id, 'edition']) - remember PDA is relative to token metadata program
	//
	// [7] = [] safetyDepositBox
	// ··········· Safety deposit box account
	//
	// [8] = [] auctionManagerStore
	// ··········· The store of the auction manager
	//
	// [9] = [] vault
	// ··········· The vault
	//
	// [10] = [] auction
	// ··········· Auction
	//
	// [11] = [] tokenProgram
	// ··········· Token program
	//
	// [12] = [] systemProgram
	// ··········· System program
	//
	// [13] = [] rentSysvar
	// ··········· Rent sysvar
	//
	// [14] = [] auctionWinnerTokenTypeTracker
	// ··········· AuctionWinnerTokenTypeTracker, pda of seed ['metaplex', program id, auction manager key, 'totals']
	//
	// [15] = [] safetyDepositConfig
	// ··········· Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewEmptyPaymentAccountInstructionBuilder creates a new `EmptyPaymentAccount` instruction builder.
func NewEmptyPaymentAccountInstructionBuilder() *EmptyPaymentAccount {
	nd := &EmptyPaymentAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 16),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *EmptyPaymentAccount) SetArgs(args EmptyPaymentAccountArgs) *EmptyPaymentAccount {
	inst.Args = &args
	return inst
}

// SetAcceptPaymentAccount sets the "acceptPayment" account.
// The accept payment account on the auction manager
func (inst *EmptyPaymentAccount) SetAcceptPaymentAccount(acceptPayment ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(acceptPayment).WRITE()
	return inst
}

// GetAcceptPaymentAccount gets the "acceptPayment" account.
// The accept payment account on the auction manager
func (inst *EmptyPaymentAccount) GetAcceptPaymentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetDestinationAccount sets the "destinationAccount" account.
// The destination account of same mint type as the accept payment account. Must be an Associated Token Account.
func (inst *EmptyPaymentAccount) SetDestinationAccount(destinationAccount ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(destinationAccount).WRITE()
	return inst
}

// GetDestinationAccount gets the "destinationAccount" account.
// The destination account of same mint type as the accept payment account. Must be an Associated Token Account.
func (inst *EmptyPaymentAccount) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetAuctionManagerAccount sets the "auctionManager" account.
// Auction manager
func (inst *EmptyPaymentAccount) SetAuctionManagerAccount(auctionManager ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(auctionManager).WRITE()
	return inst
}

// GetAuctionManagerAccount gets the "auctionManager" account.
// Auction manager
func (inst *EmptyPaymentAccount) GetAuctionManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetPayoutTicketInfoAccount sets the "payoutTicketInfo" account.
// Payout ticket info to keep track of this artist or auctioneer's payment, pda of [metaplex, auction manager, winning config index OR 'participation', safety deposit key]
func (inst *EmptyPaymentAccount) SetPayoutTicketInfoAccount(payoutTicketInfo ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(payoutTicketInfo).WRITE()
	return inst
}

// GetPayoutTicketInfoAccount gets the "payoutTicketInfo" account.
// Payout ticket info to keep track of this artist or auctioneer's payment, pda of [metaplex, auction manager, winning config index OR 'participation', safety deposit key]
func (inst *EmptyPaymentAccount) GetPayoutTicketInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPayerAccount sets the "payer" account.
// payer
func (inst *EmptyPaymentAccount) SetPayerAccount(payer ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// payer
func (inst *EmptyPaymentAccount) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMetadataAccount sets the "metadata" account.
// The metadata
func (inst *EmptyPaymentAccount) SetMetadataAccount(metadata ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// The metadata
func (inst *EmptyPaymentAccount) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMasterEditionAccount sets the "masterEdition" account.
// The master edition of the metadata (optional if exists)
// (pda of ['metadata', program id, metadata mint id, 'edition']) - remember PDA is relative to token metadata program
func (inst *EmptyPaymentAccount) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(masterEdition)
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account (optional).
// The master edition of the metadata (optional if exists)
// (pda of ['metadata', program id, metadata mint id, 'edition']) - remember PDA is relative to token metadata program
func (inst *EmptyPaymentAccount) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetSafetyDepositBoxAccount sets the "safetyDepositBox" account.
// Safety deposit box account
func (inst *EmptyPaymentAccount) SetSafetyDepositBoxAccount(safetyDepositBox ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(safetyDepositBox)
	return inst
}

// GetSafetyDepositBoxAccount gets the "safetyDepositBox" account.
// Safety deposit box account
func (inst *EmptyPaymentAccount) GetSafetyDepositBoxAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAuctionManagerStoreAccount sets the "auctionManagerStore" account.
// The store of the auction manager
func (inst *EmptyPaymentAccount) SetAuctionManagerStoreAccount(auctionManagerStore ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(auctionManagerStore)
	return inst
}

// GetAuctionManagerStoreAccount gets the "auctionManagerStore" account.
// The store of the auction manager
func (inst *EmptyPaymentAccount) GetAuctionManagerStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetVaultAccount sets the "vault" account.
// The vault
func (inst *EmptyPaymentAccount) SetVaultAccount(vault ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(vault)
	return inst
}

// GetVaultAccount gets the "vault" account.
// The vault
func (inst *EmptyPaymentAccount) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAuctionAccount sets the "auction" account.
// Auction
func (inst *EmptyPaymentAccount) SetAuctionAccount(auction ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(auction)
	return inst
}

// GetAuctionAccount gets the "auction" account.
// Auction
func (inst *EmptyPaymentAccount) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *EmptyPaymentAccount) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *EmptyPaymentAccount) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *EmptyPaymentAccount) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *EmptyPaymentAccount) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *EmptyPaymentAccount) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *EmptyPaymentAccount) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetAuctionWinnerTokenTypeTrackerAccount sets the "auctionWinnerTokenTypeTracker" account.
// AuctionWinnerTokenTypeTracker, pda of seed ['metaplex', program id, auction manager key, 'totals']
func (inst *EmptyPaymentAccount) SetAuctionWinnerTokenTypeTrackerAccount(auctionWinnerTokenTypeTracker ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(auctionWinnerTokenTypeTracker)
	return inst
}

// GetAuctionWinnerTokenTypeTrackerAccount gets the "auctionWinnerTokenTypeTracker" account.
// AuctionWinnerTokenTypeTracker, pda of seed ['metaplex', program id, auction manager key, 'totals']
func (inst *EmptyPaymentAccount) GetAuctionWinnerTokenTypeTrackerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetSafetyDepositConfigAccount sets the "safetyDepositConfig" account.
// Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
func (inst *EmptyPaymentAccount) SetSafetyDepositConfigAccount(safetyDepositConfig ag_solanago.PublicKey) *EmptyPaymentAccount {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(safetyDepositConfig)
	return inst
}

// GetSafetyDepositConfigAccount gets the "safetyDepositConfig" account.
// Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
func (inst *EmptyPaymentAccount) GetSafetyDepositConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst EmptyPaymentAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_EmptyPaymentAccount),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst EmptyPaymentAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *EmptyPaymentAccount) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AcceptPayment is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.DestinationAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AuctionManager is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PayoutTicketInfo is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Metadata is not set")
		}

		// [6] = MasterEdition is optional

		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SafetyDepositBox is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.AuctionManagerStore is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Auction is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.AuctionWinnerTokenTypeTracker is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.SafetyDepositConfig is not set")
		}
	}
	return nil
}

func (inst *EmptyPaymentAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("EmptyPaymentAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                acceptPayment", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                  destination", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               auctionManager", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             payoutTicketInfo", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                        payer", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                     metadata", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                masterEdition", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("             safetyDepositBox", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          auctionManagerStore", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                        vault", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                      auction", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                 tokenProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("                systemProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("                   rentSysvar", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("auctionWinnerTokenTypeTracker", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("          safetyDepositConfig", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj EmptyPaymentAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *EmptyPaymentAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewEmptyPaymentAccountInstruction declares a new EmptyPaymentAccount instruction with the provided parameters and accounts.
func NewEmptyPaymentAccountInstruction(
	// Parameters:
	args EmptyPaymentAccountArgs,
	// Accounts:
	acceptPayment ag_solanago.PublicKey,
	destinationAccount ag_solanago.PublicKey,
	auctionManager ag_solanago.PublicKey,
	payoutTicketInfo ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	safetyDepositBox ag_solanago.PublicKey,
	auctionManagerStore ag_solanago.PublicKey,
	vault ag_solanago.PublicKey,
	auction ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey,
	auctionWinnerTokenTypeTracker ag_solanago.PublicKey,
	safetyDepositConfig ag_solanago.PublicKey) *EmptyPaymentAccount {
	return NewEmptyPaymentAccountInstructionBuilder().
		SetArgs(args).
		SetAcceptPaymentAccount(acceptPayment).
		SetDestinationAccount(destinationAccount).
		SetAuctionManagerAccount(auctionManager).
		SetPayoutTicketInfoAccount(payoutTicketInfo).
		SetPayerAccount(payer).
		SetMetadataAccount(metadata).
		SetMasterEditionAccount(masterEdition).
		SetSafetyDepositBoxAccount(safetyDepositBox).
		SetAuctionManagerStoreAccount(auctionManagerStore).
		SetVaultAccount(vault).
		SetAuctionAccount(auction).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentSysvarAccount(rentSysvar).
		SetAuctionWinnerTokenTypeTrackerAccount(auctionWinnerTokenTypeTracker).
		SetSafetyDepositConfigAccount(safetyDepositConfig)
}
