package domain

type customEntity struct {
}

func (c customEntity) Create(model CustomModel) error {
	// todo 持久化逻辑
	println("store custom model to infra")
	return nil
}

func (c customEntity) Update(model CustomModel) error {
	// todo 持久化逻辑
	println("update custom model to infra")
	return nil
}
