// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package auction

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Place a bid on a running auction.
type PlaceBid struct {
	Args *PlaceBidArgs

	// [0] = [SIGNER] biddersPrimaryAccount
	// ··········· The bidders primary account, for PDA calculation/transit auth.
	//
	// [1] = [WRITE] biddersTokenAccount
	// ··········· The bidders token account they'll pay with
	//
	// [2] = [WRITE] pot
	// ··········· The pot, containing a reference to the stored SPL token account.
	//
	// [3] = [WRITE] potSPLAccount
	// ··········· The pot SPL account, where the tokens will be deposited.
	//
	// [4] = [WRITE] metadataAccount
	// ··········· The metadata account, storing information about the bidders actions.
	//
	// [5] = [WRITE] auctionAccount
	// ··········· Auction account, containing data about the auction and item being bid on.
	//
	// [6] = [WRITE] tokenMint
	// ··········· Token mint, for transfer instructions and verification.
	//
	// [7] = [SIGNER] transferAuthority
	// ··········· Transfer authority, for moving tokens into the bid pot.
	//
	// [8] = [SIGNER] payer
	// ··········· Payer
	//
	// [9] = [] clockSysvar
	// ··········· Clock sysvar
	//
	// [10] = [] rentSysvar
	// ··········· Rent sysvar
	//
	// [11] = [] systemProgram
	// ··········· System program
	//
	// [12] = [] splTokenProgram
	// ··········· SPL Token Program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewPlaceBidInstructionBuilder creates a new `PlaceBid` instruction builder.
func NewPlaceBidInstructionBuilder() *PlaceBid {
	nd := &PlaceBid{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *PlaceBid) SetArgs(args PlaceBidArgs) *PlaceBid {
	inst.Args = &args
	return inst
}

// SetBiddersPrimaryAccount sets the "biddersPrimaryAccount" account.
// The bidders primary account, for PDA calculation/transit auth.
func (inst *PlaceBid) SetBiddersPrimaryAccount(biddersPrimaryAccount ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(biddersPrimaryAccount).SIGNER()
	return inst
}

// GetBiddersPrimaryAccount gets the "biddersPrimaryAccount" account.
// The bidders primary account, for PDA calculation/transit auth.
func (inst *PlaceBid) GetBiddersPrimaryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetBiddersTokenAccount sets the "biddersTokenAccount" account.
// The bidders token account they'll pay with
func (inst *PlaceBid) SetBiddersTokenAccount(biddersTokenAccount ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(biddersTokenAccount).WRITE()
	return inst
}

// GetBiddersTokenAccount gets the "biddersTokenAccount" account.
// The bidders token account they'll pay with
func (inst *PlaceBid) GetBiddersTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetPotAccount sets the "pot" account.
// The pot, containing a reference to the stored SPL token account.
func (inst *PlaceBid) SetPotAccount(pot ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(pot).WRITE()
	return inst
}

// GetPotAccount gets the "pot" account.
// The pot, containing a reference to the stored SPL token account.
func (inst *PlaceBid) GetPotAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetPotSPLAccount sets the "potSPLAccount" account.
// The pot SPL account, where the tokens will be deposited.
func (inst *PlaceBid) SetPotSPLAccount(potSPLAccount ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(potSPLAccount).WRITE()
	return inst
}

// GetPotSPLAccount gets the "potSPLAccount" account.
// The pot SPL account, where the tokens will be deposited.
func (inst *PlaceBid) GetPotSPLAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetMetadataAccount sets the "metadataAccount" account.
// The metadata account, storing information about the bidders actions.
func (inst *PlaceBid) SetMetadataAccount(metadataAccount ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(metadataAccount).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadataAccount" account.
// The metadata account, storing information about the bidders actions.
func (inst *PlaceBid) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetAuctionAccount sets the "auctionAccount" account.
// Auction account, containing data about the auction and item being bid on.
func (inst *PlaceBid) SetAuctionAccount(auctionAccount ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(auctionAccount).WRITE()
	return inst
}

// GetAuctionAccount gets the "auctionAccount" account.
// Auction account, containing data about the auction and item being bid on.
func (inst *PlaceBid) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetTokenMintAccount sets the "tokenMint" account.
// Token mint, for transfer instructions and verification.
func (inst *PlaceBid) SetTokenMintAccount(tokenMint ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenMint).WRITE()
	return inst
}

// GetTokenMintAccount gets the "tokenMint" account.
// Token mint, for transfer instructions and verification.
func (inst *PlaceBid) GetTokenMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetTransferAuthorityAccount sets the "transferAuthority" account.
// Transfer authority, for moving tokens into the bid pot.
func (inst *PlaceBid) SetTransferAuthorityAccount(transferAuthority ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(transferAuthority).SIGNER()
	return inst
}

// GetTransferAuthorityAccount gets the "transferAuthority" account.
// Transfer authority, for moving tokens into the bid pot.
func (inst *PlaceBid) GetTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *PlaceBid) SetPayerAccount(payer ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *PlaceBid) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetClockSysvarAccount sets the "clockSysvar" account.
// Clock sysvar
func (inst *PlaceBid) SetClockSysvarAccount(clockSysvar ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(clockSysvar)
	return inst
}

// GetClockSysvarAccount gets the "clockSysvar" account.
// Clock sysvar
func (inst *PlaceBid) GetClockSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *PlaceBid) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *PlaceBid) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *PlaceBid) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *PlaceBid) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

