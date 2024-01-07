// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package me_launchpad

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/RoboticAgile/solana-go"
)

type InitializeCandyMachineArgs struct {
	CmBump           uint8
	LaunchStagesBump uint8
	Uuid             string
	ItemsAvailable   uint64
	Stages           []LaunchStageArgs
	IsLite           bool
	NotaryRequired   []bool
	Mip1Ruleset      *ag_solanago.PublicKey `bin:"optional"`
	IsOpenEdition    *bool                  `bin:"optional"`
}

func (obj InitializeCandyMachineArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `CmBump` param:
	err = encoder.Encode(obj.CmBump)
	if err != nil {
		return err
	}
	// Serialize `LaunchStagesBump` param:
	err = encoder.Encode(obj.LaunchStagesBump)
	if err != nil {
		return err
	}
	// Serialize `Uuid` param:
	err = encoder.Encode(obj.Uuid)
	if err != nil {
		return err
	}
	// Serialize `ItemsAvailable` param:
	err = encoder.Encode(obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Serialize `Stages` param:
	err = encoder.Encode(obj.Stages)
	if err != nil {
		return err
	}
	// Serialize `IsLite` param:
	err = encoder.Encode(obj.IsLite)
	if err != nil {
		return err
	}
	// Serialize `NotaryRequired` param:
	err = encoder.Encode(obj.NotaryRequired)
	if err != nil {
		return err
	}
	// Serialize `Mip1Ruleset` param (optional):
	{
		if obj.Mip1Ruleset == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Mip1Ruleset)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `IsOpenEdition` param (optional):
	{
		if obj.IsOpenEdition == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.IsOpenEdition)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *InitializeCandyMachineArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `CmBump`:
	err = decoder.Decode(&obj.CmBump)
	if err != nil {
		return err
	}
	// Deserialize `LaunchStagesBump`:
	err = decoder.Decode(&obj.LaunchStagesBump)
	if err != nil {
		return err
	}
	// Deserialize `Uuid`:
	err = decoder.Decode(&obj.Uuid)
	if err != nil {
		return err
	}
	// Deserialize `ItemsAvailable`:
	err = decoder.Decode(&obj.ItemsAvailable)
	if err != nil {
		return err
	}
	// Deserialize `Stages`:
	err = decoder.Decode(&obj.Stages)
	if err != nil {
		return err
	}
	// Deserialize `IsLite`:
	err = decoder.Decode(&obj.IsLite)
	if err != nil {
		return err
	}
	// Deserialize `NotaryRequired`:
	err = decoder.Decode(&obj.NotaryRequired)
	if err != nil {
		return err
	}
	// Deserialize `Mip1Ruleset` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Mip1Ruleset)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `IsOpenEdition` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.IsOpenEdition)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type InitializeConfigArgs struct {
	Gateway              string
	Cid                  string
	Uuid                 string
	CollectionName       string
	Symbol               string
	SellerFeeBasisPoints uint16
	Creators             []Creator
	IsMutable            bool
	RetainAuthority      bool
}

func (obj InitializeConfigArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Gateway` param:
	err = encoder.Encode(obj.Gateway)
	if err != nil {
		return err
	}
	// Serialize `Cid` param:
	err = encoder.Encode(obj.Cid)
	if err != nil {
		return err
	}
	// Serialize `Uuid` param:
	err = encoder.Encode(obj.Uuid)
	if err != nil {
		return err
	}
	// Serialize `CollectionName` param:
	err = encoder.Encode(obj.CollectionName)
	if err != nil {
		return err
	}
	// Serialize `Symbol` param:
	err = encoder.Encode(obj.Symbol)
	if err != nil {
		return err
	}
	// Serialize `SellerFeeBasisPoints` param:
	err = encoder.Encode(obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Serialize `Creators` param:
	err = encoder.Encode(obj.Creators)
	if err != nil {
		return err
	}
	// Serialize `IsMutable` param:
	err = encoder.Encode(obj.IsMutable)
	if err != nil {
		return err
	}
	// Serialize `RetainAuthority` param:
	err = encoder.Encode(obj.RetainAuthority)
	if err != nil {
		return err
	}
	return nil
}

func (obj *InitializeConfigArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Gateway`:
	err = decoder.Decode(&obj.Gateway)
	if err != nil {
		return err
	}
	// Deserialize `Cid`:
	err = decoder.Decode(&obj.Cid)
	if err != nil {
		return err
	}
	// Deserialize `Uuid`:
	err = decoder.Decode(&obj.Uuid)
	if err != nil {
		return err
	}
	// Deserialize `CollectionName`:
	err = decoder.Decode(&obj.CollectionName)
	if err != nil {
		return err
	}
	// Deserialize `Symbol`:
	err = decoder.Decode(&obj.Symbol)
	if err != nil {
		return err
	}
	// Deserialize `SellerFeeBasisPoints`:
	err = decoder.Decode(&obj.SellerFeeBasisPoints)
	if err != nil {
		return err
	}
	// Deserialize `Creators`:
	err = decoder.Decode(&obj.Creators)
	if err != nil {
		return err
	}
	// Deserialize `IsMutable`:
	err = decoder.Decode(&obj.IsMutable)
	if err != nil {
		return err
	}
	// Deserialize `RetainAuthority`:
	err = decoder.Decode(&obj.RetainAuthority)
	if err != nil {
		return err
	}
	return nil
}

type UpdateConfigArgs struct {
	Gateway              *string    `bin:"optional"`
	Cid                  *string    `bin:"optional"`
	CollectionName       *string    `bin:"optional"`
	Symbol               *string    `bin:"optional"`
	SellerFeeBasisPoints *uint16    `bin:"optional"`
	Creators             *[]Creator `bin:"optional"`
}

func (obj UpdateConfigArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Gateway` param (optional):
	{
		if obj.Gateway == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Gateway)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Cid` param (optional):
	{
		if obj.Cid == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Cid)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `CollectionName` param (optional):
	{
		if obj.CollectionName == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.CollectionName)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Symbol` param (optional):
	{
		if obj.Symbol == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Symbol)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `SellerFeeBasisPoints` param (optional):
	{
		if obj.SellerFeeBasisPoints == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.SellerFeeBasisPoints)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `Creators` param (optional):
	{
		if obj.Creators == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (obj *UpdateConfigArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Gateway` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Gateway)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Cid` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Cid)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `CollectionName` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.CollectionName)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Symbol` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Symbol)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `SellerFeeBasisPoints` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.SellerFeeBasisPoints)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `Creators` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.Creators)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

