package main

import (
	"context"

	"hexview/models"
	"hexview/service"
)

// App struct holds the Wails application context and service dependencies.
// It acts as a thin glue layer between the frontend bindings and the service layer.
type App struct {
	ctx       context.Context
	converter *service.Converter
}

// NewApp creates a new App application struct with initialized services.
func NewApp() *App {
	return &App{
		converter: service.NewConverter(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods.
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// ConvertHex performs all possible conversions on hex input.
// This method is exported to the frontend via Wails bindings.
func (a *App) ConvertHex(hexInput string) (*models.ConversionResult, error) {
	return a.converter.ConvertHex(hexInput)
}

// ConvertInt performs conversions from integer input to hex and binary.
// intType specifies the integer type: int8, int16, int32, int64, uint8, uint16, uint32, uint64.
// This method is exported to the frontend via Wails bindings.
func (a *App) ConvertInt(intInput string, intType string) (*models.ConversionResult, error) {
	return a.converter.ConvertInt(intInput, intType)
}

// ConvertIntAuto performs auto-detection of integer types from decimal input.
// It determines all compatible integer types based on the value range and returns
// all valid representations (e.g., int8, uint8, int16, etc.) in a single result.
// This method is exported to the frontend via Wails bindings.
func (a *App) ConvertIntAuto(intInput string) (*models.ConversionResult, error) {
	return a.converter.ConvertIntAuto(intInput)
}

// ConvertBinary performs all possible conversions on binary input.
// This method is exported to the frontend via Wails bindings.
func (a *App) ConvertBinary(binaryInput string) (*models.ConversionResult, error) {
	return a.converter.ConvertBinary(binaryInput)
}

// ConvertFloat performs conversions from float input to hex and binary.
// floatType specifies the float type: float32 or float64.
// This method is exported to the frontend via Wails bindings.
func (a *App) ConvertFloat(floatInput string, floatType string) (*models.ConversionResult, error) {
	return a.converter.ConvertFloat(floatInput, floatType)
}

// ConvertModbusRegisters converts an array of 16-bit register values.
// Input can be space/comma separated hex values (e.g., "1234 5678" or "0x1234, 0x5678")
// or decimal values with 'd' prefix (e.g., "d1000 d2000").
// This method is exported to the frontend via Wails bindings.
func (a *App) ConvertModbusRegisters(input string) (*models.ModbusResult, error) {
	return a.converter.ConvertModbusRegisters(input)
}
