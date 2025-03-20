package main

import (
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"flag"
	"fmt"
)

func main() {
	var (
		path string
	)

	flag.StringVar(&path, "path", "", "Path to CSV or JSON for student registration.")
	flag.Parse()

	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		panic(err)
	}
	var courses []*model.Course
	if err := deserializer.Deserialize(&courses); err != nil {
		panic(err)
	}

	for _, v := range courses {
		fmt.Printf("Course Id: %s, Course Name: %s, Course Optional: %t\n", v.CourseId, v.Name, v.Optional)
	}
	// var classes []*model.Class
	// if err := deserializer.Deserialize(&classes); err != nil {
	// 	panic(err)
	// }
	// for _, v := range classes {
	// 	fmt.Printf("Class Id: %s, Course Id: %s, Class Section: %d, Class Schedule: %s\n", v.ClassId, v.CourseId, v.Section, v.Schedule)
	// }
}
