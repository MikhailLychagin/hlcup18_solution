package models

import (
	"container/list"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

type StrToListIndex map[string]*list.List
type IdToListIndex []*list.List

type IndexManager struct {
	FnameIdx            StrToListIndex
	SnameIdx            StrToListIndex
	PhoneCodeIdx        IdToListIndex
	CountryIdIdx        IdToListIndex
	CityIdIdx           IdToListIndex
	HaveInterestsIdsIdx IdToListIndex // i=Interest.Id
	HaveLikeIdsIdx      IdToListIndex // i=AccountEntry.Id of likee, j=list of likers
	PremiumActiveIdx    IdToListIndex // 0=no, 1=yes
	SexIdx              IdToListIndex // 0=M, 1=F
	EmailIxd            *binaryheap.Heap
}

type IndexEntry struct {
	key   interface{}
	value interface{}
}

type EmailIxdEntry struct {
	key   *string
	value *AccountEntry
}

func StringIndexEntryComparator(a, b interface{}) int {
	return utils.StringComparator(a.(IndexEntry).key, b.(IndexEntry).key)
}

func BuildDefaultIndexManager() *IndexManager {
	var maxSmallId int = int(^SmallId(0)) + 1
	im := IndexManager{
		FnameIdx:            make(StrToListIndex, 256),
		SnameIdx:            make(StrToListIndex, 256),
		PhoneCodeIdx:        make(IdToListIndex, 900), // codes are 3 digit 100 to 999
		CountryIdIdx:        make(IdToListIndex, maxSmallId),
		CityIdIdx:           make(IdToListIndex, maxSmallId),
		HaveInterestsIdsIdx: make(IdToListIndex, maxSmallId),
		HaveLikeIdsIdx:      make(IdToListIndex, maxSmallId),
		PremiumActiveIdx:    make(IdToListIndex, 2),
		SexIdx:              make(IdToListIndex, 2),
		EmailIxd:            binaryheap.NewWith(StringIndexEntryComparator),
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
	for i := range im.PremiumActiveIdx {
		im.CountryIdIdx[i] = list.New()
	}
	for i := range im.SexIdx {
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

func (im *IndexManager) AddSex(value *uint8, acc *AccountEntry) error {
	im.SexIdx[*value].PushFront(acc)

	return nil
}

func (im *IndexManager) AddPremiumActive(value *bool, acc *AccountEntry) error {
	if *value {
		im.PremiumActiveIdx[1].PushFront(acc)
	} else {
		im.PremiumActiveIdx[0].PushFront(acc)
	}

	return nil
}

func (im *IndexManager) AddEmail(value *string, acc *AccountEntry) error {
	im.EmailIxd.Push(EmailIxdEntry{key: value, value: acc})

	return nil
}
