package ray

import "../vector"

/*Ray ... */
type Ray struct {
	A vector.Vector3D
	B vector.Vector3D
}

/*Origin ... */
func (r *Ray) Origin() vector.Vector3D {
	return r.A
}

/*Direction ... */
func (r *Ray) Direction() vector.Vector3D {
	return r.B
}

/*PointForParameter ... */
func (r *Ray) PointForParameter(t float64) vector.Vector3D {
	rBt := r.B.ScalarMultiply(t)
	return r.A.Add(rBt)
}
