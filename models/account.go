package models

//easyjson:json
type Like struct {
	Id int16
	Ts int32
}

type AccountId = uint16
type SmallId = uint8

//easyjson:json
type AccountFormEntry struct {
	Id        *uint16   // required
	Email     *string   // required
	Fname     string    // opt
	Sname     string    // opt
	Phone     string    // opt
	Sex       *string   // required
	Birth     *uint32   // required
	Country   string    // opt
	City      string    // opt
	Joined    *uint32   // required
	Status    *string   // required
	Interests *[]string // required
	Premium   *struct { // opt
		Start  *uint32 // required
		Finish *uint32 // required
	}
	Likes *[]*Like // required
}

type AccountEntry struct {
	Id            AccountId
	Email         string
	Fname         string
	Sname         string
	Phone         string
	Sex           uint8
	Birth         uint32
	Country       SmallId
	City          SmallId
	Joined        uint32
	Status        SmallId
	Interests     []SmallId
	PremiumStart  uint32
	PremiumFinish uint32
	PremiumActive bool
	Likes         []*Like
}
