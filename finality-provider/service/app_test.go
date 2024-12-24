package service_test

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/babylonlabs-io/babylon/testutil/datagen"
	bbntypes "github.com/babylonlabs-io/babylon/types"
	finalitytypes "github.com/babylonlabs-io/babylon/x/finality/types"
	sdkkeyring "github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/dapplink-labs/bbn-fp-l2/clientcontroller"
	"github.com/dapplink-labs/bbn-fp-l2/eotsmanager"
	eotscfg "github.com/dapplink-labs/bbn-fp-l2/eotsmanager/config"
	"github.com/dapplink-labs/bbn-fp-l2/finality-provider/config"
	"github.com/dapplink-labs/bbn-fp-l2/finality-provider/proto"
	"github.com/dapplink-labs/bbn-fp-l2/finality-provider/service"
	fpstore "github.com/dapplink-labs/bbn-fp-l2/finality-provider/store"
	"github.com/dapplink-labs/bbn-fp-l2/keyring"
	"github.com/dapplink-labs/bbn-fp-l2/testutil"
	"github.com/dapplink-labs/bbn-fp-l2/types"
	"github.com/dapplink-labs/bbn-fp-l2/util"
)

const (
	passphrase            = "testpass"
	hdPath                = ""
	eventuallyWaitTimeOut = 5 * time.Second
	eventuallyPollTime    = 10 * time.Millisecond
)

func FuzzCreateFinalityProvider(f *testing.F) {
	testutil.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		logger := zap.NewNop()
		// create an EOTS manager
		eotsHomeDir := filepath.Join(t.TempDir(), "eots-home")
		eotsCfg := eotscfg.DefaultConfigWithHomePath(eotsHomeDir)
		dbBackend, err := eotsCfg.DatabaseConfig.GetDBBackend()
		require.NoError(t, err)
		em, err := eotsmanager.NewLocalEOTSManager(eotsHomeDir, eotsCfg.KeyringBackend, dbBackend, logger)
		require.NoError(t, err)
		defer func() {
			dbBackend.Close()
			err = os.RemoveAll(eotsHomeDir)
			require.NoError(t, err)
		}()

		// Create mocked babylon client
		randomStartingHeight := uint64(r.Int63n(100) + 1)
		currentHeight := randomStartingHeight + uint64(r.Int63n(10)+2)
		mockClientController := testutil.PrepareMockedClientController(t, r, randomStartingHeight, currentHeight, 0)
		mockClientController.EXPECT().QueryLatestFinalizedBlocks(gomock.Any()).Return(nil, nil).AnyTimes()
		mockClientController.EXPECT().QueryFinalityProviderVotingPower(gomock.Any(),
			gomock.Any()).Return(uint64(0), nil).AnyTimes()

		// Create randomized config
		fpHomeDir := filepath.Join(t.TempDir(), "fp-home")
		fpCfg := config.DefaultConfigWithHome(fpHomeDir)
		fpCfg.PollerConfig.AutoChainScanningMode = false
		fpCfg.PollerConfig.StaticChainScanningStartHeight = randomStartingHeight
		fpdb, err := fpCfg.DatabaseConfig.GetDBBackend()
		require.NoError(t, err)
		app, err := service.NewFinalityProviderApp(&fpCfg, mockClientController, em, fpdb, logger)
		require.NoError(t, err)
		defer func() {
			err = fpdb.Close()
			require.NoError(t, err)
			err = os.RemoveAll(fpHomeDir)
			require.NoError(t, err)
		}()

		err = app.Start()
		require.NoError(t, err)
		defer func() {
			err = app.Stop()
			require.NoError(t, err)
		}()

		var eotsPk *bbntypes.BIP340PubKey
		eotsKeyName := testutil.GenRandomHexStr(r, 4)
		require.NoError(t, err)
		eotsPkBz, err := em.CreateKey(eotsKeyName, passphrase, hdPath)
		require.NoError(t, err)
		eotsPk, err = bbntypes.NewBIP340PubKey(eotsPkBz)
		require.NoError(t, err)

		// generate keyring
		keyName := testutil.GenRandomHexStr(r, 4)
		chainID := testutil.GenRandomHexStr(r, 4)

		cfg := app.GetConfig()
		_, err = testutil.CreateChainKey(cfg.BabylonConfig.KeyDirectory, cfg.BabylonConfig.ChainID, keyName, sdkkeyring.BackendTest, passphrase, hdPath, "")
		require.NoError(t, err)

		txHash := testutil.GenRandomHexStr(r, 32)
		mockClientController.EXPECT().
			RegisterFinalityProvider(
				eotsPk.MustToBTCPK(),
				gomock.Any(),
				testutil.ZeroCommissionRate(),
				gomock.Any(),
			).Return(&types.TxResponse{TxHash: txHash}, nil).AnyTimes()
		res, err := app.CreateFinalityProvider(keyName, chainID, passphrase, eotsPk, testutil.RandomDescription(r), testutil.ZeroCommissionRate())
		require.NoError(t, err)
		require.Equal(t, txHash, res.TxHash)

		fpInfo, err := app.GetFinalityProviderInfo(eotsPk)
		require.NoError(t, err)
		require.Equal(t, eotsPk.MarshalHex(), fpInfo.BtcPkHex)
	})
}

