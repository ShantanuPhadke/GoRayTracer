package vector

import (
	"testing"
)

func TestVector(t *testing.T) {

	assertFloatValuesEqual := func(t *testing.T, expected, actual float64) {
		t.Helper()
		if expected != actual {
			t.Errorf("Expected %f but got %f", expected, actual)
		}
	}

	assertFloatVectorsEqual := func(t *testing.T, expected, actual Vector3D) {
		t.Helper()
		if expected.E1 != actual.E1 {
			t.Errorf("Element 1 unexpected: Expected %f but got %f", expected.E1, actual.E1)
		}
		if expected.E2 != actual.E2 {
			t.Errorf("Element 2 unexpected: Expected %f but got %f", expected.E2, actual.E2)
		}
		if expected.E3 != actual.E3 {
			t.Errorf("Element 3 unexpected: Expected %f but got %f", expected.E3, actual.E3)
		}
	}

	t.Run("Testing the positional elements of a vector", func(t *testing.T) {
		sampleVector := Vector3D{1.0, 2.0, 3.0}

		expectedX := 1.0
		actualX := sampleVector.X()
		assertFloatValuesEqual(t, expectedX, actualX)

		expectedY := 2.0
		actualY := sampleVector.Y()
		assertFloatValuesEqual(t, expectedY, actualY)

		expectedZ := 3.0
		actualZ := sampleVector.Z()
		assertFloatValuesEqual(t, expectedZ, actualZ)
	})

	t.Run("Testing adding two vectors together", func(t *testing.T) {
		sampleVector1 := Vector3D{1.0, 1.5, 2.5}
		sampleVector2 := Vector3D{2.0, 2.1, 2.3}
		expectedResult := Vector3D{3.0, 3.6, 4.8}
		actualResult := sampleVector1.Add(sampleVector2)
		assertFloatVectorsEqual(t, expectedResult, actualResult)
	})

	t.Run("Testing scalar multiplication", func(t *testing.T) {
		sampleVector1 := Vector3D{1.0, 1.5, 2.5}
		sampleScalar := 2.0
		expectedResult := Vector3D{2.0, 3.0, 5.0}
		actualResult := sampleVector1.ScalarMultiply(sampleScalar)
		assertFloatVectorsEqual(t, expectedResult, actualResult)
	})

	t.Run("Testing vector length calculation", func(t *testing.T) {
		sampleVector1 := Vector3D{6.0, 8.0, 0.0}
		expectedLength := 10.0
		actualLength := sampleVector1.Length()
		assertFloatValuesEqual(t, expectedLength, actualLength)
	})

	t.Run("Testing vector length calculation", func(t *testing.T) {
		sampleVector1 := Vector3D{6.0, 8.0, 0.0}
		expectedSquaredLength := 100.0
		actualSquaredLength := sampleVector1.SquaredLength()
		assertFloatValuesEqual(t, expectedSquaredLength, actualSquaredLength)
	})

	t.Run("Testing unit vector calculation", func(t *testing.T) {
		sampleVector1 := Vector3D{6.0, 8.0, 0.0}
		expectedUnitVector := Vector3D{0.6, 0.8, 0.000000}
		actualUnitVector := sampleVector1.MakeUnitVector()
		assertFloatVectorsEqual(t, expectedUnitVector, actualUnitVector)
	})
}
