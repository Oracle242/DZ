package autosolon

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		if t == Inch && u.T == CM {
			value = value / 2.54
		} else if t == CM && u.T == Inch {
			value = value * 2.54
		}
	}

	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type bmw struct {
	brand, model          string
	dimensions            Dimensions
	maxSpeed, enginePower int
}
type InchesDimensions struct {
	length Unit
	width  Unit
	height Unit
}

func (d InchesDimensions) Length() Unit {
	return d.length
}

func (d InchesDimensions) Width() Unit {
	return d.width
}

func (d InchesDimensions) Height() Unit {
	return d.height
}

type CMDimensions struct {
	length Unit
	width  Unit
	height Unit
}

func (d CMDimensions) Length() Unit {
	return d.length
}

func (d CMDimensions) Width() Unit {
	return d.width
}

func (d CMDimensions) Height() Unit {
	return d.height
}

func (b bmw) Brand() string {
	return "BMW"
}
func (b bmw) Model() string {
	return b.model
}
func (b bmw) Dimensions() Dimensions {
	return b.dimensions
}
func (b bmw) MaxSpeed() int {
	return b.maxSpeed
}
func (b bmw) EnginePower() int {
	return b.enginePower
}
func NewBMWEuropean(model string, dimensions Dimensions, maxSpeed int, enginePower int) *bmw {
	return &bmw{model: model, dimensions: dimensions, maxSpeed: maxSpeed, enginePower: enginePower}
}

type mercedes struct {
	model                 string
	dimensions            Dimensions
	maxSpeed, enginePower int
}

func (m mercedes) Brand() string {
	return "Mercedes"
}
func (m mercedes) Model() string {
	return m.model
}
func (m mercedes) Dimensions() Dimensions {
	return m.dimensions
}
func (m mercedes) MaxSpeed() int {
	return m.maxSpeed
}
func (m mercedes) EnginePower() int {
	return m.enginePower
}
func NewMercedesEuropean(model string, dimensions Dimensions, maxSpeed int, enginePower int) *mercedes {
	return &mercedes{model: model, dimensions: dimensions, maxSpeed: maxSpeed, enginePower: enginePower}
}

type dodge struct {
	model                 string
	dimensions            Dimensions
	maxSpeed, enginePower int
}

func (d dodge) Brand() string {
	return "Dodge"
}
func (d dodge) Model() string {
	return d.model
}
func (d dodge) Dimensions() Dimensions {
	return d.dimensions
}

func (d dodge) MaxSpeed() int {
	return d.maxSpeed
}
func (d dodge) EnginePower() int {
	return d.enginePower
}
func NewDodgeEuropean(model string, dimensions Dimensions, maxSpeed int, enginePower int) *dodge {
	return &dodge{model: model, dimensions: dimensions, maxSpeed: maxSpeed, enginePower: enginePower}
}
func NewInchesDimensions(length, width, height float64) InchesDimensions {
	return InchesDimensions{
		length: Unit{Value: length, T: Inch},
		width:  Unit{Value: width, T: Inch},
		height: Unit{Value: height, T: Inch},
	}
}
func NewCMDimensions(length, width, height float64) CMDimensions {
	return CMDimensions{
		length: Unit{Value: length, T: CM},
		width:  Unit{Value: width, T: CM},
		height: Unit{Value: height, T: CM},
	}
}
