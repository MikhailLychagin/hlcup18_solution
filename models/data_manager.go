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
	if accForm.Id == nil {
		return nil, &ValueError{"missing field: id"}
	}

	if accForm.Email == nil {
		return nil, &ValueError{"missing field: email"}
	}
	// TODO: check email unique

	if accForm.Sex == nil {
		return nil, &ValueError{"missing field: sex"}
	}

	if accForm.Birth == nil {
		return nil, &ValueError{"missing field: birth"}
	}

	if accForm.Joined == nil {
		return nil, &ValueError{"missing field: joined"}
	}

	if accForm.Status == nil {
		return nil, &ValueError{"missing field: status"}
	}

	if accForm.Interests == nil {
		return nil, &ValueError{"missing field: interests"}
	}

	if accForm.Likes == nil {
		return nil, &ValueError{"missing field: likes"}
	}

	var sex uint8
	switch *accForm.Sex {
	case "m":
		sex = 0
	case "f":
		sex = 1
	default:
		return nil, &ValueError{"invalid field value: sex"}
	}

	var status uint8
	switch *accForm.Status {
	case "свободны":
		status = 0
	case "заняты":
		status = 1
	case "всё сложно":
		status = 2
	default:
		return nil, &ValueError{"invalid field value: status"}
	}

	var country, city uint8
	var interests []uint8
	var premiumActive bool
	var premiumStart, premiumFinish uint32
	premiumActive = false
	if accForm.Country != "" {
		country, _ = dm.GetId(accForm.Country, dm.CountryToId)
	}
	if accForm.City != "" {
		city, _ = dm.GetId(accForm.City, dm.CityToId)
	}
	if accForm.Interests != nil {
		interests, _ = dm.InterestsStrToId(*accForm.Interests, dm.InterestToId)
	}
	if accForm.Premium != nil {
		if accForm.Premium.Start == nil || accForm.Premium.Finish == nil {
			return nil, &ValueError{"invalid field value: premium should contain both start and finish"}
		}
		premiumStart = *accForm.Premium.Start
		premiumFinish = *accForm.Premium.Finish
		premiumActive = (premiumStart <= dm.CurrentDate) && (premiumFinish >= dm.CurrentDate)
	}

	acc := AccountEntry{
		Id:            *accForm.Id,
		Email:         *accForm.Email,
		Fname:         accForm.Fname,
		Sname:         accForm.Sname,
		Phone:         accForm.Phone,
		Sex:           sex,
		Birth:         *accForm.Birth,
		Country:       country,
		City:          city,
		Joined:        *accForm.Joined,
		Status:        status,
		Interests:     interests,
		PremiumStart:  premiumStart,
		PremiumFinish: premiumFinish,
		PremiumActive: premiumActive,
		Likes:         *accForm.Likes,
	}

	return &acc, nil
}
