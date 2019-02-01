package models

import (
	"container/list"
)

type StrToListIndex map[string]*list.List
type IdToListIndex []*list.List

type IndexManager struct {
	FnameIdx            StrToListIndex
	SnameIdx            StrToListIndex
	PhoneCodeIdx        StrToListIndex
	CountryIdIdx        IdToListIndex
	CityIdIdx           IdToListIndex
	HaveInterestsIdsIdx IdToListIndex // i=Interest.Id
	HaveLikeIdsIdx      IdToListIndex // i=AccountEntry.Id of likee, j=list of likers
	HavePremiumIdx      []*AccountEntry
	NoPremiumIdx        []*AccountEntry
	SexMIdx             []*AccountEntry
}

func BuildDefaultIndexManager() *IndexManager {
	var maxSmallId int = int(^SmallId(0)) + 1
	im := IndexManager{
		FnameIdx:            make(StrToListIndex),
		SnameIdx:            make(StrToListIndex),
		PhoneCodeIdx:        make(StrToListIndex),
		CountryIdIdx:        make(IdToListIndex, maxSmallId),
		CityIdIdx:           make(IdToListIndex, maxSmallId),
		HaveInterestsIdsIdx: make(IdToListIndex, maxSmallId),
		HaveLikeIdsIdx:      make(IdToListIndex, maxSmallId),
	}
	for i := range im.CountryIdIdx {
		im.CountryIdIdx[i] = list.New()
	}
	for i := range im.CityIdIdx {
		im.CountryIdIdx[i] = list.New()
	}
	for i := range im.HaveInterestsIdsIdx {
		im.CountryIdIdx[i] = list.New()
	}
	for i := range im.HaveLikeIdsIdx {
		im.CountryIdIdx[i] = list.New()
	}

	return &im
}

func (im *IndexManager) AddFname(value *string, acc *AccountEntry) error {
	if err := im.AddToStrToListIndex(value, acc, &im.FnameIdx); err != nil {
		return err
	}
	return nil
}

func (im *IndexManager) AddSname(value *string, acc *AccountEntry) error {
	if err := im.AddToStrToListIndex(value, acc, &im.FnameIdx); err != nil {
		return err
	}
	return nil
}

func (im *IndexManager) AddPhoneCode(value *string, acc *AccountEntry) error {
	if (*value) == "" {
		return nil
	}
	phoneCode := (*value)[2:6]
	if err := im.AddToStrToListIndex(&phoneCode, acc, &im.FnameIdx); err != nil {
		return err
	}
	return nil
}

func (im *IndexManager) AddToStrToListIndex(indexingValue *string, acc *AccountEntry, idx *StrToListIndex) error {
	idxEntry, exists := (*idx)[*indexingValue]
	if exists {
		idxEntry.PushFront(acc)
	} else {
		idxEntry = list.New()
		idxEntry.PushFront(acc)
		(*idx)[*indexingValue] = idxEntry
	}

	return nil
}

func (im *IndexManager) AddToIdToListIndex(indexingValue *SmallId, acc *AccountEntry, idx *IdToListIndex) error {
	(*idx)[*indexingValue].PushFront(acc)

	return nil
}
