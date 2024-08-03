package vrcosc

type AxesMoveDirection string

const (
	AxisVertical   AxesMoveDirection = "Vertical"
	AxisHorizontal AxesMoveDirection = "Horizontal"
)

// func (c *Client) AxesMove(direction AxesDirection, b bool) error {

type AxesLookDirection string

const (
	AxisLookLeft  AxesLookDirection = "LookHorizontal"
	AxisLookRight AxesLookDirection = "LookVertical"
)

// func (c *Client) AxesLook(direction AxesDirection, b bool) error {
