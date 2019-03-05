package models

import (
	"container/list"
	"log"
	"strings"

	"github.com/emirpasic/gods/trees/binaryheap"
	"github.com/emirpasic/gods/utils"
)

type StrToListIndex map[string]*list.List
type IdToListIndex []*list.List

type IsNullIndex struct {
	name           string
	values         []*AccountEntry
	pointerToEmpty int
}

func (idx IsNullIndex) Add(value *AccountEntry) {
	idx.values[idx.pointerToEmpty] = value
	idx.pointerToEmpty += 1
	if idx.pointerToEmpty == len(idx.values) {
		newCapacity := len(idx.values) + 10
		log.Printf("Increasing size of %s from %d to %d", idx.name, len(idx.values), newCapacity)
		old := idx.values
		idx.values = make([]*AccountEntry, newCapacity)
		copy(idx.values, old)
	}
}

func NewIsNullIndex(capacity int, name string) *IsNullIndex {
	return &IsNullIndex{name, make([]*AccountEntry, capacity), 0}
}

type IndexManager struct {
	FnameIdx            StrToListIndex
	SnameIdx            StrToListIndex
	PhoneCodeIdx        IdToListIndex
	PhoneNullIdx        *IsNullIndex
	CountryIdIdx        IdToListIndex
	CityIdIdx           IdToListIndex
	HaveInterestsIdsIdx IdToListIndex // i=Interest.Id
	HaveLikeIdsIdx      IdToListIndex // i=AccountEntry.Id of likee, j=list of likers
	PremiumActiveIdx    IdToListIndex // 0=no, 1=yes
	SexIdx              IdToListIndex // 0=M, 1=F
	EmailIxd            *binaryheap.Heap
	DomainIxd           StrToListIndex
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
		DomainIxd:           make(StrToListIndex, 256),
	}
	for i := range im.CountryIdIdx {
		im.CountryIdIdx[i] = list.New()
	}
	for i := range im.CityIdIdx {
		im.CityIdIdx[i] = list.New()
	}
	for i := range im.HaveInterestsIdsIdx {
		im.HaveInterestsIdsIdx[i] = list.New()
	}
	for i := range im.HaveLikeIdsIdx {
		im.HaveLikeIdsIdx[i] = list.New()
	}
	for i := range im.PremiumActiveIdx {
		im.PremiumActiveIdx[i] = list.New()
	}
	for i := range im.SexIdx {
		im.SexIdx[i] = list.New()
	}
	im.PhoneNullIdx = NewIsNullIndex(256, "PhoneNullIdx")

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

	// Email domain index
	separator_pos := strings.Index(*value, "@")
	if separator_pos == -1 {
		return &ValueError{"email missing @"}
	}
	for i := separator_pos + 1; i < len(*value); i++ {
		if (*value)[i] == '.' {
			domain := (*value)[separator_pos+1 : i]
			if err := im.AddToStrToListIndex(&domain, acc, &im.DomainIxd); err != nil {
				return err
			}
			break
		}
	}

	return nil
}

func (im *IndexManager) AddToIsNullIndex(acc *AccountEntry, idx *IsNullIndex) error {
	idx.Add(acc)

	return nil
}
