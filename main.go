package main

import (
	"fmt"
)

type BikeMetadata struct {
	Color  string
	Weight string
}

type BikeCategory struct {
	Name          string
	Description   string
	Metadata      BikeMetadata // Add metadata here
	Subcategories []BikeCategory
}

func main() {
	// Define the bike taxonomy with metadata
	bikeTaxonomy := []BikeCategory{
		{
			Name:        "Mountain Bikes",
			Description: "Off-road bikes for tackling rough terrains.",
			Metadata:    BikeMetadata{Color: "Various", Weight: "25-35 lbs"},
			Subcategories: []BikeCategory{
				{
					Name:        "Hardtail",
					Description: "Mountain bikes with front suspension only.",
					Metadata:    BikeMetadata{Color: "Various", Weight: "27-32 lbs"},
					Subcategories: []BikeCategory{
						{
							Name:        "Cross-Country",
							Description: "Hardtail bikes optimized for speed and endurance.",
							Metadata:    BikeMetadata{Color: "Red, Black, White", Weight: "28 lbs"},
						},
						{
							Name:        "Trail",
							Description: "Hardtail bikes suitable for a variety of trails.",
							Metadata:    BikeMetadata{Color: "Green, Blue, Black", Weight: "29 lbs"},
						},
					},
				},
				{
					Name:        "Full Suspension",
					Description: "Mountain bikes with both front and rear suspension.",
					Metadata:    BikeMetadata{Color: "Various", Weight: "30-40 lbs"},
					Subcategories: []BikeCategory{
						{
							Name:        "All-Mountain",
							Description: "Full suspension bikes designed for varied terrain.",
							Metadata:    BikeMetadata{Color: "Blue, Black, Orange", Weight: "35 lbs"},
						},
						{
							Name:        "Downhill",
							Description: "Full suspension bikes for fast descents.",
							Metadata:    BikeMetadata{Color: "Black, Red, Yellow", Weight: "40 lbs"},
						},
					},
				},
			},
		},
		{
			Name:        "Road Bikes",
			Description: "Bikes designed for speed on paved roads.",
			Metadata:    BikeMetadata{Color: "Various", Weight: "15-22 lbs"},
			Subcategories: []BikeCategory{
				{
					Name:        "Race",
					Description: "Lightweight road bikes optimized for racing.",
					Metadata:    BikeMetadata{Color: "Red, White, Black", Weight: "15-18 lbs"},
				},
				{
					Name:        "Endurance",
					Description: "Road bikes designed for comfort on long rides.",
					Metadata:    BikeMetadata{Color: "Blue, Gray, Black", Weight: "18-22 lbs"},
				},
			},
		},
	}

	// Print the bike taxonomy with metadata
	printBikeCategories(bikeTaxonomy, 0)
}

func printBikeCategories(categories []BikeCategory, indent int) {
	for _, category := range categories {
		fmt.Printf("%s%s: %s (Color: %s, Weight: %s)\n", getIndent(indent), category.Name, category.Description, category.Metadata.Color, category.Metadata.Weight)
		printBikeCategories(category.Subcategories, indent+1)
	}
}

func getIndent(indent int) string {
	return fmt.Sprintf("%"+fmt.Sprintf("%ds", indent*4), "")
}
