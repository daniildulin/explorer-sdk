package swap

import (
	"github.com/MinterTeam/minter-explorer-api/v2/helpers"
	"math/big"
)

func getVolumeInBip(price *big.Float, volume string) *big.Float {
	firstCoinBaseVolume := helpers.Pip2Bip(helpers.StringToBigInt(volume))
	return new(big.Float).Mul(firstCoinBaseVolume, price)
}

func computePrice(reserve1, reserve2 string) *big.Float {
	return new(big.Float).Quo(
		helpers.Pip2Bip(helpers.StringToBigInt(reserve1)),
		helpers.Pip2Bip(helpers.StringToBigInt(reserve2)),
	)
}
