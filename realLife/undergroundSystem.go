package reallife

// https://leetcode.com/problems/design-underground-system/

type TravelPass struct {
	stationName string
	time        int
}

type UndergroundSystem struct {
	TravelTime     map[string][]int
	CustomerRecord map[int]TravelPass
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		TravelTime:     make(map[string][]int),
		CustomerRecord: make(map[int]TravelPass),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	userPass := TravelPass{
		stationName: stationName,
		time:        t,
	}

	this.CustomerRecord[id] = userPass
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	userBoarding := this.CustomerRecord[id]
	key := userBoarding.stationName + " " + stationName
	time := t - userBoarding.time

	this.TravelTime[key] = append(this.TravelTime[key], time)
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	var total int
	key := startStation + " " + endStation
	station := this.TravelTime[key]
	lenVal := float64(len(station))

	for _, num := range station {
		total += num
	}

	return float64(float64(total) / lenVal)
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