type UpdateFreezeStateArgs struct {
	Expiry int64
}

func (obj UpdateFreezeStateArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Expiry` param:
	err = encoder.Encode(obj.Expiry)
	if err != nil {
		return err
	}
	return nil
}

func (obj *UpdateFreezeStateArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Expiry`:
	err = decoder.Decode(&obj.Expiry)
	if err != nil {
		return err
	}
	return nil
}

type RedeemedDuringStage struct {
	RedeemedNormal        uint8
	RedeemedRaffleTickets uint8
}

func (obj RedeemedDuringStage) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RedeemedNormal` param:
	err = encoder.Encode(obj.RedeemedNormal)
	if err != nil {
		return err
	}
	// Serialize `RedeemedRaffleTickets` param:
	err = encoder.Encode(obj.RedeemedRaffleTickets)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RedeemedDuringStage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RedeemedNormal`:
	err = decoder.Decode(&obj.RedeemedNormal)
	if err != nil {
		return err
	}
	// Deserialize `RedeemedRaffleTickets`:
	err = decoder.Decode(&obj.RedeemedRaffleTickets)
	if err != nil {
		return err
	}
	return nil
}

type LaunchStage struct {
	StageType                   LaunchStageType
	StartTime                   int64
	EndTime                     int64
	WalletLimit                 WalletLimitSpecification
	Price                       uint64
	StageSupply                 *uint32 `bin:"optional"`
	PreviousStageUnmintedSupply uint32
	MintedDuringStage           uint32
	PaymentMint                 ag_solanago.PublicKey
	PaymentAta                  ag_solanago.PublicKey
}

