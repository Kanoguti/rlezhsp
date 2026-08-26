# rlezhsp

2026 Kanoguti

## 注意

- このライブラリは開発中の段階であることをご了承ください。
- rlez.dllでは、本来`float`型で扱うべき引数等を、`double`型で扱っております。rlez.dll内で実行されるraylibの関数に値を渡す際は`double`型から`float`型にキャストされます。これはHSP3言語では「`float`型」の扱いが少し難しいためです。小数の精度が異なってしまう点にご注意ください。

## このソフトウェアについて

C/C++言語向けのゲーム制作ライブラリのraylibを、rlezhspの作者にとって扱いやすく機能を厳選と簡易化し、HSP3言語で使えるようにしたライブラリです。rlez.dllはHSP3言語以外でも適切にヘッダーを記述したら使用できます。

## 動作確認環境

- OS:Windows 11 64bit
- HSP3言語のバージョン:3.8 beta1(64bit版)

## rlezhspの使用例

以下のスクリプトファイルを作成し、`src/build/opengl33`フォルダ内にある`rlez.dll`ファイルと、`src`フォルダ内にある`rlez.hsp`ファイルをスクリプトファイルと同じフォルダにコピーして実行すると、raylibのウィンドウが表示され、そのウィンドウ内に「`Hello, world!`」という文字列が表示されます。

```
#packopt hide 1

#include "hsp3_64.as"
#include "rlez.hsp"

gsel 0,-1
rlez_open_window 640,480
rlez_set_window_focus

while rlez_check_window_close()==0
rlez_begin_window

rlez_background 255,255,255,255

rlez_begin_2d 0,0,0,0,0,1

rlez_color 0,0,0,255
rlez_draw_text -1,"Hello, world!",0,0,30,5

rlez_end_2d

rlez_end_draw
wend

rlez_end
end
```

## rlez.dllの関数リファレンス

rlez.dllの関数リファレンスは`Reference.md`ファイルまたは`Reference.html`ファイルに記述されています。

## rlez.dllの作成

作者がrlez.dllを作成する際、以下の環境で行いました。

- Go言語
    - バージョン:1.27
- MSYS2(MinGW-w64)
    - gccコマンドやmakeコマンド等を使用できるようにしてください。
    - ANGLEバックエンドを用いたrlez.dllを作成する場合は、ANGLEの開発環境もインストールしてください。

まずは、コマンドプロンプト等でrlezhspの`src`フォルダ内に移動し、以下のコマンドでraylib-goをインストールします。

```
make download_raylib
```

環境が整ったら、以下のコマンドを`src`フォルダ内で実行すると、`build/opengl33`フォルダ内にOpenGL 3.3バックエンドを使用する`rlez.dll`ファイルが作成されます。

```
make build_opengl33
```

`build/angle`フォルダにANGLEバックエンドを用いた`rlez.dll`を作成する場合は、以下のコマンドを実行してください。

```
make build_angle
```

ANGLEバックエンドを用いた`rlez.dll`をWindowsで使用する場合、`libEGL.dll`ファイルと`libGLESv2.dll`ファイルが最低限必要で、環境によっては`d3dcompiler_47.dll`ファイルも必要です。ANGLEのWindowsの64bit版バイナリは、例えば以下のURLからダウンロード可能です。

`https://packages.msys2.org/packages/mingw-w64-ucrt-x86_64-angleproject`

全てのバックエンドの`rlez.dll`を一括で作成する場合は、以下のコマンドを実行してください。

```
make build_all
```

## rlez.hspの作成

`src`フォルダ内にある`make_rlezhsp.hsp`ファイルをHSPスクリプトエディタ等で実行すると、`export_list.h`ファイルの情報を基に`rlez.hsp`ファイルが自動生成されます。

## このソフトウェアを開発するにあたって使用させていただいたソフトウェアについて

- raylib<br>`https://www.raylib.com/`<br>zlib License
- raylib-go<br>`https://github.com/gen2brain/raylib-go`<br>zlib License
- The Go Programming Language<br>`https://go.dev/`<br>BSD 3-Clause "New" or "Revised" License
- ANGLE<br>`http://angleproject.org/`<br>BSD 3-Clause License

各ソフトウェアの開発者様に感謝申し上げます。各ソフトウェアのライセンス文書は「licenses」フォルダ内に含まれております。

## rlezhspのライセンス

MIT License