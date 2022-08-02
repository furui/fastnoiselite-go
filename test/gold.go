package main

//#cgo CFLAGS: -I fastnoiselite/C/
/*
#define FNL_IMPL

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include "FastNoiseLite.h"

typedef struct {
    float oSimplex2GetNoise2D;
	float oSimplex2GetNoise3D;
    float oSimplex2SGetNoise2D;
	float oSimplex2SGetNoise3D;
    float oCellularGetNoise2D;
	float oCellularGetNoise3D;
	float oPerlinGetNoise2D;
	float oPerlinGetNoise3D;
	float oValueCubicNoise2D;
	float oValueCubicNoise3D;
	float oValueNoise2D;
	float oValueNoise3D;

	float oFractalNone2D;
	float oFractalNone3D;
	float oFractalRidged2D;
	float oFractalRidged3D;
	float oFractalPingPong2D;
	float oFractalPingPong3D;
} GoldOutput;

int test(GoldOutput *out)
{
  fnl_state noise = fnlCreateState();
  noise.noise_type = FNL_NOISE_OPENSIMPLEX2;
  noise.fractal_type = FNL_FRACTAL_FBM;
  noise.frequency = 0.02f;
  noise.octaves = 5;

  out->oSimplex2GetNoise2D = fnlGetNoise2D(&noise, 0.1f, 0.1f);
  out->oSimplex2GetNoise3D = fnlGetNoise3D(&noise, 0.1f, 0.1f, 0.1f);
  noise.noise_type = FNL_NOISE_OPENSIMPLEX2S;
  out->oSimplex2SGetNoise2D = fnlGetNoise2D(&noise, 0.1f, 0.1f);
  out->oSimplex2SGetNoise3D = fnlGetNoise3D(&noise, 0.1f, 0.1f, 0.1f);
  noise.noise_type = FNL_NOISE_CELLULAR;
  out->oCellularGetNoise2D = fnlGetNoise2D(&noise, 0.1f, 0.1f);
  out->oCellularGetNoise3D = fnlGetNoise3D(&noise, 0.1f, 0.1f, 0.1f);
  noise.noise_type = FNL_NOISE_PERLIN;
  out->oPerlinGetNoise2D = fnlGetNoise2D(&noise, 0.1f, 0.1f);
  out->oPerlinGetNoise3D = fnlGetNoise3D(&noise, 0.1f, 0.1f, 0.1f);
  noise.noise_type = FNL_NOISE_VALUE_CUBIC;
  out->oValueCubicNoise2D = fnlGetNoise2D(&noise, 0.1f, 0.1f);
  out->oValueCubicNoise3D = fnlGetNoise3D(&noise, 0.1f, 0.1f, 0.1f);
  noise.noise_type = FNL_NOISE_VALUE;
  out->oValueNoise2D = fnlGetNoise2D(&noise, 0.1f, 0.1f);
  out->oValueNoise3D = fnlGetNoise3D(&noise, 0.1f, 0.1f, 0.1f);

  noise.noise_type = FNL_NOISE_PERLIN;
  noise.fractal_type = FNL_FRACTAL_NONE;
  out->oFractalNone2D = fnlGetNoise2D(&noise, 0.1f, 0.1f);
  out->oFractalNone3D = fnlGetNoise3D(&noise, 0.1f, 0.1f, 0.1f);
  noise.fractal_type = FNL_FRACTAL_RIDGED;
  out->oFractalRidged2D = fnlGetNoise2D(&noise, 0.1f, 0.1f);
  out->oFractalRidged3D = fnlGetNoise3D(&noise, 0.1f, 0.1f, 0.1f);
  noise.fractal_type = FNL_FRACTAL_PINGPONG;
  out->oFractalPingPong2D = fnlGetNoise2D(&noise, 0.1f, 0.1f);
  out->oFractalPingPong3D = fnlGetNoise3D(&noise, 0.1f, 0.1f, 0.1f);

  return 0;
}

float CubicLerp(float a, float b, float c, float d, float t)
{
    float p = (d - c) - (a - b);
    return t * t * t * p + t * t * ((a - b) - p) + t * (c - a) + b;
}

int Hash2D(int seed, int xPrimed, int yPrimed)
{
    int hash = seed ^ xPrimed ^ yPrimed;

    hash *= 0x27d4eb2d;
    return hash;
}

float ValCoord2D(int seed, int xPrimed, int yPrimed)
{
    int hash = Hash2D(seed, xPrimed, yPrimed);
    hash *= hash;
    hash ^= hash << 19;
    return hash * (1 / 2147483648.0f);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type GoldOutput C.GoldOutput

func GoldLerp(a, b, c, d, t float32) float32 {
	return float32(C.CubicLerp(C.float(a), C.float(b), C.float(c), C.float(d), C.float(t)))
}

func GoldValCoord2D(seed, xPrimed, yPrimed int) float32 {
	return float32(C.ValCoord2D(C.int(seed), C.int(xPrimed), C.int(yPrimed)))
}

func GetGold() string {
	gold := GoldOutput{}
	C.test((*C.GoldOutput)(unsafe.Pointer(&gold)))
	s := "{"
	s += fmt.Sprintf("\"oSimplex2GetNoise2D\": %f,", gold.oSimplex2GetNoise2D)
	s += fmt.Sprintf("\"oSimplex2GetNoise3D\": %f,", gold.oSimplex2GetNoise3D)
	s += fmt.Sprintf("\"oSimplex2SGetNoise2D\": %f,", gold.oSimplex2SGetNoise2D)
	s += fmt.Sprintf("\"oSimplex2SGetNoise3D\": %f,", gold.oSimplex2SGetNoise3D)
	s += fmt.Sprintf("\"oCellularGetNoise2D\": %f,", gold.oCellularGetNoise2D)
	s += fmt.Sprintf("\"oCellularGetNoise3D\": %f,", gold.oCellularGetNoise3D)
	s += fmt.Sprintf("\"oPerlinGetNoise2D\": %f,", gold.oPerlinGetNoise2D)
	s += fmt.Sprintf("\"oPerlinGetNoise3D\": %f,", gold.oPerlinGetNoise3D)
	s += fmt.Sprintf("\"oValueCubicNoise2D\": %f,", gold.oValueCubicNoise2D)
	s += fmt.Sprintf("\"oValueCubicNoise3D\": %f,", gold.oValueCubicNoise3D)
	s += fmt.Sprintf("\"oValueNoise2D\": %f,", gold.oValueNoise2D)
	s += fmt.Sprintf("\"oValueNoise3D\": %f,", gold.oValueNoise3D)
	s += fmt.Sprintf("\"oFractalNone2D\": %f,", gold.oFractalNone2D)
	s += fmt.Sprintf("\"oFractalNone3D\": %f,", gold.oFractalNone3D)
	s += fmt.Sprintf("\"oFractalRidged2D\": %f,", gold.oFractalRidged2D)
	s += fmt.Sprintf("\"oFractalRidged3D\": %f,", gold.oFractalRidged3D)
	s += fmt.Sprintf("\"oFractalPingPong2D\": %f,", gold.oFractalPingPong2D)
	s += fmt.Sprintf("\"oFractalPingPong3D\": %f", gold.oFractalPingPong3D)
	s += "}"
	return s
}