func FuzzSyncFinalityProviderStatus(f *testing.F) {
	testutil.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		randomStartingHeight := uint64(r.Int63n(100) + 1)
		currentHeight := randomStartingHeight + uint64(r.Int63n(10)+2)
		mockClientController := testutil.PrepareMockedClientController(t, r, randomStartingHeight, currentHeight, 0)

		blkInfo := &types.BlockInfo{Height: currentHeight}

		mockClientController.EXPECT().QueryLastCommittedPublicRand(gomock.Any(), uint64(1)).Return(nil, nil).AnyTimes()
		mockClientController.EXPECT().QueryLatestFinalizedBlocks(gomock.Any()).Return(nil, nil).AnyTimes()
		mockClientController.EXPECT().QueryBestBlock().Return(blkInfo, nil).Return(blkInfo, nil).AnyTimes()
		mockClientController.EXPECT().QueryBlock(gomock.Any()).Return(nil, errors.New("chain not online")).AnyTimes()

		noVotingPowerTable := r.Int31n(10) > 5
		if noVotingPowerTable {
			allowedErr := fmt.Sprintf("failed to query Finality Voting Power at Height %d: rpc error: code = Unknown desc = %s: unknown request",
				currentHeight, finalitytypes.ErrVotingPowerTableNotUpdated.Wrapf("height: %d", currentHeight).Error())
			mockClientController.EXPECT().QueryFinalityProviderVotingPower(gomock.Any(), gomock.Any()).Return(uint64(0), nil).AnyTimes()
			mockClientController.EXPECT().QueryActivatedHeight().Return(uint64(0), errors.New(allowedErr)).AnyTimes()
		} else {
			mockClientController.EXPECT().QueryActivatedHeight().Return(currentHeight, nil).AnyTimes()
			mockClientController.EXPECT().QueryFinalityProviderVotingPower(gomock.Any(), gomock.Any()).Return(uint64(2), nil).AnyTimes()
		}
		mockClientController.EXPECT().QueryFinalityProviderHighestVotedHeight(gomock.Any()).Return(uint64(0), nil).AnyTimes()
		var isSlashedOrJailed int
		if noVotingPowerTable {
			// 0 means is slashed, 1 means is jailed, 2 means neither slashed nor jailed
			isSlashedOrJailed = r.Intn(3)
			switch isSlashedOrJailed {
			case 0:
				mockClientController.EXPECT().QueryFinalityProviderSlashedOrJailed(gomock.Any()).Return(true, false, nil).AnyTimes()
			case 1:
				mockClientController.EXPECT().QueryFinalityProviderSlashedOrJailed(gomock.Any()).Return(false, true, nil).AnyTimes()
			case 2:
				mockClientController.EXPECT().QueryFinalityProviderSlashedOrJailed(gomock.Any()).Return(false, false, nil).AnyTimes()
			}
		}

		// Create randomized config
		pathSuffix := datagen.GenRandomHexStr(r, 10)
		fpHomeDir := filepath.Join(t.TempDir(), "fp-home", pathSuffix)
		fpCfg := config.DefaultConfigWithHome(fpHomeDir)
		// no need for other intervals to run
		fpCfg.StatusUpdateInterval = time.Minute * 10
		fpCfg.SubmissionRetryInterval = time.Minute * 10

		// Create fp app
		app, fpPk, cleanup := startFPAppWithRegisteredFp(t, r, fpHomeDir, &fpCfg, mockClientController)
		defer cleanup()

		fpInfo, err := app.GetFinalityProviderInfo(fpPk)
		require.NoError(t, err)

		expectedStatus := proto.FinalityProviderStatus_ACTIVE
		if noVotingPowerTable {
			switch isSlashedOrJailed {
			case 0:
				expectedStatus = proto.FinalityProviderStatus_SLASHED
			case 1:
				expectedStatus = proto.FinalityProviderStatus_JAILED
			case 2:
				expectedStatus = proto.FinalityProviderStatus_REGISTERED
			}
		}

		require.Equal(t, fpInfo.Status, expectedStatus.String())
	})
}

