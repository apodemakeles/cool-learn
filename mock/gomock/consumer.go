package gomock

type Consumer struct{
    faker Faker
}

func (c Consumer) Double(key string) int{
   return 2 * c.faker.Get(key)
}

