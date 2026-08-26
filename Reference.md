# rlez.dll - リファレンス

rlez.dll内の関数の定義と解説です。

## HSP3言語で使用する際の注意

HSP3言語では、各関数の名前は全て小文字のスネークケースに置き換えてください。返り値がない場合は「命令」として記述し、返り値がある場合は「関数」として記述してください。

```
// C言語での記述例
RlezBegin2D(0,0,0,0,0,1);
// HSP3言語での記述例
rlez_begin_2d 0,0,0,0,0,1

// C言語での記述例
int x=RlezGetMouseX();
// HSP3言語での記述例
x=rlez_get_mouse_x()
```

引数に文字列型(`const char*`)を除いたポインタ型がある場合は、varptr関数等で取得したアドレスではなく、変数を入力してください。

## 関数一覧

### rlez.dllの初期化と終了

| | |
| :---: | :---: |
| [RlezInit](#RlezInit) | rlez.dllの初期化 |
| [RlezEnd](#RlezEnd) | rlez.dllの終了 |
| | |

### リソース管理

| | |
| :---: | :---: |
| [RlezGroup](#RlezGroup) | リソースグループ名の設定 |
| [RlezDelete](#RlezDelete) | リソースの削除 |
| [RlezDeleteGroup](#RlezDeleteGroup) | リソースをリソースグループ名を指定してまとめて削除 |
| [RlezDeleteAll](#RlezDeleteAll) | リソースを全て削除 |
| | |

### テンポラリフォルダ関連

| | |
| :---: | :---: |
| [RlezLoadTempDir](#RlezLoadTempDir) | テンポラリフォルダの作成 |
| [RlezDeleteDir](#RlezDeleteDir) | フォルダを削除 |
| | |

### ウィンドウ関連

| | |
| :---: | :---: |
| [RlezOpenWindow](#RlezOpenWindow) | raylibのウィンドウを作成 |
| [RlezSetWindowFocus](#RlezSetWindowFocus) | raylibのウィンドウにフォーカスさせる |
| [RlezCheckWindowFocus](#RlezCheckWindowFocus) | raylibのウィンドウにフォーカスされているかを取得 |
| [RlezSetWindowSize](#RlezSetWindowSize) | raylibのウィンドウをリサイズ |
| [RlezSetWindowPosition](#RlezSetWindowPosition) | raylibのウィンドウの位置を設定 |
| [RlezGetWindowX](#RlezGetWindowX) | raylibのウィンドウのX座標を取得 |
| [RlezGetWindowY](#RlezGetWindowY) | raylibのウィンドウのY座標を取得 |
| [RlezGetWindowHandle](#RlezGetWindowHandle) | raylibのウィンドウのウィンドウハンドルを取得 |
| [RlezSetWindowIcon](#RlezSetWindowIcon) | raylibのウィンドウのアイコンを書き換える |
| [RlezGetWindowWidth](#RlezGetWindowWidth) | raylibのウィンドウの横幅のサイズを取得 |
| [RlezGetWindowHeight](#RlezGetWindowHeight) | raylibのウィンドウの縦幅のサイズを取得 |
| [RlezGetWidth](#RlezGetWidth) | テクスチャの横幅サイズを取得 |
| [RlezGetHeight](#RlezGetHeight) | テクスチャの縦幅サイズを取得 |
| [RlezCheckWindowState](#RlezCheckWindowState) | raylibのウィンドウのウィンドウステートを確認 |
| [RlezSetWindowState](#RlezSetWindowState) | raylibのウィンドウのウィンドウステートの設定をする |
| [RlezSetWindowTitle](#RlezSetWindowTitle) | raylibのウィンドウのタイトルを設定 |
| [RlezCheckWindowClose](#RlezCheckWindowClose) | raylibのウィンドウの閉じるボタンが押されたかを取得 |
| | |

### フレームレート関連

| | |
| :---: | :---: |
| [RlezGetFrameRate](#RlezGetFrameRate) | フレームレートを取得 |
| [RlezSetFrameRate](#RlezSetFrameRate) | フレームレートを設定 |
| | |

### ディスプレイ関連

| | |
| :---: | :---: |
| [RlezGetDisplayCount](#RlezGetDisplayCount) | ディスプレイの数を取得 |
| [RlezGetCurrentDisplay](#RlezGetCurrentDisplay) | 現在raylibのウィンドウが表示されているディスプレイを取得 |
| [RlezGetDisplayX](#RlezGetDisplayX) | ディスプレイのX座標を取得 |
| [RlezGetDisplayY](#RlezGetDisplayY) | ディスプレイのY座標を取得 |
| [RlezGetDisplayWidth](#RlezGetDisplayWidth) | ディスプレイの横幅サイズを取得 |
| [RlezGetDisplayHeight](#RlezGetDisplayHeight) | ディスプレイの縦幅サイズを取得 |
| | |

### システム関連

| | |
| :---: | :---: |
| [RlezGetBackend](#RlezGetBackend) | バックエンドレンダラー名を取得 |
| [RlezGetTime](#RlezGetTime) | raylibのウィンドウを作成してから現在までの時間を取得 |
| | |

### 開始と終了処理関連

| | |
| :---: | :---: |
| [RlezBeginDraw](#RlezBeginDraw) | テクスチャに描画開始 |
| [RlezBeginWindow](#RlezBeginWindow) | raylibのウィンドウに描画開始 |
| [RlezEndDraw](#RlezEndDraw) | 描画の終了 |
| [RlezBegin2D](#RlezBegin2D) | 2D描画の開始 |
| [RlezEnd2D](#RlezEnd2D) | 2D描画の終了 |
| [RlezBegin3D](#RlezBegin3D) | 3D描画の開始 |
| [RlezEnd3D](#RlezEnd3D) | 3D描画の終了 |
| [RlezBeginBlend](#RlezBeginBlend) | ブレンドモードを用いての描画を開始 |
| [RlezBeginBlendCustom](#RlezBeginBlendCustom) | ブレンドファクターを用いての描画を開始 |
| [RlezBeginBlendCustomSeparate](#RlezBeginBlendCustomSeparate) | ブレンドファクターを用いての描画を開始 |
| [RlezEndBlend](#RlezEndBlend) | ブレンドモードまたはブレンドファクターを用いての描画を終了 |
| [RlezBeginShader](#RlezBeginShader) | シェーダーを用いての描画を開始 |
| [RlezEndShader](#RlezEndShader) | シェーダーを用いての描画を終了 |
| | |

### 座標変換関連

| | |
| :---: | :---: |
| [RlezPerspective](#RlezPerspective) | 透視投影行列を透視投影図法に設定 |
| [RlezFrustum](#RlezFrustum) | 透視投影行列を透視投影図法に設定 |
| [RlezOrtho](#RlezOrtho) | 透視投影行列を平行投影図法に設定 |
| [RlezPushMatrix](#RlezPushMatrix) | 現在のモデル行列を記憶 |
| [RlezPopMatrix](#RlezPopMatrix) | 記憶したモデル行列を復元 |
| [RlezTranslate](#RlezTranslate) | 平行移動行列を適用 |
| [RlezScale](#RlezScale) | 拡大行列を適用 |
| [RlezRotateAxis](#RlezRotateAxis) | 任意軸の回転行列を適用 |
| [RlezRotateX](#RlezRotateX) | X軸の回転行列を適用 |
| [RlezRotateY](#RlezRotateY) | Y軸の回転行列を適用 |
| [RlezRotateZ](#RlezRotateZ) | Z軸の回転行列を適用 |
| [RlezRotate](#RlezRotate) | 2D回転行列を適用 |
| [RlezLocalToWorld](#RlezLocalToWorld) | ローカル座標からワールド座標へ変換 |
| [RlezWorldToLocal](#RlezWorldToLocal) | ワールド座標からローカル座標へ変換 |
| [RlezWorldToScreen](#RlezWorldToScreen) | ワールド座標からスクリーン座標へ変換 |
| | |

### テクスチャ関連

| | |
| :---: | :---: |
| [RlezLoadRenderTexture](#RlezLoadRenderTexture) | 空のテクスチャを作成 |
| [RlezLoadTexture](#RlezLoadTexture) | 画像ファイルからテクスチャを作成 |
| [RlezLoadTextureFromMemory](#RlezLoadTextureFromMemory) | メモリ上の画像ファイルからテクスチャを作成 |
| [RlezSetTextureFilter](#RlezSetTextureFilter) | テクスチャのフィルターの設定 |
| [RlezSetTextureMipmaps](#RlezSetTextureMipmaps) | テクスチャのミップマップを更新 |
| | |

### フォント関連

| | |
| :---: | :---: |
| [RlezLoadFont](#RlezLoadFont) | フォントファイルからフォントデータを作成 |
| [RlezLoadFontFromMemory](#RlezLoadFontFromMemory) | メモリ上のフォントファイルからフォントデータを作成 |
| | |

### ピクセル操作関連

| | |
| :---: | :---: |
| [RlezLoadPixels](#RlezLoadPixels) | ピクセル情報を取得 |
| [RlezCopyPixels](#RlezCopyPixels) | 取得したピクセル情報をコピー |
| [RlezRestorePixels](#RlezRestorePixels) | 取得したピクセル情報に書き込む |
| [RlezGetPixel](#RlezGetPixel) | 取得したピクセル情報からピクセルの色を取得 |
| [RlezSetPixel](#RlezSetPixel) | 取得したピクセル情報のピクセルの色を設定 |
| [RlezUpdatePixels](#RlezUpdatePixels) | 変更されたピクセル情報を反映 |
| [RlezSavePixels](#RlezSavePixels) | 取得したピクセル情報を画像ファイルに保存 |
| [RlezUnloadPixels](#RlezUnloadPixels) | 取得したピクセル情報を破棄 |
| | |

### 描画関連

| | |
| :---: | :---: |
| [RlezBackground](#RlezBackground) | 背景色を設定 |
| [RlezBeginShape](#RlezBeginShape) | 形状の描画を開始 |
| [RlezVertex](#RlezVertex) | 形状の頂点を設定 |
| [RlezEndShape](#RlezEndShape) | 形状の描画を実際に行う |
| [RlezEndMesh](#RlezEndMesh) | 作成済みの形状からメッシュを作成 |
| [RlezSetMeshTexture](#RlezSetMeshTexture) | メッシュのテクスチャを変更 |
| [RlezLoadModel](#RlezLoadModel) | 3Dモデルを読み込み |
| [RlezFromMeshToModel](#RlezFromMeshToModel) |メッシュから3Dモデルを作成 |
| [RlezSetModelTexture](#RlezSetModelTexture) | 3Dモデルのテクスチャを変更 |
| [RlezGetAnimationCount](#RlezGetAnimationCount) | 3Dモデルのアニメーションの数を取得 |
| [RlezGetAnimationId](#RlezGetAnimationId) | 3Dモデルのアニメーションのインデックス値を文字列から取得 |
| [RlezGetAnimationFrames](#RlezGetAnimationFrames) | 3Dモデルのアニメーションのキーフレーム数を取得 |
| [RlezSetModelAnimation](#RlezSetModelAnimation) | 3Dモデルにアニメーションを適用 |
| [RlezSetModelAnimationBlend](#RlezSetModelAnimationBlend) | 3Dモデルのアニメーション2つをブレンドして適用 |
| [RlezColor](#RlezColor) | 使用する色の設定 |
| [RlezGetColorR](#RlezGetColorR) | 現在の描画色の赤色の値を取得 |
| [RlezGetColorG](#RlezGetColorG) | 現在の描画色の緑色の値を取得 |
| [RlezGetColorB](#RlezGetColorB) | 現在の描画色の青色の値を取得 |
| [RlezGetColorA](#RlezGetColorA) | 現在の描画色のアルファ値を取得 |
| [RlezDrawMesh](#RlezDrawMesh) | メッシュを描画 |
| [RlezDrawModel](#RlezDrawModel) | 3Dモデルを描画 |
| [RlezDrawText](#RlezDrawText) | 文字列を描画 |
| [RlezDrawLine](#RlezDrawLine) | 線を描画 |
| [RlezDrawRect](#RlezDrawRect) | 長方形を描画 |
| [RlezDrawEllipse](#RlezDrawEllipse) | 楕円を描画 |
| [RlezDrawBox](#RlezDrawBox) | 直方体を描画 |
| [RlezDrawSphere](#RlezDrawSphere) | 球を描画 |
| [RlezDrawCylinder](#RlezDrawCylinder) | 円柱を描画 |
| [RlezDrawCapsule](#RlezDrawCapsule) | カプセルを描画 |
| [RlezDrawTexture](#RlezDrawTexture) | テクスチャを長方形に貼り付けて描画 |
| | |

### シェーダー関連

| | |
| :---: | :---: |
| [RlezLoadShaderFromMemory](#RlezLoadShaderFromMemory) | シェーダーをメモリ上の文字列から作成 |
| [RlezLoadShader](#RlezLoadShader) | シェーダーをファイルから作成 |
| | |

### サウンド関連

| | |
| :---: | :---: |
| [RlezLoadSoundFromMemory](#RlezLoadSoundFromMemory) | サウンドをメモリ上のバイナリデータから作成 |
| [RlezLoadSound](#RlezLoadSound) | サウンドをファイルから作成 |
| [RlezStopSound](#RlezStopSound) | サウンドを停止 |
| [RlezPauseSound](#RlezPauseSound) | サウンドを一時停止 |
| [RlezResumeSound](#RlezResumeSound) | サウンドの再生を再開 |
| [RlezPlaySound](#RlezPlaySound) | サウンドの再生 |
| [RlezGetSoundStatus](#RlezGetSoundStatus) | サウンドが再生中か取得 |
| [RlezGetSoundTime](#RlezGetSoundTime) | サウンドの再生位置を取得 |
| [RlezSetSoundTime](#RlezSetSoundTime) | サウンドの再生位置を設定 |
| [RlezGetSoundLength](#RlezGetSoundLength) | サウンドの総再生時間を取得 |
| [RlezSetSoundPitch](#RlezSetSoundPitch) | サウンドの再生スピードを設定 |
| [RlezSetSoundVolume](#RlezSetSoundVolume) | サウンドの音量を設定 |
| [RlezSetSoundPan](#RlezSetSoundPan) | サウンドのパンニングを設定 |
| | |

### 入力関連

| | |
| :---: | :---: |
| [RlezGetKey](#RlezGetKey) | キーボードのキーが押されているかを取得 |
| [RlezGetMouseButton](#RlezGetMouseButton) | マウスのボタンが押されているかを取得 |
| [RlezGetMouseX](#RlezGetMouseX) | マウスカーソルのX座標を取得 |
| [RlezGetMouseY](#RlezGetMouseY) | マウスカーソルのY座標を取得 |
| [RlezSetMousePosition](#RlezSetMousePosition) | マウスカーソルを移動させる |
| [RlezSetMouseVisible](#RlezSetMouseVisible) | マウスカーソルの表示設定 |
| [RlezCheckMouseInWindow](#RlezCheckMouseInWindow) | マウスカーソルがウィンドウ内にあるかを取得 |
| [RlezGetMouseWheelX](#RlezGetMouseWheelX) | マウスホイールのX値を取得 |
| [RlezGetMouseWheelY](#RlezGetMouseWheelY) | マウスホイールのY値を取得 |
| [RlezCheckGamepad](#RlezCheckGamepad) | ゲームパッドが接続されているかを取得 |
| [RlezGetGamepadButton](#RlezGetGamepadButton) | ゲームパッドのボタンが押されているかを取得 |
| [RlezGetAxisCount](#RlezGetAxisCount) | ゲームパッドの軸の数を取得 |
| [RlezGetGamepadAxis](#RlezGetGamepadAxis) | ゲームパッドの軸の値を取得 |
| [RlezSetGamepadVibration](#RlezSetGamepadVibration) | ゲームパッドを振動させる |
| | |

### HSP3言語でのみ使用できる命令
| | |
| :---: | :---: |
| [rlez_load_file_from_hsp](#rlez_load_file_from_hsp) | ファイルのデータを変数に書き込む |
| [rlez_load_texture_from_hsp](#rlez_load_texture_from_hsp) | ファイルからテクスチャを作成 |
| [rlez_load_font_from_hsp](#rlez_load_font_from_hsp) | ファイルからフォントを作成 |
| [rlez_load_sound_from_hsp](#rlez_load_sound_from_hsp) | ファイルからサウンドを作成 |
| | |

---

<div id="RlezInit"></div>

### RlezInit

```
void RlezInit(void);
```

rlez.dllの内部処理を初期化します。この関数は必ず１回は実行してください。

---

<div id="RlezEnd"></div>

### RlezEnd

```
void RlezEnd(void);
```

rlez.dllの内部処理を終了します。この関数はアプリケーションの終了時に必ず実行してください。この関数を実行した際には、内部で[RlezDeleteAll](#RlezDeleteAll)関数も実行されます。

---

<div id="RlezGroup"></div>

### RlezGroup

```
void RlezGroup(int id,const char *name);
```

`id`:リソースID

`name`:リソースグループ名の文字列

指定したリソースIDのリソースグループ名を設定します。リソースグループ名を設定することで、[RlezDeleteGroup](#RlezDeleteGroup)関数を使用した際に指定したリソースグループ名に一致するリソースIDをまとめて削除することができます。リソースグループ名のデフォルト値は空の文字列です。

---

<div id="RlezDelete"></div>

### RlezDelete

```
void RlezDelete(int id);
```

`id`:リソースID

指定したリソースIDを削除します。

---

<div id="RlezDeleteGroup"></div>

### RlezDeleteGroup

```
void RlezDeleteGroup(const char *name);
```

`name`:リソースグループ名

指定したリソースグループ名に一致するリソースIDを全て削除します。リソースグループ名については[RlezGroup](#RlezGroup)関数の説明をご覧ください。

---

<div id="RlezDeleteAll"></div>

### RlezDeleteAll

```
void RlezDeleteAll(void);
```

全てのリソースIDに対して[RlezDelete](#RlezDelete)関数を実行します。

---

<div id="RlezLoadTempDir"></div>

### RlezLoadTempDir

```
void RlezLoadTempDir(char *return_pointer);
```

`return_pointer`:作成されたテンポラリフォルダのパスの文字列を代入させる変数のポインタ

一時的に使うテンポラリフォルダを作成し、そのフォルダのパスを文字列で取得します。結果を代入させる変数のサイズは、わからなければあらかじめOSのファイルパスの最大文字数分のサイズを確保しておくと良いです。作成されたフォルダは自動的には削除されないので、必要がなくなったら[RlezDeleteDir](#RlezDeleteDir)関数を使用して削除してください。

---

<div id="RlezDeleteDir"></div>

### RlezDeleteDir

```
void RlezDeleteDir(const char *path);
```

`path`:削除するフォルダのパスの文字列

指定したフォルダとそのフォルダの内部のファイルやフォルダを全て削除します。[RlezLoadTempDir](#RlezLoadTempDir)関数で作成されたテンポラリフォルダが必要なくなった際等に使用すると便利です。

---

<div id="RlezOpenWindow"></div>

### RlezOpenWindow

```
void RlezOpenWindow(int width,int height);
```

`width`:作成するウィンドウの横幅のサイズ

`height`:作成するウィンドウの縦幅のサイズ

raylibのウィンドウを指定したサイズで作成します。`width`引数または`height`引数を0以下にすると、メインディスプレイのサイズに設定されます。

---

<div id="RlezSetWindowFocus"></div>

### RlezSetWindowFocus

```
void RlezSetWindowFocus(void);
```

raylibのウィンドウにフォーカスさせます。raylibのウィンドウに現在フォーカスされているかは[RlezCheckWindowFocus](#RlezCheckWindowFocus)関数で確認できます。

---

<div id="RlezCheckWindowFocus"></div>

```
int RlezCheckWindowFocus(void);
```

`(return)`:0だとフォーカスはされていない、1だとされている

raylibのウィンドウに現在フォーカスされているかを取得します。

---

<div id="RlezSetWindowSize"></div>

### RlezSetWindowSize

```
void RlezSetWindowSize(int width,int height);
```

`width`:設定するウィンドウの横幅のサイズ

`height`:設定するウィンドウの縦幅のサイズ

作成済みのraylibのウィンドウのサイズを設定します。

---

<div id="RlezSetWindowPosition"></div>

### RlezSetWindowPosition

```
void RlezSetWindowPosition(int x,int y);
```

`x`:設定するウィンドウのX座標

`y`:設定するウィンドウのY座標

作成済みのraylibのウィンドウの座標を設定します。

---

<div id="RlezGetWindowX"></div>

### RlezGetWindowX

```
int RlezGetWindowX(void);
```

`(return)`:現在のウィンドウのX座標

作成済みのraylibのウィンドウのX座標を取得します。

---

<div id="RlezGetWindowX"></div>

### RlezGetWindowX

```
int RlezGetWindowY(void);
```

`(return)`:現在のウィンドウのY座標

作成済みのraylibのウィンドウのY座標を取得します。

---

<div id="RlezGetWindowHandle"></div>

### RlezGetWindowHandle

```
void *RlezGetWindowHandle(void);
```

`(return)`:raylibのウィンドウのウィンドウハンドル

作成済みのraylibのウィンドウのウィンドウハンドルを取得します。

---

<div id="RlezSetWindowIcon"></div>

### RlezSetWindowIcon

```
void RlezSetWindowIcon(int texture);
```

`texture`:テクスチャのリソースID

作成済みのraylibのウィンドウのアイコンを、指定したテクスチャに書き換えます。

---

<div id="RlezGetWindowWidth"></div>

### RlezGetWindowWidth

```
int RlezGetWindowWidth(void);
```

`(return)`:現在のウィンドウの横幅のサイズ

作成済みのraylibのウィンドウの横幅のサイズを取得します。ピクセル単位のサイズを取得したい場合は[RlezGetWidth](#RlezGetWidth)関数の`resource`引数に`-1`を指定してください。

---

<div id="RlezGetWindowHeight"></div>

### RlezGetWindowHeight

```
int RlezGetWindowHeight(void);
```

`(return)`:現在のウィンドウの縦幅のサイズ

作成済みのraylibのウィンドウの縦幅のサイズを取得します。ピクセル単位のサイズを取得したい場合は[RlezGetHeight](#RlezGetHeight)関数の`resource`引数に`-1`を指定してください。

---

<div id="RlezGetWidth"></div>

### RlezGetWidth

```
int RlezGetWidth(int resource);
```

`resource`:テクスチャのリソースID

`(return)`:テクスチャの横幅のピクセルサイズ

指定したテクスチャの横幅のピクセルサイズを取得します。`resource`引数の値を`-1`以下にすると、作成済みのraylibのウィンドウの横幅のピクセルサイズを取得します。[RlezGetWindowWidth](#RlezGetWindowWidth)関数とは違う挙動をすることにご注意ください。

---

<div id="RlezGetHeight"></div>

### RlezGetHeight

```
int RlezGetHeight(int resource);
```

`resource`:テクスチャのリソースID

`(return)`:テクスチャの縦幅のピクセルサイズ

指定したテクスチャの縦幅のピクセルサイズを取得します。`resource`引数の値を`-1`以下にすると、作成済みのraylibのウィンドウの縦幅のピクセルサイズを取得します。[RlezGetWindowHeight](#RlezGetWindowHeight)関数とは違う挙動をすることにご注意ください。

---

<div id="RlezCheckWindowState"></div>

### RlezCheckWindowState

```
int RlezCheckWindowState(const char *flag);
```

`flag`:現在の状態を確認するウィンドウステート名

`(return)`:0だと無効を示し、1だと有効を示す

指定したウィンドウステートの現在の状態を取得します。指定できるウィンドウステート名は以下の通りです。

```
"FLAG_VSYNC_HINT"
"FLAG_FULLSCREEN_MODE"
"FLAG_WINDOW_RESIZABLE"
"FLAG_WINDOW_UNDECORATED"
"FLAG_WINDOW_HIDDEN"
"FLAG_WINDOW_MINIMIZED"
"FLAG_WINDOW_MAXIMIZED"
"FLAG_WINDOW_UNFOCUSED"
"FLAG_WINDOW_TOPMOST"
"FLAG_WINDOW_ALWAYS_RUN"
"FLAG_WINDOW_TRANSPARENT"
"FLAG_WINDOW_HIGHDPI"
"FLAG_WINDOW_MOUSE_PASSTHROUGH"
"FLAG_BORDERLESS_WINDOWED_MODE"
"FLAG_MSAA_4X_HINT"
"FLAG_INTERLACED_HINT"
```

各ウィンドウステート名の意味はraylibの「`raylib.h`」ファイル内をご覧ください。

---

<div id="RlezSetWindowState"></div>

### RlezSetWindowState

```
void RlezSetWindowState(const char *flag,int value);
```

`flag`:設定するウィンドウステート名

`value`:0だと無効化、0以外だと有効化

指定したウィンドウステートを設定します。`flag`引数に設定できるウィンドウステート名については[RlezCheckWindowState](#RlezCheckWindowState)関数の説明をご覧ください。一部のウィンドウステートは[RlezOpenWindow](#RlezOpenWindow)関数を呼び出す前にも設定できます。

---

<div id="RlezSetWindowTitle"></div>

### RlezSetWindowTitle

```
void RlezSetWindowTitle(const char *title);
```

`title`:ウィンドウのタイトルに設定する文字列

作成済みのraylibのウィンドウのタイトルの文字列を設定します。

---

<div id="RlezCheckWindowClose"></div>

### RlezCheckWindowClose

```
int RlezCheckWindowClose(void);
```

`(return)`:0だと閉じるボタンが押されていない、1だと閉じるボタンが押されている

作成済みのraylibのウィンドウの閉じるボタンが押されたかを取得します。

---

<div id="RlezGetFrameRate"></div>

### RlezGetFrameRate

```
int RlezGetFrameRate(void);
```

`(return)`:現在のフレームレート

現在のフレームレートを取得します。

---

<div id="RlezSetFrameRate"></div>

### RlezSetFrameRate

```
void RlezSetFrameRate(int fps);
```

`fps`:目標フレームレート

目標とするフレームレートを設定します。

---

<div id="RlezGetDisplayCount"></div>

### RlezGetDisplayCount

```
int RlezGetDisplayCount(void);
```

`(return)`:ディスプレイの数

現在有効なディスプレイの数を取得します。

---

<div id="RlezGetCurrentDisplay"></div>

### RlezGetCurrentDisplay

```
int RlezGetCurrentDisplay(void);
```

`(return)`:ディスプレイID

現在raylibのウィンドウが表示されているディスプレイのIDを取得します。

---

<div id="RlezGetDisplayX"></div>

### RlezGetDisplayX

```
int RlezGetDisplayX(int display);
```

`display`:情報を取得するディスプレイID

`(return)`:ディスプレイのX座標

指定したディスプレイのX座標を取得します。

---

<div id="RlezGetDisplayY"></div>

### RlezGetDisplayY

```
int RlezGetDisplayY(int display);
```

`display`:情報を取得するディスプレイID

`(return)`:ディスプレイのY座標

指定したディスプレイのY座標を取得します。

---

<div id="RlezGetDisplayWidth"></div>

### RlezGetDisplayWidth

```
int RlezGetDisplayWidth(int display);
```

`display`:情報を取得するディスプレイID

`(return)`:ディスプレイの横幅のサイズ

指定したディスプレイの横幅のサイズを取得します。

---

<div id="RlezGetDisplayHeight"></div>

### RlezGetDisplayHeight

```
int RlezGetDisplayHeight(int display);
```

`display`:情報を取得するディスプレイID

`(return)`:ディスプレイの縦幅のサイズ

指定したディスプレイの縦幅のサイズを取得します。

---

<div id="RlezGetBackend"></div>

### RlezGetBackend

```
void RlezGetBackend(char *return_pointer);
```

`return_pointer`:バックエンド名を代入させる文字列型変数のポインタ

現在raylibの機能が使用しているレンダラーのバックエンド名を取得します。`return_pointer`引数で指定された変数には、レンダラーのバックエンドが`OpenGL 3.3`の場合は「`"opengl33"`」、`ANGLE`の場合は「`"angle"`」という文字列が代入されます。`return_pointer`引数で指定する変数のサイズは、わからなければ32文字分確保しておくと良いです。

---

<div id="RlezGetTime"></div>

### RlezGetTime

```
double RlezGetTime(void);
```

`(return)`:[RlezOpenWindow](#RlezOpenWindow)関数を実行してからの時間

[RlezOpenWindow](#RlezOpenWindow)関数を実行してからの時間を秒単位で取得します。

---

<div id="RlezBeginDraw"></div>

### RlezBeginDraw

```
void RlezBeginDraw(int resource);
```

`resource`:テクスチャのリソースID

指定したテクスチャに対して描画を開始します。テクスチャへの描画が完了したら必ず[RlezEndDraw](#RlezEndDraw)関数を実行してください。

---

<div id="RlezBeginWindow"></div>

### RlezBeginWindow

```
void RlezBeginWindow(void);
```

作成済みのraylibのウィンドウに描画を開始します。raylibのウィンドウへの描画が完了したら必ず[RlezEndDraw](#RlezEndDraw)関数を実行してください。さらに、このRlezBeginWindow関数と[RlezEndDraw](#RlezEndDraw)関数の組み合わせは必ず毎フレーム実行する必要があります。もしこの組み合わせを実行しなかったら、その間raylibのウィンドウの処理はフリーズしているように見えるでしょう。

---

<div id="RlezEndDraw"></div>

### RlezEndDraw

```
void RlezEndDraw(void);
```

テクスチャやraylibのウィンドウへの描画を完了させます。[RlezBeginDraw](#RlezBeginDraw)関数や[RlezBeginWindow](#RlezBeginWindow)関数で描画を開始した後、描画を完了させるために必ずこの関数を実行する必要があります。

---

<div id="RlezBegin2D"></div>

### RlezBegin2D

```
void RlezBegin2D(double offset_x,double offset_y,double target_x,double target_y,double rotation,double zoom);
```

`offset_x`,`offset_y`:カメラの基準点の座標

`target_x`,`target_y`:カメラの注視点の座標

`rotation`:カメラの角度(度単位)

`zoom`:カメラの拡大値(1.0で等倍)

2D描画を開始します。モデル行列は初期化され、ビュー行列と透視投影行列が設定されます。2D描画を終了する際は必ず[RlezEnd2D](#RlezEnd2D)関数を実行してください。この関数は[RlezBeginDraw](#RlezBeginDraw)関数または[RlezBeginWindow](#RlezBeginWindow)関数から[RlezEndDraw](#RlezEndDraw)関数の間に記述してください。

---

<div id="RlezEnd2D"></div>

### RlezEnd2D

```
void RlezEnd2D(void);
```

[RlezBegin2D](#RlezBegin2D)関数で開始した2D描画を終了します。この関数は[RlezBeginDraw](#RlezBeginDraw)関数または[RlezBeginWindow](#RlezBeginWindow)関数から[RlezEndDraw](#RlezEndDraw)関数の間に記述してください。

---

<div id="RlezBegin3D"></div>

```
void RlezBegin3D(double position_x,double position_y,double position_z,double target_x,double target_y,double target_z,double up_x,double up_y,double up_z,double fovy,int projection);
```

`position_x`,`position_y`,`position_z`:カメラの位置の座標

`target_x`,`target_y`,`target_z`:カメラの注視点の座標

`up_x`,`up_y`,`up_z`:カメラの上方向とするベクトル(長さは1にすることをおすすめします)

`fovy`:カメラの縦の視野角(度単位)

`projection`:0にすると透視投影、0以外だと平行投影

3D描画を開始します。モデル行列は初期化され、ビュー行列と透視投影行列が設定されます。3D描画を終了する際は必ず[RlezEnd3D](#RlezEnd3D)関数を実行してください。この関数は[RlezBeginDraw](#RlezBeginDraw)関数または[RlezBeginWindow](#RlezBeginWindow)関数から[RlezEndDraw](#RlezEndDraw)関数の間に記述してください。

---

<div id="RlezEnd3D"></div>

### RlezEnd3D

```
void RlezEnd3D(void);
```

[RlezBegin3D](#RlezBegin3D)関数で開始した3D描画を終了します。この関数は[RlezBeginDraw](#RlezBeginDraw)関数または[RlezBeginWindow](#RlezBeginWindow)関数から[RlezEndDraw](#RlezEndDraw)関数の間に記述してください。

---

<div id="RlezBeginBlend"></div>

### RlezBeginBlend

```
void RlezBeginBlend(const char *mode);
```

`mode`:ブレンドモードを記述した文字列

指定したブレンドモードでの描画を開始します。指定できるブレンドモードは以下の通りです。

```
"BLEND_ALPHA"
"BLEND_ADDITIVE"
"BLEND_MULTIPLIED"
"BLEND_ADD_COLORS"
"BLEND_SUBTRACT_COLORS"
"BLEND_ALPHA_PREMULTIPLY"
```

各ブレンドモードの意味はraylibの「`raylib.h`」ファイル内をご覧ください。ブレンドモードでの描画を終了する際は必ず[RlezEndBlend](#RlezEndBlend)関数を実行してください。この関数は[RlezBeginDraw](#RlezBeginDraw)関数または[RlezBeginWindow](#RlezBeginWindow)関数から[RlezEndDraw](#RlezEndDraw)関数の間に記述してください。

---

<div id="RlezBeginBlendCustom"></div>

### RlezBeginBlendCustom

```
void RlezBeginBlendCustom(const char *src_factor,const char *dest_factor,const char *equation);
```

`src_factor`:描画元のブレンドファクター名の文字列

`dest_factor`:描画先のブレンドファクター名の文字列

`equation`:合成方法の名前の文字列

指定したブレンドファクターを用いて描画を開始します。ブレンドファクター名の引数に指定できるブレンドファクター名は以下の通りです。

```
"ZERO"
"ONE"
"SRC_COLOR"
"ONE_MINUS_SRC_COLOR"
"SRC_ALPHA"
"ONE_MINUS_SRC_ALPHA"
"DST_ALPHA"
"ONE_MINUS_DST_ALPHA"
"DST_COLOR"
"ONE_MINUS_DST_COLOR"
"SRC_ALPHA_SATURATE"
"CONSTANT_COLOR"
"ONE_MINUS_CONSTANT_COLOR"
"CONSTANT_ALPHA"
"ONE_MINUS_CONSTANT_ALPHA"
```

合成方法の名前の引数に指定できる合成方法の名前は以下の通りです。

```
"FUNC_ADD"
"MIN"
"MAX"
"FUNC_SUBTRACT"
"FUNC_REVERSE_SUBTRACT"
"BLEND_EQUATION"
"BLEND_EQUATION_RGB"
"BLEND_EQUATION_ALPHA"
"BLEND_DST_RGB"
"BLEND_SRC_RGB"
"BLEND_DST_ALPHA"
"BLEND_SRC_ALPHA"
"BLEND_COLOR"
```

各名前の意味はraylibの「`rlgl.h`」ファイルの内容をご覧ください。例えば以下の様に入力すると、図形等を描画した際に描画元の色が描画先の色を完全に上書きするようになります。

```
RlezBeginBlendCustom("ONE","ZERO","FUNC_ADD")
```

ブレンドファクターを用いての描画を終了する際は必ず[RlezEndBlend](#RlezEndBlend)関数を実行してください。この関数は[RlezBeginDraw](#RlezBeginDraw)関数または[RlezBeginWindow](#RlezBeginWindow)関数から[RlezEndDraw](#RlezEndDraw)関数の間に記述してください。

---

<div id="RlezBeginBlendCustomSeparate"></div>

### RlezBeginBlendCustomSeparate

```
void RlezBeginBlendCustomSeparate(const char *src_rgb,const char *dest_rgb,const char *src_alpha,const char *dest_alpha,const char *eq_rgb,const char *eq_alpha);
```

`src_rgb`,`dest_rgb`:描画元と描画先のRGB値のブレンドファクター名の文字列

`src_alpha`,`dest_alpha`:描画元と描画先のアルファ値のブレンドファクター名の文字列

`eq_rgb`,`eq_alpha`:RGB値とアルファ値の合成方法の名前の文字列

指定したブレンドファクターを用いて描画を開始します。[RlezBeginBlendCustom](#RlezBeginBlendCustom)関数との違いは、ブレンドファクターをRGB値とアルファ値それぞれに設定できるようになった点です。指定できるブレンドファクター名と合成方法の名前は[RlezBeginBlendCustom](#RlezBeginBlendCustom)関数の説明をご覧ください。ブレンドファクターを用いての描画を終了する際は必ず[RlezEndBlend](#RlezEndBlend)関数を実行してください。この関数は[RlezBeginDraw](#RlezBeginDraw)関数または[RlezBeginWindow](#RlezBeginWindow)関数から[RlezEndDraw](#RlezEndDraw)関数の間に記述してください。

---

<div id="RlezEndBlend"></div>

### RlezEndBlend

```
void RlezEndBlend(void);
```

[RlezBeginBlend](#RlezBeginBlend)関数や[RlezBeginBlendCustom](#RlezBeginBlendCustom)関数と[RlezBeginBlendCustomSeparate](#RlezBeginBlendCustomSeparate)関数で開始したブレンドモードやブレンドファクターでの描画を終了します。この関数は[RlezBeginDraw](#RlezBeginDraw)関数または[RlezBeginWindow](#RlezBeginWindow)関数から[RlezEndDraw](#RlezEndDraw)関数の間に記述してください。

---

<div id="RlezBeginShader"></div>

### RlezBeginShader

```
void RlezBeginShader(int resource);
```

`resource`:シェーダーのリソースID

指定したシェーダーを使用する描画を開始します。指定したシェーダーを使用する描画を終了する際は必ず[RlezEndShader](#RlezEndShader)関数を実行してください。この関数は[RlezBeginDraw](#RlezBeginDraw)関数または[RlezBeginWindow](#RlezBeginWindow)関数から[RlezEndDraw](#RlezEndDraw)関数の間に記述してください。

---

<div id="RlezEndShader"></div>

### RlezEndShader

```
void RlezEndShader(void);
```

[RlezBeginShader](#RlezBeginShader)関数で指定したシェーダーを使用する描画を終了します。この関数は[RlezBeginDraw](#RlezBeginDraw)関数または[RlezBeginWindow](#RlezBeginWindow)関数から[RlezEndDraw](#RlezEndDraw)関数の間に記述してください。

---

<div id="RlezPerspective"></div>

### RlezPerspective

```
void RlezPerspective(double fovy,double aspect,double near,double far);
```

`fovy`:縦の視野角(度単位)

`aspect`:アスペクト比

`near`:クリッピング面の手前の距離

`far`:クリッピング面の奥の距離

現在の透視投影行列を視野角とアスペクト比を用いて透視投影図法に書き換えます。物体の位置がカメラから遠くなるほど物体が小さく見えます。

---

<div id="RlezFrustum"></div>

### RlezFrustum

```
void RlezFrustum(double left,double right,double bottom,double top,double near,double far);
```

`left`:手前のクリッピング面の左の位置

`right`:手前のクリッピング面の右の位置

`bottom`:手前のクリッピング面の下の位置

`top`:手前のクリッピング面の上の位置

`near`:クリッピング面の手前の距離

`far`:クリッピング面の奥の距離

現在の透視投影行列を手前のクリッピング面の位置の設定を用いて透視投影図法に書き換えます。物体の位置がカメラから遠くなるほど物体が小さく見えます。

---

<div id="RlezOrtho"></div>

### RlezOrtho

```
void RlezOrtho(double left,double right,double bottom,double top,double near,double far);
```

`left`:クリッピング面の左の位置

`right`:クリッピング面の右の位置

`bottom`:クリッピング面の下の位置

`top`:クリッピング面の上の位置

`near`:クリッピング面の手前の距離

`far`:クリッピング面の奥の距離

現在の透視投影行列を平行投影図法に書き換えます。物体の位置がカメラから遠くなっても物体の大きさは変わりません。

---

<div id="RlezPushMatrix"></div>

### RlezPushMatrix

```
void RlezPushMatrix(void);
```

現在のモデル行列を記憶します。この関数の後に[RlezPopMatrix](#RlezPopMatrix)関数を実行することで、記憶したモデル行列を復元できます。RlezPushMatrix関数を実行した後はRlezPushMatrix関数を実行した回数と同じ回数分[RlezPopMatrix](#RlezPopMatrix)関数を実行するようにしてください。raylibの仕様上モデル行列を記憶できる回数には制限があるためです。

---

<div id="RlezPopMatrix"></div>

### RlezPopMatrix

```
void RlezPopMatrix(void);
```

[RlezPushMatrix](#RlezPushMatrix)関数で記憶したモデル行列を現在のモデル行列に復元した後破棄します。[RlezPushMatrix](#RlezPushMatrix)関数で何もモデル行列を記憶していない状態でRlezPopMatrix関数を呼び出さないようにしてください。

---

<div id="RlezTranslate"></div>

### RlezTranslate

```
void RlezTranslate(double x,double y,double z);
```

`x`,`y`,`z`:平行移動する距離

平行移動行列を作成し、現在のモデル行列に掛けます。描画の基準点の座標が移動します。

---

<div id="RlezScale"></div>

### RlezScale

```
void RlezScale(double x,double y,double z);
```

`x`,`y`,`z`:拡大率(それぞれ1.0で等倍)

拡大行列を作成し、現在のモデル行列に掛けます。描画の基準点からの拡大率が乗算されます。

---

<div id="RlezRotateAxis"></div>

### RlezRotateAxis

```
void RlezRotateAxis(double angle,double x,double y,double z);
```

`angle`:回転させる角度(度単位)

`x`,`y`,`z`:回転させる軸

指定した角度と軸から回転行列を作成し、現在のモデル行列に掛けます。描画の基準点が回転します。

---

<div id="RlezRotateX"></div>

### RlezRotateX

```
void RlezRotateX(double angle);
```

`angle`:回転させる角度(度単位)

描画の基準点のX軸を回転させます。行っている処理は以下と同等です。

```
RlezRotateAxis(angle,1,0,0);
```

---

<div id="RlezRotateY"></div>

### RlezRotateY

```
void RlezRotateY(double angle);
```

`angle`:回転させる角度(度単位)

描画の基準点のY軸を回転させます。行っている処理は以下と同等です。

```
RlezRotateAxis(angle,1,0,0);
```

---

<div id="RlezRotateZ"></div>

### RlezRotateZ

```
void RlezRotateZ(double angle);
```

`angle`:回転させる角度(度単位)

描画の基準点のZ軸を回転させます。行っている処理は以下と同等です。

```
RlezRotateAxis(angle,1,0,0);
```

---

<div id="RlezRotate"></div>

### RlezRotate

```
void RlezRotate(double angle);
```

`angle`:回転させる角度(度単位)

2D回転させます。行っている処理は[RlezRotateZ](#RlezRotateZ)関数と同じです。

---

<div id="RlezLocalToWorld"></div>

### RlezLocalToWorld

```
void RlezLocalToWorld(double sx,double sy,double sz,double *dx,double *dy,double *dz);
```

`sx`,`sy`,`sz`:変換前の座標(ローカル座標)

`dx`,`dy`,`dz`:変換後の座標(ワールド座標)を代入させる変数へのポインタ

現在のモデル行列を基に、ローカル座標からワールド座標に変換します。

---

<div id="RlezWorldToLocal"></div>

### RlezWorldToLocal

```
void RlezWorldToLocal(double sx,double sy,double sz,double *dx,double *dy,double *dz);
```

`sx`,`sy`,`sz`:変換前の座標(ワールド座標)

`dx`,`dy`,`dz`:変換後の座標(ローカル座標)を代入させる変数へのポインタ

現在のモデル行列を基に、ワールド座標からローカル座標に変換します。

---

<div id="RlezWorldToScreen"></div>

### RlezWorldToScreen

```
void RlezWorldToScreen(double screen_x,double screen_y,double screen_w,double screen_h,double sx,double sy,double sz,double *dx,double *dy,double *dz,double *dw);
```

`screen_x`,`screen_y`:スクリーンの中心の座標

`screen_w`,`screen_h`:スクリーンの横幅と縦幅

`sx`,`sy`,`sz`:変換前の座標(ワールド座標)

`dx`,`dy`,`dz`,`dw`:変換後の座標(スクリーン座標)を代入させる変数へのポインタ

現在のビュー行列と透視投影行列を基に、ワールド座標からスクリーン座標に変換します。

---

<div id="RlezLoadRenderTexture"></div>

### RlezLoadRenderTexture

```
int RlezLoadRenderTexture(int width,int height);
```

`width`:テクスチャの横幅のサイズ

`height`:テクスチャの縦幅のサイズ

`(return)`:作成されたテクスチャのリソースID

空のテクスチャを作成します。rlez.dllでは、レンダーテクスチャもテクスチャとして扱います。

---

<div id="RlezLoadTexture"></div>

### RlezLoadTexture

```
int RlezLoadTexture(const char *path);
```

`path`:画像ファイルへのパスの文字列

`(return)`:作成されたテクスチャのリソースID

画像ファイルからテクスチャを作成します。

---

<div id="RlezLoadTextureFromMemory"></div>

### RlezLoadTextureFromMemory

```
int RlezLoadTextureFromMemory(const char *file_type,void *data,int size);
```

`file_type`:画像ファイルの拡張子の文字列(例えば`".png"`)(ドットを含めてください)

`data`:画像ファイルのバイナリデータが書き込まれている変数へのポインタ

`size`:画像ファイルのバイナリデータのサイズ

`(return)`:作成されたテクスチャのリソースID

メモリ上の画像ファイルのバイナリデータからテクスチャを作成します。

---

<div id="RlezSetTextureFilter"></div>

### RlezSetTextureFilter

```
void RlezSetTextureFilter(int texture,const char *filter_type);
```

`texture`:テクスチャのリソースID

`filter_type`:フィルターのタイプ名の文字列

指定したテクスチャを拡大や縮小をして描画する際のフィルターの設定を行います。`filter_type`引数に設定できるフィルターのタイプ名は以下の通りです。

```
"TEXTURE_FILTER_POINT"
"TEXTURE_FILTER_BILINEAR"
"TEXTURE_FILTER_TRILINEAR"(*)
"TEXTURE_FILTER_ANISOTROPIC_4X"(*)
"TEXTURE_FILTER_ANISOTROPIC_8X"(*)
"TEXTURE_FILTER_ANISOTROPIC_16X"(*)
```

上記のうち、「`(*)`」がついているタイプ名を使用する際は、テクスチャを読み込んだ後や書き換えた後に[RlezSetTextureMipmaps](#RlezSetTextureMipmaps)関数を実行してミップマップを更新する必要があります。

---

<div id="RlezSetTextureMipmaps"></div>

### RlezSetTextureMipmaps

```
void RlezSetTextureMipmaps(int texture);
```

`texture`:テクスチャのリソースID

指定したテクスチャのミップマップを更新します。[RlezSetTextureFilter](#RlezSetTextureFilter)関数の説明に記述されている一部のフィルターを使用しているテクスチャを読み込んだ後や書き換えた後にこの関数を実行する必要があります。

---

<div id="RlezLoadFont"></div>

### RlezLoadFont

```
int RlezLoadFont(const char *path,int font_size,const char *target_string,int target_image_width,int target_image_height);
```

`path`:フォントファイルのパスの文字列

`font_size`:生成されるフォントデータ内のフォントテクスチャに描画されるフォントサイズ

`target_string`:生成されるフォントデータ内のフォントテクスチャに描画される文字群の文字列

`target_image_width`,`target_image_height`:内部で生成されるフォントテクスチャの目標サイズの横幅と縦幅

`(return)`:フォントデータのリソースID

フォントファイルからフォントデータを作成します。フォントデータの内部で作成されるフォントテクスチャに`target_string`引数に指定された文字群を描画します。[RlezDrawText](#RlezDrawText)関数で文字列を描画する際、そのフォントテクスチャを使用して描画します。`target_strint`引数を空「`""`」に設定すると、下記の通りの基本的な英数字に設定されます。

    ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890\"!`?'.,;:()[]{}<>|/@\\^$-%+=#_&~*

フォントテクスチャ作成時、大まかに以下の処理が行われます。

1. `target_image_width`引数と`target_image_height`引数で指定した目標フォントテクスチャサイズを基に、大雑把に`target_string`引数の一部の文字を切り取る。
2. 切り取った`target_string`引数の一部の文字を使ってフォントテクスチャを作成する。
3. 「2.」で作成したフォントテクスチャのサイズが`target_image_width`引数と`target_image_height`引数以上だったら「5.」に進む。そうでない場合は次へ。
4. 切り取った`target_string`引数の一部に残りの文字を1文字ずつ足して「2.」に戻る。
5. 切り取った`target_string`引数の一部を使ったフォントテクスチャの作成を完了し、フォントデータに登録する。もし`target_string`引数の文字の中でフォントテクスチャに描画していない文字があったらその残りの文字の一部を切り取って「2.」に戻る。もし残りの文字が空だった場合、次へ。
6. フォントデータの作成を終了する。

---

<div id="RlezLoadFontFromMemory"></div>

### RlezLoadFontFromMemory

```
int RlezLoadFontFromMemory(void *font_data,int font_data_size,int font_size,const char *target_string,int target_image_width,int target_image_height);
```

`font_data`:フォントファイルのバイナリデータが入った変数へのポインタ

`font_data_size`:フォントファイルのサイズ

`font_size`:生成されるフォントデータ内のフォントテクスチャに描画されるフォントサイズ

`target_string`:生成されるフォントデータ内のフォントテクスチャに描画される文字群の文字列

`target_image_width`,`target_image_height`:内部で生成されるフォントテクスチャの目標サイズの横幅と縦幅

`(return)`:フォントデータのリソースID

メモリ上のフォントファイルのバイナリデータからフォントデータを作成します。`font_size`引数と`target_string`引数と`target_image_width`引数と`target_image_height`引数に関しては[RlezLoadFont](#RlezLoadFont)関数と同じです。フォントデータを作成する過程については[RlezLoadFont](#RlezLoadFont)関数の説明をご覧ください。

---

<div id="RlezLoadPixels"></div>

### RlezLoadPixels

```
void RlezLoadPixels(int texture,const char *format);
```

`texture`:テクスチャのリソースID

`format`:ピクセルデータのフォーマット名の文字列

指定したテクスチャのピクセル情報を取得します。この関数を実行する際、内部では[RlezUnloadPixels](#RlezUnloadPixels)関数が最初に実行されます。ピクセル情報を取得した後、必要がなくなったら[RlezUnloadPixels](#RlezUnloadPixels)関数を実行してください。`format`引数で指定できるフォーマット名は以下の通りです。

```
"RGBA"(1ピクセルあたり4バイト)
"RGB"(1ピクセルあたり3バイト)
"BGR"(1ピクセルあたり3バイト)
```

取得したピクセル情報のサイズの計算は`テクスチャの横幅*テクスチャの縦幅*1ピクセルあたりのバイト数`になります。

---

<div id="RlezCopyPixels"></div>

### RlezCopyPixels

```
void RlezCopyPixels(int src_offset,int src_length,void *dest_pointer,int dest_offset);
```

`src_offset`:読み込み元のピクセル情報の読み込み開始位置(バイト単位)

`src_length`:読み込み元のピクセル情報を読み込む長さ(バイト単位)

`dest_pointer`:ピクセル情報の書き込み先の変数のポインタ

`dest_offset`:ピクセル情報の書き込み先の変数の書き込み開始位置

[RlezLoadPixels](#RlezLoadPixels)関数で取得したピクセル情報を指定した変数にコピーします。読み込み元のピクセル情報から書き込み先の変数に書き込む際、オーバーフローしないように気を付けてください。

---

<div id="RlezRestorePixels"></div>

### RlezRestorePixels

```
void RlezRestorePixels(void *src_pointer,int src_offset,int src_length,int dest_offset);
```

`src_pointer`:読み込み元のピクセル情報の変数のポインタ

`src_offset`:読み込み元のピクセル情報の読み込み開始位置(バイト単位)

`src_length`:読み込み元のピクセル情報を読み込む長さ(バイト単位)

`dest_offset`:ピクセル情報の書き込み先の書き込み開始位置

[RlezLoadPixels](#RlezLoadPixels)関数で取得したピクセル情報のデータを書き換えます。読み込み元のピクセル情報の変数から書き込み先のピクセル情報のデータに書き込む際、オーバーフローしないように気を付けてください。ピクセル情報を書き換えたら、[RlezUpdatePixels](#RlezUpdatePixels)関数を実行するまで元のテクスチャにピクセル情報は反映されません。

---

<div id="RlezGetPixel"></div>

### RlezGetPixel

```
void RlezGetPixel(int x,int y);
```

`x`:ピクセル情報のデータのピクセルのX座標

`y`:ピクセル情報のデータのピクセルのY座標

[RlezLoadPixels](#RlezLoadPixels)関数で取得したピクセル情報のデータから指定した座標のピクセルの色を取得します。取得した色情報は、[RlezGetColorR](#RlezGetColorR)関数、[RlezGetColorG](#RlezGetColorG)関数、[RlezGetColorB](#RlezGetColorB)関数、[RlezGetColorA](#RlezGetColorA)関数で取得できます。つまり、[RlezColor](#RlezColor)関数の設定が書き換えられます。

---

<div id="RlezSetPixel"></div>

### RlezSetPixel

```
void RlezSetPixel(int x,int y);
```

`x`:ピクセル情報のデータのピクセルのX座標

`y`:ピクセル情報のデータのピクセルのY座標

[RlezLoadPixels](#RlezLoadPixels)関数で取得したピクセル情報のデータの指定した座標のピクセルの色を、現在[RlezColor](#RlezColor)関数で設定されている色に書き換えます。ピクセル情報を書き換えたら、[RlezUpdatePixels](#RlezUpdatePixels)関数を実行するまで元のテクスチャにピクセル情報は反映されません。

---

<div id="RlezUpdatePixels"></div>

### RlezUpdatePixels

```
void RlezUpdatePixels(void);
```

[RlezLoadPixels](#RlezLoadPixels)関数で取得したピクセル情報のデータを[RlezRestorePixels](#RlezRestorePixels)関数等で書き換えた後にこの関数を実行することで、元のテクスチャにピクセル情報を反映させます。

---

<div id="RlezSavePixels"></div>

### RlezSavePixels

```
void RlezSavePixels(const char *path);
```

`path`:保存先の画像ファイルのパスの文字列

[RlezLoadPixels](#RlezLoadPixels)関数で取得したピクセル情報のデータを指定した画像ファイルのパスに保存します。`path`引数には「`"image.png"`」のように拡張子まで入力してください。

---

<div id="RlezUnloadPixels"></div>

### RlezUnloadPixels

```
void RlezUnloadPixels(void);
```

[RlezLoadPixels](#RlezLoadPixels)関数で取得したピクセル情報のデータを破棄します。[RlezLoadPixels](#RlezLoadPixels)関数でピクセル情報を取得した後、必要がなくなったらこの関数を実行してください。

---

<div id="RlezBackground"></div>

### RlezBackground

```
void RlezBackground(int r,int g,int b,int a);
```

`r`:赤色の値(0～255)

`g`:緑色の値(0～255)

`b`:青色の値(0～255)

`a`:透明度の値(0～255)

現在の描画先のピクセル情報を全て削除し、指定された色で全体を塗りつぶします。

---

<div id="RlezBeginShape"></div>

### RlezBeginShape

```
void RlezBeginShape(const char *mode,int auto_normal,int resource);
```

`mode`:描画する形状の名前

`auto_normal`:法線ベクトルを自動設定するか(0だと手動、0以外だと自動)

`resource`:形状のテクスチャにするテクスチャのリソースID(0未満だと無効化)

形状の描画を開始します。形状の描画の手順は以下の通りです。

1. [RlezBeginShape](#RlezBeginShape)関数で形状の描画を開始する。
2. [RlezVertex](#RlezVertex)関数を任意回数実行し、形状の頂点情報を設定する。
3. [RlezEndShape](#RlezEndShape)関数を実行して形状の描画を行うか、[RlezEndMesh](#RlezEndMesh)関数を実行して形状からメッシュを取得して完了。

`mode`引数で設定できる形状の名前は以下の通りです。

```
"TRIANGLES"(1->2->3,4->5->6,7->8->9,...)
"TRIANGLE_FAN"(1->2->3,1->3->4,1->4->5,...)
"TRIANGLE_STRIP"(1->2->3,2->3->4,3->4->5,...)
"LINES"(1->2,3->4,5->6,...)
"LINE_LOOP"(1->2,2->3,3->4,...,N->1)
"LINE_STRIP"(1->2,2->3,3->4,...)
```

上記の形状の名前の右に記述されている数列は、例えば「`1->2->3`」は「1回目、2回目、3回目の[RlezVertex](#RlezVertex)関数で設定した3つの頂点を繋いだ三角形を描画する」ということを表し、「`1->2`」は「1回目、2回目の[RlezVertex](#RlezVertex)関数で設定した2つの頂点を繋いだ線を描画する」ということを表します。`auto_normal`引数を0以外にすると、[RlezVertex](#RlezVertex)関数を呼び出す際に自動で法線ベクトルを計算するようになります。この状態の場合、[RlezVertex](#RlezVertex)関数の法線ベクトルを設定する引数は無視されます。形状にテクスチャを貼る場合、`resource`引数にはテクスチャのリソースIDを設定してください。もしテクスチャを貼らない場合は、「`-1`」のように0未満の値を設定してください。

---

<div id="RlezVertex"></div>

### RlezVertex

```
void RlezVertex(double x,double y,double z,double u,double v,int r,int g,int b,int a,double nx,double ny,double nz);
```

`x`,`y`,`z`:追加する頂点の座標

`u`,`v`:追加するテクスチャ座標(ピクセル単位)

`r`,`g`,`b`,`a`:追加する頂点の色(それぞれ0～255)

`nx`,`ny`,`nz`:追加する頂点の法線ベクトル

[RlezBeginShape](#RlezBeginShape)関数で形状の描画を開始後にこの関数を任意回数実行することで、形状の頂点情報を設定します。形状にテクスチャを貼り付ける設定をした場合、テクスチャ座標の値を入力する際はピクセル単位で入力することにご注意ください。つまり`u`引数はテクスチャ内のX座標、`v`引数はテクスチャ内のY座標を入力することになります。[RlezBeginShape](#RlezBeginShape)関数を実行した際に法線ベクトルを自動計算する設定を行った場合は、`nx`引数と`ny`引数と`nz`引数に入力した値は無視されます。

---

<div id="RlezEndShape"></div>

### RlezEndShape

```
void RlezEndShape();
```

[RlezBeginShape](#RlezBeginShape)関数と[RlezVertex](#RlezVertex)関数で形状の設定を行った後、この関数で形状の描画を実際に行います。

---

<div id="RlezEndMesh"></div>

### RlezEndMesh

```
int RlezEndMesh();
```

`(return)`:メッシュのリソースID

[RlezBeginShape](#RlezBeginShape)関数と[RlezVertex](#RlezVertex)関数で形状の設定を行った後、この関数で形状のメッシュを取得します。

---

<div id="RlezSetMeshTexture"></div>

### RlezSetMeshTexture

```
void RlezSetMeshTexture(int mesh,int texture);
```

`mesh`:メッシュのリソースID

`texture`:メッシュに貼り付けるテクスチャのリソースID

[RlezEndMesh](#RlezEndMesh)関数で作成したメッシュのテクスチャを書き換えます。

---

<div id="RlezLoadModel"></div>

### RlezLoadModel

```
int RlezLoadModel(const char *path,int load_animation);
```

`path`:読み込む3Dモデルファイルのパスの文字列

`load_animation`:3Dモデルのアニメーションも読み込むかの設定(0だと読み込まない、0以外だと読み込む)

`(return)`:3DモデルのリソースID

指定した3Dモデルファイルから3DモデルのリソースIDを作成します。

---

<div id="RlezFromMeshToModel"></div>

### RlezFromMeshToModel

```
void RlezFromMeshToModel(int mesh);
```

`mesh`:メッシュのリソースID

指定したメッシュから3Dモデルに変換します。この関数を実行後、`mesh`引数で指定したメッシュのリソースIDは、3DモデルのリソースIDに変わります。

---

<div id="RlezSetModelTexture"></div>

### RlezSetModelTexture

```
void RlezSetModelTexture(int model,int material_index,const char *map_name,int texture);
```

`model`:3DモデルのリソースID

`material_index`:3Dモデル内にあるマテリアルの書き換える対象とするインデックス値

`map_name`:マテリアルマップ名の文字列

`texture`:テクスチャのリソースID

指定した3Dモデルのテクスチャを書き換えます。`map_name`引数で使用できるマテリアルマップ名は以下の通りです。

```
"MATERIAL_MAP_ALBEDO"
"MATERIAL_MAP_METALNESS"
"MATERIAL_MAP_NORMAL"
"MATERIAL_MAP_ROUGHNESS"
"MATERIAL_MAP_OCCLUSION"
"MATERIAL_MAP_EMISSION"
"MATERIAL_MAP_HEIGHT"
"MATERIAL_MAP_CUBEMAP"
"MATERIAL_MAP_IRRADIANCE"
"MATERIAL_MAP_PREFILTER"
"MATERIAL_MAP_BRDF"
"MATERIAL_MAP_DIFFUSE"
"MATERIAL_MAP_SPECULAR"
```

各マテリアルマップ名の意味はraylibの「`raylib.h`」ファイル内をご覧ください。

---

<div id="RlezGetAnimationCount"></div>

### RlezGetAnimationCount

```
int RlezGetAnimationCount(int model);
```

`model`:3DモデルのリソースID

`(return)`:3Dモデル内のアニメーションの数

指定した3Dモデル内に存在するアニメーションの数を取得します。

---

<div id="RlezGetAnimationId"></div>

### RlezGetAnimationId

```
int RlezGetAnimationId(int model,const char *name);
```

`model`:3DモデルのリソースID

`name`:3Dモデル内のアニメーション名

`(return)`:3Dモデルのアニメーションのインデックス値

指定した3Dモデル内のアニメーション名からアニメーションのインデックス値を取得します。見つからなかった場合は`(return)`に`-1`が入ります。

---

<div id="RlezGetAnimationFrames"></div>

### RlezGetAnimationFrames

```
int RlezGetAnimationFrames(int model,int id);
```

`model`:3DモデルのリソースID

`id`:3Dモデルのアニメーションのインデックス値

`(return)`:3Dモデルのキーフレーム数

指定した3Dモデルのアニメーションのキーフレーム数を取得します。

---

<div id="RlezSetModelAnimation"></div>

### RlezSetModelAnimation

```
void RlezSetModelAnimation(int model,int id,double frame);
```

`model`:3DモデルのリソースID

`id`:3Dモデルのアニメーションのインデックス値

`frame`:3Dモデルのフレーム値

指定した3Dモデルにアニメーションを適用します。

---

<div id="RlezSetModelAnimationBlend"></div>

### RlezSetModelAnimationBlend

```
void RlezSetModelAnimationBlend(int model,int a_id,double a_frame,int b_id,double b_frame,double blend);
```

`model`:3DモデルのリソースID

`a_id`:1つめのアニメーションのインデックス値

`a_frame`:1つめのアニメーションのフレーム値

`b_id`:2つめのアニメーションのインデックス値

`b_frame`:2つめのアニメーションのフレーム値

`blend`:アニメーションのブレンド値(0.0～1.0)

指定した3Dモデルの2つのアニメーションをブレンドして適用します。`blend`引数の値が`0.0`の場合は1つめのアニメーションが適用され、`1.0`の場合は2つめのアニメーションが適用され、`0.5`の場合は1つめのアニメーションと2つめのアニメーションの間のアニメーションが適用されます。

---

<div id="RlezColor"></div>

### RlezColor

```
void RlezColor(int r,int g,int b,int a);
```

`r`:赤色の値(0～255)

`g`:緑色の値(0～255)

`b`:青色の値(0～255)

`a`:アルファ値(0～255)

描画系関数の描画色を設定します。現在設定されている描画色を取得する場合は、[RlezGetColorR](#RlezGetColorR)関数、[RlezGetColorG](#RlezGetColorG)関数、[RlezGetColorB](#RlezGetColorB)関数、[RlezGetColorA](#RlezGetColorA)関数を使います。

---

<div id="RlezGetColorR"></div>

### RlezGetColorR

```
int RlezGetColorR(void);
```

`(return)`:現在の描画色の赤色の値

[RlezColor](#RlezColor)関数で現在設定されている描画色の赤色の値を取得します。

---

<div id="RlezGetColorG"></div>

### RlezGetColorG

```
int RlezGetColorG(void);
```

`(return)`:現在の描画色の緑色の値

[RlezColor](#RlezColor)関数で現在設定されている描画色の緑色の値を取得します。

---

<div id="RlezGetColorB"></div>

### RlezGetColorB

```
int RlezGetColorB(void);
```

`(return)`:現在の描画色の青色の値

[RlezColor](#RlezColor)関数で現在設定されている描画色の青色の値を取得します。

---

<div id="RlezGetColorA"></div>

### RlezGetColorA

```
int RlezGetColorA(void);
```

`(return)`:現在の描画色のアルファ値

[RlezColor](#RlezColor)関数で現在設定されている描画色のアルファ値を取得します。

---

<div id="RlezDrawMesh"></div>

### RlezDrawMesh

```
void RlezDrawMesh(int mesh);
```

`mesh`:描画するメッシュのリソースID

現在の原点の位置に指定したメッシュを描画します。

---

<div id="RlezDrawModel"></div>

### RlezDrawModel

```
void RlezDrawModel(int model);
```

`model`:描画する3DモデルのリソースID

現在の原点の位置に指定した3Dモデルを描画します。

---

<div id="RlezDrawText"></div>

### RlezDrawText

```
void RlezDrawText(int font,const char *text,double x,double y,double size,double spacing);
```

`font`:フォントのリソースID(0未満だとraylibのデフォルトフォントを使用)

`text`:描画する文字列

`x`,`y`:描画する位置

`size`:描画サイズ

`spacing`:文字間の距離

指定したフォントを使用して文字列を描画します。`font`引数を`-1`等の0未満の値にするとraylibのデフォルトフォントを使用します。

---

<div id="RlezDrawLine"></div>

### RlezDrawLine

```
void RlezDrawLine(double x1,double y1,double z1,double x2,double y2,double z2);
```

`x1`,`y1`,`z1`:線の始点の座標

`x2`,`y2`,`z2`:線の終点の座標

線を描画します。

---

<div id="RlezDrawRect"></div>

### RlezDrawRect

```
void RlezDrawRect(double x,double y,double width,double height,int fill);
```

`x`,`y`:長方形の描画開始座標

`width`,`height`:長方形の横幅と縦幅のサイズ

`fill`:0だと線で描画、それ以外だと塗りつぶして描画

長方形を描画します。

---

<div id="RlezDrawEllipse"></div>

### RlezDrawEllipse

```
void RlezDrawEllipse(double x,double y,double width,double height,int fill);
```

`x`,`y`:楕円の中心座標

`width`,`height`:楕円の横幅と縦幅のサイズ

`fill`:0だと線で描画、それ以外だと塗りつぶして描画

楕円を描画します。

---

<div id="RlezDrawBox"></div>

### RlezDrawBox

```
void RlezDrawBox(double x,double y,double z,double width,double height,double depth,int fill);
```

`x`,`y`,`z`:直方体の中心座標

`width`,`height`,`depth`:直方体の横幅と高さと奥行きのサイズ

直方体を描画します。

---

<div id="RlezDrawSphere"></div>

### RlezDrawSphere

```
void RlezDrawSphere(double x,double y,double z,double size,int rings,int slices,int fill);
```

`x`,`y`,`z`:球の中心座標

`size`:球の直径のサイズ

`slices`:球の表面の分割数

`fill`:0だと線で描画、それ以外だと塗りつぶして描画

球を描画します。

---

<div id="RlezDrawCylinder"></div>

### RlezDrawCylinder

```
void RlezDrawCylinder(double x1,double y1,double z1,double x2,double y2,double z2,double size1,double size2,int sides,int fill);
```

`x1`,`y1`,`z1`:円柱の始点の座標

`x2`,`y2`,`z2`:円柱の終点の座標

`size1`:円柱の始点のサイズ

`size2`:円柱の終点のサイズ

`sides`:円柱の分割数

`fill`:0だと線で描画、それ以外だと塗りつぶして描画

円柱を描画します。

---

<div id="RlezDrawCapsule"></div>

### RlezDrawCapsule

```
void RlezDrawCapsule(double x1,double y1,double z1,double x2,double y2,double z2,double size,int slices,int rings,int fill);
```

`x1`,`y1`,`z1`:カプセルの始点の座標

`x2`,`y2`,`z2`:カプセルの終点の座標

`size`:カプセルの太さ

`slices`:カプセルの円柱部分の分割数

`rings`:カプセルの半球部分の分割数

`fill`:0だと線で描画、それ以外だと塗りつぶして描画

カプセルを描画します。

---

<div id="RlezDrawTexture"></div>

### RlezDrawTexture

```
void RlezDrawTexture(int texture,int src_x,int src_y,int src_w,int src_h,double dest_x,double dest_y,double dest_w,double dest_h);
```

`texture`:テクスチャのリソースID

`src_x`,`src_y`:描画元のテクスチャの読み込み開始座標

`src_w`,`src_h`:描画元のテクスチャを読み込む横幅と縦幅のサイズ

`dest_x`,`dest_y`:描画開始座標

`dest_w`,`dest_h`:描画する横幅と縦幅のサイズ

テクスチャを長方形に貼り付けて描画します。

---

<div id="RlezLoadShaderFromMemory"></div>

### RlezLoadShaderFromMemory

```
int RlezLoadShaderFromMemory(const char *vertex_code,const char *fragment_code);
```

`vertex_code`:頂点シェーダーのコードの文字列

`fragment_code`:フラグメントシェーダーのコードの文字列

`(return)`:シェーダーのリソースID

シェーダーをメモリ上のコードの文字列から作成します。

---

<div id="RlezLoadShader"></div>

### RlezLoadShader

```
int RlezLoadShader(const char *vertex_code_path,const char *fragment_code_path);
```

`vertex_code_path`:頂点シェーダーのコードが記述されたファイルのパスの文字列

`fragment_code_path`:フラグメントシェーダーのコードが記述されたファイルのパスの文字列

`(return)`:シェーダーのリソースID

シェーダーをファイルから作成します。

---

<div id="RlezLoadSoundFromMemory"></div>

### RlezLoadSoundFromMemory

```
int RlezLoadSoundFromMemory(const char *file_type,void *sound_data,int sound_data_size,int is_music,double music_update_time,int music_update_samples);
```

`file_type`:サウンドファイルの拡張子の文字列(例えば`".wav"`)(ドットを含めてください)

`sound_data`:サウンドファイルのバイナリデータが入った変数のポインタ

`sound_data_size`:サウンドファイルのバイナリデータのサイズ

`is_music`:0だとサウンドを効果音として扱い、それ以外だと音楽(BGM)として扱う

`music_update_time`:サウンドを音楽として扱う際、サウンド情報を更新する時間の間隔(秒単位)

`music_update_samples`:サウンドを音楽として扱う際、サウンド情報を1回更新するときに読み込むサンプル数

`(return)`:サウンドのリソースID

サウンドをメモリ上のバイナリデータから作成します。`is_music`引数でサウンドを「効果音」として扱うか、「音楽」として扱うかでrlez.dllのサウンド関連の関数で使用できる関数が異なります。「効果音」として扱うと、サウンドファイルのデータを全てメモリ上にデコードします。「音楽」として扱うと、再生中に随時サウンドファイルのデータを部分的にデコードしながら再生します。「音楽」として扱う場合、`music_update_time`引数でサウンド情報を更新するタイミングを設定し、`music_update_samples`引数でサウンド情報を1回更新するときに読み込むサンプル数を設定できます。よくわからなければ、`music_update_time`引数の値は`1.0/60.0`にしたり、`music_update_samples`引数の値は`4096`にすると良いでしょう。もしサウンドを再生中音が途切れてしまったりしたら、この2つの引数を調整してみると良いでしょう。

---

<div id="RlezLoadSound"></div>

### RlezLoadSound

```
int RlezLoadSound(const char *path,int is_music,double music_update_time,int music_update_samples);
```

`path`:サウンドファイルのパスの文字列

`is_music`:0だとサウンドを効果音として扱い、それ以外だと音楽(BGM)として扱う

`music_update_time`:サウンドを音楽として扱う際、サウンド情報を更新する時間の間隔(秒単位)

`music_update_samples`:サウンドを音楽として扱う際、サウンド情報を1回更新するときに読み込むサンプル数

`(return)`:サウンドのリソースID

サウンドをファイルから作成します。`is_music`引数以降の引数については[RlezLoadSoundFromMemory](#RlezLoadSoundFromMemory)関数の説明をご覧ください。

---

<div id="RlezStopSound"></div>

### RlezStopSound

```
void RlezStopSound(int sound);
```

`sound`:サウンドのリソースID

サウンドの再生を停止します。

---

<div id="RlezPauseSound"></div>

### RlezPauseSound

```
void RlezPauseSound(int sound);
```

`sound`:サウンドのリソースID

サウンドの再生を一時停止します。一時停止したサウンドは、[RlezResumeSound](#RlezResumeSound)関数で再開できます。

---

<div id="RlezResumeSound"></div>

### RlezResumeSound

```
void RlezResumeSound(int sound);
```

`sound`:サウンドのリソースID

[RlezPauseSound](#RlezPauseSound)関数で一時停止したサウンドの再生を再開します。

---

<div id="RlezPlaySound"></div>

### RlezPlaySound

```
void RlezPlaySound(int sound,int music_loop);
```

`sound`:サウンドのリソースID

`music_loop`:0だとループ再生を行わない、0以外だとループ再生を行う

サウンドの再生を行います。`music_loop`引数の設定は、[RlezLoadSoundFromMemory](#RlezLoadSoundFromMemory)関数や[RlezLoadSound](#RlezLoadSound)関数でサウンドを作成した際にサウンドが「音楽」として扱われている場合のみ有効です。

---

<div id="RlezGetSoundStatus"></div>

### RlezGetSoundStatus

```
int RlezGetSoundStatus(int sound);
```

`sound`:サウンドのリソースID

`(return)`:1だと再生中、0だと再生されていない

指定したサウンドが現在再生中かを取得します。

---

<div id="RlezGetSoundTime"></div>

### RlezGetSoundTime

```
double RlezGetSoundTime(int sound);
```

`sound`:サウンドのリソースID

`(return)`:現在の再生位置(秒単位)

指定したサウンドの現在の再生位置を取得します。この関数は、[RlezLoadSoundFromMemory](#RlezLoadSoundFromMemory)関数や[RlezLoadSound](#RlezLoadSound)関数でサウンドを作成した際にサウンドが「音楽」として扱われている場合のみ有効です。

---

<div id="RlezSetSoundTime"></div>

### RlezSetSoundTime

```
void RlezSetSoundTime(int sound,double set_time);
```

`sound`:サウンドのリソースID

指定したサウンドの現在の再生位置を設定します。この関数は、[RlezLoadSoundFromMemory](#RlezLoadSoundFromMemory)関数や[RlezLoadSound](#RlezLoadSound)関数でサウンドを作成した際にサウンドが「音楽」として扱われている場合のみ有効です。

---

<div id="RlezGetSoundLength"></div>

### RlezGetSoundLength

```
double RlezGetSoundLength(int sound);
```

`sound`:サウンドのリソースID

指定したサウンドの現在の総再生時間を取得します。この関数は、[RlezLoadSoundFromMemory](#RlezLoadSoundFromMemory)関数や[RlezLoadSound](#RlezLoadSound)関数でサウンドを作成した際にサウンドが「音楽」として扱われている場合のみ有効です。

---

<div id="RlezSetSoundPitch"></div>

### RlezSetSoundPitch

```
void RlezSetSoundPitch(int sound,double pitch);
```

`sound`:サウンドのリソースID

`pitch`:再生スピード(1.0=通常速度)

サウンドの再生スピードを設定します。

---

<div id="RlezSetSoundVolume"></div>

### RlezSetSoundVolume

```
void RlezSetSoundVolume(int sound,double volume);
```

`sound`:サウンドのリソースID

`volume`:音量(0.0～1.0)

サウンドの音量を設定します。

---

<div id="RlezSetSoundPan"></div>

### RlezSetSoundPan

```
void RlezSetSoundPan(int sound,double pan);
```

`sound`:サウンドのリソースID

`pan`:パンニングの値(-1.0～1.0)

サウンドのパンニング(定位)を設定します。

---

<div id="RlezGetKey"></div>

### RlezGetKey

```
int RlezGetKey(const char *name);
```

`name`:キーボードのキーの名前の文字列

`(return)`:0だと押されていない、1だと押されている

指定したキーボードのキーが現在押されているかを取得します。`name`引数で指定できるキーボードのキーの名前は以下の通りです。

```
"KEY_NULL"
"KEY_APOSTROPHE"
"KEY_COMMA"
"KEY_MINUS"
"KEY_PERIOD"
"KEY_SLASH"
"KEY_ZERO"
"KEY_ONE"
"KEY_TWO"
"KEY_THREE"
"KEY_FOUR"
"KEY_FIVE"
"KEY_SIX"
"KEY_SEVEN"
"KEY_EIGHT"
"KEY_NINE"
"KEY_SEMICOLON"
"KEY_EQUAL"
"KEY_A"
"KEY_B"
"KEY_C"
"KEY_D"
"KEY_E"
"KEY_F"
"KEY_G"
"KEY_H"
"KEY_I"
"KEY_J"
"KEY_K"
"KEY_L"
"KEY_M"
"KEY_N"
"KEY_O"
"KEY_P"
"KEY_Q"
"KEY_R"
"KEY_S"
"KEY_T"
"KEY_U"
"KEY_V"
"KEY_W"
"KEY_X"
"KEY_Y"
"KEY_Z"
"KEY_LEFT_BRACKET"
"KEY_BACKSLASH"
"KEY_RIGHT_BRACKET"
"KEY_GRAVE"
"KEY_SPACE"
"KEY_ESCAPE"
"KEY_ENTER"
"KEY_TAB"
"KEY_BACKSPACE"
"KEY_INSERT"
"KEY_DELETE"
"KEY_RIGHT"
"KEY_LEFT"
"KEY_DOWN"
"KEY_UP"
"KEY_PAGE_UP"
"KEY_PAGE_DOWN"
"KEY_HOME"
"KEY_END"
"KEY_CAPS_LOCK"
"KEY_SCROLL_LOCK"
"KEY_NUM_LOCK"
"KEY_PRINT_SCREEN"
"KEY_PAUSE"
"KEY_F1"
"KEY_F2"
"KEY_F3"
"KEY_F4"
"KEY_F5"
"KEY_F6"
"KEY_F7"
"KEY_F8"
"KEY_F9"
"KEY_F10"
"KEY_F11"
"KEY_F12"
"KEY_LEFT_SHIFT"
"KEY_LEFT_CONTROL"
"KEY_LEFT_ALT"
"KEY_LEFT_SUPER"
"KEY_RIGHT_SHIFT"
"KEY_RIGHT_CONTROL"
"KEY_RIGHT_ALT"
"KEY_RIGHT_SUPER"
"KEY_KB_MENU"
"KEY_KP_0"
"KEY_KP_1"
"KEY_KP_2"
"KEY_KP_3"
"KEY_KP_4"
"KEY_KP_5"
"KEY_KP_6"
"KEY_KP_7"
"KEY_KP_8"
"KEY_KP_9"
"KEY_KP_DECIMAL"
"KEY_KP_DIVIDE"
"KEY_KP_MULTIPLY"
"KEY_KP_SUBTRACT"
"KEY_KP_ADD"
"KEY_KP_ENTER"
"KEY_KP_EQUAL"
```

---

<div id="RlezGetMouseButton"></div>

### RlezGetMouseButton

```
int RlezGetMouseButton(const char *name);
```

`name`:マウスのボタンの名前の文字列

`(return)`:0だと押されていない、1だと押されている

指定したマウスのボタンが現在押されているかを取得します。`name`引数で指定できるマウスのボタンの名前は以下の通りです。

```
"MOUSE_BUTTON_LEFT"
"MOUSE_BUTTON_RIGHT"
"MOUSE_BUTTON_MIDDLE"
"MOUSE_BUTTON_SIDE"
"MOUSE_BUTTON_EXTRA"
"MOUSE_BUTTON_FORWARD"
"MOUSE_BUTTON_BACK"
```

---

<div id="RlezGetMouseX"></div>

### RlezGetMouseX

```
int RlezGetMouseX(void);
```

`(return)`:マウスカーソルのX座標

現在のraylibのウィンドウ内のマウスカーソルのX座標を取得します。

---

<div id="RlezGetMouseY"></div>

### RlezGetMouseY

```
int RlezGetMouseY(void);
```

`(return)`:マウスカーソルのY座標

現在のraylibのウィンドウ内のマウスカーソルのY座標を取得します。

---

<div id="RlezSetMousePosition"></div>

### RlezSetMousePosition

```
void RlezSetMousePosition(int x,int y);
```

`x`:マウスカーソルの移動先のX座標

`y`:マウスカーソルの移動先のY座標

raylibのウィンドウの左上を基準にマウスカーソルの座標を移動させます。

---

<div id="RlezSetMouseVisible"></div>

### RlezSetMouseVisible

```
void RlezSetMouseVisible(int visible);
```

`visible`:0だとマウスカーソルを非表示、0以外だとマウスカーソルを表示

raylibのウィンドウ内でマウスカーソルを表示するか非表示にするかを設定します。

---

<div id="RlezCheckMouseInWindow"></div>

### RlezCheckMouseInWindow

```
int RlezCheckMouseInWindow(void);
```

`(return)`:0だとウィンドウ外、1だとウィンドウ内

raylibのウィンドウ内にマウスカーソルがあるかを取得します。

---

<div id="RlezGetMouseWheelX"></div>

### RlezGetMouseWheelX

```
double RlezGetMouseWheelX(void);
```

`(return)`:マウスホイールのX値

マウスホイールのX値を取得します。

---

<div id="RlezGetMouseWheelY"></div>

### RlezGetMouseWheelY

```
double RlezGetMouseWheelY(void);
```

`(return)`:マウスホイールのY値

マウスホイールのY値を取得します。

---

<div id="RlezCheckGamepad"></div>

### RlezCheckGamepad

```
int RlezCheckGamepad(int gamepad);
```

`gamepad`:ゲームパッドのインデックス値

`(return)`:0だと接続されていない、1だと接続されている

指定したゲームパッドが現在接続されているかを取得します。

---

<div id="RlezGetGamepadButton"></div>

### RlezGetGamepadButton

```
int RlezGetGamepadButton(int gamepad,const char *button);
```

`gamepad`:ゲームパッドのインデックス値

`button`:ゲームパッドのボタンの名前の文字列

`(return)`:0だと押されていない、1だと押されている

指定したゲームパッドのボタンが押されているかを取得します。`button`引数で指定できるボタンの名前は以下の通りです。

```
"GAMEPAD_BUTTON_UNKNOWN"
"GAMEPAD_BUTTON_LEFT_FACE_UP"
"GAMEPAD_BUTTON_LEFT_FACE_RIGHT"
"GAMEPAD_BUTTON_LEFT_FACE_DOWN"
"GAMEPAD_BUTTON_LEFT_FACE_LEFT"
"GAMEPAD_BUTTON_RIGHT_FACE_UP"
"GAMEPAD_BUTTON_RIGHT_FACE_RIGHT"
"GAMEPAD_BUTTON_RIGHT_FACE_DOWN"
"GAMEPAD_BUTTON_RIGHT_FACE_LEFT"
"GAMEPAD_BUTTON_LEFT_TRIGGER_1"
"GAMEPAD_BUTTON_LEFT_TRIGGER_2"
"GAMEPAD_BUTTON_RIGHT_TRIGGER_1"
"GAMEPAD_BUTTON_RIGHT_TRIGGER_2"
"GAMEPAD_BUTTON_MIDDLE_LEFT"
"GAMEPAD_BUTTON_MIDDLE"
"GAMEPAD_BUTTON_MIDDLE_RIGHT"
"GAMEPAD_BUTTON_LEFT_THUMB"
"GAMEPAD_BUTTON_RIGHT_THUMB"
```

---

<div id="RlezGetAxisCount"></div>

### RlezGetAxisCount

```
int RlezGetAxisCount(int gamepad);
```

`gamepad`:ゲームパッドのインデックス値

`(return)`:ゲームパッドの軸の数

指定したゲームパッドのスティック等の軸の数を取得します。

---

<div id="RlezGetGamepadAxis"></div>

### RlezGetGamepadAxis

```
double RlezGetGamepadAxis(int gamepad,int axis);
```

`gamepad`:ゲームパッドのインデックス値

`axis`:ゲームパッドの軸のインデックス値

`(return)`:ゲームパッドの軸の値(-1.0～1.0)

指定したゲームパッドのスティック等の軸の値を取得します。

---

<div id="RlezSetGamepadVibration"></div>

### RlezSetGamepadVibration

```
void RlezSetGamepadVibration(int gamepad,double left,double right,double duration);
```

`gamepad`:ゲームパッドのインデックス値

`left`:左側のモーターの振動の強さ(0.0～1.0)

`right`:右側のモーターの振動の強さ(0.0～1.0)

`duration`:振動の長さ(秒単位)

指定したゲームパッドを振動させます。

---

<div id="rlez_load_file_from_hsp"></div>

### rlez_load_file_from_hsp

```
#deffunc rlez_load_file_from_hsp str p_file,var p_data
```

`p_file`:ファイルのパスの文字列

`p_data`:書き込む文字列型変数

`(stat)`:ファイルのサイズ

ファイルのデータを変数に書き込みます。この命令を実行する際、`p_data`引数で指定した文字列型変数は全て空の状態にした後、ファイルのデータのサイズに拡張されます。この命令で指定する`p_file`引数には、DPMファイル内のファイルや実行ファイル内に埋め込まれたファイルのパスも指定できます。

---

<div id="rlez_load_texture_from_hsp"></div>

### rlez_load_texture_from_hsp

```
#defcfunc rlez_load_texture_from_hsp str p1
```

`p1`:ファイルのパスの文字列

`(return)`:テクスチャのリソースID

指定したファイルからテクスチャを作成します。この命令で指定する`p1`引数には、DPMファイル内のファイルや実行ファイル内に埋め込まれたファイルのパスも指定できます。

---

<div id="rlez_load_font_from_hsp"></div>

### rlez_load_font_from_hsp

```
#defcfunc rlez_load_font_from_hsp str p1,int p2,str p3,int p4,int p5
```

`p1`:ファイルのパスの文字列

`p2`:生成されるフォントデータ内のフォントテクスチャに描画されるフォントサイズ

`p3`:生成されるフォントデータ内のフォントテクスチャに描画される文字群の文字列

`p4`,`p5`:内部で生成されるフォントテクスチャの目標サイズの横幅と縦幅

`(return)`:フォントデータのリソースID

指定したファイルからフォントを作成します。この命令で指定する`p1`引数には、DPMファイル内のファイルや実行ファイル内に埋め込まれたファイルのパスも指定できます。

---

<div id="rlez_load_sound_from_hsp"></div>

### rlez_load_sound_from_hsp

```
#defcfunc rlez_load_sound_from_hsp str p1,int p2,double p3,int p4
```

`p1`:ファイルのパスの文字列

`p2`:0だとサウンドを効果音として扱い、それ以外だと音楽(BGM)として扱う

`p3`:サウンドを音楽として扱う際、サウンド情報を更新する時間の間隔(秒単位)

`p4`:サウンドを音楽として扱う際、サウンド情報を1回更新するときに読み込むサンプル数

`(return)`:サウンドのリソースID

指定したファイルからサウンドを作成します。この命令で指定する`p1`引数には、DPMファイル内のファイルや実行ファイル内に埋め込まれたファイルのパスも指定できます。

---