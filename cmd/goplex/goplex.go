package main

import (
	"log"

	. "github.com/tagnard/goplex"
)

func main() {

	plex := NewPlexClient("rn1-vm-pms01.raeven.net", 32400, "mBCsqJx6v8DVk8fptdVA", true, true)

	/*
	err, session_metrics := plex.SessionMetrics()
	if err != nil {
		log.Println(err)
	}

	log.Println("Session Metrics")
	log.Printf("Total Sessions: %v", session_metrics.Total)
	log.Printf("Active Sessions: %v", session_metrics.Active)
	log.Printf("Inactive Sessions: %v", session_metrics.Inactive)

	*/

	err, section_metrics := plex.SectionMetrics(2)
	if err != nil {
		log.Println(err)
	}

	log.Println("")
	log.Println("Section Metrics")
	log.Printf("Total items: %v", section_metrics.Total)
	log.Printf("Watched items: %v", section_metrics.Watched)
	log.Printf("Unwatched items: %v", section_metrics.Unwatched)
	log.Printf("Watched percentage: %v", section_metrics.WatchedPercent)

	/*

	err, sections := plex.LibrarySections()
	if err != nil {
		log.Println(err)
	}

	for _, section := range sections.Directory {
		err, items := plex.LibrarySection(section.Location.Id, filters.All)
		if err != nil {
			panic(err)
		}

		err, unwatched := plex.LibrarySection(section.Location.Id, filters.Unwatched)
		if err != nil {
			panic(err)
		}

		log.Printf("%v has %v items, %v is unwatched", section.Title, items.Size, unwatched.Size)
	}
	*/
}