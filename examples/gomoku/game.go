package main

import (
	"log"
	"math/rand"
)

type Point struct {
	net   *Net
	x     int
	y     int
	slots []*Slot
	r     [3]int
	s     int
}

func NewPoint(net *Net, x int, y int) *Point {
	point := &Point{net: net, x: x, y: y}
	//log.Println(*point)
	return point
}

func (point *Point) isValidScp(d int) bool {
	x := point.x - 7
	y := point.y - 7
	// 0 - vert, 1 - horiz, 2 - up, 3 - down
	if d == 0 && y > -6 && y < 6 {
		return true
	}
	if d == 1 && x > -6 && x < 6 {
		return true
	}
	if d == 2 && (x > -6 && y < 6) && (x < 6 && y > -6) {
		return true
	}
	if d == 3 && (x > -6 && y > -6) && (x < 6 && y < 6) {
		return true
	}
	return false
}

func (point *Point) addSlot(s *Slot) {
	point.slots = append(point.slots, s)
	point.r[s.s] += 1
}

type Slot struct {
	net    *Net
	scp    *Point
	d      int
	points [5]*Point
	r      int
	s      int
}

func NewSlot(net *Net, scp *Point, d int) *Slot {
	slot := &Slot{net: net, scp: scp, d: d}
	slot.points[2] = net.getPoint(scp.x, scp.y)
	//log.Println(*slot)
	return slot
}

func (slot *Slot) Init() {
	slot.points[2] = slot.net.getPoint(slot.scp.x, slot.scp.y)
	if slot.d == 0 {
		slot.points[0] = slot.net.getPoint(slot.scp.x, slot.scp.y-2)
		slot.points[1] = slot.net.getPoint(slot.scp.x, slot.scp.y-1)
		slot.points[3] = slot.net.getPoint(slot.scp.x, slot.scp.y+1)
		slot.points[4] = slot.net.getPoint(slot.scp.x, slot.scp.y+2)
	} else if slot.d == 1 {
		slot.points[0] = slot.net.getPoint(slot.scp.x-2, slot.scp.y)
		slot.points[1] = slot.net.getPoint(slot.scp.x-1, slot.scp.y)
		slot.points[3] = slot.net.getPoint(slot.scp.x+1, slot.scp.y)
		slot.points[4] = slot.net.getPoint(slot.scp.x+2, slot.scp.y)
	} else if slot.d == 2 {
		slot.points[0] = slot.net.getPoint(slot.scp.x-2, slot.scp.y-2)
		slot.points[1] = slot.net.getPoint(slot.scp.x-1, slot.scp.y-1)
		slot.points[3] = slot.net.getPoint(slot.scp.x+1, slot.scp.y+1)
		slot.points[4] = slot.net.getPoint(slot.scp.x+2, slot.scp.y+2)
	} else if slot.d == 3 {
		slot.points[0] = slot.net.getPoint(slot.scp.x-2, slot.scp.y+2)
		slot.points[1] = slot.net.getPoint(slot.scp.x-1, slot.scp.y+1)
		slot.points[3] = slot.net.getPoint(slot.scp.x+1, slot.scp.y-1)
		slot.points[4] = slot.net.getPoint(slot.scp.x+2, slot.scp.y-2)
	}
	for _, p := range slot.points {
		p.addSlot(slot)
	}
}

type Net struct {
	all_slots    []*Slot
	active_slots [3][]*Slot
	all_points   []*Point
	empty_points []*Point
	steps        []Step
}

func NewNet(steps []Step) *Net {
	net := &Net{}
	net.steps = steps
	net.all_slots = []*Slot{}
	net.active_slots = [3][]*Slot{}
	net.all_points = make([]*Point, 225)
	net.empty_points = make([]*Point, 225)

	for i := 0; i < 225; i++ {
		p := NewPoint(net, int(i/15), i%15)
		net.all_points[i] = p
		for d := 0; d < 4; d++ {
			if p.isValidScp(d) {
				s := NewSlot(net, p, d)
				//log.Printf("%v - %v,%v", d, s.scp.x, s.scp.y)
				net.all_slots = append(net.all_slots, s)
			}
		}
	}
	copy(net.empty_points, net.all_points)
	net.active_slots[0] = make([]*Slot, len(net.all_slots))
	copy(net.active_slots[0], net.all_slots)
	//log.Printf("All Slots: %v", len(net.all_slots))
	//log.Print(net.active_slots[0])
	//log.Print(net.all_slots)
	//log.Printf("Active Slots 0: %v", len(net.active_slots[0]))
	//log.Printf("Active Slots 1: %v", len(net.active_slots[1]))
	//log.Printf("Active Slots 2: %v", len(net.active_slots[2]))
	for _, s := range net.all_slots {
		s.Init()
	}

	//for _, p := range net.all_points {
	//	log.Printf("%v,%v %v", p.x, p.y, len(p.slots))
	//}

	for i, st := range steps {
		net.addStep(i, st)
		/*
			c := i%2 + 1
			p := net.getPoint(int(st[0]), int(st[1]))
			p.s = c
			//net.empty_points.remove(p)
			n := net.findPoint(net.empty_points[:], p)
			net.empty_points = append(net.empty_points[:n], net.empty_points[n+1:]...)

			for _, s := range p.slots {
				if s.s == 0 {
					p.r[0] -= 1
					p.r[c] += 1
					s.s = c
					s.r = 1
					//net.active_slots[0].remove(s)
					m := net.findSlot(net.active_slots[0][:], s)
					//k := net.findSlot(net.all_slots, s)
					//log.Print(net.active_slots[0][m], net.all_slots[k])
					net.active_slots[0] = append(net.active_slots[0][:m], net.active_slots[0][m+1:]...)
					//log.Print(net.active_slots[c][l], net.all_slots[k])
				} else if s.s == c {
					p.r[c] += 1
					s.r += 1
				} else if s.s != 3 {
					p.r[c] -= 1
					//net.active_slots[s.s].remove(s)
					m := net.findSlot(net.active_slots[s.s][:], s)
					net.active_slots[s.s] = append(net.active_slots[s.s][:m], net.active_slots[s.s][m+1:]...)
					s.s = 3
				}
			}
		*/
	}
	//log.Print(net)
	log.Printf("Active Slots 0: %v (%v)", len(net.active_slots[0]), net.countSlots(0))
	log.Printf("Active Slots 1: %v (%v)", len(net.active_slots[1]), net.countSlots(1))
	log.Printf("Active Slots 2: %v (%v)", len(net.active_slots[2]), net.countSlots(2))
	log.Printf("Slots 3: (%v)", net.countSlots(3))
	log.Printf("Empty Points: %v", len(net.empty_points))
	return net
}

