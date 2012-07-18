package gfx

// #cgo LDFLAGS:-lcsfml-graphics
// #include <SFML/Graphics/Export.h>
// #include <SFML/Graphics/Rect.h>
// #include <SFML/Graphics/Types.h>
// #include <SFML/System/Vector2.h>
// #include <SFML/Graphics/Transform.h>
import "C"

type Transform struct {
	Cref *C.sfTransform
}

// \brief Create a new transform
// This function creates an identity transform.
// \return A new sfTransform object
// sfTransform* sfTransform_create(void);
func NewTransform() Transform { 
    return Transform{C.sfTransform_create()}
}

// \brief Create a new transform from a matrix
// \param a00 Element (0, 0) of the matrix
// \param a01 Element (0, 1) of the matrix
// \param a02 Element (0, 2) of the matrix
// \param a10 Element (1, 0) of the matrix
// \param a11 Element (1, 1) of the matrix
// \param a12 Element (1, 2) of the matrix
// \param a20 Element (2, 0) of the matrix
// \param a21 Element (2, 1) of the matrix
// \param a22 Element (2, 2) of the matrix
// \return A new sfTransform object
// sfTransform* sfTransform_createFromMatrix(	float a00, float a01, float a02, 
// float a10, float a11, float a12,
// float a20, float a21, float a22);
func NewFromMatrix(a00, a01, a02, a10, a11, a12, a20, a21, a22 float32) Transform {	
    return Transform {
		C.sfTransform_createFromMatrix(
			C.float(a00), C.float(a01), C.float(a02), 
			C.float(a10), C.float(a11), C.float(a12),
			C.float(a20), C.float(a21), C.float(a22),
		)}
}
            
// \brief Copy an existing transform
// \param transform Transform to copy
// \return Copied object
// sfTransform* sfTransform_copy(sfTransform* transform);
func (self Transform) Copy() Transform { 
    return Transform{C.sfTransform_copy(self.Cref)}
}

// \brief Destroy an existing transform
// \param transform Transform to delete
// void sfTransform_destroy(sfTransform* transform);
func (self Transform) Destroy() { 
    C.sfTransform_destroy(self.Cref);
}

/*                       
            
// \brief Return the 4x4 matrix of a transform
// This function returns a pointer to an array of 16 floats
// containing the transform elements as a 4x4 matrix, which
// is directly compatible with OpenGL functions.
// \code
// sfTransform* transform = ...;
// glLoadMatrixf(sfTransform_getMatrix(transform));
// \endcode
// \param transform Transform object
// \return Pointer to a 4x4 matrix
// const float* sfTransform_getMatrix(const sfTransform* transform);

func (self *float) *float(Transform_getMatrix)  { 
    return C.sf*float(self.Cref);
}
            
// \brief Return the inverse of a transform
// If the inverse cannot be computed, a new identity transform
// is returned.
// \param transform Transform object
// \param result Returned inverse matrix
// void sfTransform_getInverse(const sfTransform* transform, sfTransform* result);

func (self Transform) Getinverse(result *Transform) void { 
    return C.sfTransform_getInverse(self.Cref, sf*transform(result));
}
            
// \brief Apply a transform to a 2D point
// \param transform Transform object
// \param point     Point to transform
// \return Transformed point
// sfVector2f sfTransform_transformPoint(const sfTransform* transform, sfVector2f point);

func (self Transform) Transformpoint(point Vector2f) Vector2f { 
    return C.sfTransform_transformPoint(self.Cref, sfVector2f(point));
}
            
// \brief Apply a transform to a rectangle
// Since SFML doesn't provide support for oriented rectangles,
// the result of this function is always an axis-aligned
// rectangle. Which means that if the transform contains a
// rotation, the bounding rectangle of the transformed rectangle
// is returned.
// \param transform Transform object
// \param rectangle Rectangle to transform
// \return Transformed rectangle
// sfFloatRect sfTransform_transformRect(const sfTransform* transform, sfFloatRect rectangle);

func (self Transform) Transformrect(rectangle FloatRect) FloatRect { 
    return C.sfTransform_transformRect(self.Cref, sfFloatrect(rectangle));
}
            
// \brief Combine two transforms
// The result is a transform that is equivalent to applying
// \a transform followed by \a other. Mathematically, it is
// equivalent to a matrix multiplication.
// \param transform Transform object
// \param right     Transform to combine to \a transform
// void sfTransform_combine(sfTransform* transform, const sfTransform* other);

func (self Transform) Combine(other *Transform ) void { 
    return C.sfTransform_combine(self.Cref, sf(*Transform));
}
            
// \brief Combine a transform with a translation
// \param transform Transform object
// \param x         Offset to apply on X axis
// \param y         Offset to apply on Y axis
// void sfTransform_translate(sfTransform* transform, float x, float y);

func (self Transform) Translate(x float, y float) void { 
    return C.sfTransform_translate(self.Cref, sfFloat(x), sfFloat(y));
}
            
// \brief Combine the current transform with a rotation
// \param transform Transform object
// \param angle     Rotation angle, in degrees
// void sfTransform_rotate(sfTransform* transform, float angle);

func (self Transform) Rotate(angle float) void { 
    return C.sfTransform_rotate(self.Cref, sfFloat(angle));
}
            
// \brief Combine the current transform with a rotation
// The center of rotation is provided for convenience as a second
// argument, so that you can build rotations around arbitrary points
// more easily (and efficiently) than the usual
// [translate(-center), rotate(angle), translate(center)].
// \param transform Transform object
// \param angle     Rotation angle, in degrees
// \param centerX   X coordinate of the center of rotation
// \param centerY   Y coordinate of the center of rotation
// void sfTransform_rotateWithCenter(sfTransform* transform, float angle, float centerX, float centerY);

func (self Transform) Rotatewithcenter(angle float, centerX float, centerY float) void { 
    return C.sfTransform_rotateWithCenter(self.Cref, sfFloat(angle), sfFloat(centerX), sfFloat(centerY));
}
            
// \brief Combine the current transform with a scaling
// \param transform Transform object
// \param scaleX    Scaling factor on the X axis
// \param scaleY    Scaling factor on the Y axis
// void sfTransform_scale(sfTransform* transform, float scaleX, float scaleY);

func (self Transform) Scale(scaleX float, scaleY float) void { 
    return C.sfTransform_scale(self.Cref, sfFloat(scaleX), sfFloat(scaleY));
}
            
// \brief Combine the current transform with a scaling
// The center of scaling is provided for convenience as a second
// argument, so that you can build scaling around arbitrary points
// more easily (and efficiently) than the usual
// [translate(-center), scale(factors), translate(center)]
// \param transform Transform object
// \param scaleX    Scaling factor on X axis
// \param scaleY    Scaling factor on Y axis
// \param centerX   X coordinate of the center of scaling
// \param centerY   Y coordinate of the center of scaling
// void sfTransform_scaleWithCenter(sfTransform* transform, float scaleX, float scaleY, float centerX, float centerY);

func (self Transform) Scalewithcenter(scaleX float, scaleY float, centerX float, centerY float) void { 
    return C.sfTransform_scaleWithCenter(self.Cref, sfFloat(scaleX), sfFloat(scaleY), sfFloat(centerX), sfFloat(centerY));
}
            
*/