func (obj LaunchStage) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `StageType` param:
	err = encoder.Encode(obj.StageType)
	if err != nil {
		return err
	}
	// Serialize `StartTime` param:
	err = encoder.Encode(obj.StartTime)
	if err != nil {
		return err
	}
	// Serialize `EndTime` param:
	err = encoder.Encode(obj.EndTime)
	if err != nil {
		return err
	}
	// Serialize `WalletLimit` param:
	{
		tmp := walletLimitSpecificationContainer{}
		// switch realvalue := obj.WalletLimit.(type) {
		// case *WalletLimitSpecificationNoLimit:
		// 	tmp.Enum = 0
		// 	tmp.NoLimit = *realvalue
		// case *WalletLimitSpecificationFixedLimit:
		// 	tmp.Enum = 1
		// 	tmp.FixedLimit = *realvalue
		// case *WalletLimitSpecificationVariableLimit:
		// 	tmp.Enum = 2
		// 	tmp.VariableLimit = *realvalue
		// }
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `StageSupply` param (optional):
	{
		if obj.StageSupply == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.StageSupply)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `PreviousStageUnmintedSupply` param:
	err = encoder.Encode(obj.PreviousStageUnmintedSupply)
	if err != nil {
		return err
	}
	// Serialize `MintedDuringStage` param:
	err = encoder.Encode(obj.MintedDuringStage)
	if err != nil {
		return err
	}
	// Serialize `PaymentMint` param:
	err = encoder.Encode(obj.PaymentMint)
	if err != nil {
		return err
	}
	// Serialize `PaymentAta` param:
	err = encoder.Encode(obj.PaymentAta)
	if err != nil {
		return err
	}
	return nil
}

func (obj *LaunchStage) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `StageType`:
	err = decoder.Decode(&obj.StageType)
	if err != nil {
		return err
	}
	// Deserialize `StartTime`:
	err = decoder.Decode(&obj.StartTime)
	if err != nil {
		return err
	}
	// Deserialize `EndTime`:
	err = decoder.Decode(&obj.EndTime)
	if err != nil {
		return err
	}
	// Deserialize `WalletLimit`:
	{
		tmp := new(walletLimitSpecificationContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.WalletLimit = (*NoLimit)(&tmp.Enum)
		case 1:
			obj.WalletLimit = &tmp.FixedLimit
		case 2:
			obj.WalletLimit = (*VariableLimit)(&tmp.Enum)
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `StageSupply` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.StageSupply)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `PreviousStageUnmintedSupply`:
	err = decoder.Decode(&obj.PreviousStageUnmintedSupply)
	if err != nil {
		return err
	}
	// Deserialize `MintedDuringStage`:
	err = decoder.Decode(&obj.MintedDuringStage)
	if err != nil {
		return err
	}
	// Deserialize `PaymentMint`:
	err = decoder.Decode(&obj.PaymentMint)
	if err != nil {
		return err
	}
	// Deserialize `PaymentAta`:
	err = decoder.Decode(&obj.PaymentAta)
	if err != nil {
		return err
	}
	return nil
}

type LaunchStageArgs struct {
	StageType          LaunchStageType
	StartTime          int64
	EndTime            int64
	WalletLimit        WalletLimitSpecification
	Price              uint64
	StageSupply        *uint32 `bin:"optional"`
	PaymentMintIndex   uint8
	PaymentMintAtaBump uint8
}

func (obj LaunchStageArgs) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `StageType` param:
	err = encoder.Encode(obj.StageType)
	if err != nil {
		return err
	}
	// Serialize `StartTime` param:
	err = encoder.Encode(obj.StartTime)
	if err != nil {
		return err
	}
	// Serialize `EndTime` param:
	err = encoder.Encode(obj.EndTime)
	if err != nil {
		return err
	}
	// Serialize `WalletLimit` param:
	{
		tmp := walletLimitSpecificationContainer{}
		// switch realvalue := obj.WalletLimit.(type) {
		// case *WalletLimitSpecificationNoLimit:
		// 	tmp.Enum = 0
		// 	tmp.NoLimit = *realvalue
		// case *WalletLimitSpecificationFixedLimit:
		// 	tmp.Enum = 1
		// 	tmp.FixedLimit = *realvalue
		// case *WalletLimitSpecificationVariableLimit:
		// 	tmp.Enum = 2
		// 	tmp.VariableLimit = *realvalue
		// }
		err := encoder.Encode(tmp)
		if err != nil {
			return err
		}
	}
	// Serialize `Price` param:
	err = encoder.Encode(obj.Price)
	if err != nil {
		return err
	}
	// Serialize `StageSupply` param (optional):
	{
		if obj.StageSupply == nil {
			err = encoder.WriteBool(false)
			if err != nil {
				return err
			}
		} else {
			err = encoder.WriteBool(true)
			if err != nil {
				return err
			}
			err = encoder.Encode(obj.StageSupply)
			if err != nil {
				return err
			}
		}
	}
	// Serialize `PaymentMintIndex` param:
	err = encoder.Encode(obj.PaymentMintIndex)
	if err != nil {
		return err
	}
	// Serialize `PaymentMintAtaBump` param:
	err = encoder.Encode(obj.PaymentMintAtaBump)
	if err != nil {
		return err
	}
	return nil
}

