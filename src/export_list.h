void RlezInit(void);
void RlezEnd(void);

int RlezGetReturnSize(void);
int RlezGetReturnTypeNameSize(void);
void RlezGetReturnType(void *p);
void RlezGetReturnData(void *p);

void RlezGroup(int id,const char *name);
void RlezDelete(int id);
void RlezDeleteGroup(const char *name);
void RlezDeleteAll(void);

void RlezLoadTempDir(char *return_pointer);
void RlezDeleteDir(const char *path);

void RlezOpenWindow(int width,int height);
void RlezSetWindowFocus(void);
int RlezCheckWindowFocus(void);
void RlezSetWindowSize(int width,int height);
void RlezSetWindowPosition(int x,int y);
int RlezGetWindowX(void);
int RlezGetWindowY(void);
void *RlezGetWindowHandle(void);
void RlezSetWindowIcon(int texture);
int RlezGetWindowWidth(void);
int RlezGetWindowHeight(void);
int RlezGetWidth(int resource);
int RlezGetHeight(int resource);
int RlezCheckWindowState(const char *flag);
void RlezSetWindowState(const char *flag,int value);
void RlezSetWindowTitle(const char *title);
int RlezCheckWindowClose(void);

int RlezGetFrameRate(void);
void RlezSetFrameRate(int fps);

int RlezGetDisplayCount(void);
int RlezGetCurrentDisplay(void);
int RlezGetDisplayX(int display);
int RlezGetDisplayY(int display);
int RlezGetDisplayWidth(int display);
int RlezGetDisplayHeight(int display);

void RlezGetBackend(char *return_pointer);
double RlezGetTime(void);

void RlezBeginDraw(int resource);
void RlezBeginWindow(void);
void RlezEndDraw(void);
void RlezBegin2D(double offset_x,double offset_y,double target_x,double target_y,double rotation,double zoom);
void RlezEnd2D(void);
void RlezBegin3D(double position_x,double position_y,double position_z,double target_x,double target_y,double target_z,double up_x,double up_y,double up_z,double fovy,int projection);
void RlezEnd3D(void);
void RlezBeginBlend(const char *mode);
void RlezBeginBlendCustom(const char *src_factor,const char *dest_factor,const char *equation);
void RlezBeginBlendCustomSeparate(const char *src_rgb,const char *dest_rgb,const char *src_alpha,const char *dest_alpha,const char *eq_rgb,const char *eq_alpha);
void RlezEndBlend(void);
void RlezBeginShader(int resource);
void RlezEndShader(void);

void RlezPerspective(double fovy,double aspect,double near,double far);
void RlezFrustum(double left,double right,double bottom,double top,double near,double far);
void RlezOrtho(double left,double right,double bottom,double top,double near,double far);
void RlezPushMatrix(void);
void RlezPopMatrix(void);
void RlezTranslate(double x,double y,double z);
void RlezScale(double x,double y,double z);
void RlezRotateAxis(double angle,double x,double y,double z);
void RlezRotateX(double angle);
void RlezRotateY(double angle);
void RlezRotateZ(double angle);
void RlezRotate(double angle);
void RlezLocalToWorld(double sx,double sy,double sz,double *dx,double *dy,double *dz);
void RlezWorldToLocal(double sx,double sy,double sz,double *dx,double *dy,double *dz);
void RlezWorldToScreen(double screen_x,double screen_y,double screen_w,double screen_h,double sx,double sy,double sz,double *dx,double *dy,double *dz,double *dw);

int RlezLoadRenderTexture(int width,int height);
int RlezLoadTexture(const char *path);
int RlezLoadTextureFromMemory(const char *file_type,void *data,int size);
void RlezSetTextureFilter(int texture,const char *filter_type);
void RlezSetTextureMipmaps(int texture);

int RlezLoadFont(const char *path,int font_size,const char *target_string,int target_image_width,int target_image_height);
int RlezLoadFontFromMemory(void *font_data,int font_data_size,int font_size,const char *target_string,int target_image_width,int target_image_height);

