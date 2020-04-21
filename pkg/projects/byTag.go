package projects

func FilterProjectsByTag(projects []Project, tags []string) []Project {
	if len(tags) == 0 {
		return projects
	}

	output := []Project{}
	tagsMap := map[string]bool{}

	for _, tag := range tags {
		tagsMap[tag] = true
	}

	for _, project := range projects {
		match := false

		for _, tag := range project.Tags {
			if _, ok := tagsMap[tag]; ok {
				match = true
			}
		}

		if match {
			output = append(output, project)
		}
	}

	return output
}
