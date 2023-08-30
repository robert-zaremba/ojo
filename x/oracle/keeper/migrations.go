package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ojo-network/ojo/x/oracle/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper *Keeper
}

// NewMigrator creates a Migrator.
func NewMigrator(keeper *Keeper) Migrator {
	return Migrator{keeper: keeper}
}

// MigrateValidatorSet fixes the validator set being stored as map
// causing non determinism by storing it as a list.
func (m Migrator) MigrateValidatorSet(ctx sdk.Context) {
	m.keeper.SetValidatorRewardSet(ctx)
}

// MigratePriceFeederCurrencyPairProviders adds the price feeder
// currency pair provider list.
func (m Migrator) MigratePriceFeederCurrencyPairProviders(ctx sdk.Context) {
	priceFeederCurrencyPairProviders := types.CurrencyPairProvidersList{
		types.CurrencyPairProviders{
			BaseDenom:  "USDT",
			QuoteDenom: "USD",
			Providers: []string{
				"kraken",
				"coinbase",
				"crypto",
				"gate",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "ATOM",
			QuoteDenom: "USDT",
			Providers: []string{
				"okx",
				"bitget",
				"gate",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "ATOM",
			QuoteDenom: "USD",
			Providers: []string{
				"kraken",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "ETH",
			QuoteDenom: "USDT",
			Providers: []string{
				"okx",
				"bitget",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "ETH",
			QuoteDenom: "USD",
			Providers: []string{
				"kraken",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "BTC",
			QuoteDenom: "USDT",
			Providers: []string{
				"okx",
				"gate",
				"bitget",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "BTC",
			QuoteDenom: "USD",
			Providers: []string{
				"coinbase",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "OSMO",
			QuoteDenom: "USDT",
			Providers: []string{
				"bitget",
				"gate",
				"huobi",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "OSMO",
			QuoteDenom: "ATOM",
			Providers: []string{
				"osmosisv2",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "stATOM",
			QuoteDenom: "ATOM",
			Providers: []string{
				"osmosisv2",
				"crescent",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "stOSMO",
			QuoteDenom: "OSMO",
			Providers: []string{
				"osmosisv2",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "DAI",
			QuoteDenom: "USDT",
			Providers: []string{
				"okx",
				"bitget",
				"huobi",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "DAI",
			QuoteDenom: "USD",
			Providers: []string{
				"kraken",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "JUNO",
			QuoteDenom: "USDT",
			Providers: []string{
				"bitget",
				"mexc",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "JUNO",
			QuoteDenom: "USD",
			Providers: []string{
				"kraken",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "JUNO",
			QuoteDenom: "ATOM",
			Providers: []string{
				"osmosisv2",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "stJUNO",
			QuoteDenom: "JUNO",
			Providers: []string{
				"osmosisv2",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "SCRT",
			QuoteDenom: "USD",
			Providers: []string{
				"kraken",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "SCRT",
			QuoteDenom: "USDT",
			Providers: []string{
				"mexc",
				"gate",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "WBTC",
			QuoteDenom: "USDT",
			Providers: []string{
				"okx",
				"bitget",
				"crypto",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "USDC",
			QuoteDenom: "USDT",
			Providers: []string{
				"okx",
				"bitget",
				"kraken",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "USDC",
			QuoteDenom: "USD",
			Providers: []string{
				"kraken",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "IST",
			QuoteDenom: "OSMO",
			Providers: []string{
				"osmosisv2",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "IST",
			QuoteDenom: "USDC",
			Providers: []string{
				"crescent",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "BNB",
			QuoteDenom: "USDT",
			Providers: []string{
				"mexc",
				"bitget",
				"okx",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "LUNA",
			QuoteDenom: "USDT",
			Providers: []string{
				"okx",
				"gate",
				"huobi",
				"bitget",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "DOT",
			QuoteDenom: "USD",
			Providers: []string{
				"kraken",
				"coinbase",
				"crypto",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "DOT",
			QuoteDenom: "USDT",
			Providers: []string{
				"gate",
				"bitget",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "AXL",
			QuoteDenom: "USD",
			Providers: []string{
				"coinbase",
				"crypto",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "AXL",
			QuoteDenom: "OSMO",
			Providers: []string{
				"osmosisv2",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "STARS",
			QuoteDenom: "ATOM",
			Providers: []string{
				"osmosisv2",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "STARS",
			QuoteDenom: "OSMO",
			Providers: []string{
				"osmosisv2",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "XRP",
			QuoteDenom: "USD",
			Providers: []string{
				"kraken",
				"coinbase",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "XRP",
			QuoteDenom: "USDT",
			Providers: []string{
				"gate",
				"mexc",
				"bitget",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "USK",
			QuoteDenom: "USDC",
			Providers: []string{
				"kujira",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "KUJI",
			QuoteDenom: "USDC",
			Providers: []string{
				"kujira",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:  "MNTA",
			QuoteDenom: "USDC",
			Providers: []string{
				"kujira",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:       "RETH",
			QuoteDenom:      "WETH",
			Address:         "0xa4e0faA58465A2D369aa21B3e42d43374c6F9613",
			AddressProvider: "eth-uniswap",
			Providers: []string{
				"eth-uniswap",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:       "WETH",
			QuoteDenom:      "USDC",
			Address:         "0x88e6a0c2ddd26feeb64f039a2c41296fcb3f5640",
			AddressProvider: "eth-uniswap",
			Providers: []string{
				"eth-uniswap",
			},
		},
		types.CurrencyPairProviders{
			BaseDenom:       "CBETH",
			QuoteDenom:      "WETH",
			Address:         "0x840deeef2f115cf50da625f7368c24af6fe74410",
			AddressProvider: "eth-uniswap",
			Providers: []string{
				"eth-uniswap",
			},
		},
	}
	m.keeper.SetPriceFeederCurrencyPairProvidersList(ctx, priceFeederCurrencyPairProviders)
}

// MigratePriceFeederCurrencyDeviationThresholds adds the price feeder
// currency deviation threshold list.
func (m Migrator) MigratePriceFeederCurrencyDeviationThresholds(ctx sdk.Context) {
	priceFeederCurrencyDeviationThresholds := types.CurrencyDeviationThresholdList{
		types.CurrencyDeviationThreshold{
			BaseDenom: "USDT",
			Threshold: "1.5",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "ATOM",
			Threshold: "1.5",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "ETH",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "BTC",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "OSMO",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "stATOM",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "stOSMO",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "DAI",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "JUNO",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "stJUNO",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "SCRT",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "WBTC",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "USDC",
			Threshold: "1.5",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "IST",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "BNB",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "LUNA",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "DOT",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "AXL",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "STARS",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "XRP",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "USK",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "KUJI",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "MNTA",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "RETH",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "WETH",
			Threshold: "2",
		},
		types.CurrencyDeviationThreshold{
			BaseDenom: "CBETH",
			Threshold: "2",
		},
	}
	m.keeper.SetPriceFeederCurrencyDeviationThresholdList(ctx, priceFeederCurrencyDeviationThresholds)
}
