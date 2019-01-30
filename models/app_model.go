package models

//easyjson:json
type Like struct {
	Id int16
	Ts int32
}

//easyjson:json
type AccountFormEntry struct {
	Id        uint16
	Email     string
	Fname     string
	Sname     string
	Phone     string
	Sex       string
	Birth     uint32
	Country   string
	City      string
	Joined    uint32
	Status    string
	Interests []string
	Premium   struct {
		Start  uint32
		Finish uint32
	}
	Likes []*Like
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
