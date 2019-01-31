package models

//easyjson:json
type Like struct {
	Id int16
	Ts int32
}

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
	Id            uint16
	Email         string
	Fname         string
	Sname         string
	Phone         string
	Sex           uint8
	Birth         uint32
	Country       uint8
	City          uint8
	Joined        uint32
	Status        uint8
	Interests     []uint8
	PremiumStart  uint32
	PremiumFinish uint32
	PremiumActive bool
	Likes         []*Like
}
