package flightJourney

type flightJourneySvc struct{}

func NewFlightJourneyService() *flightJourneySvc {
	return &flightJourneySvc{}
}

var (
	airporttoNumericMap = make(map[string]int)
)

var (
	numerictoAirportMap = make(map[int]string)
)

// GetFlightStartingAndEndingAirportCode searches for possible flight path and return starting and ending airport code.
func (svc *flightJourneySvc) GetFlightStartingAndEndingAirportCode(tickets [][]string) []string {
	numerictickets := make([][]int, len(tickets))
	for i := range numerictickets {
		numerictickets[i] = make([]int, 2)
	}

	AirportstoNumericCode(tickets)

	for x, b := range tickets {
		for y, _ := range b {
			numerictickets[x][y] = airporttoNumericMap[tickets[x][y]]
		}
	}

	NumericCodetoAirports(numerictickets)

	h := Search(NumberofAirports(numerictickets), numerictickets)

	airporttoNumericMap = make(map[string]int)

	x, y := numerictoAirportMap[h[0]], numerictoAirportMap[h[len(h)-1]]

	numerictoAirportMap = make(map[int]string)

	return []string{x, y}
}

// Search searches for possible paths from a given source
func Search(numberofAirports int, tickets [][]int) []int {
	graph := map[int][]int{}
	inDegree := make([]int, numberofAirports)
	for i := range tickets {
		graph[tickets[i][0]] = append(graph[tickets[i][0]], tickets[i][1])
		inDegree[tickets[i][1]]++
	}
	queue := make([]int, 0)
	print(len(queue))
	for i := range inDegree {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	ret := make([]int, 0)
	for len(queue) != 0 {
		head := queue[0]
		queue = queue[1:]
		for i := range graph[head] {
			inDegree[graph[head][i]]--
			if inDegree[graph[head][i]] == 0 {
				queue = append(queue, graph[head][i])
			}
		}
		ret = append(ret, head)
	}
	if len(ret) != numberofAirports {
		return []int{}
	}
	return ret
}

func AirportstoNumericCode(tickets [][]string) {
	i := 0
	for x, b := range tickets {
		for y, _ := range b {
			if _, ok := airporttoNumericMap[tickets[x][y]]; !ok {
				airporttoNumericMap[tickets[x][y]] = i
				i++
			}
		}
	}
}

func NumericCodetoAirports(a [][]int) {
	for k, v := range airporttoNumericMap {
		numerictoAirportMap[v] = k
	}
}

func NumberofAirports(a [][]int) int {
	s := map[int]bool{}
	for x, b := range a {
		for y, _ := range b {
			s[a[x][y]] = true
		}
	}
	return len(s)

}
