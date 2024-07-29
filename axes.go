package vrcinput

type AxesMoveDirection string

const (
	AxisVertical   AxesMoveDirection = "Vertical"
	AxisHorizontal AxesMoveDirection = "Horizontal"
)

// func (c *localOscClient) AxesMove(direction AxesDirection, b bool) error {

type AxesLookDirection string

const (
	AxisLookLeft  AxesLookDirection = "LookHorizontal"
	AxisLookRight AxesLookDirection = "LookVertical"
)

// func (c *localOscClient) AxesLook(direction AxesDirection, b bool) error {
