package model

import orm "happyD/util/database"

type Odd struct {
	DateTime string  `json:"date_time"` // 列名为 `id`
	MatchId  string  `json:"match_id"`  // 列名为 `username`
	CompId   string  `json:"comp_id"`   // 列名为 `password`
	HomeOdd  float32 `json:"home_odd"`  // 列名为 `password`
	GuestOdd float32 `json:"guest_odd"` // 列名为 `password`
	OddTerm  string  `json:"odd_term"`  // 列名为 `password`

}

//var Odds = []Odd

func (odd *Odd) Odds() (odds []Odd, err error) {
	if err = orm.Eloquent.Find(&odds).Error; err != nil {
		return
	}
	return
}
