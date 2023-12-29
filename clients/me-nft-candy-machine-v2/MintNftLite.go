// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/desperatee/solana-go"
	ag_format "github.com/desperatee/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// * Lite version of mint_nft. This is intended to be used to support dynamic price mints (Strata Protocol)\n     * Caveats:\n     * - cannot use Crossmint\n     * - no randomized order, should always use post-mint reveal\n     * - no bot tax on any stage\n     * - a candy machine should NEVER be used for both lite and non-lite mints, see is_lite field in CandyMachine
type MintNftLite struct {
	UserLimit *uint8 `bin:"optional"`

	// [0] = [] config
	//
	// [1] = [WRITE] candyMachine
	//
	// [2] = [WRITE, SIGNER] payer
	//
	// [3] = [WRITE] launchStagesInfo
	//
	// [4] = [WRITE] payFrom
	//
	// [5] = [WRITE] payTo
	//
	// [6] = [WRITE] metadata
	//
	// [7] = [] notary
	//
	// [8] = [WRITE, SIGNER] mint
	//
	// [9] = [WRITE] tokenAta
	//
	// [10] = [WRITE] masterEdition
	//
	// [11] = [WRITE] walletLimitInfo
	//
	// [12] = [] tokenMetadataProgram
	//
	// [13] = [] tokenProgram
	//
	// [14] = [] systemProgram
	//
	// [15] = [] associatedTokenProgram
	//
	// [16] = [] rent
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewMintNftLiteInstructionBuilder creates a new `MintNftLite` instruction builder.
func NewMintNftLiteInstructionBuilder() *MintNftLite {
	nd := &MintNftLite{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 17),
	}
	return nd
}

// SetUserLimit sets the "userLimit" parameter.
func (inst *MintNftLite) SetUserLimit(userLimit uint8) *MintNftLite {
	inst.UserLimit = &userLimit
	return inst
}

// SetConfigAccount sets the "config" account.
func (inst *MintNftLite) SetConfigAccount(config ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(config)
	return inst
}

// GetConfigAccount gets the "config" account.
func (inst *MintNftLite) GetConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetCandyMachineAccount sets the "candyMachine" account.
func (inst *MintNftLite) SetCandyMachineAccount(candyMachine ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(candyMachine).WRITE()
	return inst
}

// GetCandyMachineAccount gets the "candyMachine" account.
func (inst *MintNftLite) GetCandyMachineAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPayerAccount sets the "payer" account.
func (inst *MintNftLite) SetPayerAccount(payer ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(payer).WRITE().SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
func (inst *MintNftLite) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetLaunchStagesInfoAccount sets the "launchStagesInfo" account.
func (inst *MintNftLite) SetLaunchStagesInfoAccount(launchStagesInfo ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(launchStagesInfo).WRITE()
	return inst
}

// GetLaunchStagesInfoAccount gets the "launchStagesInfo" account.
func (inst *MintNftLite) GetLaunchStagesInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetPayFromAccount sets the "payFrom" account.
func (inst *MintNftLite) SetPayFromAccount(payFrom ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payFrom).WRITE()
	return inst
}

// GetPayFromAccount gets the "payFrom" account.
func (inst *MintNftLite) GetPayFromAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPayToAccount sets the "payTo" account.
func (inst *MintNftLite) SetPayToAccount(payTo ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(payTo).WRITE()
	return inst
}

// GetPayToAccount gets the "payTo" account.
func (inst *MintNftLite) GetPayToAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMetadataAccount sets the "metadata" account.
func (inst *MintNftLite) SetMetadataAccount(metadata ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(metadata).WRITE()
	return inst
}

// GetMetadataAccount gets the "metadata" account.
func (inst *MintNftLite) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetNotaryAccount sets the "notary" account.
func (inst *MintNftLite) SetNotaryAccount(notary ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(notary)
	return inst
}

// GetNotaryAccount gets the "notary" account.
func (inst *MintNftLite) GetNotaryAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetMintAccount sets the "mint" account.
func (inst *MintNftLite) SetMintAccount(mint ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(mint).WRITE().SIGNER()
	return inst
}

// GetMintAccount gets the "mint" account.
func (inst *MintNftLite) GetMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenAtaAccount sets the "tokenAta" account.
func (inst *MintNftLite) SetTokenAtaAccount(tokenAta ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenAta).WRITE()
	return inst
}

