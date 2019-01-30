package models

import (
	"sort"
)

type ReprToIdMap map[string]uint8
type ValueError struct {
	errMsg string
}

func (e *ValueError) Error() string {
	return e.errMsg
}

type DataManager struct {
	Accounts     map[uint16]*AccountEntry
	CountryToId  ReprToIdMap
	IdToCountry  []*string
	CityToId     ReprToIdMap
	IdToCity     []*string
	InterestToId ReprToIdMap
	IdToInterest []*string
	CurrentDate  uint32
}

func (dm *DataManager) AccountAdd(acc *AccountEntry) error {
	dm.Accounts[acc.Id] = acc

	return nil
}

func (dm *DataManager) AccountDel(acc *AccountEntry) error {
	delete(dm.Accounts, acc.Id)

	return nil
}

func (dm *DataManager) GetId(obj string, idsStorage ReprToIdMap) (uint8, error) {
	var id uint8
	id, exists := idsStorage[obj]
	if !exists {
		id = uint8(len(idsStorage))
		idsStorage[obj] = id
	}

	return id, nil
}

func (dm *DataManager) InterestsStrToId(arr []string, idsStorage ReprToIdMap) ([]uint8, error) {
	res := make([]uint8, len(arr))
	for i, val := range arr {
		res[i], _ = dm.GetId(val, dm.InterestToId)
	}
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res, nil
}

func (dm *DataManager) FormToAccount(accForm *AccountFormEntry) (*AccountEntry, error) {
	var sex uint8
	switch accForm.Sex {
	case "m":
		sex = 0
	case "f":
		sex = 1
	default:
		return nil, &ValueError{"invalid sex"}
	}
	var status uint8
	switch accForm.Status {
	case "свободны":
		status = 0
	case "заняты":
		status = 1
	case "всё сложно":
		status = 2
	default:
		return nil, &ValueError{"invalid status"}
	}
	country, _ := dm.GetId(accForm.Country, dm.CountryToId)
	city, _ := dm.GetId(accForm.City, dm.CityToId)
	interests, _ := dm.InterestsStrToId(accForm.Interests, dm.InterestToId)
	premiumActive := (accForm.Premium.Start <= dm.CurrentDate) && (accForm.Premium.Finish >= dm.CurrentDate)
	acc := AccountEntry{
		Id:            accForm.Id,
		Email:         accForm.Email,
		Fname:         accForm.Fname,
		Sname:         accForm.Sname,
		Phone:         accForm.Phone,
		Sex:           sex,
		Birth:         accForm.Birth,
		Country:       country,
		City:          city,
		Joined:        accForm.Joined,
		Status:        status,
		Interests:     interests,
		PremiumStart:  accForm.Premium.Start,
		PremiumFinish: accForm.Premium.Finish,
		PremiumActive: premiumActive,
		Likes:         accForm.Likes,
	}

	return &acc, nil
}
