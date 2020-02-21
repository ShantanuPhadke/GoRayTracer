package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	makeSimpleImage()
	makeRainbowImage()
}

func makeSimpleImage() {
	f, err := os.Create("simple_image.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}
	simpleImageString := ""
	simpleImageString += "P3\n"
	simpleImageString += "3 2\n"
	simpleImageString += "255\n"
	simpleImageString += "255 0 0\n0 255 0\n0 0 255\n"
	simpleImageString += "255 255 0\n255 255 255\n0 0 0 "

	l, err := f.WriteString(simpleImageString)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func makeRainbowImage() {
	cols := 200
	rows := 100

	outputString := "P3\n" + strconv.Itoa(cols) + " " + strconv.Itoa(rows) + "\n255\n"

	f, err := os.Create("rainbow_image.ppm")
	if err != nil {
		fmt.Println(err)
		return
	}

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

	l, err := f.WriteString(outputString + " ")
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

}
