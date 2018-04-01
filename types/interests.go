package types

import "log"

type AllInterests struct {
	Sport   []*Interests `json:"sport"`
	Music   []*Interests `json:"music"`
	Outdoor []*Interests `json:"outdoor"`
	Films   []*Interests `json:"films"`
	IT      []*Interests `json:"IT"`
}

func NewAllInterests() *AllInterests {
	return &AllInterests{
		Sport:   make([]*Interests, 0),
		Music:   make([]*Interests, 0),
		Outdoor: make([]*Interests, 0),
		Films:   make([]*Interests, 0),
		IT:      make([]*Interests, 0),
	}
}

type Interests struct {
	ID         int64  `json:"id"`
	CategoryID int64  `json:"-"`
	Name       string `json:"name"`
}

func Categorize(is []*Interests) *AllInterests {
	ic := NewAllInterests()

	for _, i := range is {
		switch i.CategoryID {
		case 1:
			ic.Sport = append(ic.Sport, i)
		case 2:
			ic.Music = append(ic.Music, i)
		case 3:
			ic.Outdoor = append(ic.Outdoor, i)
		case 4:
			ic.Films = append(ic.Films, i)
		case 5:
			ic.IT = append(ic.IT, i)
		default:
			log.Println("unspecified interest", i)
		}
	}
	return ic
}

func (ic *AllInterests) ToArray() []*Interests {
	var is []*Interests
	for _, v := range ic.Sport {
		is = append(is, v)
	}

	for _, v := range ic.Music {
		is = append(is, v)
	}

	for _, v := range ic.Outdoor {
		is = append(is, v)
	}

	for _, v := range ic.Films {
		is = append(is, v)
	}

	for _, v := range ic.IT {
		is = append(is, v)
	}

	return is
}

func (ic *AllInterests) Validate() {
	empty := make([]*Interests, 0)
	if ic.Films == nil {
		ic.Films = empty
	}

	if ic.Sport == nil {
		ic.Sport = empty
	}

	if ic.Outdoor == nil {
		ic.Outdoor = empty
	}

	if ic.Music == nil {
		ic.Music = empty
	}

	if ic.IT == nil {
		ic.IT = empty
	}
}

func (ic *AllInterests) ToMapByID() map[int64]bool {
	m := make(map[int64]bool)
	for _, v := range ic.Sport {
		m[v.ID] = true
	}

	for _, v := range ic.Music {
		m[v.ID] = true
	}

	for _, v := range ic.Outdoor {
		m[v.ID] = true
	}

	for _, v := range ic.Films {
		m[v.ID] = true
	}

	for _, v := range ic.IT {
		m[v.ID] = true
	}

	return m
}