func FuzzUnjailFinalityProvider(f *testing.F) {
	testutil.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		randomStartingHeight := uint64(r.Int63n(100) + 1)
		currentHeight := randomStartingHeight + uint64(r.Int63n(10)+2)
		mockClientController := testutil.PrepareMockedClientController(t, r, randomStartingHeight, currentHeight, 0)

		// Create randomized config
		pathSuffix := datagen.GenRandomHexStr(r, 10)
		fpHomeDir := filepath.Join(t.TempDir(), "fp-home", pathSuffix)
		fpCfg := config.DefaultConfigWithHome(fpHomeDir)
		// use shorter interval for the test to end faster
		fpCfg.StatusUpdateInterval = time.Millisecond * 10
		fpCfg.SubmissionRetryInterval = time.Millisecond * 10

		blkInfo := &types.BlockInfo{Height: currentHeight}

		mockClientController.EXPECT().QueryLastCommittedPublicRand(gomock.Any(), uint64(1)).Return(nil, nil).AnyTimes()
		mockClientController.EXPECT().QueryLatestFinalizedBlocks(gomock.Any()).Return(nil, nil).AnyTimes()
		mockClientController.EXPECT().QueryBestBlock().Return(blkInfo, nil).Return(blkInfo, nil).AnyTimes()
		mockClientController.EXPECT().QueryBlock(gomock.Any()).Return(nil, errors.New("chain not online")).AnyTimes()

		// set voting power to be positive so that the fp should eventually become ACTIVE
		mockClientController.EXPECT().QueryFinalityProviderVotingPower(gomock.Any(), gomock.Any()).Return(uint64(0), nil).AnyTimes()
		mockClientController.EXPECT().QueryActivatedHeight().Return(uint64(1), nil).AnyTimes()
		mockClientController.EXPECT().QueryFinalityProviderSlashedOrJailed(gomock.Any()).Return(false, true, nil).AnyTimes()
		mockClientController.EXPECT().QueryFinalityProviderHighestVotedHeight(gomock.Any()).Return(uint64(0), nil).AnyTimes()

		// Create fp app
		app, fpPk, cleanup := startFPAppWithRegisteredFp(t, r, fpHomeDir, &fpCfg, mockClientController)
		defer cleanup()

		expectedTxHash := datagen.GenRandomHexStr(r, 32)
		mockClientController.EXPECT().UnjailFinalityProvider(fpPk.MustToBTCPK()).Return(&types.TxResponse{TxHash: expectedTxHash}, nil)
		res, err := app.UnjailFinalityProvider(fpPk)
		require.NoError(t, err)
		require.Equal(t, expectedTxHash, res.TxHash)
		fpInfo, err := app.GetFinalityProviderInfo(fpPk)
		require.NoError(t, err)
		require.Equal(t, proto.FinalityProviderStatus_INACTIVE.String(), fpInfo.GetStatus())
	})
}