void RlezLoadPixels(int texture,const char *format);
void RlezCopyPixels(int src_offset,int src_length,void *dest_pointer,int dest_offset);
void RlezRestorePixels(void *src_pointer,int src_offset,int src_length,int dest_offset);
void RlezGetPixel(int x,int y);
void RlezSetPixel(int x,int y);
void RlezUpdatePixels(void);
void RlezSavePixels(const char *path);
void RlezUnloadPixels(void);

void RlezBackground(int r,int g,int b,int a);
void RlezBeginShape(const char *mode,int auto_normal,int resource);
void RlezVertex(double x,double y,double z,double u,double v,int r,int g,int b,int a,double nx,double ny,double nz);
void RlezEndShape();
int RlezEndMesh();
void RlezSetMeshTexture(int mesh,int texture);
int RlezLoadModel(const char *path,int load_animation);
void RlezFromMeshToModel(int mesh);
void RlezSetModelTexture(int model,int material_index,const char *map_name,int texture);
int RlezGetAnimationCount(int model);
int RlezGetAnimationId(int model,const char *name);
int RlezGetAnimationFrames(int model,int id);
void RlezSetModelAnimation(int model,int id,double frame);
void RlezSetModelAnimationBlend(int model,int a_id,double a_frame,int b_id,double b_frame,double blend);
void RlezColor(int r,int g,int b,int a);
int RlezGetColorR(void);
int RlezGetColorG(void);
int RlezGetColorB(void);
int RlezGetColorA(void);
void RlezDrawMesh(int mesh);
void RlezDrawModel(int model);
void RlezDrawText(int font,const char *text,double x,double y,double size,double spacing);
void RlezDrawLine(double x1,double y1,double z1,double x2,double y2,double z2);
void RlezDrawRect(double x,double y,double width,double height,int fill);
void RlezDrawEllipse(double x,double y,double width,double height,int fill);
void RlezDrawBox(double x,double y,double z,double width,double height,double depth,int fill);
void RlezDrawSphere(double x,double y,double z,double size,int rings,int slices,int fill);
void RlezDrawCylinder(double x1,double y1,double z1,double x2,double y2,double z2,double size1,double size2,int sides,int fill);
void RlezDrawCapsule(double x1,double y1,double z1,double x2,double y2,double z2,double size,int slices,int rings,int fill);
void RlezDrawTexture(int texture,int src_x,int src_y,int src_w,int src_h,double dest_x,double dest_y,double dest_w,double dest_h);

int RlezLoadShaderFromMemory(const char *vertex_code,const char *fragment_code);
int RlezLoadShader(const char *vertex_code_path,const char *fragment_code_path);

int RlezLoadSoundFromMemory(const char *file_type,void *sound_data,int sound_data_size,int is_music,double music_update_time,int music_update_samples);
int RlezLoadSound(const char *path,int is_music,double music_update_time,int music_update_samples);
void RlezStopSound(int sound);
void RlezPauseSound(int sound);
void RlezResumeSound(int sound);
void RlezPlaySound(int sound,int music_loop);
int RlezGetSoundStatus(int sound);
double RlezGetSoundTime(int sound);
void RlezSetSoundTime(int sound,double set_time);
double RlezGetSoundLength(int sound);
void RlezSetSoundPitch(int sound,double pitch);
void RlezSetSoundVolume(int sound,double volume);
void RlezSetSoundPan(int sound,double pan);

int RlezGetKey(const char *name);
int RlezGetMouseButton(const char *name);
int RlezGetMouseX(void);
int RlezGetMouseY(void);
void RlezSetMousePosition(int x,int y);
void RlezSetMouseVisible(int visible);
int RlezCheckMouseInWindow(void);
double RlezGetMouseWheelX(void);
double RlezGetMouseWheelY(void);
int RlezCheckGamepad(int gamepad);
int RlezGetGamepadButton(int gamepad,const char *button);
int RlezGetAxisCount(int gamepad);
double RlezGetGamepadAxis(int gamepad,int axis);
void RlezSetGamepadVibration(int gamepad,double left,double right,double duration);
