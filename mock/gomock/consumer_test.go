package gomock

import (
    "github.com/golang/mock/gomock"
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestConsumerDouble(t *testing.T){
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    faker := NewMockFaker(ctrl)
    faker.EXPECT().Get("a").Return(1)

    consumer := Consumer{faker: faker}
    actual := consumer.Double("a")
    assert.Equal(t, 2, actual)
}

