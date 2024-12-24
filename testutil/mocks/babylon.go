// Code generated by MockGen. DO NOT EDIT.
// Source: clientcontroller/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	types0 "github.com/babylonlabs-io/babylon/x/finality/types"
	types1 "github.com/dapplink-labs/bbn-fp-l2/types"
	btcec "github.com/btcsuite/btcd/btcec/v2"
	schnorr "github.com/btcsuite/btcd/btcec/v2/schnorr"
	gomock "github.com/golang/mock/gomock"
)

// MockClientController is a mock of ClientController interface.
type MockClientController struct {
	ctrl     *gomock.Controller
	recorder *MockClientControllerMockRecorder
}

// MockClientControllerMockRecorder is the mock recorder for MockClientController.
type MockClientControllerMockRecorder struct {
	mock *MockClientController
}

// NewMockClientController creates a new mock instance.
func NewMockClientController(ctrl *gomock.Controller) *MockClientController {
	mock := &MockClientController{ctrl: ctrl}
	mock.recorder = &MockClientControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientController) EXPECT() *MockClientControllerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockClientController) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientControllerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientController)(nil).Close))
}

// CommitPubRandList mocks base method.
func (m *MockClientController) CommitPubRandList(fpPk *btcec.PublicKey, startHeight, numPubRand uint64, commitment []byte, sig *schnorr.Signature) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitPubRandList", fpPk, startHeight, numPubRand, commitment, sig)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CommitPubRandList indicates an expected call of CommitPubRandList.
func (mr *MockClientControllerMockRecorder) CommitPubRandList(fpPk, startHeight, numPubRand, commitment, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitPubRandList", reflect.TypeOf((*MockClientController)(nil).CommitPubRandList), fpPk, startHeight, numPubRand, commitment, sig)
}

