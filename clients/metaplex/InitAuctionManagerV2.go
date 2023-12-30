// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initializes an Auction Manager V2
//
// NOTE: It is not possible to use MasterEditionV1s for participation nfts with these managers.
type InitAuctionManagerV2 struct {
	Args *InitAuctionManagerV2Args

	// [0] = [WRITE] uninitializedUnallocatedAuctionManager
	// ··········· Uninitialized, unallocated auction manager account with pda of ['metaplex', auction_key from auction referenced below]
	//
	// [1] = [WRITE] auctionWinnerTokenTypeTracker
	// ··········· AuctionWinnerTokenTypeTracker, pda of seed ['metaplex', program id, auction manager key, 'totals']
	//
	// [2] = [] combinedVault
	// ··········· Combined vault account with authority set to auction manager account (this will be checked)
	// ··········· Note in addition that this vault account should have authority set to this program's pda of ['metaplex', auction_key]
	//
	// [3] = [] auction
	// ··········· Auction with auctioned item being set to the vault given and authority set to this program's pda of ['metaplex', auction_key]
	//
	// [4] = [] auctionManagerAuthority
	// ··········· Authority for the Auction Manager
	//
	// [5] = [SIGNER] payer
	// ··········· Payer
	//
	// [6] = [] acceptPayment
	// ··········· Accept payment account of same token mint as the auction for taking payment for open editions, owner should be auction manager key
	//
	// [7] = [] store
	// ··········· Store that this auction manager will belong to
	//
	// [8] = [] systemSysvar
	// ··········· System sysvar
	//
	// [9] = [] rentSysvar
	// ··········· Rent sysvar
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitAuctionManagerV2InstructionBuilder creates a new `InitAuctionManagerV2` instruction builder.
func NewInitAuctionManagerV2InstructionBuilder() *InitAuctionManagerV2 {
	nd := &InitAuctionManagerV2{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *InitAuctionManagerV2) SetArgs(args InitAuctionManagerV2Args) *InitAuctionManagerV2 {
	inst.Args = &args
	return inst
}

// SetUninitializedUnallocatedAuctionManagerAccount sets the "uninitializedUnallocatedAuctionManager" account.
// Uninitialized, unallocated auction manager account with pda of ['metaplex', auction_key from auction referenced below]
func (inst *InitAuctionManagerV2) SetUninitializedUnallocatedAuctionManagerAccount(uninitializedUnallocatedAuctionManager ag_solanago.PublicKey) *InitAuctionManagerV2 {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(uninitializedUnallocatedAuctionManager).WRITE()
	return inst
}

// GetUninitializedUnallocatedAuctionManagerAccount gets the "uninitializedUnallocatedAuctionManager" account.
// Uninitialized, unallocated auction manager account with pda of ['metaplex', auction_key from auction referenced below]
func (inst *InitAuctionManagerV2) GetUninitializedUnallocatedAuctionManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAuctionWinnerTokenTypeTrackerAccount sets the "auctionWinnerTokenTypeTracker" account.
// AuctionWinnerTokenTypeTracker, pda of seed ['metaplex', program id, auction manager key, 'totals']
func (inst *InitAuctionManagerV2) SetAuctionWinnerTokenTypeTrackerAccount(auctionWinnerTokenTypeTracker ag_solanago.PublicKey) *InitAuctionManagerV2 {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(auctionWinnerTokenTypeTracker).WRITE()
	return inst
}

// GetAuctionWinnerTokenTypeTrackerAccount gets the "auctionWinnerTokenTypeTracker" account.
// AuctionWinnerTokenTypeTracker, pda of seed ['metaplex', program id, auction manager key, 'totals']
func (inst *InitAuctionManagerV2) GetAuctionWinnerTokenTypeTrackerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetCombinedVaultAccount sets the "combinedVault" account.
// Combined vault account with authority set to auction manager account (this will be checked)
// Note in addition that this vault account should have authority set to this program's pda of ['metaplex', auction_key]
func (inst *InitAuctionManagerV2) SetCombinedVaultAccount(combinedVault ag_solanago.PublicKey) *InitAuctionManagerV2 {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(combinedVault)
	return inst
}

// GetCombinedVaultAccount gets the "combinedVault" account.
// Combined vault account with authority set to auction manager account (this will be checked)
// Note in addition that this vault account should have authority set to this program's pda of ['metaplex', auction_key]
func (inst *InitAuctionManagerV2) GetCombinedVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAuctionAccount sets the "auction" account.
// Auction with auctioned item being set to the vault given and authority set to this program's pda of ['metaplex', auction_key]
func (inst *InitAuctionManagerV2) SetAuctionAccount(auction ag_solanago.PublicKey) *InitAuctionManagerV2 {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(auction)
	return inst
}

// GetAuctionAccount gets the "auction" account.
// Auction with auctioned item being set to the vault given and authority set to this program's pda of ['metaplex', auction_key]
func (inst *InitAuctionManagerV2) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAuctionManagerAuthorityAccount sets the "auctionManagerAuthority" account.
// Authority for the Auction Manager
func (inst *InitAuctionManagerV2) SetAuctionManagerAuthorityAccount(auctionManagerAuthority ag_solanago.PublicKey) *InitAuctionManagerV2 {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(auctionManagerAuthority)
	return inst
}

// GetAuctionManagerAuthorityAccount gets the "auctionManagerAuthority" account.
// Authority for the Auction Manager
func (inst *InitAuctionManagerV2) GetAuctionManagerAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *InitAuctionManagerV2) SetPayerAccount(payer ag_solanago.PublicKey) *InitAuctionManagerV2 {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *InitAuctionManagerV2) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAcceptPaymentAccount sets the "acceptPayment" account.
// Accept payment account of same token mint as the auction for taking payment for open editions, owner should be auction manager key
func (inst *InitAuctionManagerV2) SetAcceptPaymentAccount(acceptPayment ag_solanago.PublicKey) *InitAuctionManagerV2 {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(acceptPayment)
	return inst
}

// GetAcceptPaymentAccount gets the "acceptPayment" account.
// Accept payment account of same token mint as the auction for taking payment for open editions, owner should be auction manager key
func (inst *InitAuctionManagerV2) GetAcceptPaymentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetStoreAccount sets the "store" account.
// Store that this auction manager will belong to
func (inst *InitAuctionManagerV2) SetStoreAccount(store ag_solanago.PublicKey) *InitAuctionManagerV2 {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(store)
	return inst
}

// GetStoreAccount gets the "store" account.
// Store that this auction manager will belong to
func (inst *InitAuctionManagerV2) GetStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemSysvarAccount sets the "systemSysvar" account.
// System sysvar
func (inst *InitAuctionManagerV2) SetSystemSysvarAccount(systemSysvar ag_solanago.PublicKey) *InitAuctionManagerV2 {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemSysvar)
	return inst
}

