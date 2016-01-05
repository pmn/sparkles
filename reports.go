package main

import (
	"time"
)

// SparkleStats datatype
type SparkleStats struct {
	TopGiven     []string  `json:"top_given,omitempty"`
	TopGivers    []string  `json:"top_givers,omitempty"`
	FirstSparkle time.Time `json:"first_sparkle"`
	SimilarUsers []string  `json:"similar_users,omitempty"`
	Categories   []string  `json:"categories"`
}

type SparkleGraph struct {
	Edges []SparkleGraphEdge `json:"edges"`
}

type SparkleGraphEdge struct {
	Sparkler string `json:"sparkler"`
	Sparklee string `json:"sparklee"`
	Weight   int    `json:"weight"`
}

func (s *SparkleFileDatabase) Graph() SparkleGraph {
	var sg SparkleGraph
	for _, v := range s.Sparkles {
		// See if the edge already exists. If it does, increment the score.
		// Otherwise, insert it.
		exists := false
		for k1, v1 := range sg.Edges {
			if v.Sparkler == v1.Sparkler && v.Sparklee == v1.Sparklee {
				sg.Edges[k1].Weight++
				exists = true
			}
		}

		if !exists {
			// Construct an edge
			var e SparkleGraphEdge
			e.Sparkler = v.Sparkler
			e.Sparklee = v.Sparklee
			e.Weight = 1
			sg.Edges = append(sg.Edges, e)
		}
	}
	return sg
}

func StatsForUser(user string) SparkleStats {
	var ss SparkleStats
	return ss
}
