package goraytracer

import "math"

/*Vector3D ... A type that represents 3D vectors as arrays of length 3 */
type Vector3D struct {
	e1 float64
	e2 float64
	e3 float64
}

// The following three functions are for getting the x,y, and z values of
// a 3D positional vector.
func (v *Vector3D) x() float64 {
	return v.e1
}

func (v *Vector3D) y() float64 {
	return v.e2
}

func (v *Vector3D) z() float64 {
	return v.e3
}

// The following three functions are for getting the red, green and blue
// intensity values for a 3D color vector.
func (v *Vector3D) r() float64 {
	return v.e1
}

func (v *Vector3D) g() float64 {
	return v.e2
}

func (v *Vector3D) b() float64 {
	return v.e3
}

// All the functions below are generalized functions for vectors regardless
// of what the vectors are being used for.

//  Basic Mathematical Functions
func (v *Vector3D) add(v2 Vector3D) {
	v.e1 += v2.e1
	v.e2 += v2.e2
	v.e3 += v2.e3
}

func (v *Vector3D) subtract(v2 Vector3D) {
	v.e1 -= v2.e1
	v.e2 -= v2.e2
	v.e3 -= v2.e3
}

func (v *Vector3D) multiply(v2 Vector3D) {
	v.e1 *= v2.e1
	v.e2 *= v2.e2
	v.e3 *= v2.e3
}

func (v *Vector3D) divide(v2 Vector3D) {
	v.e1 /= v2.e1
	v.e2 /= v2.e2
	v.e3 /= v2.e3
}

func (v *Vector3D) dotProduct(v2 Vector3D) float64 {
	return v.e1*v2.e1 + v.e2*v2.e2 + v.e3*v2.e3
}

func (v *Vector3D) crossProduct(v2 Vector3D) Vector3D {
	// Using the following equation for the cross product
	// = ((a_2*b_3-a_3b_2),(-(a_1*b_3-a_3*b_1)), (a_1*b_2-a_2*b_1))
	crossProduct := Vector3D{}
	crossProduct.e1 = v.e2*v2.e3 - v.e3*v2.e2
	crossProduct.e2 = -1.0 * (v.e1*v2.e3 - v.e3*v2.e1)
	crossProduct.e3 = v.e1*v2.e2 - v.e2*v2.e1
	return crossProduct
}

func (v *Vector3D) scalarMultiply(c float64) {
	v.e1 *= c
	v.e2 *= c
	v.e3 *= c
}

func (v *Vector3D) scalarDivide(c float64) {
	v.e1 /= c
	v.e2 /= c
	v.e3 /= c
}

// Vector Length Function
func (v *Vector3D) length() float64 {
	return math.Sqrt(math.Pow(v.e1, 2.0) + math.Pow(v.e2, 2.0) + math.Pow(v.e3, 2.0))
}

// Vector Length Squared Function
func (v *Vector3D) squaredLength() float64 {
	return math.Pow(v.e1, 2.0) + math.Pow(v.e2, 2.0) + math.Pow(v.e3, 2.0)
}

// Function for making Unit Vector
func (v *Vector3D) makeUnitVector() Vector3D {
	denominator := v.length()
	unitVector := Vector3D{}
	unitVector.e1 = v.e1 / denominator
	unitVector.e2 = v.e2 / denominator
	unitVector.e3 = v.e3 / denominator
	return unitVector
}
