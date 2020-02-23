package ray

import (
	"testing"

	"../vector"
)

func TestRays(t *testing.T) {

	assertFloatVectorsEqual := func(t *testing.T, expected, actual vector.Vector3D) {
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

	t.Run("Testing ray origin and direction parameters", func(t *testing.T) {
		vectorA := vector.Vector3D{E1: 1.0, E2: 1.0, E3: 1.0}
		vectorB := vector.Vector3D{E1: 2.0, E2: 2.0, E3: 2.0}
		rayAB := Ray{vectorA, vectorB}

		expectedDirection := vector.Vector3D{E1: 2.0, E2: 2.0, E3: 2.0}
		actualDirection := rayAB.Direction()
		assertFloatVectorsEqual(t, expectedDirection, actualDirection)

		expectedOrigin := vector.Vector3D{E1: 1.0, E2: 1.0, E3: 1.0}
		actualOrigin := rayAB.Origin()
		assertFloatVectorsEqual(t, expectedOrigin, actualOrigin)
	})

	t.Run("Testing the PointForParameter function", func(t *testing.T) {
		vectorA := vector.Vector3D{E1: 1.0, E2: 1.5, E3: 2.0}
		vectorB := vector.Vector3D{E1: 2.0, E2: 2.5, E3: 3.0}
		rayAB := Ray{vectorA, vectorB}

		expectedPoint := vector.Vector3D{E1: 3.0, E2: 4.0, E3: 5.0}
		actualPoint := rayAB.PointForParameter(1.0)
		assertFloatVectorsEqual(t, expectedPoint, actualPoint)
	})

}
