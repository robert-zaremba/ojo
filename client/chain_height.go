package client

import (
	"context"
	"errors"
	"sync"

	"github.com/rs/zerolog"
	tmrpcclient "github.com/tendermint/tendermint/rpc/client"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

var (
	errParseEventDataNewBlockHeader = errors.New("error parsing EventDataNewBlockHeader")
	queryEventNewBlockHeader        = tmtypes.QueryForEvent(tmtypes.EventNewBlockHeader)
)

// ChainHeight is used to cache the chain height of the
// current node which is being updated each time the
// node sends an event of EventNewBlockHeader.
// It starts a goroutine to subscribe to blockchain new block event and update the cached height.
type ChainHeight struct {
	Logger zerolog.Logger

	mtx               sync.RWMutex
	errGetChainHeight error
	lastChainHeight   int64
	HeightChanged     chan (int64)
}

// NewChainHeight returns a new ChainHeight struct that
// starts a new goroutine subscribed to EventNewBlockHeader.
func NewChainHeight(
	ctx context.Context,
	rpcClient tmrpcclient.Client,
	logger zerolog.Logger,
) (*ChainHeight, error) {
	if !rpcClient.IsRunning() {
		if err := rpcClient.Start(); err != nil {
			return nil, err
		}
	}

	newBlockHeaderSubscription, err := rpcClient.Subscribe(
		ctx, tmtypes.EventNewBlockHeader, queryEventNewBlockHeader.String())
	if err != nil {
		return nil, err
	}

	chainHeight := &ChainHeight{
		Logger: logger.With().Str("oracle_client", "chain_height").Logger(),
	}
	chainHeight.HeightChanged = make(chan int64)

	go chainHeight.subscribe(ctx, rpcClient, newBlockHeaderSubscription)

	return chainHeight, nil
}

// updateChainHeight receives the data to be updated thread safe.
func (chainHeight *ChainHeight) updateChainHeight(blockHeight int64, err error) {
	chainHeight.mtx.Lock()
	defer chainHeight.mtx.Unlock()

	if chainHeight.lastChainHeight != blockHeight {
		select {
		case chainHeight.HeightChanged <- blockHeight:
		default:
		}
	}
	chainHeight.lastChainHeight = blockHeight
	chainHeight.errGetChainHeight = err
}

// subscribe listens to new blocks being made
// and updates the chain height.
func (chainHeight *ChainHeight) subscribe(
	ctx context.Context,
	eventsClient tmrpcclient.EventsClient,
	newBlockHeaderSubscription <-chan tmctypes.ResultEvent,
) {
	for {
		select {
		case <-ctx.Done():
			err := eventsClient.Unsubscribe(ctx, tmtypes.EventNewBlockHeader, queryEventNewBlockHeader.String())
			if err != nil {
				chainHeight.Logger.Err(err)
				chainHeight.updateChainHeight(chainHeight.lastChainHeight, err)
			}
			chainHeight.Logger.Info().Msg("closing the ChainHeight subscription")
			return

		case resultEvent := <-newBlockHeaderSubscription:
			eventDataNewBlockHeader, ok := resultEvent.Data.(tmtypes.EventDataNewBlockHeader)
			if !ok {
				chainHeight.Logger.Err(errParseEventDataNewBlockHeader)
				chainHeight.updateChainHeight(chainHeight.lastChainHeight, errParseEventDataNewBlockHeader)
				continue
			}
			chainHeight.updateChainHeight(eventDataNewBlockHeader.Header.Height, nil)
		}
	}
}

// GetChainHeight returns the last chain height available.
func (chainHeight *ChainHeight) GetChainHeight() (int64, error) {
	chainHeight.mtx.RLock()
	defer chainHeight.mtx.RUnlock()

	return chainHeight.lastChainHeight, chainHeight.errGetChainHeight
}