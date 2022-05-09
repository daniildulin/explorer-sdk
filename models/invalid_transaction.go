package models

import "time"

type InvalidTransaction struct {
	ID                  uint64            `json:"id" bun:",pk"`
	FromAddressID       uint64            `json:"from_address_id"`
	BlockID             uint64            `json:"block_id"`
	CreatedAt           time.Time         `json:"created_at"`
	Type                uint8             `json:"type"`
	Hash                string            `json:"hash"`
	TxData              string            `json:"tx_data"`
	Log                 string            `json:"log"`
	Nonce               uint64            `json:"nonce"`
	GasPrice            uint64            `json:"gas_price"`
	Gas                 uint64            `json:"gas"`
	Commission          string            `json:"commission"`
	GasCoinID           uint64            `json:"gas_coin_id" pg:",use_zero"`
	ServiceData         string            `json:"service_data"`
	Tags                map[string]string `json:"tags"`
	Payload             []byte            `json:"payload"`
	RawTx               []byte            `json:"raw_tx"`
	Block               *Block            `bun:"rel:has-one"`                    //Relation has one to Blocks
	FromAddress         *Address          `bun:"rel:has-one,fk:from_address_id"` //Relation has one to Addresses
	GasCoin             *Coin             `json:"gas_coin"                 bun:"rel:has-one,join:gas_coin_id=id"`
	CommissionPriceCoin interface{}       `json:"commission_price_coin"    bun:"-"`
}

// GetHash Return transactions hash with prefix
func (t InvalidTransaction) GetHash() string {
	return `Mt` + t.Hash
}