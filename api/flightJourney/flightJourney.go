package flightJourney

import (
	"errors"
	"sort"
)

type flightJourneySvc struct{}

func NewFlightJourneyService() *flightJourneySvc {
	return &flightJourneySvc{}
}

// GetFlightStartingAndEndingAirportCode searches for possible flight path and return starting and ending airport code.
func (svc *flightJourneySvc) GetFlightStartingAndEndingAirportCode(tickets [][]string) ([]string, error) {

	trace, err := Search_Best_Time(tickets)

	if err != nil {
		return nil, err

	}

	return trace, nil
}

// Search searches for possible paths from a given source
/***
In a connected graph with an Eulerian path, there are only two cases:
1) There is one node with one more outdegree than indegree and another node with one more indegree than out degree. In this case, one have to pick the node with more outdegree than indegree as the starting node, an Eulerian path ends at the node with more bigger indegree eventually.
2) All nodes with indegree == outdegree. In this case, one can pick the starting node randomly (here its is lexicographic first airport), an Eulerian path ends at the starting node eventually.
***/
func Search_Best_Time(tickets [][]string) ([]string, error) {
	cnt := make(map[string]int)
	for _, p := range tickets {
		cnt[p[0]]++
		cnt[p[1]]--
	}

	flag := false

	start := tickets[0][0]
	cntInOutdegreepositive := 0
	cntInOutdegreenegative := 0

	for _, inOutDegree := range cnt {

		if inOutDegree == 0 {
			flag = true
		} else {
			flag = false
		}
	}

	if flag {
		airports := airportsinlexicographicOrder(tickets)
		start = airports[0]
	}

	for _, inOutDegree := range cnt {
		if inOutDegree > 1 {
			return nil, errors.New("in a flight path graph inoutdegree greater than 1 is not possible")
		}

		if inOutDegree < -1 {
			return nil, errors.New("in a flight path graph inoutdegree less than -1 is not possible")
		}

		if inOutDegree == 1 {
			cntInOutdegreepositive++

		}

		if inOutDegree == -1 {
			cntInOutdegreenegative++

		}

		if cntInOutdegreepositive > 1 {
			return nil, errors.New("in a flight path graph two nodes can't have inoutdegree equal to 1")
		}

		if cntInOutdegreenegative > 1 {
			return nil, errors.New("in a flight path graph two nodes can't have inoutdegree equal to -1")

		}

	}
	for vertex, inOutDegree := range cnt {

		if inOutDegree > 0 {
			start = vertex
			break
		}
	}

	m := make(map[string][]string, len(tickets)+1)
	var routes []string

	for _, t := range tickets {
		m[t[0]] = append(m[t[0]], t[1])
	}

	for k := range m {
		sort.Strings(m[k])
	}

	DFS(start, m, &routes)

	// reverse ans array

	i, j := 0, len(routes)-1
	for i < j {
		routes[i], routes[j] = routes[j], routes[i]
		i++
		j--
	}

	return []string{routes[0], routes[len(routes)-1]}, nil
}

func DFS(start string, m map[string][]string, routes *[]string) {

	for len(m[start]) > 0 {
		cur := m[start][0]
		m[start] = m[start][1:]
		DFS(cur, m, routes)
	}

	*routes = append(*routes, start)

}

func Search_Best_Memory(tickets [][]string) ([]string, error) {
	cnt := make(map[string]int)
	for _, p := range tickets {
		cnt[p[0]]++
		cnt[p[1]]--
	}

	flag := false

	start := tickets[0][0]
	cntInOutdegreepositive := 0
	cntInOutdegreenegative := 0

	for _, inOutDegree := range cnt {

		if inOutDegree == 0 {
			flag = true
		} else {
			flag = false
		}
	}

	if flag {
		airports := airportsinlexicographicOrder(tickets)
		start = airports[0]
	}

	for _, inOutDegree := range cnt {
		if inOutDegree > 1 {
			return nil, errors.New("in a flight path graph inoutdegree greater than 1 is not possible")
		}

		if inOutDegree < -1 {
			return nil, errors.New("in a flight path graph inoutdegree less than -1 is not possible")
		}

		if inOutDegree == 1 {
			cntInOutdegreepositive++

		}

		if inOutDegree == -1 {
			cntInOutdegreenegative++

		}

		if cntInOutdegreepositive > 1 {
			return nil, errors.New("in a flight path graph two nodes can't have inoutdegree equal to 1")
		}

		if cntInOutdegreenegative > 1 {
			return nil, errors.New("in a flight path graph two nodes can't have inoutdegree equal to -1")

		}

	}
	for vertex, inOutDegree := range cnt {

		if inOutDegree > 0 {
			start = vertex
			break
		}
	}

	sort.Slice(tickets, func(i, j int) bool {
		if tickets[i][0] != tickets[j][0] {
			return tickets[i][0] < tickets[j][0]
		} else {
			return tickets[i][1] < tickets[j][1]
		}
	})
	edges := make(map[string][]int)
	for i := 0; i < len(tickets); {
		j := i + 1
		for j < len(tickets) && tickets[j][0] == tickets[i][0] {
			j++
		}
		edges[tickets[i][0]] = []int{i, j}
		i = j
	}

	routes := make([]string, len(tickets)+1)
	head := len(tickets)

	visit(start, routes, &head, tickets, edges)

	return []string{routes[0], routes[len(routes)-1]}, nil
}

func visit(airport string, results []string, head *int, tickets [][]string, edges map[string][]int) {

	if e, ok := edges[airport]; ok {
		for e[0] < e[1] {
			i := e[0]
			e[0]++
			visit(tickets[i][1], results, head, tickets, edges)
		}
	}
	results[*head] = airport
	*head--
}

func airportsinlexicographicOrder(tickets [][]string) []string {
	s := map[string]bool{}
	for x, b := range tickets {
		for y, _ := range b {
			s[tickets[x][y]] = true
		}
	}
	keys := make([]string, 0, len(s))

	for k := range s {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return keys

}
