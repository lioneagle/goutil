package sdl

/**
 *  \brief  The structure that defines a point
 *
 *  \sa SDL_EnclosePoints
 *  \sa SDL_PointInRect
 */
type SDL_Point struct {
	X int32
	Y int32
}

/**
 *  \brief A rectangle, with the origin at the upper left.
 *
 *  \sa SDL_RectEmpty
 *  \sa SDL_RectEquals
 *  \sa SDL_HasIntersection
 *  \sa SDL_IntersectRect
 *  \sa SDL_UnionRect
 *  \sa SDL_EnclosePoints
 */
type SDL_Rect struct {
	X, Y int32
	W, H int32
}
