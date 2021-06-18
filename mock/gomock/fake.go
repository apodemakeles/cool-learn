package gomock

//go:generate mockgen -source=fake.go -destination=fake_mock.go -package=gomock

type Faker interface {
    Get(key string) int
    DoSth()
    Print(x interface{})
}
