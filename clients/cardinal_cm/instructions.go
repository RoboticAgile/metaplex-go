// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package candy_machine

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_text "github.com/RoboticAgile/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey = ag_solanago.MustPublicKeyFromBase58("ccmpgw68x3NJmNPePFrTm6TsKCEYUVhF8rEAVL9rSDd")

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "CandyMachine"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_InitializeCandyMachine = ag_binary.TypeID([8]byte{142, 137, 167, 107, 47, 39, 240, 124})

	Instruction_UpdateCandyMachine = ag_binary.TypeID([8]byte{243, 251, 124, 156, 211, 211, 118, 239})

	Instruction_UpdateAuthority = ag_binary.TypeID([8]byte{32, 46, 64, 28, 149, 75, 243, 88})

	Instruction_AddConfigLines = ag_binary.TypeID([8]byte{223, 50, 224, 227, 151, 8, 115, 106})

	Instruction_SetCollection = ag_binary.TypeID([8]byte{192, 254, 206, 76, 168, 182, 59, 223})

	Instruction_RemoveCollection = ag_binary.TypeID([8]byte{223, 52, 106, 217, 61, 220, 36, 160})

	Instruction_MintNft = ag_binary.TypeID([8]byte{211, 57, 6, 167, 15, 219, 35, 251})

	Instruction_SetCollectionDuringMint = ag_binary.TypeID([8]byte{103, 17, 200, 25, 118, 95, 125, 61})

	Instruction_SetLockupSettings = ag_binary.TypeID([8]byte{124, 128, 161, 215, 10, 153, 211, 196})

	Instruction_CloseLockupSettings = ag_binary.TypeID([8]byte{111, 166, 221, 21, 169, 96, 201, 123})

	Instruction_SetPermissionedSettings = ag_binary.TypeID([8]byte{254, 107, 24, 140, 97, 62, 219, 233})

	Instruction_ClosePermissionedSettings = ag_binary.TypeID([8]byte{23, 185, 11, 107, 101, 221, 97, 242})

	Instruction_WithdrawFunds = ag_binary.TypeID([8]byte{241, 36, 29, 111, 208, 31, 104, 217})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_InitializeCandyMachine:
		return "InitializeCandyMachine"
	case Instruction_UpdateCandyMachine:
		return "UpdateCandyMachine"
	case Instruction_UpdateAuthority:
		return "UpdateAuthority"
	case Instruction_AddConfigLines:
		return "AddConfigLines"
	case Instruction_SetCollection:
		return "SetCollection"
	case Instruction_RemoveCollection:
		return "RemoveCollection"
	case Instruction_MintNft:
		return "MintNft"
	case Instruction_SetCollectionDuringMint:
		return "SetCollectionDuringMint"
	case Instruction_SetLockupSettings:
		return "SetLockupSettings"
	case Instruction_CloseLockupSettings:
		return "CloseLockupSettings"
	case Instruction_SetPermissionedSettings:
		return "SetPermissionedSettings"
	case Instruction_ClosePermissionedSettings:
		return "ClosePermissionedSettings"
	case Instruction_WithdrawFunds:
		return "WithdrawFunds"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			"initialize_candy_machine", (*InitializeCandyMachine)(nil),
		},
		{
			"update_candy_machine", (*UpdateCandyMachine)(nil),
		},
		{
			"update_authority", (*UpdateAuthority)(nil),
		},
		{
			"add_config_lines", (*AddConfigLines)(nil),
		},
		{
			"set_collection", (*SetCollection)(nil),
		},
		{
			"remove_collection", (*RemoveCollection)(nil),
		},
		{
			"mint_nft", (*MintNft)(nil),
		},
		{
			"set_collection_during_mint", (*SetCollectionDuringMint)(nil),
		},
		{
			"set_lockup_settings", (*SetLockupSettings)(nil),
		},
		{
			"close_lockup_settings", (*CloseLockupSettings)(nil),
		},
		{
			"set_permissioned_settings", (*SetPermissionedSettings)(nil),
		},
		{
			"close_permissioned_settings", (*ClosePermissionedSettings)(nil),
		},
		{
			"withdraw_funds", (*WithdrawFunds)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := DecodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func DecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}
