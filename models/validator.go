package models

import (
	"time"
)

const ValidatorStatusNotReady = 1
const ValidatorStatusReady = 2

type Validator struct {
	ID                   uint                  `json:"id" bun:",pk"`
	RewardAddressID      *uint                 `json:"reward_address_id"`
	OwnerAddressID       *uint                 `json:"owner_address_id"`
	ControlAddressID     *uint                 `json:"control_address_id"`
	CreatedAtBlockID     *uint                 `json:"created_at_block_id"`
	PublicKey            string                `json:"public_key"           bun:"type:varchar(64)"`
	Status               *uint8                `json:"status"`
	Commission           *uint64               `json:"commission"`
	TotalStake           *string               `json:"total_stake"          bun:"type:numeric(70)"`
	Name                 *string               `json:"name"`
	SiteUrl              *string               `json:"site_url"`
	IconUrl              *string               `json:"icon_url"`
	Description          *string               `json:"description"`
	MetaUpdatedAtBlockID *uint64               `json:"meta_updated_at_block_id"`
	UpdateAt             *time.Time            `json:"update_at"`
	ControlAddress       *Address              `json:"control_address" bun:"rel:belongs-to,join:control_address_id=id"`
	RewardAddress        *Address              `json:"reward_address"  bun:"rel:belongs-to,join:reward_address_id=id"`
	OwnerAddress         *Address              `json:"owner_address"   bun:"rel:belongs-to,join:owner_address_id=id"`
	Stakes               []*Stake              `json:"stakes"          bun:"rel:has-many"`
	PublicKeys           []ValidatorPublicKeys `json:"public_keys"     bun:"rel:has-many"`
	Bans                 []ValidatorBan        `json:"bans"            bun:"rel:has-many"`
}

//Return validators PK with prefix
func (v Validator) GetPublicKey() string {
	return `Mp` + v.PublicKey
}