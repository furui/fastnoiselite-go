# fastnoiselite-go

Port of [fastnoiselite](https://github.com/Auburn/FastNoiseLite) to golang.

Use at your own risk until better tests exist.

## Using

```
	noise := fastnoiselite.NewNoise()
	noise.SetNoiseType(fastnoiselite.NoiseTypeOpenSimplex2)
	noise.FractalType = fastnoiselite.FractalTypeFBm
	noise.Frequency = 0.02
	noise.SetFractalOctaves(5)

	val := noise.GetNoise2D(0.1, 0.1)
```

## Tests

The tests use cgo to build the original fastnoiselite library and make the assumption that the original fastnoiselite library is mathematically and algorithmically correct since it compares the golang version with the C version.

To run tests, at the command line:

```
git submodule init
git submodule update
cd test
go test
```

Tests are partially implemented but do not cover the entire codebase yet. If you would like to contribute tests, that would be greatly appreciated it.

## License

The original fastnoiselite library is under the MIT License by Jordan Peck:

```
MIT License

Copyright(c) 2020 Jordan Peck (jordan.me2@gmail.com)
Copyright(c) 2020 Contributors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

This port is also under the MIT License. See: [LICENSE](./LICENSE)
