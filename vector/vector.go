package vector

import "math"

/*Vector3D ... A type that represents 3D vectors as arrays of length 3 */
type Vector3D struct {
	E1 float64
	E2 float64
	E3 float64
}

// The following three functions are for getting the x,y, and z values of
// a 3D positional vector.

/*X ...  */
func (v *Vector3D) X() float64 {
	return v.E1
}

/*Y ... */
func (v *Vector3D) Y() float64 {
	return v.E2
}

/*Z ... */
func (v *Vector3D) Z() float64 {
	return v.E3
}

// The following three functions are for getting the red, green and blue
// intensity values for a 3D color vector.

/*R ... */
func (v *Vector3D) R() float64 {
	return v.E1
}

/*G ... */
func (v *Vector3D) G() float64 {
	return v.E2
}

/*B ... */
func (v *Vector3D) B() float64 {
	return v.E3
}

// All the functions below are generalized functions for vectors regardless
// of what the vectors are being used for.

//  Basic Mathematical Functions

/*Add ... */
func (v *Vector3D) Add(v2 Vector3D) Vector3D {
	return Vector3D{v.E1 + v2.E1, v.E2 + v2.E2, v.E3 + v2.E3}
}

/*Subtract ... */
func (v *Vector3D) Subtract(v2 Vector3D) Vector3D {
	return Vector3D{v.E1 - v.E1, v.E2 - v2.E2, v.E3 - v2.E3}
}

/*Multiply ... */
func (v *Vector3D) Multiply(v2 Vector3D) Vector3D {
	return Vector3D{v.E1 * v2.E1, v.E2 * v2.E2, v.E3 * v2.E3}
}

/*Divide ... */
func (v *Vector3D) Divide(v2 Vector3D) Vector3D {
	return Vector3D{v.E1 / v2.E1, v.E2 / v2.E2, v.E3 / v2.E3}
}

/*DotProduct ... */
func (v *Vector3D) DotProduct(v2 Vector3D) float64 {
	return v.E1*v2.E1 + v.E2*v2.E2 + v.E3*v2.E3
}

/*CrossProduct ... */
func (v *Vector3D) CrossProduct(v2 Vector3D) Vector3D {
	// Using the following equation for the cross product
	// = ((a_2*b_3-a_3b_2),(-(a_1*b_3-a_3*b_1)), (a_1*b_2-a_2*b_1))
	crossProduct := Vector3D{}
	crossProduct.E1 = v.E2*v2.E3 - v.E3*v2.E2
	crossProduct.E2 = -1.0 * (v.E1*v2.E3 - v.E3*v2.E1)
	crossProduct.E3 = v.E1*v2.E2 - v.E2*v2.E1
	return crossProduct
}

/*ScalarMultiply ... */
func (v *Vector3D) ScalarMultiply(c float64) Vector3D {
	return Vector3D{v.E1 * c, v.E2 * c, v.E3 * c}
}

/*ScalarDivide ... */
func (v *Vector3D) ScalarDivide(c float64) Vector3D {
	return Vector3D{v.E1 / c, v.E2 / c, v.E3 / c}
}

/*Length ... */
func (v *Vector3D) Length() float64 {
	return math.Sqrt(math.Pow(v.E1, 2.0) + math.Pow(v.E2, 2.0) + math.Pow(v.E3, 2.0))
}

/*SquaredLength ... */
func (v *Vector3D) SquaredLength() float64 {
	return math.Pow(v.E1, 2.0) + math.Pow(v.E2, 2.0) + math.Pow(v.E3, 2.0)
}

/*MakeUnitVector ... */
func (v *Vector3D) MakeUnitVector() Vector3D {
	denominator := v.Length()
	unitVector := Vector3D{}
	unitVector.E1 = v.E1 / denominator
	unitVector.E2 = v.E2 / denominator
	unitVector.E3 = v.E3 / denominator
	return unitVector
}
