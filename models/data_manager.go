package models

type DataManager struct {
	Accounts map[int16]*AccountEntry
}

func (dm *DataManager) AccountAdd(acc *AccountEntry) error {
	dm.Accounts[acc.Id] = acc

	return nil
}

func (dm *DataManager) AccountDel(acc *AccountEntry) error {
	delete(dm.Accounts, acc.Id)

	return nil
}
