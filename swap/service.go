package swap

import (
	"errors"
	"github.com/MinterTeam/minter-explorer-extender/v2/models"
	"github.com/go-pg/pg/v10"
	"github.com/starwander/goraph"
	"math/big"
)

type Service struct {
	db *pg.DB
}

func NewService(db *pg.DB) *Service {
	return &Service{db}
}

func (s *Service) GetPoolLiquidity(pools []models.LiquidityPool, p models.LiquidityPool) *big.Float {
	if p.FirstCoinId == 0 {
		return getVolumeInBip(big.NewFloat(2), p.FirstCoinVolume)
	}

	currentVolume := p.FirstCoinVolume

	paths, err := s.FindSwapRoutePathsByGraph(pools, p.FirstCoinId, uint64(0), 4)
	if err != nil {
		paths, err = s.FindSwapRoutePathsByGraph(pools, p.SecondCoinId, uint64(0), 4)
		if err != nil {
			return big.NewFloat(0)
		}

		currentVolume = p.SecondCoinVolume
	}

	path := paths[0]

	var price *big.Float
	for i := len(path) - 2; i >= 0; i-- {
		secondCoinId := path[i+1].(uint64)
		firstCoinId := path[i].(uint64)

		for _, pool := range pools {
			if (pool.FirstCoinId == firstCoinId && pool.SecondCoinId == secondCoinId) || (pool.FirstCoinId == secondCoinId && pool.SecondCoinId == firstCoinId) {
				cprice := big.NewFloat(0)
				if pool.FirstCoinId == firstCoinId {
					cprice = computePrice(pool.SecondCoinVolume, pool.FirstCoinVolume)
				} else {
					cprice = computePrice(pool.FirstCoinVolume, pool.SecondCoinVolume)
				}

				if price == nil {
					price = cprice
				} else {
					price.Mul(price, cprice)
				}

				break
			}
		}
	}

	liquidity := getVolumeInBip(price, currentVolume)
	return liquidity.Mul(liquidity, big.NewFloat(2))
}

func (s *Service) FindSwapRoutePathsByGraph(pools []models.LiquidityPool, fromCoinId, toCoinId uint64, depth int) ([][]goraph.ID, error) {
	graph := goraph.NewGraph()
	for _, pool := range pools {
		graph.AddVertex(pool.FirstCoinId, pool.FirstCoin)
		graph.AddVertex(pool.SecondCoinId, pool.SecondCoin)
		graph.AddEdge(pool.FirstCoinId, pool.SecondCoinId, 1, nil)
		graph.AddEdge(pool.SecondCoinId, pool.FirstCoinId, 1, nil)
	}

	_, paths, err := graph.Yen(fromCoinId, toCoinId, 1000)
	if err != nil {
		return nil, err
	}

	if len(paths[0]) == 0 {
		return nil, errors.New("path not found")
	}

	if depth == 0 {
		return paths, nil
	}

	var result [][]goraph.ID
	for _, path := range paths {
		if len(path) > depth+1 || len(path) == 0 {
			break
		}

		result = append(result, path)
	}

	if len(result) == 0 {
		return nil, errors.New("path not found")
	}

	return result, nil
}
