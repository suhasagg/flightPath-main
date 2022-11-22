package flightJourney

import (
	"fmt"
	"testing"
)

func Test_flightJourneySvc_GetFlightStartingAndEndingAirportCode_Positive(t *testing.T) {

	tests := []struct {
		name  string
		args  [][]string
		want1 string
		want2 string
	}{
		{name: "one flight",
			args:  [][]string{{"SFO", "EWR"}},
			want1: "SFO",
			want2: "EWR",
		},

		{name: "two flights",
			args:  [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}},
			want1: "SFO",
			want2: "EWR",
		},

		{name: "multiple flights",
			args:  [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
			want1: "SFO",
			want2: "EWR",
		},

		{name: "multiple flights",
			args:  [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}},
			want1: "JFK",
			want2: "SJC",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &flightJourneySvc{}
			got := svc.GetFlightStartingAndEndingAirportCode(tt.args)
			/*
				if err != nil {
					t.Errorf("GetFlightStartingAndEndingAirportCode() returned unexpected err = %s", err)
				}
			*/
			if got[0] != tt.want1 {
				t.Errorf("GetFlightStartingAndEndingAirportCode() got = %v, want1 %v", got[0], tt.want1)
			}
			if got[1] != tt.want2 {
				t.Errorf("GetFlightStartingAndEndingAirportCode() got1 = %v, want1 %v", got[1], tt.want2)
			}
		})
	}
}

func Test_flightJourneySvc_GetFlightStartingAndEndingAirportCode_Negative(t *testing.T) {

	tests := []struct {
		name  string
		args  [][]string
		want1 string
		want2 string
	}{

		{name: "two flights",
			args: [][]string{{"SFO", "EWR"}, {"SFO", "ATL"}},
		},

		{name: "multiple flights",
			args: [][]string{{"IND", "EWR"}, {"SFO", "IND"}, {"ATL", "GSO"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &flightJourneySvc{}
			got := svc.GetFlightStartingAndEndingAirportCode(tt.args)
			/*
				if err == nil {
					t.Errorf("GetFlightStartingAndEndingAirportCode() must return error")
				}

			*/
			fmt.Println(got)
		})
	}
}
