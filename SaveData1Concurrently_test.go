package MoewDB

import (
	"io/ioutil"
	"sync"
	"testing"
)

func TestSaveData1Concurrently(t *testing.T) {
	testData := []byte("Hello, Go!")

	var wg sync.WaitGroup

	const numGoroutines = 10

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()

			testFilePath := "testFile.txt"

			err := SaveData1(testFilePath, testData)
			if err != nil {
				t.Errorf("SaveData1 failed with error on goroutine %d: %v", idx, err)
				return
			}

			// Read the saved file content
			savedData, err := ioutil.ReadFile(testFilePath)
			if err != nil {
				t.Errorf("Failed to read the saved file on goroutine %d: %v", idx, err)
				return
			}

			// Compare the saved file content with the expected test data
			if string(savedData) != string(testData) {
				t.Errorf("Expected %s but got %s on goroutine %d", testData, savedData, idx)
				return
			}

		}(i)
	}
	wg.Wait()

}
