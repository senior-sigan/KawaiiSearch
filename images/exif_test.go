package images

import (
"image/color"
"os"
"strconv"
"testing"
)

func TestDecode(t *testing.T) {
	for i := 1; i < 9; i++ {
		path := "test_data/f" + strconv.Itoa(i) + "t.jpg"
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		img, _, err := Decode(file)
		if err != nil {
			panic(err)
		}
		bounds := img.Bounds()
		minP := bounds.Min
		maxP := bounds.Max
		leftUpper := img.At(minP.X, minP.Y)
		rightUpper := img.At(maxP.X-1, minP.Y)
		leftLower := img.At(minP.X, maxP.Y-1)
		rightLower := img.At(maxP.X-1, maxP.Y-1)
		if !(isBlack(leftUpper) && isBlack(rightUpper) && isBlack(leftLower) && !isBlack(rightLower)) {
			t.Errorf(`
				The path is %s.
				isBlack(leftUpper) && isBlack(rightUpper) && isBlack(leftLower) && !isBlack(rightLower):
				Expected 	%t, %t, %t, %t.
				Got			%t, %t, %t, %t`,
				path,
				true, true, true, true,
				isBlack(leftUpper), isBlack(rightUpper), isBlack(leftLower), !isBlack(rightLower),
			)
		}
	}
}

func isBlack(point color.Color) bool {
	r, g, b, _ := point.RGBA()
	return r == 0 && g == 0 && b == 0
}
