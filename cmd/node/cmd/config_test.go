package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_LoadConfigFromFile(t *testing.T) {
	configFileName := "/testConfig/nodeConfigTest.json"
	currentProjectPath, err := os.Getwd()
	assert.Equal(t, err, nil, "1")
	configFilePath := filepath.Join(currentProjectPath, configFileName)

	config, err := LoadConfigFromFile(configFilePath)

	assert.Equal(t, err, nil, "2")
	assert.Equal(t, config.BasicConfig.Name, "seele node2", "3")
	assert.Equal(t, config.BasicConfig.Version, "1.0", "4")
	assert.Equal(t, config.BasicConfig.RPCAddr, "0.0.0.0:55028", "5")
	assert.Equal(t, config.BasicConfig.Coinbase, "0x4dd6881d13ab5152127533c5954e4e062eb4bb2dcd93becf4f4e9b1d2d69f1363eea0395e8e76a2716b033d1e3cc8da2bf24811b1e31a86ac8bcacca4c4b29bd", "6")

	assert.Equal(t, config.HTTPServer.HTTPCors[0], "*", "6")
	assert.Equal(t, config.HTTPServer.HTTPCors[0], "*", "7")
	assert.Equal(t, config.HTTPServer.HTTPAddr, "127.0.0.1:65027", "8")

	assert.Equal(t, config.P2PConfig.ListenAddr, "0.0.0.0:39008", "9")
	assert.Equal(t, (int)(config.P2PConfig.NetworkID), 1, "10")
	assert.Equal(t, len(config.P2PConfig.StaticNodes), 1, "10")
	assert.Equal(t, config.P2PConfig.StaticNodes[0].UDPPort, 39007, "11")
	assert.Equal(t, len(config.P2PConfig.StaticNodes[0].IP), 16, "12")
	assert.Equal(t, config.P2PConfig.StaticNodes[0].TCPPort, 0, "13")

	assert.Equal(t, len(config.SeeleConfig.GenesisConfig.Accounts), 2, "14")
	assert.Equal(t, config.SeeleConfig.GenesisConfig.Difficult, int64(22), "15")
	assert.Equal(t, config.SeeleConfig.GenesisConfig.ShardNumber, uint(12), "16")
}