func (net *Net) addStep(i int, st Step) {
	//log.Println(i, st)
	c := i%2 + 1
	p := net.getPoint(st.x, st.y)
	//log.Println(c, p)
	p.s = c
	//net.empty_points.remove(p)
	n := net.findPoint(net.empty_points[:], p)
	net.empty_points = append(net.empty_points[:n], net.empty_points[n+1:]...)

	for _, s := range p.slots {
		//log.Println(s)
		if s.s == 0 {
			p.r[0] -= 1
			p.r[c] += 1
			s.s = c
			s.r = 1
			//net.active_slots[0].remove(s)
			m := net.findSlot(net.active_slots[0][:], s)
			//k := net.findSlot(net.all_slots, s)
			//log.Print(net.active_slots[0][m], net.all_slots[k])
			net.active_slots[0] = append(net.active_slots[0][:m], net.active_slots[0][m+1:]...)
			//log.Print(net.active_slots[c][l], net.all_slots[k])
			net.active_slots[c] = append(net.active_slots[c][:], s)
		} else if s.s == c {
			p.r[c] += 1
			s.r += 1
		} else if s.s != 3 {
			p.r[c] -= 1
			//net.active_slots[s.s].remove(s)
			m := net.findSlot(net.active_slots[s.s][:], s)
			net.active_slots[s.s] = append(net.active_slots[s.s][:m], net.active_slots[s.s][m+1:]...)
			s.s = 3
		}
	}
}

func (net *Net) countSlots(s int) int {
	c := 0
	for _, e := range net.all_slots {
		//if s == 1 {
		//	log.Println(e)
		//}
		if e.s == s {
			c += 1
		}
	}

	return c
}

func (net *Net) checkWin() bool {
	for _, s := range net.active_slots[1] {
		if s.r == 5 {
			//net.mes = net.name_c[c] + " :: win!!!"
			return true
		}
	}
	for _, s := range net.active_slots[2] {
		if s.r == 5 {
			//net.mes = net.name_c[c] + " :: win!!!"
			return true
		}
	}
	return false
}

func (net *Net) checkDraw() bool {
	if len(net.active_slots[0]) == 0 && len(net.active_slots[1]) == 0 && len(net.active_slots[2]) == 0 {
		//net.mes = " draw :("
		return true
	} else {
		return false
	}
}

func (net *Net) calcPoint() Step {
	a := Step{}
	c := len(net.steps)%2 + 1
	ret := net.findSlot4(c)
	if len(ret) == 0 {
		ret = net.findSlot4(3 - c)
	}

	if len(ret) == 0 {
		ret = net.findPointX(c, 2, 1)
	}
	if len(ret) == 0 {
		ret = net.findPointX(3-c, 2, 1)
	}

	if len(ret) == 0 {
		ret = net.findPointX(c, 1, 5)
	}
	if len(ret) == 0 {
		ret = net.findPointX(3-c, 1, 5)
	}

	if len(ret) == 0 {
		ret = net.findPointX(c, 1, 4)
	}
	if len(ret) == 0 {
		ret = net.findPointX(3-c, 1, 4)
	}

	if len(ret) == 0 {
		ret = net.findPointX(c, 1, 3)
	}
	if len(ret) == 0 {
		ret = net.findPointX(3-c, 1, 3)
	}

	if len(ret) == 0 {
		ret = net.findPointX(c, 1, 2)
	}
	if len(ret) == 0 {
		ret = net.findPointX(3-c, 1, 2)
	}

	if len(ret) == 0 {
		ret = net.findPointX(c, 1, 1)
	}
	if len(ret) == 0 {
		ret = net.findPointX(3-c, 1, 1)
	}

	if len(ret) == 0 {
		ret = net.findPointX(c, 0, 10)
	}
	if len(ret) == 0 {
		ret = net.findPointX(3-c, 0, 10)
	}

	if len(ret) == 0 {
		ret = net.findPointX(c, 0, 9)
	}
	if len(ret) == 0 {
		ret = net.findPointX(3-c, 0, 9)
	}

	if len(ret) == 0 {
		ret = net.findPointX(c, 0, 8)
	}
	if len(ret) == 0 {
		ret = net.findPointX(3-c, 0, 8)
	}

	if len(ret) == 0 {
		ret = net.findPointX(c, 0, 7)
	}
	if len(ret) == 0 {
		ret = net.findPointX(3-c, 0, 7)
	}

	if len(ret) == 0 {
		ret = net.calcPointMaxRate(c)
	}

	n := rand.Intn(len(ret))
	a.x = ret[n][0]
	a.y = ret[n][1]
	return a
}

