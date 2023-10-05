package tagcloud

import (
	"sort"
)

// TagCloud aggregates statistics about used tags
type TagCloud struct {
	kv map[string]int
}

// TagStat represents statistics regarding single tag
type TagStat struct {
	Tag             string
	OccurrenceCount int
}

// New should create a valid TagCloud instance
// TODO: You decide whether this function should return a pointer or a value
func New() TagCloud {
	// TODO: Implement this
	//return TagCloud{}
	return TagCloud{make(map[string]int)}
}

// AddTag should add a tag to the cloud if it wasn't present and increase tag occurrence count
// thread-safety is not needed
// TODO: You decide whether receiver should be a pointer or a value
func (tagCloud *TagCloud) AddTag(tag string) {
	// TODO: Implement this
	tagCloud.kv[tag] += 1
}

// TopN should return top N most frequent tags ordered in descending order by occurrence count
// if there are multiple tags with the same occurrence count then the order is defined by implementation
// if n is greater that TagCloud size then all elements should be returned
// thread-safety is not needed
// there are no restrictions on time complexity
// TODO: You decide whether receiver should be a pointer or a value
func (tagCloud TagCloud) TopN(n int) []TagStat {
	// TODO: Implement this
	var kv = tagCloud.kv
	keys := make([]string, 0, len(kv))
	for key := range kv {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return kv[keys[i]] > kv[keys[j]]
	})
	tagStat := make([]TagStat, 0, n)
	for i, v := range keys {
		tagStat = append(tagStat, TagStat{v, kv[v]})
		if i == n-1 {
			break
		}
	}
	return tagStat
}
