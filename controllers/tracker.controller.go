package controllers

import (
	"fmt"
	"my-assets-be/database"
	"strings"
	"sync"
	"time"

	"github.com/umbracle/go-web3"
	"github.com/umbracle/go-web3/blocktracker"
	"github.com/umbracle/go-web3/jsonrpc"
)

type TrackerClient struct {
	Client    *jsonrpc.Client
	ChainHost string
	Chain     string
	Symbol    string
}

var trackerClients map[string]TrackerClient
var trackerSync sync.Once

func GetTrackers() map[string]TrackerClient {
	trackerSync.Do(func() {
		if trackerClients == nil {
			trackerClients = make(map[string]TrackerClient)
		}
	})
	return trackerClients
}

func GetTracker(symbol string) TrackerClient {
	trackerSync.Do(func() {
		if trackerClients == nil {
			trackerClients = make(map[string]TrackerClient)
		}
	})
	return trackerClients[symbol]
}

func NewTracker(host string, symbol string) (*TrackerClient, error) {
	clients := GetTrackers()
	client, err := jsonrpc.NewClient(host)
	if err != nil {
		return nil, err
	}
	var tracker = TrackerClient{
		Client:    client,
		ChainHost: host,
		Symbol:    symbol,
		Chain:     symbol,
	}
	clients[symbol] = tracker
	return &tracker, nil
}

func (client *TrackerClient) Track() bool {
	if client == nil {
		return false
	}

	tracker := blocktracker.NewBlockTracker(client.Client.Eth())

	if err := tracker.Init(); err != nil {
		panic(err)
	}

	go tracker.Start()

	sub := tracker.Subscribe()
	go func() {
		for {
			select {
			case evnt := <-sub:
				if evnt.Type == 0 {
					OnNewBlockFound(tracker, client)
				}
			}
		}
	}()

	return true
}

func OnNewBlockFound(tracker *blocktracker.BlockTracker, client *TrackerClient) {
	lastBlockNumber := tracker.LastBlocked().Number
	addMap := database.GetAddressMap()

	block, _ := client.Client.Eth().GetBlockByNumber(web3.BlockNumber(lastBlockNumber), true)

	for _, tx := range block.Transactions {

		if addMap[strings.ToLower(tx.From.String())] {
			fmt.Println("New block of", client.Symbol, ":", lastBlockNumber)
			fmt.Println("Found an transaction in this block", tx.Hash.String())
			time.Sleep(3 * time.Second)
			// SyncTx(strings.ToLower(tx.From.String()), )
		} else if tx.To != nil && addMap[strings.ToLower(tx.To.String())] {
			fmt.Println("New block of", client.Symbol, ":", lastBlockNumber)
			fmt.Println("Found an transaction in this block", tx.Hash.String())
			time.Sleep(3 * time.Second)
			// SyncTx(strings.ToLower(tx.To.String()))
		}
	}
}