// SetSplTokenProgramAccount sets the "splTokenProgram" account.
// SPL Token Program
func (inst *PlaceBid) SetSplTokenProgramAccount(splTokenProgram ag_solanago.PublicKey) *PlaceBid {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(splTokenProgram)
	return inst
}

// GetSplTokenProgramAccount gets the "splTokenProgram" account.
// SPL Token Program
func (inst *PlaceBid) GetSplTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[12]
}

func (inst PlaceBid) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_PlaceBid),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst PlaceBid) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *PlaceBid) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.BiddersPrimaryAccount is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.BiddersTokenAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Pot is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.PotSPLAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.MetadataAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AuctionAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TransferAuthority is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.ClockSysvar is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.SplTokenProgram is not set")
		}
	}
	return nil
}

func (inst *PlaceBid) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("PlaceBid")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   biddersPrimary", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("     biddersToken", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("              pot", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("           potSPL", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("         metadata", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("          auction", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("        tokenMint", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("transferAuthority", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("            payer", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("      clockSysvar", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("       rentSysvar", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta("    systemProgram", inst.AccountMetaSlice[11]))
						accountsBranch.Child(ag_format.Meta("  splTokenProgram", inst.AccountMetaSlice[12]))
					})
				})
		})
}

func (obj PlaceBid) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *PlaceBid) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewPlaceBidInstruction declares a new PlaceBid instruction with the provided parameters and accounts.
func NewPlaceBidInstruction(
	// Parameters:
	args PlaceBidArgs,
	// Accounts:
	biddersPrimaryAccount ag_solanago.PublicKey,
	biddersTokenAccount ag_solanago.PublicKey,
	pot ag_solanago.PublicKey,
	potSPLAccount ag_solanago.PublicKey,
	metadataAccount ag_solanago.PublicKey,
	auctionAccount ag_solanago.PublicKey,
	tokenMint ag_solanago.PublicKey,
	transferAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	clockSysvar ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	splTokenProgram ag_solanago.PublicKey) *PlaceBid {
	return NewPlaceBidInstructionBuilder().
		SetArgs(args).
		SetBiddersPrimaryAccount(biddersPrimaryAccount).
		SetBiddersTokenAccount(biddersTokenAccount).
		SetPotAccount(pot).
		SetPotSPLAccount(potSPLAccount).
		SetMetadataAccount(metadataAccount).
		SetAuctionAccount(auctionAccount).
		SetTokenMintAccount(tokenMint).
		SetTransferAuthorityAccount(transferAuthority).
		SetPayerAccount(payer).
		SetClockSysvarAccount(clockSysvar).
		SetRentSysvarAccount(rentSysvar).
		SetSystemProgramAccount(systemProgram).
		SetSplTokenProgramAccount(splTokenProgram)
}
