package example

// 司机
type DriverService interface {
	Info() string
}

// 完成率
type FinishRaterService interface {
	Query() string
	Save() error
}

type FinishRate struct {
	driver DriverService
}

func (p FinishRate) Query() string {
	return "finish_rate || " + p.driver.Info()
}

func (p FinishRate) Save() error {
	return nil
}