// GetTokenAtaAccount gets the "tokenAta" account.
func (inst *MintNftLite) GetTokenAtaAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetMasterEditionAccount sets the "masterEdition" account.
func (inst *MintNftLite) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
func (inst *MintNftLite) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetWalletLimitInfoAccount sets the "walletLimitInfo" account.
func (inst *MintNftLite) SetWalletLimitInfoAccount(walletLimitInfo ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(walletLimitInfo).WRITE()
	return inst
}

// GetWalletLimitInfoAccount gets the "walletLimitInfo" account.
func (inst *MintNftLite) GetWalletLimitInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
func (inst *MintNftLite) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
func (inst *MintNftLite) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
func (inst *MintNftLite) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
func (inst *MintNftLite) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(13)
}

// SetSystemProgramAccount sets the "systemProgram" account.
func (inst *MintNftLite) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
func (inst *MintNftLite) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(14)
}

// SetAssociatedTokenProgramAccount sets the "associatedTokenProgram" account.
func (inst *MintNftLite) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associatedTokenProgram" account.
func (inst *MintNftLite) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(15)
}

// SetRentAccount sets the "rent" account.
func (inst *MintNftLite) SetRentAccount(rent ag_solanago.PublicKey) *MintNftLite {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
func (inst *MintNftLite) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(16)
}

func (inst MintNftLite) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_MintNftLite,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst MintNftLite) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *MintNftLite) Validate() error {
	// Check whether all (required) parameters are set:
	{
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Config is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.CandyMachine is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.LaunchStagesInfo is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.PayFrom is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PayTo is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Notary is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Mint is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenAta is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.MasterEdition is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.WalletLimitInfo is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.Rent is not set")
		}
	}
	return nil
}

func (inst *MintNftLite) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("MintNftLite")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("UserLimit (OPT)", inst.UserLimit))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=17]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                config", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("          candyMachine", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                 payer", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("      launchStagesInfo", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("               payFrom", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                 payTo", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("              metadata", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                notary", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                  mint", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("              tokenAta", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("         masterEdition", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("       walletLimitInfo", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("  tokenMetadataProgram", inst.AccountMetaSlice.Get(12)))
						accountsBranch.Child(ag_format.Meta("          tokenProgram", inst.AccountMetaSlice.Get(13)))
						accountsBranch.Child(ag_format.Meta("         systemProgram", inst.AccountMetaSlice.Get(14)))
						accountsBranch.Child(ag_format.Meta("associatedTokenProgram", inst.AccountMetaSlice.Get(15)))
						accountsBranch.Child(ag_format.Meta("                  rent", inst.AccountMetaSlice.Get(16)))
					})
				})
		})
}

func (obj MintNftLite) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `UserLimit` param (optional):
	{
		if obj.UserLimit == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.UserLimit)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
func (obj *MintNftLite) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `UserLimit` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.UserLimit)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// NewMintNftLiteInstruction declares a new MintNftLite instruction with the provided parameters and accounts.
func NewMintNftLiteInstruction(
	// Parameters:
	userLimit uint8,
	// Accounts:
	config ag_solanago.PublicKey,
	candyMachine ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	launchStagesInfo ag_solanago.PublicKey,
	payFrom ag_solanago.PublicKey,
	payTo ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	notary ag_solanago.PublicKey,
	mint ag_solanago.PublicKey,
	tokenAta ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	walletLimitInfo ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey,
	rent ag_solanago.PublicKey) *MintNftLite {
	return NewMintNftLiteInstructionBuilder().
		SetUserLimit(userLimit).
		SetConfigAccount(config).
		SetCandyMachineAccount(candyMachine).
		SetPayerAccount(payer).
		SetLaunchStagesInfoAccount(launchStagesInfo).
		SetPayFromAccount(payFrom).
		SetPayToAccount(payTo).
		SetMetadataAccount(metadata).
		SetNotaryAccount(notary).
		SetMintAccount(mint).
		SetTokenAtaAccount(tokenAta).
		SetMasterEditionAccount(masterEdition).
		SetWalletLimitInfoAccount(walletLimitInfo).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram).
		SetRentAccount(rent)
}
