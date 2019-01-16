package model

import "encoding/json"

type Profile struct {
	Name 		string
	Age 		string
	Sex			string
	Height 		string
	Income 		string
	Marriage 	string
	Education 	string
	City		string
	ImageUrl	string
}

func FromJsonObj(o interface{})(Profile,error){
	var profile Profile
	s,err := json.Marshal(o)
	if err!= nil {
		return profile,err
	}

	err = json.Unmarshal(s,&profile)
	return profile,err
}
