package framework

var Registry = Reg{
	days: map[int]Day{},
}

type Reg struct {
	days map[int]Day
}

func (r *Reg) Get(day int) Day {
	return r.days[day]
}

func (r *Reg) Register(i int, day Day) {
	r.days[i] = day
}