func (obj *LaunchStageArgs) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `StageType`:
	err = decoder.Decode(&obj.StageType)
	if err != nil {
		return err
	}
	// Deserialize `StartTime`:
	err = decoder.Decode(&obj.StartTime)
	if err != nil {
		return err
	}
	// Deserialize `EndTime`:
	err = decoder.Decode(&obj.EndTime)
	if err != nil {
		return err
	}
	// Deserialize `WalletLimit`:
	{
		tmp := new(walletLimitSpecificationContainer)
		err := decoder.Decode(tmp)
		if err != nil {
			return err
		}
		switch tmp.Enum {
		case 0:
			obj.WalletLimit = (*NoLimit)(&tmp.Enum)
		case 1:
			obj.WalletLimit = &tmp.FixedLimit
		case 2:
			obj.WalletLimit = (*VariableLimit)(&tmp.Enum)
		default:
			return fmt.Errorf("unknown enum index: %v", tmp.Enum)
		}
	}
	// Deserialize `Price`:
	err = decoder.Decode(&obj.Price)
	if err != nil {
		return err
	}
	// Deserialize `StageSupply` (optional):
	{
		ok, err := decoder.ReadBool()
		if err != nil {
			return err
		}
		if ok {
			err = decoder.Decode(&obj.StageSupply)
			if err != nil {
				return err
			}
		}
	}
	// Deserialize `PaymentMintIndex`:
	err = decoder.Decode(&obj.PaymentMintIndex)
	if err != nil {
		return err
	}
	// Deserialize `PaymentMintAtaBump`:
	err = decoder.Decode(&obj.PaymentMintAtaBump)
	if err != nil {
		return err
	}
	return nil
}

type Creator struct {
	Address  ag_solanago.PublicKey
	Verified bool
	Share    uint8
}

func (obj Creator) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Address` param:
	err = encoder.Encode(obj.Address)
	if err != nil {
		return err
	}
	// Serialize `Verified` param:
	err = encoder.Encode(obj.Verified)
	if err != nil {
		return err
	}
	// Serialize `Share` param:
	err = encoder.Encode(obj.Share)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Creator) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Address`:
	err = decoder.Decode(&obj.Address)
	if err != nil {
		return err
	}
	// Deserialize `Verified`:
	err = decoder.Decode(&obj.Verified)
	if err != nil {
		return err
	}
	// Deserialize `Share`:
	err = decoder.Decode(&obj.Share)
	if err != nil {
		return err
	}
	return nil
}

type LaunchStageType ag_binary.BorshEnum

const (
	LaunchStageTypeInvalid LaunchStageType = iota
	LaunchStageTypeNormalSale
	LaunchStageTypeRaffle
)

func (value LaunchStageType) String() string {
	switch value {
	case LaunchStageTypeInvalid:
		return "Invalid"
	case LaunchStageTypeNormalSale:
		return "NormalSale"
	case LaunchStageTypeRaffle:
		return "Raffle"
	default:
		return ""
	}
}

type WalletLimitSpecification interface {
	isWalletLimitSpecification()
}

type walletLimitSpecificationContainer struct {
	Enum          ag_binary.BorshEnum `borsh_enum:"true"`
	NoLimit       NoLimit
	FixedLimit    WalletLimitSpecificationFixedLimit
	VariableLimit VariableLimit
}

type NoLimit uint8

func (obj NoLimit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (obj *NoLimit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

func (_ *NoLimit) isWalletLimitSpecification() {}

type WalletLimitSpecificationFixedLimit struct {
	Limit uint8
}

func (obj WalletLimitSpecificationFixedLimit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Limit` param:
	err = encoder.Encode(obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WalletLimitSpecificationFixedLimit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Limit`:
	err = decoder.Decode(&obj.Limit)
	if err != nil {
		return err
	}
	return nil
}

func (_ *WalletLimitSpecificationFixedLimit) isWalletLimitSpecification() {}

type VariableLimit uint8

func (obj VariableLimit) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}

func (obj *VariableLimit) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

func (_ *VariableLimit) isWalletLimitSpecification() {}
