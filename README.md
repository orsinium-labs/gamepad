# gamepad

Cross-platform Go library to read input from an [Xbox 360 controller](https://en.wikipedia.org/wiki/Xbox_360_controller) (or another compatible gamepad).

## Installation

```bash
go get github.com/orsinium-labs/gamepad
```

## Usage

```go
gamepad, err := gamepad.NewGamepad(0)
// ...
state, err := gamepad.State()
fmt.Println(state.A())
```
