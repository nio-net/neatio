package vm

import (
	"math/big"

	"github.com/neatlab/neatio/chain/core/types"
	"github.com/neatlab/neatio/utilities/common"
)

func NoopCanTransfer(db StateDB, from common.Address, balance *big.Int) bool {
	return true
}
func NoopTransfer(db StateDB, from, to common.Address, amount *big.Int) {}

type NoopEVMCallContext struct{}

func (NoopEVMCallContext) Call(caller ContractRef, addr common.Address, data []byte, gas, value *big.Int) ([]byte, error) {
	return nil, nil
}
func (NoopEVMCallContext) CallCode(caller ContractRef, addr common.Address, data []byte, gas, value *big.Int) ([]byte, error) {
	return nil, nil
}
func (NoopEVMCallContext) Create(caller ContractRef, data []byte, gas, value *big.Int) ([]byte, common.Address, error) {
	return nil, common.Address{}, nil
}
func (NoopEVMCallContext) DelegateCall(me ContractRef, addr common.Address, data []byte, gas *big.Int) ([]byte, error) {
	return nil, nil
}

type NoopStateDB struct{}

func (NoopStateDB) CreateAccount(common.Address)                                       {}
func (NoopStateDB) SubBalance(common.Address, *big.Int)                                {}
func (NoopStateDB) AddBalance(common.Address, *big.Int)                                {}
func (NoopStateDB) GetBalance(common.Address) *big.Int                                 { return nil }
func (NoopStateDB) GetNonce(common.Address) uint64                                     { return 0 }
func (NoopStateDB) SetNonce(common.Address, uint64)                                    {}
func (NoopStateDB) GetCodeHash(common.Address) common.Hash                             { return common.Hash{} }
func (NoopStateDB) GetCode(common.Address) []byte                                      { return nil }
func (NoopStateDB) SetCode(common.Address, []byte)                                     {}
func (NoopStateDB) GetCodeSize(common.Address) int                                     { return 0 }
func (NoopStateDB) AddRefund(uint64)                                                   {}
func (NoopStateDB) GetRefund() uint64                                                  { return 0 }
func (NoopStateDB) GetState(common.Address, common.Hash) common.Hash                   { return common.Hash{} }
func (NoopStateDB) SetState(common.Address, common.Hash, common.Hash)                  {}
func (NoopStateDB) Suicide(common.Address) bool                                        { return false }
func (NoopStateDB) HasSuicided(common.Address) bool                                    { return false }
func (NoopStateDB) Exist(common.Address) bool                                          { return false }
func (NoopStateDB) Empty(common.Address) bool                                          { return false }
func (NoopStateDB) RevertToSnapshot(int)                                               {}
func (NoopStateDB) Snapshot() int                                                      { return 0 }
func (NoopStateDB) AddLog(*types.Log)                                                  {}
func (NoopStateDB) AddPreimage(common.Hash, []byte)                                    {}
func (NoopStateDB) ForEachStorage(common.Address, func(common.Hash, common.Hash) bool) {}
