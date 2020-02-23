package main

import (
	"fmt"
	"os"
	"strconv"

	"./ray"
	"./vector"
)

func main() {
	writeToFile("simple_image.ppm", makeSimpleImage())
	writeToFile("rainbow_image.ppm", makeRainbowImage())
	writeToFile("background_image.ppm", makeBackgroundImage())
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
			// Current Vector Value
			currentBaseIntensities := vector.Vector3D{E1: (float64(redIntensityBase) / float64(cols)), E2: (float64(greenIntensityBase) / float64(rows)), E3: 0.2}
			// Calculating the appropriate Red, Green, Blue values
			redIntensity := 255.99 * currentBaseIntensities.R()
			greenIntensity := 255.99 * currentBaseIntensities.G()
			blueIntensity := 255.99 * currentBaseIntensities.B()

			redIntensityInteger := int(redIntensity)
			greenIntensityInteger := int(greenIntensity)
			blueIntensityInteger := int(blueIntensity)

			currentIntensitiesString := strconv.Itoa(redIntensityInteger) + " " + strconv.Itoa(greenIntensityInteger) + " " + strconv.Itoa(blueIntensityInteger) + "\n"
			outputString += currentIntensitiesString
		}
	}
	return outputString
}

func color(r *ray.Ray) vector.Vector3D {
	// Calculate the unit direction vector of the inputted Ray
	directionVector := r.Direction()
	unitDirectionVector := directionVector.MakeUnitVector()
	// Computation for linearly blending in white and blue colors
	t := 0.5 * (unitDirectionVector.Y() + 1.0)
	vectorA := vector.Vector3D{E1: 1.0, E2: 1.0, E3: 1.0}
	vectorB := vector.Vector3D{E1: 0.5, E2: 0.7, E3: 1.0}
	vectorAt := vectorA.ScalarMultiply(1 - t)
	vectorBt := vectorB.ScalarMultiply(t)
	return vectorAt.Add(vectorBt)
}

func makeBackgroundImage() string {
	rows := 100
	cols := 200

	outputString := "P3\n" + strconv.Itoa(cols) + " " + strconv.Itoa(rows) + "\n255\n"

	// The 4 main 3D Vectors we will be using
	lowerLeftCorner := vector.Vector3D{E1: -2.0, E2: -1.0, E3: -1.0}
	horizontal := vector.Vector3D{E1: 4.0, E2: 0.0, E3: 0.0}
	vertical := vector.Vector3D{E1: 0.0, E2: 2.0, E3: 0.0}
	origin := vector.Vector3D{E1: 0.0, E2: 0.0, E3: 0.0}

	for currentCol := cols - 1; currentCol >= 0; currentCol-- {
		for currentRow := 0; currentRow < rows; currentRow++ {
			horizontalVectorCoefficient := float64(currentRow) / float64(rows)
			verticalVectorCoefficient := float64(currentCol) / float64(cols)

			totalHorizontalVector := horizontal.ScalarMultiply(horizontalVectorCoefficient)
			totalVerticalVector := vertical.ScalarMultiply(verticalVectorCoefficient)

			rayAParameter := origin
			rayBParameter := lowerLeftCorner.Add(totalHorizontalVector.Add(totalVerticalVector))
			r := ray.Ray{A: rayAParameter, B: rayBParameter}

			colorVector := color(&r)

			redIntensity := int(255.99 * colorVector.E1)
			greenIntensity := int(255.99 * colorVector.E2)
			blueIntensity := int(255.99 * colorVector.E3)

			currentIntensitiesString := strconv.Itoa(redIntensity) + " " + strconv.Itoa(greenIntensity) + " " + strconv.Itoa(blueIntensity) + "\n"
			outputString += currentIntensitiesString
		}
	}
	return outputString
}
