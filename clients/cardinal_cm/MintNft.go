// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_format "github.com/RoboticAgile/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// MintNft is the `mintNft` instruction.
type MintNft struct {
	CreatorBump *uint8

	// [0] = [WRITE] candyMachine
	//
	// [1] = [] candyMachineCreator
	//
	// [2] = [SIGNER] payer
	//
	// [3] = [WRITE] wallet
	//
	// [4] = [WRITE] metadata
	//
	// [5] = [WRITE,SIGNER] mint
	//
	// [6] = [SIGNER] mintAuthority
	//
	// [7] = [SIGNER] updateAuthority
	//
	// [8] = [WRITE] masterEdition
	//
	// [9] = [] tokenMetadataProgram
	//
	// [10] = [] tokenProgram
	//
	// [11] = [] systemProgram
	//
	// [12] = [] rent
	//
	// [13] = [] clock
	//
	// [14] = [] recentBlockhashes
	//
	// [15] = [] instructionSysvarAccount
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMintNftInstructionBuilder creates a new `MintNft` instruction builder.
func NewMintNftInstructionBuilder() *MintNft {
	nd := &MintNft{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 30),
	}
	return nd
}

// SetCreatorBump sets the "creatorBump" parameter.
func (inst *MintNft) SetCreatorBump(creatorBump uint8) *MintNft {
	inst.CreatorBump = &creatorBump
	return inst
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *MintNft) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *MintNft) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineCreatorAccount sets the "candyMachineCreator" account.
func (inst *MintNft) SetCandyMachineCreatorAccount(candyMachineCreator ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachineCreator)
	return inst
}

// GetCandyMachineCreatorAccount gets the "candyMachineCreator" account.
func (inst *MintNft) GetCandyMachineCreatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *MintNft) SetPayerAccount(payer ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *MintNft) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetWalletAccount sets the "wallet" account.
func (inst *MintNft) SetWalletAccount(wallet ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(wallet).WRITE()
	return inst
}

// GetWalletAccount gets the "wallet" account.
func (inst *MintNft) GetWalletAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *MintNft) SetMetadataAccount(metadata ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *MintNft) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetMintAccount sets the "mint" account.
func (inst *MintNft) SetMintAccount(mint ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(mint).WRITE().SIGNER()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *MintNft) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
func (inst *MintNft) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(mintAuthority).SIGNER()
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
func (inst *MintNft) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
func (inst *MintNft) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
func (inst *MintNft) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMasterEditionAccount sets the "masterEdition" account.
func (inst *MintNft) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
func (inst *MintNft) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *MintNft) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *MintNft) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *MintNft) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *MintNft) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *MintNft) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *MintNft) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetRentAccount sets the "rent" account.
func (inst *MintNft) SetRentAccount(rent ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *MintNft) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetClockAccount sets the "clock" account.
func (inst *MintNft) SetClockAccount(clock ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(clock)
	return inst
}

// GetClockAccount gets the "clock" account.
func (inst *MintNft) GetClockAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetRecentBlockhashesAccount sets the "recentBlockhashes" account.
func (inst *MintNft) SetRecentBlockhashesAccount(recentBlockhashes ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(recentBlockhashes)
	return inst
}

// GetRecentBlockhashesAccount gets the "recentBlockhashes" account.
func (inst *MintNft) GetRecentBlockhashesAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetInstructionSysvarAccountAccount sets the "instructionSysvarAccount" account.
func (inst *MintNft) SetInstructionSysvarAccountAccount(instructionSysvarAccount ag_solanago.PublicKey) *MintNft {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(instructionSysvarAccount)
	return inst
}

// GetInstructionSysvarAccountAccount gets the "instructionSysvarAccount" account.
func (inst *MintNft) GetInstructionSysvarAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

func (inst *MintNft) SetRemainingAccounts(pk []ag_solanago.AccountMeta) *MintNft {
	amount := len(pk)
	if amount == 0 {
		return inst
	}
	for i := 0; i < amount; i++ {
		inst.AccountMetaSlice[16+i] = &pk[i]
	}
	return inst
}

func (inst MintNft) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MintNft,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MintNft) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MintNft) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.CreatorBump == nil {
			return errors.New("CreatorBump parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CandyMachineCreator is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Wallet is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.Clock is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.RecentBlockhashes is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.InstructionSysvarAccount is not set")
		}
	}
	return nil
}

func (inst *MintNft) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MintNft")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("CreatorBump", *inst.CreatorBump))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=16]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("        candyMachine", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta(" candyMachineCreator", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("               payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              wallet", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("            metadata", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                mint", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("       mintAuthority", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("     updateAuthority", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("       masterEdition", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("tokenMetadataProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("        tokenProgram", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("                rent", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("               clock", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("   recentBlockhashes", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("   instructionSysvar", inst.AccountMetaSlice.Get(15)))
					})
				})
		})
}

func (obj MintNft) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CreatorBump` param:
	err = encoder.Encode(obj.CreatorBump)
	if err != nil {
		return err
	}
	return nil
}
func (obj *MintNft) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CreatorBump`:
	err = decoder.Decode(&obj.CreatorBump)
	if err != nil {
		return err
	}
	return nil
}

// NewMintNftInstruction declares a new MintNft instruction with the provided parameters and accounts.
func NewMintNftInstruction(
	// Parameters:
	creatorBump uint8,
	// Accounts:
	candyMachine ag_solanago.PublicKey,
	candyMachineCreator ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	wallet ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	clock ag_solanago.PublicKey,
	recentBlockhashes ag_solanago.PublicKey,
	instructionSysvarAccount ag_solanago.PublicKey,
	remainingAccounts []ag_solanago.AccountMeta) *MintNft {
	return NewMintNftInstructionBuilder().
		SetCreatorBump(creatorBump).
		SetCandyMachineAccount(candyMachine).
		SetCandyMachineCreatorAccount(candyMachineCreator).
		SetPayerAccount(payer).
		SetWalletAccount(wallet).
		SetMetadataAccount(metadata).
		SetMintAccount(mint).
		SetMintAuthorityAccount(mintAuthority).
		SetUpdateAuthorityAccount(updateAuthority).
		SetMasterEditionAccount(masterEdition).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentAccount(rent).
		SetClockAccount(clock).
		SetRecentBlockhashesAccount(recentBlockhashes).
		SetInstructionSysvarAccountAccount(instructionSysvarAccount).
		SetRemainingAccounts(remainingAccounts)
}
