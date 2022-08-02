package main

import (
	"encoding/json"
	"testing"

	"github.com/furui/fastnoiselite-go"
)

const errorAllowed = 0.001

type goldValues struct {
	vals map[string]float64
}

func loadValues(data string) (*goldValues, error) {
	output := &goldValues{}
	output.vals = make(map[string]float64)
	err := json.Unmarshal([]byte(data), &output.vals)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func Compare(t *testing.T, values *(map[string]float64), name string, actual float64) {
	expected := (*values)[name]
	diff := expected - actual
	if diff > (errorAllowed) || diff < (0-errorAllowed) {
		t.Errorf("for %s, expected: %f, actual: %f", name, expected, actual)
	}
}

func TestLerp(t *testing.T) {
	goldValue := GoldLerp(0.123, 0.223, 0.323, 0.423, 0.023)
	testValue := float32(fastnoiselite.CubicLerp(0.123, 0.223, 0.323, 0.423, 0.023))
	diff := goldValue - testValue
	if diff > (errorAllowed) || diff < (0-errorAllowed) {
		t.Errorf("for CubicLerp, expected: %f, actual: %f", goldValue, testValue)
	}
}

func TestValCoord2D(t *testing.T) {
	goldValue := GoldValCoord2D(1337, 1, 1)
	testValue := float32(fastnoiselite.ValCoord2D(1337, 1, 1))
	diff := goldValue - testValue
	if diff > (errorAllowed) || diff < (0-errorAllowed) {
		t.Errorf("for ValCoord2D, expected: %f, actual: %f", goldValue, testValue)
	}
}

func TestValCoord3D(t *testing.T) {
	goldValue := GoldValCoord3D(1337, 1, 1, 1)
	testValue := float32(fastnoiselite.ValCoord3D(1337, 1, 1, 1))
	diff := goldValue - testValue
	if diff > (errorAllowed) || diff < (0-errorAllowed) {
		t.Errorf("for ValCoord2D, expected: %f, actual: %f", goldValue, testValue)
	}
}

func TestNoise(t *testing.T) {
	out := GetGold()
	if out == "" {
		t.Error("got empty output")
	}
	gold, err := loadValues(out)
	if err != nil {
		t.Error(err)

	}

	noise := fastnoiselite.NewNoise()
	noise.SetNoiseType(fastnoiselite.NoiseTypeOpenSimplex2)
	noise.FractalType = fastnoiselite.FractalTypeFBm
	noise.Frequency = 0.02
	noise.SetFractalOctaves(5)

	val := noise.GetNoise2D(0.1, 0.1)
	Compare(t, &gold.vals, "oSimplex2GetNoise2D", val)
	val = noise.GetNoise3D(0.1, 0.1, 0.1)
	Compare(t, &gold.vals, "oSimplex2GetNoise3D", val)
	noise.SetNoiseType(fastnoiselite.NoiseTypeOpenSimplex2S)
	val = noise.GetNoise2D(0.1, 0.1)
	Compare(t, &gold.vals, "oSimplex2SGetNoise2D", val)
	val = noise.GetNoise3D(0.1, 0.1, 0.1)
	Compare(t, &gold.vals, "oSimplex2SGetNoise3D", val)
	noise.SetNoiseType(fastnoiselite.NoiseTypeCellular)
	val = noise.GetNoise2D(0.1, 0.1)
	Compare(t, &gold.vals, "oCellularGetNoise2D", val)
	val = noise.GetNoise3D(0.1, 0.1, 0.1)
	Compare(t, &gold.vals, "oCellularGetNoise3D", val)
	noise.SetNoiseType(fastnoiselite.NoiseTypePerlin)
	val = noise.GetNoise2D(0.1, 0.1)
	Compare(t, &gold.vals, "oPerlinGetNoise2D", val)
	val = noise.GetNoise3D(0.1, 0.1, 0.1)
	Compare(t, &gold.vals, "oPerlinGetNoise3D", val)
	noise.SetNoiseType(fastnoiselite.NoiseTypeValueCubic)
	val = noise.GetNoise2D(0.1, 0.1)
	Compare(t, &gold.vals, "oValueCubicNoise2D", val)
	val = noise.GetNoise3D(0.1, 0.1, 0.1)
	Compare(t, &gold.vals, "oValueCubicNoise3D", val)
	noise.SetNoiseType(fastnoiselite.NoiseTypeValue)
	val = noise.GetNoise2D(0.1, 0.1)
	Compare(t, &gold.vals, "oValueNoise2D", val)
	val = noise.GetNoise3D(0.1, 0.1, 0.1)
	Compare(t, &gold.vals, "oValueNoise3D", val)

	noise.SetNoiseType(fastnoiselite.NoiseTypePerlin)
	noise.FractalType = fastnoiselite.FractalTypeNone
	val = noise.GetNoise2D(0.1, 0.1)
	Compare(t, &gold.vals, "oFractalNone2D", val)
	val = noise.GetNoise3D(0.1, 0.1, 0.1)
	Compare(t, &gold.vals, "oFractalNone3D", val)
	noise.FractalType = fastnoiselite.FractalTypeRidged
	val = noise.GetNoise2D(0.1, 0.1)
	Compare(t, &gold.vals, "oFractalRidged2D", val)
	val = noise.GetNoise3D(0.1, 0.1, 0.1)
	Compare(t, &gold.vals, "oFractalRidged3D", val)
	noise.FractalType = fastnoiselite.FractalTypePingPong
	val = noise.GetNoise2D(0.1, 0.1)
	Compare(t, &gold.vals, "oFractalPingPong2D", val)
	val = noise.GetNoise3D(0.1, 0.1, 0.1)
	Compare(t, &gold.vals, "oFractalPingPong3D", val)
}
