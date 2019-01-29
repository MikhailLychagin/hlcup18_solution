package models

//easyjson:json
type Like struct {
	Id int16
	Ts int32
}

//easyjson:json
type AccountEntry struct {
	Id            int16
	Email         string
	Fname         string
	Sname         string
	Phone         string
	Sex           string
	Birth         int32
	Country       string
	City          string
	Joined        int32
	Status        string
	Interests     []string
	PremiumStart  int32
	PremiumFinish int32
	PremiumActive bool
	Likes         []*Like
}
