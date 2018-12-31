package Singleton

import "testing"

func TestGetInstance(t *testing.T) {
	//测试GetInstance函数能否正确返回对象
	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("excepted pointer to Singleton after calling GetInstance(), not nil")
	}
	exceptedCounter := counter1

	//测试单例引用计数是否正确
	correntCount := counter1.AddOne()
	if correntCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", correntCount)

	}

	//测试两次调用获取到的实例是否一致
	counter2 := GetInstance()
	if counter2 != exceptedCounter {
		t.Error("Excepted same instance in counter2 but it got a different instance")
	}
}
