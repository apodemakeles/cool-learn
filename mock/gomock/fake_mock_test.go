package gomock

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// 成功的例子
func TestCallSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)
	faker.EXPECT().Get("a").Return(1)

	value := faker.Get("a")

	assert.Equal(t, 1, value)
}

// fail! mock了但没call
func TestNoCall(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)
	faker.EXPECT().Get("a").Return(1)
}

// fail! call但没mock
func TestNoMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)

	assert.Equal(t, 1, faker.Get("a"))
}

// fail! 参数不匹配
func TestArgMissMatching(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)
	faker.EXPECT().Get("a").Return(1)

	assert.Equal(t, 2, faker.Get("b"))
}

// fail！次数不匹配
func TestCallTimesMissMatching(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)
	faker.EXPECT().Get("a").Return(1) // .Times(2)

	assert.Equal(t, 1, faker.Get("a"))
	assert.Equal(t, 1, faker.Get("a"))
}

// mock的次数与调用次数一致
func TestCallTimes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)
	faker.EXPECT().Get("a").Return(1).AnyTimes()
	faker.EXPECT().DoSth().MinTimes(2).MaxTimes(3)

	faker.Get("a")
	faker.Get("a")
	faker.Get("a")
	faker.DoSth()
	faker.DoSth()
	faker.DoSth()

}

// 匹配任意参数
func TestAnyArg(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)
	faker.EXPECT().Get(gomock.Any()).Return(1).AnyTimes()

	faker.Get("a")
	faker.Get("b")
	faker.Get("")
}

// 匹配nil的参数
func TestNilMatcher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)
	faker.EXPECT().Print(nil).Do(func(_ interface{}) {
		fmt.Println("it's nil")
	})

	var s *struct{ Name string }
	assert.True(t, s == nil)
	assert.False(t, func(x interface{}) bool {
		return x == nil
	}(s))
	faker.Print(s) // 打印 "it's nil"
}

func TestAfter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)
	callDoSth := faker.EXPECT().DoSth()
	// faker.EXPECT().Print(1)
	faker.EXPECT().Get("a").Return(1).After(callDoSth)

	faker.DoSth()
	// faker.Print(1)
	faker.Get("a")
}

func TestInOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)

	gomock.InOrder(
		faker.EXPECT().Get("a").Return(1),
		faker.EXPECT().Get("b").Return(2),
	)

	faker.Get("a")
	faker.Get("b")
}

func TestDo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	faker := NewMockFaker(ctrl)
	var x interface{}
	faker.EXPECT().Print(gomock.Any()).Do(func(arg interface{}) {
		x = arg
	}).AnyTimes()

	faker.Print(1)
	assert.Equal(t, 1, x)
	faker.Print("a")
	assert.Equal(t, "a", x)
	faker.Print([]int{1, 2, 3})
	assert.Equal(t, []int{1, 2, 3}, x)

	called := false
	faker.EXPECT().DoSth().Do(func() {
		called = true
	}).AnyTimes()

	faker.DoSth()
	assert.True(t, called)
}
