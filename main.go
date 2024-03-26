package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
)

// red makes text red
func red(s string) string {
	return "\033[31m" + s + "\033[0m"
}

func main() {
	// Initialize the enforcer
	e, err := casbin.NewEnforcer("./model.conf", "./policy.csv")
	if err != nil {
		panic(red(err.Error()))
	}

	// Retrieve all the grouping policies named "g"
	g := e.GetNamedGroupingPolicy("g")

	// Iterate over each grouping policy in the retrieved list
	for _, gp := range g {
		// Check if the grouping policy has exactly 5 parameters
		// Due to the specific structure expected by AddNamedDomainLinkConditionFunc
		if len(gp) != 5 {
			panic("g parameters' num isn't 5")
		}

		// Adds a domain link condition function for the "g" type policies
		// This function is used to match or evaluate conditions based on time constraints
		e.AddNamedDomainLinkConditionFunc("g", gp[0], gp[1], gp[2], util.TimeMatchFunc)
	}

	// Test cases
	testCases := []struct {
		sub  string
		obj  string
		act  string
		want bool
	}{
		{"alice", "data1", "read", true},
		{"alice", "data1", "write", true},
		{"alice", "data2", "read", false},
		{"alice", "data2", "write", false},
		{"bob", "data1", "read", false},
		{"bob", "data1", "write", false},
		{"bob", "data2", "read", true},
		{"bob", "data2", "write", false},
		{"carol", "data1", "read", false},
		{"carol", "data1", "write", false},
		{"carol", "data2", "read", false},
		{"carol", "data2", "write", false},
	}

	for _, tc := range testCases {
		res, err := e.Enforce(tc.sub, tc.obj, tc.act)
		if err != nil {
			fmt.Println(red("Error: "), red(err.Error()))
			continue
		}
		if res != tc.want {
			fmt.Println(red(fmt.Sprintf("Test failed: Enforce(%s, %s, %s): %t, want: %t", tc.sub, tc.obj, tc.act, res, tc.want)))
		} else {
			fmt.Printf("Enforce(%s, %s, %s): %t, want: %t\n", tc.sub, tc.obj, tc.act, res, tc.want)
		}
	}
}
