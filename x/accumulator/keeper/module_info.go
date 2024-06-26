package keeper

//type ModuleInfo struct {
//	address string
//	amount  uint64
//}

//func NewModuleInfo(address string, amount uint64) (ModuleInfo, error) {
//	module := ModuleInfo{address: address, amount: amount}
//	if err := module.ValidateBasic(); err != nil {
//		return ModuleInfo{}, errors.Wrap(err, "failed to create module info")
//	}
//
//	return module, nil
//}
//
//func (m ModuleInfo) ValidateBasic() error {
//	if m.address == "" {
//		return fmt.Errorf("address cannot be empty")
//	}
//
//	if m.amount == 0 {
//		return fmt.Errorf("amount cannot be zero")
//	}
//	return nil
//}
