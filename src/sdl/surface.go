package sdl

type SDL_Surface struct {
	Flags  uint32           /**< Read-only */
	Format *SDL_PixelFormat /**< Read-only */
	W      int32            /**< Read-only */
	H      int32            /**< Read-only */
	Pitch  int32            /**< Read-only */
	Pixels *byte            /**< Read-write */

	/** Application data associated with the surface */
	Userdata *byte /**< Read-write */

	/** information needed for surfaces requiring locks */
	Locked    int32 /**< Read-only */
	Lock_data *byte /**< Read-only */

	/** clipping information */
	Clip_rect SDL_Rect /**< Read-only */

	/** info for fast blit mapping to other surfaces */
	Map *SDL_BlitMap /**< Private */

	/** Reference count -- used when freeing surface */
	Refcount int32 /**< Read-mostly */

}