func FuzzStatusUpdate(f *testing.F) {
	testutil.AddRandomSeedsToFuzzer(f, 10)
	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))

		randomStartingHeight := uint64(r.Int63n(100) + 1)
		currentHeight := randomStartingHeight + uint64(r.Int63n(10)+2)
		mockClientController := testutil.PrepareMockedClientController(t, r, randomStartingHeight, currentHeight, 0)

		// setup mocks
		votingPower := uint64(r.Intn(2))
		mockClientController.EXPECT().QueryFinalityProviderVotingPower(gomock.Any(), currentHeight).Return(votingPower, nil).AnyTimes()
		mockClientController.EXPECT().Close().Return(nil).AnyTimes()
		mockClientController.EXPECT().QueryLatestFinalizedBlocks(gomock.Any()).Return(nil, nil).AnyTimes()
		mockClientController.EXPECT().QueryFinalityProviderHighestVotedHeight(gomock.Any()).Return(uint64(0), nil).AnyTimes()
		mockClientController.EXPECT().QueryLastCommittedPublicRand(gomock.Any(), uint64(1)).Return(nil, nil).AnyTimes()
		mockClientController.EXPECT().SubmitFinalitySig(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(&types.TxResponse{TxHash: ""}, nil).AnyTimes()
		mockClientController.EXPECT().QueryFinalityProviderSlashedOrJailed(gomock.Any()).Return(false, false, nil).AnyTimes()

		// Create randomized config
		pathSuffix := datagen.GenRandomHexStr(r, 10)
		fpHomeDir := filepath.Join(t.TempDir(), "fp-home", pathSuffix)
		fpCfg := config.DefaultConfigWithHome(fpHomeDir)
		// use shorter interval for the test to end faster
		fpCfg.StatusUpdateInterval = time.Millisecond * 10
		fpCfg.SubmissionRetryInterval = time.Millisecond * 10

		// Create fp app
		app, fpPk, cleanup := startFPAppWithRegisteredFp(t, r, fpHomeDir, &fpCfg, mockClientController)
		defer cleanup()

		err := app.StartFinalityProvider(fpPk, passphrase)
		require.NoError(t, err)
		fpIns, err := app.GetFinalityProviderInstance()
		require.NoError(t, err)

		if votingPower > 0 {
			waitForStatus(t, fpIns, proto.FinalityProviderStatus_ACTIVE)
		} else if fpIns.GetStatus() == proto.FinalityProviderStatus_ACTIVE {
			waitForStatus(t, fpIns, proto.FinalityProviderStatus_INACTIVE)
		}
	})
}

func waitForStatus(t *testing.T, fpIns *service.FinalityProviderInstance, s proto.FinalityProviderStatus) {
	require.Eventually(t,
		func() bool {
			return fpIns.GetStatus() == s
		}, eventuallyWaitTimeOut, eventuallyPollTime)
}

func startFPAppWithRegisteredFp(t *testing.T, r *rand.Rand, homePath string, cfg *config.Config, cc clientcontroller.ClientController) (*service.FinalityProviderApp, *bbntypes.BIP340PubKey, func()) {
	logger := zap.NewNop()
	// create an EOTS manager
	eotsHomeDir := filepath.Join(t.TempDir(), "eots-home")
	eotsCfg := eotscfg.DefaultConfigWithHomePath(eotsHomeDir)
	eotsdb, err := eotsCfg.DatabaseConfig.GetDBBackend()
	require.NoError(t, err)
	em, err := eotsmanager.NewLocalEOTSManager(eotsHomeDir, eotsCfg.KeyringBackend, eotsdb, logger)
	require.NoError(t, err)

	// create finality-provider app with randomized config
	input := strings.NewReader("")
	require.NoError(t, err)
	err = util.MakeDirectory(config.DataDir(homePath))
	require.NoError(t, err)
	db, err := cfg.DatabaseConfig.GetDBBackend()
	require.NoError(t, err)
	fpStore, err := fpstore.NewFinalityProviderStore(db)
	require.NoError(t, err)
	app, err := service.NewFinalityProviderApp(cfg, cc, em, db, logger)
	require.NoError(t, err)

	// create registered finality-provider
	keyName := datagen.GenRandomHexStr(r, 10)
	chainID := datagen.GenRandomHexStr(r, 10)
	kr, err := keyring.CreateKeyring(
		cfg.BabylonConfig.KeyDirectory,
		cfg.BabylonConfig.ChainID,
		cfg.BabylonConfig.KeyringBackend,
		input,
	)
	require.NoError(t, err)
	kc, err := keyring.NewChainKeyringControllerWithKeyring(kr, keyName, input)
	require.NoError(t, err)
	btcPkBytes, err := em.CreateKey(keyName, passphrase, hdPath)
	require.NoError(t, err)
	btcPk, err := bbntypes.NewBIP340PubKey(btcPkBytes)
	require.NoError(t, err)
	keyInfo, err := kc.CreateChainKey(passphrase, hdPath, "")
	require.NoError(t, err)
	fpAddr := keyInfo.AccAddress

	err = fpStore.CreateFinalityProvider(
		fpAddr,
		btcPk.MustToBTCPK(),
		testutil.RandomDescription(r),
		testutil.ZeroCommissionRate(),
		chainID,
	)
	require.NoError(t, err)
	err = app.Start()
	require.NoError(t, err)

	cleanUp := func() {
		err = app.Stop()
		require.NoError(t, err)
		err = eotsdb.Close()
		require.NoError(t, err)
		err = db.Close()
		require.NoError(t, err)
		err = os.RemoveAll(eotsHomeDir)
		require.NoError(t, err)
		err = os.RemoveAll(homePath)
		require.NoError(t, err)
	}

	return app, btcPk, cleanUp
}