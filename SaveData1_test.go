package MoewDB

import (
	"os"
	"testing"
)

func TestSaveData1(t *testing.T) {

	testData := []byte("你好呀！")
	testFilePath := "testFile.txt"

	// 调用SaveData1函数保存测试数据
	err := SaveData1(testFilePath, testData)
	if err != nil {
		t.Fatalf("SaveData1 failed with error: %v", err)
	}

	// 读取保存的文件内容
	savedData, err := os.ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("Failed to read the saved file: %v", err)
	}
	if string(savedData) != string(testData) {
		t.Fatalf("Expected %s but got %s", testData, savedData)
	}

	// 清理测试文件
	os.Remove(testFilePath)
}
