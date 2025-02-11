package linkedList

import (
	"testing"

	"github.com/stretchr/testify/assert" // Usando testify para assertions
)

type MyStruct struct {
	Name string
	Age  int
}

func equalsMyStruct(a, b MyStruct) bool {
	return a.Name == b.Name && a.Age == b.Age
}

func TestLinkedList(t *testing.T) {
	// Test Case 1: Append and Search with integers
	listInt := LinkedList[int]{}
	listInt.Append(1)
	listInt.Append(2)
	listInt.Append(3)

	foundInt, errInt := listInt.Search(2, func(a, b int) bool { return a == b })
	assert.NoError(t, errInt, "Search for existing int should not return an error")
	assert.Equal(t, 2, foundInt.Value, "Found int value should be correct")

	notFoundInt, errNotFoundInt := listInt.Search(4, func(a, b int) bool { return a == b })
	assert.Error(t, errNotFoundInt, ErrNodeNotFound, "Search for non-existing int should return ErrNodeNotFound")
	assert.Nil(t, notFoundInt, "Search for non-existing int should return nil node")

	// Test Case 2: Append and Search with custom struct
	listStruct := LinkedList[MyStruct]{}
	struct1 := MyStruct{Name: "Alice", Age: 30}
	struct2 := MyStruct{Name: "Bob", Age: 25}
	struct3 := MyStruct{Name: "Charlie", Age: 35}

	listStruct.Append(struct1)
	listStruct.Append(struct2)
	listStruct.Append(struct3)

	foundStruct, errStruct := listStruct.Search(struct2, equalsMyStruct)
	assert.NoError(t, errStruct, "Search for existing struct should not return an error")
	assert.Equal(t, struct2, foundStruct.Value, "Found struct value should be correct")

	notFoundStruct, errNotFoundStruct := listStruct.Search(MyStruct{Name: "David", Age: 40}, equalsMyStruct)
	assert.Error(t, errNotFoundStruct, ErrNodeNotFound, "Search for non-existing struct should return ErrNodeNotFound")
	assert.Nil(t, notFoundStruct, "Search for non-existing struct should return nil node")

	// Test Case 3: RemoveElement - Head
	listRemoveHead := LinkedList[int]{}
	listRemoveHead.Append(1)
	listRemoveHead.Append(2)

	errRemoveHead := listRemoveHead.RemoveElement(1, func(a, b int) bool { return a == b })
	assert.NoError(t, errRemoveHead, "Removing head should not return an error")
	assert.Equal(t, 2, listRemoveHead.Head.Value, "Head should be updated after removing the original head")

	// Test Case 4: RemoveElement - Middle
	listRemoveMiddle := LinkedList[int]{}
	listRemoveMiddle.Append(1)
	listRemoveMiddle.Append(2)
	listRemoveMiddle.Append(3)

	errRemoveMiddle := listRemoveMiddle.RemoveElement(2, func(a, b int) bool { return a == b })
	assert.NoError(t, errRemoveMiddle, "Removing middle element should not return an error")

	foundMiddle, _ := listRemoveMiddle.Search(2, func(a, b int) bool { return a == b })
	assert.Nil(t, foundMiddle, "Middle element should be removed")

	// Test Case 5: RemoveElement - Last
	listRemoveLast := LinkedList[int]{}
	listRemoveLast.Append(1)
	listRemoveLast.Append(2)
	listRemoveLast.Append(3)

	errRemoveLast := listRemoveLast.RemoveElement(3, func(a, b int) bool { return a == b })
	assert.NoError(t, errRemoveLast, "Removing last element should not return an error")

	foundLast, _ := listRemoveLast.Search(3, func(a, b int) bool { return a == b })
	assert.Nil(t, foundLast, "Last element should be removed")

	// Test Case 6: RemoveElement - Not Found
	listRemoveNotFound := LinkedList[int]{}
	listRemoveNotFound.Append(1)
	listRemoveNotFound.Append(2)

	errRemoveNotFound := listRemoveNotFound.RemoveElement(3, func(a, b int) bool { return a == b })
	assert.NoError(t, errRemoveNotFound, "Removing non-existent element should not return an error") // Should not return an error, just do nothing

	// Test Case 7: Empty List
	emptyList := LinkedList[int]{}
	_, errEmpty := emptyList.Search(1, func(a, b int) bool { return a == b })
	assert.Error(t, errEmpty, ErrNodeNotFound, "Search in empty list should return ErrNodeNotFound")

	errRemoveEmpty := emptyList.RemoveElement(1, func(a, b int) bool { return a == b })
	assert.NoError(t, errRemoveEmpty, "Removing from empty list should not return an error")

}
