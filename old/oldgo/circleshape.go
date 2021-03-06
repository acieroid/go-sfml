////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////
/// \brief Create a new circle shape
///
/// \return A new sfCircleShape object, or NULL if it failed
///
////////////////////////////////////////////////////////////
//  sfCircleShape* sfCircleShape_create(void);

func (self CircleShape) Create() *CircleShape { 
    return C.sfCircleShape_create();
}
            
////////////////////////////////////////////////////////////
/// \brief Copy an existing circle shape
///
/// \param shape Shape to copy
///
/// \return Copied object
///
////////////////////////////////////////////////////////////
//  sfCircleShape* sfCircleShape_copy(sfCircleShape* shape);

func (self CircleShape) Copy() *CircleShape { 
    return C.sfCircleShape_copy();
}
            
////////////////////////////////////////////////////////////
/// \brief Destroy an existing circle Shape
///
/// \param Shape Shape to delete
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_destroy(sfCircleShape* shape);

func (self CircleShape) Destroy() void { 
    return C.sfCircleShape_destroy();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the position of a circle shape
///
/// This function completely overwrites the previous position.
/// See sfCircleShape_move to apply an offset based on the previous position instead.
/// The default position of a circle Shape object is (0, 0).
///
/// \param shape    Shape object
/// \param position New position
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setPosition(sfCircleShape* shape, sfVector2f position);

func (self CircleShape) Setposition(position Vector2f) void { 
    return C.sfCircleShape_setPosition();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the orientation of a circle shape
///
/// This function completely overwrites the previous rotation.
/// See sfCircleShape_rotate to add an angle based on the previous rotation instead.
/// The default rotation of a circle Shape object is 0.
///
/// \param shape Shape object
/// \param angle New rotation, in degrees
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setRotation(sfCircleShape* shape, float angle);

func (self CircleShape) Setrotation(angle float) void { 
    return C.sfCircleShape_setRotation();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the scale factors of a circle shape
///
/// This function completely overwrites the previous scale.
/// See sfCircleShape_scale to add a factor based on the previous scale instead.
/// The default scale of a circle Shape object is (1, 1).
///
/// \param shape Shape object
/// \param scale New scale factors
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setScale(sfCircleShape* shape, sfVector2f scale);

func (self CircleShape) Setscale(scale Vector2f) void { 
    return C.sfCircleShape_setScale();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the local origin of a circle shape
///
/// The origin of an object defines the center point for
/// all transformations (position, scale, rotation).
/// The coordinates of this point must be relative to the
/// top-left corner of the object, and ignore all
/// transformations (position, scale, rotation).
/// The default origin of a circle Shape object is (0, 0).
///
/// \param shape  Shape object
/// \param origin New origin
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setOrigin(sfCircleShape* shape, sfVector2f origin);

func (self CircleShape) Setorigin(origin Vector2f) void { 
    return C.sfCircleShape_setOrigin();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the position of a circle shape
///
/// \param shape Shape object
///
/// \return Current position
///
////////////////////////////////////////////////////////////
//  sfVector2f sfCircleShape_getPosition(const sfCircleShape* shape);

func (self CircleShape) Getposition() Vector2f { 
    return C.sfCircleShape_getPosition();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the orientation of a circle shape
///
/// The rotation is always in the range [0, 360].
///
/// \param shape Shape object
///
/// \return Current rotation, in degrees
///
////////////////////////////////////////////////////////////
//  float sfCircleShape_getRotation(const sfCircleShape* shape);

func (self CircleShape) Getrotation() float { 
    return C.sfCircleShape_getRotation();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the current scale of a circle shape
///
/// \param shape Shape object
///
/// \return Current scale factors
///
////////////////////////////////////////////////////////////
//  sfVector2f sfCircleShape_GetScale(const sfCircleShape* shape);

func (self CircleShape) Getscale() Vector2f { 
    return C.sfCircleShape_GetScale();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the local origin of a circle shape
///
/// \param shape Shape object
///
/// \return Current origin
///
////////////////////////////////////////////////////////////
//  sfVector2f sfCircleShape_GetOrigin(const sfCircleShape* shape);

func (self CircleShape) Getorigin() Vector2f { 
    return C.sfCircleShape_GetOrigin();
}
            
////////////////////////////////////////////////////////////
/// \brief Move a circle shape by a given offset
///
/// This function adds to the current position of the object,
/// unlike sfCircleShape_setPosition which overwrites it.
///
/// \param shape  Shape object
/// \param offset Offset
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_move(sfCircleShape* shape, sfVector2f offset);

func (self CircleShape) Move(offset Vector2f) void { 
    return C.sfCircleShape_move();
}
            
////////////////////////////////////////////////////////////
/// \brief Rotate a circle shape
///
/// This function adds to the current rotation of the object,
/// unlike sfCircleShape_setRotation which overwrites it.
///
/// \param shape Shape object
/// \param angle Angle of rotation, in degrees
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_rotate(sfCircleShape* shape, float angle);

func (self CircleShape) Rotate(angle float) void { 
    return C.sfCircleShape_rotate();
}
            
////////////////////////////////////////////////////////////
/// \brief Scale a circle shape
///
/// This function multiplies the current scale of the object,
/// unlike sfCircleShape_setScale which overwrites it.
///
/// \param shape   Shape object
/// \param factors Scale factors
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_scale(sfCircleShape* shape, sfVector2f factors);

func (self CircleShape) Scale(factors Vector2f) void { 
    return C.sfCircleShape_scale();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the combined transform of a circle shape
///
/// \param shape Shape object
///
/// \return Transform combining the position/rotation/scale/origin of the object
///
////////////////////////////////////////////////////////////
//  const sfTransform* sfCircleShape_getTransform(const sfCircleShape* shape);

func (self *Transform) *Transform(CircleShape_getTransform)  { 
    return C.sf*Transform();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the inverse of the combined transform of a circle shape
///
/// \param shape Shape object
///
/// \return Inverse of the combined transformations applied to the object
///
////////////////////////////////////////////////////////////
//  const sfTransform* sfCircleShape_getInverseTransform(const sfCircleShape* shape);

func (self *Transform) *Transform(CircleShape_getInverseTransform)  { 
    return C.sf*Transform();
}
            
////////////////////////////////////////////////////////////
/// \brief Change the source texture of a circle shape
///
/// The \a texture argument refers to a texture that must
/// exist as long as the shape uses it. Indeed, the shape
/// doesn't store its own copy of the texture, but rather keeps
/// a pointer to the one that you passed to this function.
/// If the source texture is destroyed and the shape tries to
/// use it, the behaviour is undefined.
/// \a texture can be NULL to disable texturing.
/// If \a resetRect is true, the TextureRect property of
/// the shape is automatically adjusted to the size of the new
/// texture. If it is false, the texture rect is left unchanged.
///
/// \param shape     Shape object
/// \param texture   New texture
/// \param resetRect Should the texture rect be reset to the size of the new texture?
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setTexture(sfCircleShape* shape, const sfTexture* texture, sfBool resetRect);

func (self CircleShape) Settexture(texture *Texture , resetRect Bool) void { 
    return C.sfCircleShape_setTexture();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the sub-rectangle of the texture that a circle shape will display
///
/// The texture rect is useful when you don't want to display
/// the whole texture, but rather a part of it.
/// By default, the texture rect covers the entire texture.
///
/// \param shape Shape object
/// \param rect  Rectangle defining the region of the texture to display
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setTextureRect(sfCircleShape* shape, sfIntRect rect);

func (self CircleShape) Settexturerect(rect IntRect) void { 
    return C.sfCircleShape_setTextureRect();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the fill color of a circle shape
///
/// This color is modulated (multiplied) with the shape's
/// texture if any. It can be used to colorize the shape,
/// or change its global opacity.
/// You can use sfTransparent to make the inside of
/// the shape transparent, and have the outline alone.
/// By default, the shape's fill color is opaque white.
///
/// \param shape Shape object
/// \param color New color of the shape
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setFillColor(sfCircleShape* shape, sfColor color);

func (self CircleShape) Setfillcolor(color Color) void { 
    return C.sfCircleShape_setFillColor();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the outline color of a circle shape
///
/// You can use sfTransparent to disable the outline.
/// By default, the shape's outline color is opaque white.
///
/// \param shape Shape object
/// \param color New outline color of the shape
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setOutlineColor(sfCircleShape* shape, sfColor color);

func (self CircleShape) Setoutlinecolor(color Color) void { 
    return C.sfCircleShape_setOutlineColor();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the thickness of a circle shape's outline
///
/// This number cannot be negative. Using zero disables
/// the outline.
/// By default, the outline thickness is 0.
///
/// \param shape     Shape object
/// \param thickness New outline thickness
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setOutlineThickness(sfCircleShape* shape, float thickness);

func (self CircleShape) Setoutlinethickness(thickness float) void { 
    return C.sfCircleShape_setOutlineThickness();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the source texture of a circle shape
///
/// If the shape has no source texture, a NULL pointer is returned.
/// The returned pointer is const, which means that you can't
/// modify the texture when you retrieve it with this function.
///
/// \param shape Shape object
///
/// \return Pointer to the shape's texture
///
////////////////////////////////////////////////////////////
//  const sfTexture* sfCircleShape_getTexture(const sfCircleShape* shape);

func (self *Texture) *Texture(CircleShape_getTexture)  { 
    return C.sf*Texture();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the sub-rectangle of the texture displayed by a circle shape
///
/// \param shape Shape object
///
/// \return Texture rectangle of the shape
///
////////////////////////////////////////////////////////////
//  sfIntRect sfCircleShape_getTextureRect(const sfCircleShape* shape);

func (self CircleShape) Gettexturerect() IntRect { 
    return C.sfCircleShape_getTextureRect();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the fill color of a circle shape
///
/// \param shape Shape object
///
/// \return Fill color of the shape
///
////////////////////////////////////////////////////////////
//  sfColor sfCircleShape_getFillColor(const sfCircleShape* shape);

func (self CircleShape) Getfillcolor() Color { 
    return C.sfCircleShape_getFillColor();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the outline color of a circle shape
///
/// \param shape Shape object
///
/// \return Outline color of the shape
///
////////////////////////////////////////////////////////////
//  sfColor sfCircleShape_getOutlineColor(const sfCircleShape* shape);

func (self CircleShape) Getoutlinecolor() Color { 
    return C.sfCircleShape_getOutlineColor();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the outline thickness of a circle shape
///
/// \param shape Shape object
///
/// \return Outline thickness of the shape
///
////////////////////////////////////////////////////////////
//  float sfCircleShape_getOutlineThickness(const sfCircleShape* shape);

func (self CircleShape) Getoutlinethickness() float { 
    return C.sfCircleShape_getOutlineThickness();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the total number of points of a circle shape
///
/// \param shape Shape object
///
/// \return Number of points of the shape
///
////////////////////////////////////////////////////////////
//  unsigned int sfCircleShape_getPointCount(const sfCircleShape* shape);

func (self int) int(CircleShape_getPointCount)  { 
    return C.sfint();
}
            
////////////////////////////////////////////////////////////
/// \brief Get a point of a circle shape
///
/// The result is undefined if \a index is out of the valid range.
///
/// \param shape Shape object
/// \param index Index of the point to get, in range [0 .. getPointCount() - 1]
///
/// \return Index-th point of the shape
///
////////////////////////////////////////////////////////////
//  sfVector2f sfCircleShape_getPoint(const sfCircleShape* shape, unsigned int index);

func (self CircleShape) Getpoint(index int ) Vector2f { 
    return C.sfCircleShape_getPoint();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the radius of a circle
///
/// \param shape  Shape object
/// \param radius New radius of the circle
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setRadius(sfCircleShape* shape, float radius);

func (self CircleShape) Setradius(radius float) void { 
    return C.sfCircleShape_setRadius();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the radius of a circle
///
/// \param shape Shape object
///
/// \return Radius of the circle
///
////////////////////////////////////////////////////////////
//  float sfCircleShape_getRadius(const sfCircleShape* shape);

func (self CircleShape) Getradius() float { 
    return C.sfCircleShape_getRadius();
}
            
////////////////////////////////////////////////////////////
/// \brief Set the number of points of a circle
///
/// \param shape Shape object
/// \param count New number of points of the circle
///
////////////////////////////////////////////////////////////
//  void sfCircleShape_setPointCount(sfCircleShape* shape, unsigned int count);

func (self CircleShape) Setpointcount(count int ) void { 
    return C.sfCircleShape_setPointCount();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the local bounding rectangle of a circle shape
///
/// The returned rectangle is in local coordinates, which means
/// that it ignores the transformations (translation, rotation,
/// scale, ...) that are applied to the entity.
/// In other words, this function returns the bounds of the
/// entity in the entity's coordinate system.
///
/// \param shape Shape object
///
/// \return Local bounding rectangle of the entity
///
////////////////////////////////////////////////////////////
//  sfFloatRect sfCircleShape_getLocalBounds(const sfCircleShape* shape);

func (self CircleShape) Getlocalbounds() FloatRect { 
    return C.sfCircleShape_getLocalBounds();
}
            
////////////////////////////////////////////////////////////
/// \brief Get the global bounding rectangle of a circle shape
///
/// The returned rectangle is in global coordinates, which means
/// that it takes in account the transformations (translation,
/// rotation, scale, ...) that are applied to the entity.
/// In other words, this function returns the bounds of the
/// sprite in the global 2D world's coordinate system.
///
/// \param shape Shape object
///
/// \return Global bounding rectangle of the entity
///
////////////////////////////////////////////////////////////
//  sfFloatRect sfCircleShape_getGlobalBounds(const sfCircleShape* shape);

func (self CircleShape) Getglobalbounds() FloatRect { 
    return C.sfCircleShape_getGlobalBounds();
}
            
