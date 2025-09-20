package leetcode

// 3508. Implement Router
// https://leetcode.com/problems/implement-router/

type Packet struct {
	ts, src int
}

type Destination struct {
	packets []Packet
	start   int
	end     int
	dupSet  map[int64]bool
}

func newDestination() *Destination {
	return &Destination{
		packets: make([]Packet, 0),
		start:   -1,
		end:     -1,
		dupSet:  make(map[int64]bool),
	}
}

func (d *Destination) addPacket(ts, src int) {
	if d.start == -1 {
		d.start = 0
	}

	if len(d.packets) > d.end+1 {
		d.packets[d.end+1] = Packet{ts, src}
		d.end++
	} else {
		d.packets = append(d.packets, Packet{ts, src})
		d.end++
	}

	key := int64(ts)*1000000 + int64(src)
	d.dupSet[key] = true
}

func (d *Destination) removePacket() {
	if d.start == -1 {
		return
	}

	pkt := d.packets[d.start]
	key := int64(pkt.ts)*1000000 + int64(pkt.src)
	delete(d.dupSet, key)

	if d.start == d.end {
		d.start = -1
		d.end = -1
	} else {
		d.start++
	}
}

func (d *Destination) getCount(startTime, endTime int) int {
	if d.start == -1 {
		return 0
	}

	lb := d.end + 1
	ub := -1

	// Binary search for lower bound
	left, right := d.start, d.end
	for left <= right {
		mid := left + (right-left)/2
		if d.packets[mid].ts >= startTime {
			lb = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// Binary search for upper bound
	left, right = d.start, d.end
	for left <= right {
		mid := left + (right-left)/2
		if d.packets[mid].ts <= endTime {
			ub = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if lb <= ub {
		return ub - lb + 1
	}
	return 0
}

func (d *Destination) contains(ts, src int) bool {
	if d.start == -1 {
		return false
	}
	key := int64(ts)*1000000 + int64(src)
	return d.dupSet[key]
}

type Router struct {
	memLim int
	queue  [][3]int
	head   int
	size   int
	dests  map[int]*Destination
}

func Constructor(memoryLimit int) Router {
	return Router{
		memLim: memoryLimit,
		queue:  make([][3]int, memoryLimit),
		head:   0,
		size:   0,
		dests:  make(map[int]*Destination),
	}
}

func (r *Router) AddPacket(source int, destination int, timestamp int) bool {
	var dest *Destination
	if d, exists := r.dests[destination]; exists {
		dest = d
		if dest.contains(timestamp, source) {
			return false
		}
	} else {
		dest = newDestination()
		r.dests[destination] = dest
	}

	if r.size == r.memLim {
		firstPkt := r.queue[r.head]
		r.dests[firstPkt[1]].removePacket()
		r.head = (r.head + 1) % r.memLim
		r.size--
	}

	newIdx := (r.head + r.size) % r.memLim
	r.queue[newIdx] = [3]int{source, destination, timestamp}
	dest.addPacket(timestamp, source)
	r.size++
	return true
}

func (r *Router) ForwardPacket() []int {
	if r.size == 0 {
		return []int{}
	}

	firstPkt := r.queue[r.head]
	r.dests[firstPkt[1]].removePacket()
	r.head = (r.head + 1) % r.memLim
	r.size--

	return []int{firstPkt[0], firstPkt[1], firstPkt[2]}
}

func (r *Router) GetCount(destination int, startTime int, endTime int) int {
	if dest, exists := r.dests[destination]; exists {
		return dest.getCount(startTime, endTime)
	}
	return 0
}

/**
 * Your Router object will be instantiated and called as such:
 * obj := Constructor(memoryLimit);
 * param_1 := obj.AddPacket(source,destination,timestamp);
 * param_2 := obj.ForwardPacket();
 * param_3 := obj.GetCount(destination,startTime,endTime);
 */
