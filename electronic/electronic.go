package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}
type StationPhone interface {
	ButtonsCount() int
}
type Smartphone interface {
	OS() string
}

type applePhone struct {
	model string
}

func (a applePhone) Brand() string {
	return "Apple"
}
func (a applePhone) Model() string {
	return a.model
}
func (a applePhone) Type() string {
	return "smartphone"
}
func (a applePhone) OS() string {
	return "IOS"
}

func NewApplePhone(model string) *applePhone {
	return &applePhone{model: model}
}

type androidPhone struct {
	brand, model string
}

func (a androidPhone) Brand() string {
	return a.brand
}
func (a androidPhone) Model() string {
	return a.model
}
func (a androidPhone) Type() string {
	return "smartphone"
}
func (a androidPhone) OS() string {
	return "Android"
}

func NewAndroidPhone(brand, model string) *androidPhone {
	return &androidPhone{brand: brand, model: model}
}

type radioPhone struct {
	model        string
	buttonsCount int
}

func (a radioPhone) Brand() string {
	return "Noname"
}
func (a radioPhone) Model() string {
	return a.model
}
func (a radioPhone) Type() string {
	return "station"
}
func (a radioPhone) ButtonsCount() int {
	return a.buttonsCount
}

func NewRadioPhone(model string, buttonsCount int) *radioPhone {
	return &radioPhone{model: model, buttonsCount: buttonsCount}
}
