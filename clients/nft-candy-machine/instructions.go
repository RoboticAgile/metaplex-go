// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package nft_candy_machine

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
	ag_text "github.com/RoboticAgile/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(pubkey ag_solanago.PublicKey) {
	ProgramID = pubkey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "NftCandyMachine"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_MintNft = ag_binary.TypeID([8]byte{211, 57, 6, 167, 15, 219, 35, 251})

	Instruction_UpdateCandyMachine = ag_binary.TypeID([8]byte{243, 251, 124, 156, 211, 211, 118, 239})

	Instruction_InitializeConfig = ag_binary.TypeID([8]byte{208, 127, 21, 1, 194, 190, 196, 70})

	Instruction_AddConfigLines = ag_binary.TypeID([8]byte{223, 50, 224, 227, 151, 8, 115, 106})

	Instruction_InitializeCandyMachine = ag_binary.TypeID([8]byte{142, 137, 167, 107, 47, 39, 240, 124})

	Instruction_UpdateAuthority = ag_binary.TypeID([8]byte{32, 46, 64, 28, 149, 75, 243, 88})

	Instruction_WithdrawFunds = ag_binary.TypeID([8]byte{241, 36, 29, 111, 208, 31, 104, 217})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_MintNft:
		return "MintNft"
	case Instruction_UpdateCandyMachine:
		return "UpdateCandyMachine"
	case Instruction_InitializeConfig:
		return "InitializeConfig"
	case Instruction_AddConfigLines:
		return "AddConfigLines"
	case Instruction_InitializeCandyMachine:
		return "InitializeCandyMachine"
	case Instruction_UpdateAuthority:
		return "UpdateAuthority"
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
			"mint_nft", (*MintNft)(nil),
		},
		{
			"update_candy_machine", (*UpdateCandyMachine)(nil),
		},
		{
			"initialize_config", (*InitializeConfig)(nil),
		},
		{
			"add_config_lines", (*AddConfigLines)(nil),
		},
		{
			"initialize_candy_machine", (*InitializeCandyMachine)(nil),
		},
		{
			"update_authority", (*UpdateAuthority)(nil),
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
