/*
   Copyright 2021 Erigon contributors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package txpool

import (
	"context"
	"encoding/hex"
	"sync"
	"testing"

	"github.com/ledgerwatch/erigon-lib/direct"
	"github.com/ledgerwatch/erigon-lib/gointerfaces/sentry"
	"github.com/ledgerwatch/erigon-lib/log"
	"github.com/stretchr/testify/require"
)

func TestFetch(t *testing.T) {
	logger := log.NewTest(t)

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	var genesisHash [32]byte
	var networkId uint64 = 1
	forks := []uint64{1, 5, 10}

	mock := NewMockSentry(ctx)
	sentryClient := direct.NewSentryClientDirect(direct.ETH66, mock)
	pool := &PoolMock{
		IdHashKnownFunc: func(hash []byte) bool { return false },
	}

	fetch := NewFetch(ctx, []sentry.SentryClient{sentryClient}, genesisHash, networkId, forks, pool, logger)
	var wg sync.WaitGroup
	fetch.SetWaitGroup(&wg)
	mock.StreamWg.Add(2)
	fetch.Start()
	mock.StreamWg.Wait()
	// Send one transaction id
	wg.Add(1)
	data, _ := hex.DecodeString("e1a0595e27a835cd79729ff1eeacec3120eeb6ed1464a04ec727aaca734ead961328")
	errs := mock.Send(&sentry.InboundMessage{
		Id:     sentry.MessageId_NEW_POOLED_TRANSACTION_HASHES_66,
		Data:   data,
		PeerId: PeerId,
	})
	for i, err := range errs {
		if err != nil {
			t.Errorf("sending new pool tx hashes 66 (%d): %v", i, err)
		}
	}
	wg.Wait()
}

func TestSendTxPropagate(t *testing.T) {
	logger := log.NewTest(t)

	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn()

	mock := NewMockSentry(ctx)
	send := NewSend(ctx, []sentry.SentryClient{direct.NewSentryClientDirect(direct.ETH66, mock)}, nil, logger)
	// Send one transaction id
	send.BroadcastRemotePooledTxs([][32]byte{[32]byte{1}, [32]byte{42}})
	require.Equal(t, sentry.MessageId_NEW_POOLED_TRANSACTION_HASHES_66, mock.sentMessages[0].Id)
}
