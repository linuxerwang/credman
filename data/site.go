package data

import "encoding/json"

type Site map[string]string // fields

const SpecialFieldTags = "$$_TAGS"

func IsSpecialField(name string) bool {
	return name == SpecialFieldTags
}

func (s Site) GetTags() (tags []string) {
	tagsList, ok := s[SpecialFieldTags]
	tags = make([]string, 0)
	if !ok {
		return
	}

	json.Unmarshal([]byte(tagsList), &tags)

	return
}

func (s Site) SetTags(tags []string) {
	for _, siteTag := range s.GetTags() {
		found := false
		for _, tag := range tags {
			if tag == siteTag {
				found = true
			}
		}
		if !found {
			tags = append(tags, siteTag)
		}
	}

	tagsBytes, err := json.Marshal(tags)
	if err != nil {
		return
	}
	s[SpecialFieldTags] = string(tagsBytes)
}

func (s Site) HasTag(tag string) bool {
	siteTags := s.GetTags()
	for _, siteTag := range siteTags {
		if tag == siteTag {
			return true
		}
	}
	return false
}

func (s Site) HasTags(tags []string) (found bool, foundTags []string) {
	foundTags = make([]string, 0)
	siteTags := s.GetTags()
	for _, tag := range siteTags {
		for _, askedTag := range tags {
			if tag == askedTag {
				foundTags = append(foundTags, tag)
			}
		}
	}
	found = len(foundTags) > 0
	return
}
func (s Site) RemoveTags(tags []string) {
	//
}