// GetSystemSysvarAccount gets the "systemSysvar" account.
// System sysvar
func (inst *InitAuctionManagerV2) GetSystemSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *InitAuctionManagerV2) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *InitAuctionManagerV2 {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *InitAuctionManagerV2) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst InitAuctionManagerV2) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_InitAuctionManagerV2),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitAuctionManagerV2) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitAuctionManagerV2) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.UninitializedUnallocatedAuctionManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.AuctionWinnerTokenTypeTracker is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.CombinedVault is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Auction is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AuctionManagerAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.AcceptPayment is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Store is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemSysvar is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
	}
	return nil
}

func (inst *InitAuctionManagerV2) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitAuctionManagerV2")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("uninitializedUnallocatedAuctionManager", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("         auctionWinnerTokenTypeTracker", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                         combinedVault", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                               auction", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("               auctionManagerAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                                 payer", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                         acceptPayment", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                                 store", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                          systemSysvar", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                            rentSysvar", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj InitAuctionManagerV2) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *InitAuctionManagerV2) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewInitAuctionManagerV2Instruction declares a new InitAuctionManagerV2 instruction with the provided parameters and accounts.
func NewInitAuctionManagerV2Instruction(
	// Parameters:
	args InitAuctionManagerV2Args,
	// Accounts:
	uninitializedUnallocatedAuctionManager ag_solanago.PublicKey,
	auctionWinnerTokenTypeTracker ag_solanago.PublicKey,
	combinedVault ag_solanago.PublicKey,
	auction ag_solanago.PublicKey,
	auctionManagerAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	acceptPayment ag_solanago.PublicKey,
	store ag_solanago.PublicKey,
	systemSysvar ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey) *InitAuctionManagerV2 {
	return NewInitAuctionManagerV2InstructionBuilder().
		SetArgs(args).
		SetUninitializedUnallocatedAuctionManagerAccount(uninitializedUnallocatedAuctionManager).
		SetAuctionWinnerTokenTypeTrackerAccount(auctionWinnerTokenTypeTracker).
		SetCombinedVaultAccount(combinedVault).
		SetAuctionAccount(auction).
		SetAuctionManagerAuthorityAccount(auctionManagerAuthority).
		SetPayerAccount(payer).
		SetAcceptPaymentAccount(acceptPayment).
		SetStoreAccount(store).
		SetSystemSysvarAccount(systemSysvar).
		SetRentSysvarAccount(rentSysvar)
}