// EditFinalityProvider mocks base method.
func (m *MockClientController) EditFinalityProvider(fpPk *btcec.PublicKey, commission *math.LegacyDec, description []byte) (*types.MsgEditFinalityProvider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditFinalityProvider", fpPk, commission, description)
	ret0, _ := ret[0].(*types.MsgEditFinalityProvider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditFinalityProvider indicates an expected call of EditFinalityProvider.
func (mr *MockClientControllerMockRecorder) EditFinalityProvider(fpPk, commission, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditFinalityProvider", reflect.TypeOf((*MockClientController)(nil).EditFinalityProvider), fpPk, commission, description)
}

// QueryActivatedHeight mocks base method.
func (m *MockClientController) QueryActivatedHeight() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryActivatedHeight")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryActivatedHeight indicates an expected call of QueryActivatedHeight.
func (mr *MockClientControllerMockRecorder) QueryActivatedHeight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryActivatedHeight", reflect.TypeOf((*MockClientController)(nil).QueryActivatedHeight))
}

// QueryBestBlock mocks base method.
func (m *MockClientController) QueryBestBlock() (*types1.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBestBlock")
	ret0, _ := ret[0].(*types1.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBestBlock indicates an expected call of QueryBestBlock.
func (mr *MockClientControllerMockRecorder) QueryBestBlock() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBestBlock", reflect.TypeOf((*MockClientController)(nil).QueryBestBlock))
}

// QueryBlock mocks base method.
func (m *MockClientController) QueryBlock(height uint64) (*types1.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBlock", height)
	ret0, _ := ret[0].(*types1.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBlock indicates an expected call of QueryBlock.
func (mr *MockClientControllerMockRecorder) QueryBlock(height interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBlock", reflect.TypeOf((*MockClientController)(nil).QueryBlock), height)
}

// QueryBlocks mocks base method.
func (m *MockClientController) QueryBlocks(startHeight, endHeight uint64, limit uint32) ([]*types1.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryBlocks", startHeight, endHeight, limit)
	ret0, _ := ret[0].([]*types1.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryBlocks indicates an expected call of QueryBlocks.
func (mr *MockClientControllerMockRecorder) QueryBlocks(startHeight, endHeight, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryBlocks", reflect.TypeOf((*MockClientController)(nil).QueryBlocks), startHeight, endHeight, limit)
}

// QueryFinalityActivationBlockHeight mocks base method.
func (m *MockClientController) QueryFinalityActivationBlockHeight() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFinalityActivationBlockHeight")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFinalityActivationBlockHeight indicates an expected call of QueryFinalityActivationBlockHeight.
func (mr *MockClientControllerMockRecorder) QueryFinalityActivationBlockHeight() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFinalityActivationBlockHeight", reflect.TypeOf((*MockClientController)(nil).QueryFinalityActivationBlockHeight))
}

// QueryFinalityProviderHighestVotedHeight mocks base method.
func (m *MockClientController) QueryFinalityProviderHighestVotedHeight(fpPk *btcec.PublicKey) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFinalityProviderHighestVotedHeight", fpPk)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFinalityProviderHighestVotedHeight indicates an expected call of QueryFinalityProviderHighestVotedHeight.
func (mr *MockClientControllerMockRecorder) QueryFinalityProviderHighestVotedHeight(fpPk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFinalityProviderHighestVotedHeight", reflect.TypeOf((*MockClientController)(nil).QueryFinalityProviderHighestVotedHeight), fpPk)
}

// QueryFinalityProviderSlashedOrJailed mocks base method.
func (m *MockClientController) QueryFinalityProviderSlashedOrJailed(fpPk *btcec.PublicKey) (bool, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFinalityProviderSlashedOrJailed", fpPk)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// QueryFinalityProviderSlashedOrJailed indicates an expected call of QueryFinalityProviderSlashedOrJailed.
func (mr *MockClientControllerMockRecorder) QueryFinalityProviderSlashedOrJailed(fpPk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFinalityProviderSlashedOrJailed", reflect.TypeOf((*MockClientController)(nil).QueryFinalityProviderSlashedOrJailed), fpPk)
}

// QueryFinalityProviderVotingPower mocks base method.
func (m *MockClientController) QueryFinalityProviderVotingPower(fpPk *btcec.PublicKey, blockHeight uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryFinalityProviderVotingPower", fpPk, blockHeight)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryFinalityProviderVotingPower indicates an expected call of QueryFinalityProviderVotingPower.
func (mr *MockClientControllerMockRecorder) QueryFinalityProviderVotingPower(fpPk, blockHeight interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryFinalityProviderVotingPower", reflect.TypeOf((*MockClientController)(nil).QueryFinalityProviderVotingPower), fpPk, blockHeight)
}

// QueryLastCommittedPublicRand mocks base method.
func (m *MockClientController) QueryLastCommittedPublicRand(fpPk *btcec.PublicKey, count uint64) (map[uint64]*types0.PubRandCommitResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLastCommittedPublicRand", fpPk, count)
	ret0, _ := ret[0].(map[uint64]*types0.PubRandCommitResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLastCommittedPublicRand indicates an expected call of QueryLastCommittedPublicRand.
func (mr *MockClientControllerMockRecorder) QueryLastCommittedPublicRand(fpPk, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLastCommittedPublicRand", reflect.TypeOf((*MockClientController)(nil).QueryLastCommittedPublicRand), fpPk, count)
}

// QueryLatestFinalizedBlocks mocks base method.
func (m *MockClientController) QueryLatestFinalizedBlocks(count uint64) ([]*types1.BlockInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryLatestFinalizedBlocks", count)
	ret0, _ := ret[0].([]*types1.BlockInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryLatestFinalizedBlocks indicates an expected call of QueryLatestFinalizedBlocks.
func (mr *MockClientControllerMockRecorder) QueryLatestFinalizedBlocks(count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryLatestFinalizedBlocks", reflect.TypeOf((*MockClientController)(nil).QueryLatestFinalizedBlocks), count)
}

// RegisterFinalityProvider mocks base method.
func (m *MockClientController) RegisterFinalityProvider(fpPk *btcec.PublicKey, pop []byte, commission *math.LegacyDec, description []byte) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterFinalityProvider", fpPk, pop, commission, description)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RegisterFinalityProvider indicates an expected call of RegisterFinalityProvider.
func (mr *MockClientControllerMockRecorder) RegisterFinalityProvider(fpPk, pop, commission, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterFinalityProvider", reflect.TypeOf((*MockClientController)(nil).RegisterFinalityProvider), fpPk, pop, commission, description)
}

// SubmitBatchFinalitySigs mocks base method.
func (m *MockClientController) SubmitBatchFinalitySigs(fpPk *btcec.PublicKey, blocks []*types1.BlockInfo, pubRandList []*btcec.FieldVal, proofList [][]byte, sigs []*btcec.ModNScalar) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitBatchFinalitySigs", fpPk, blocks, pubRandList, proofList, sigs)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitBatchFinalitySigs indicates an expected call of SubmitBatchFinalitySigs.
func (mr *MockClientControllerMockRecorder) SubmitBatchFinalitySigs(fpPk, blocks, pubRandList, proofList, sigs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitBatchFinalitySigs", reflect.TypeOf((*MockClientController)(nil).SubmitBatchFinalitySigs), fpPk, blocks, pubRandList, proofList, sigs)
}

// SubmitFinalitySig mocks base method.
func (m *MockClientController) SubmitFinalitySig(fpPk *btcec.PublicKey, block *types1.BlockInfo, pubRand *btcec.FieldVal, proof []byte, sig *btcec.ModNScalar) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitFinalitySig", fpPk, block, pubRand, proof, sig)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitFinalitySig indicates an expected call of SubmitFinalitySig.
func (mr *MockClientControllerMockRecorder) SubmitFinalitySig(fpPk, block, pubRand, proof, sig interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitFinalitySig", reflect.TypeOf((*MockClientController)(nil).SubmitFinalitySig), fpPk, block, pubRand, proof, sig)
}

// UnjailFinalityProvider mocks base method.
func (m *MockClientController) UnjailFinalityProvider(fpPk *btcec.PublicKey) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnjailFinalityProvider", fpPk)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnjailFinalityProvider indicates an expected call of UnjailFinalityProvider.
func (mr *MockClientControllerMockRecorder) UnjailFinalityProvider(fpPk interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnjailFinalityProvider", reflect.TypeOf((*MockClientController)(nil).UnjailFinalityProvider), fpPk)
}