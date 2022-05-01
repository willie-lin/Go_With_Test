package maps

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {

	return d[word], nil
}

//
//func Search(dictionary map[string]string, word string) string {
//	return dictionary[word]
//}