func (net *Net) findSlot4(c int) [][]int {
	ret := [][]int{}
	//msg := fmt.Sprintf("%v :: find_slot_4(%v,%v)", c)
	for _, s := range net.active_slots[c] {
		if s.r == 4 {
			for _, p := range s.points {
				if p.s == 0 {
					elm := []int{p.x, p.y}
					ret = append(ret, elm)
					//msg := fmt.Sprintf("%v :: find_slot_4 ->(%v,%v)", c)
				}
			}
		}
	}
	if len(ret) > 0 {
		log.Printf("%v :: find_slot4 -> %v", c, ret)
	}
	return ret
}

func (net *Net) findPointX(c int, r int, b int) [][]int {
	ret := [][]int{}
	//msg := fmt.Sprintf("%v :: find_point_x(%v,%v)", c, r, b)
	for _, p := range net.empty_points {
		i := 0
		for _, s := range p.slots {
			if s.s == c && s.r > r {
				i += 1
				if i > b {
					elm := []int{p.x, p.y}
					ret = append(ret, elm)
					//msg = fmt.Sprintf("%v :: find_point_x(%v,%v) -> (%v, %v)", c, r, b, elm[0], elm[1])
				}
			}
		}
	}
	if len(ret) > 0 {
		log.Printf("%v :: find_point_x(%v,%v) -> %v", c, r, b, ret)
	}
	return ret
}

func (net *Net) calcPointMaxRate(c int) [][]int {
	ret := [][]int{}
	r := -1
	d := 0
	i := 0
	//msg := fmt.Sprintf("%v :: point_max_rate(%v,%v)", c, i, r)
	for _, p := range net.empty_points {
		d = 0
		for _, s := range p.slots {
			if s.s == 0 {
				d += 1
			} else if s.s == c {
				d += (1 + s.r) * (1 + s.r)
			} else if s.s != 3 {
				d += (1 + s.r) * (1 + s.r)
			}
		}
		if d > r {
			i = 1
			r = d
			ret = [][]int{}
			elm := []int{p.x, p.y}
			ret = append(ret, elm)
			//msg = fmt.Sprintf("%v :: point_max_rate(%v,%v) -> (%v, %v)", c, i, r, elm[0], elm[1])
		} else if d == r {
			i += 1
			elm := []int{p.x, p.y}
			ret = append(ret, elm)
			//msg = fmt.Sprintf("%v :: point_max_rate(%v,%v) -> (%v, %v)", c, i, r, elm[0], elm[1])
		}
	}
	log.Printf("%v :: point_max_rate(%v,%v) -> %v", c, i, r, ret)
	return ret
}

func (net *Net) getPoint(x int, y int) *Point {
	return net.all_points[x*15+y]
}

func (net *Net) findPoint(a []*Point, p *Point) int {
	for i, e := range a {
		if e == p {
			return i
		}
	}
	return -1
}

func (net *Net) findSlot(a []*Slot, s *Slot) int {
	for i, e := range a {
		if e == s {
			return i
		}
	}
	return -1
}

//type Step [2]float64
type Step struct {
	x int
	y int
}

//type Steps map[string]interface{}
//type Pack struct {
//	Game   Steps  `json:"game"`
//	Status string `json:"status"`
//}

func calcStep(steps []Step) ([]Step, string) {
	//log.Println(string(body))
	//var pack Pack

	//json.Unmarshal(body, &pack)
	//log.Println(steps)
	//arStep := pack.Game
	status := ""
	net := NewNet(steps)
	if !net.checkWin() && !net.checkDraw() {
		newStep := net.calcPoint()
		net.addStep(len(steps), newStep)
		steps = append(steps, newStep) //[2]float64{8, 8})
		if !net.checkWin() && !net.checkDraw() {
			status = "play"
		} else {
			status = "over"
		}
	} else {
		status = "over"
	}

	//newGame, _ := json.Marshal(pack) //"{\"game\":[[7,7], [8,8]]}"
	//response := string(newGame)
	//log.Println(response)
	return steps, status
}
