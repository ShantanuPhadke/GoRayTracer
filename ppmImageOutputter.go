package goraytracer

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	writeToFile("simple_image.ppm", makeSimpleImage())
	writeToFile("rainbow_image.ppm", makeRainbowImage())
}

/* writeToFile ... Generalized function that just writes given content
   to a given file
		INPUTS:
			filename = file name as a string
			content = the content to write to the file as a string
		OUTPUTS:
			bool = Signifies whether or not the function was successful
			error = Error (if any) that is returned by any operations
*/
func writeToFile(filename string, content string) (bool, error) {
	f, err := os.Create(filename)
	if err != nil {
		return false, err
	}

	l, err := f.WriteString(content)
	if err != nil {
		f.Close()
		return false, err
	}

	fmt.Println(l, "bytes written successfully to", filename)

	err = f.Close()
	if err != nil {
		return false, err
	}

	return true, nil
}

func makeSimpleImage() string {
	simpleImageString := ""
	simpleImageString += "P3\n"
	simpleImageString += "3 2\n"
	simpleImageString += "255\n"
	simpleImageString += "255 0 0\n0 255 0\n0 0 255\n"
	simpleImageString += "255 255 0\n255 255 255\n0 0 0 "
	return simpleImageString
}

func makeRainbowImage() string {
	cols := 200
	rows := 100

	outputString := "P3\n" + strconv.Itoa(cols) + " " + strconv.Itoa(rows) + "\n255\n"

	for greenIntensityBase := rows - 1; greenIntensityBase >= 0; greenIntensityBase-- {
		for redIntensityBase := 0; redIntensityBase < cols; redIntensityBase++ {
			// Calculating the appropriate Red, Green, Blue values
			redIntensity := 255.99 * (float64(redIntensityBase) / float64(cols))
			greenIntensity := 255.99 * (float64(greenIntensityBase) / float64(rows))
			blueIntensity := 255.99 * 0.2

			redIntensityInteger := int(redIntensity)
			greenIntensityInteger := int(greenIntensity)
			blueIntensityInteger := int(blueIntensity)

			currentIntensitiesString := strconv.Itoa(redIntensityInteger) + " " + strconv.Itoa(greenIntensityInteger) + " " + strconv.Itoa(blueIntensityInteger) + "\n"
			outputString += currentIntensitiesString
		}
	}
	return outputString
}
