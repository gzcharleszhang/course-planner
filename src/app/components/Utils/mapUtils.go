package Utils

func InMap(val interface{}, mp map[interface{}]interface{}) bool {
	_, ok := mp[val]
	return ok
}
