
# zsh

https://github.com/zsh-users/antigen

```
curl -L git.io/antigen > antigen.zsh
```


difference from bash? Time to start switching? [[2023-01-29#Zsh]]

- [x] autocompletion like recently seen in Kali?

## i3

- [x] **i3** document `.config/i3/config`, bindkey, bindcode etc #tools #workshop

This is just a start.

`nano .config/i3/config`

```
# .config/i3/config

bindsym $mod+space exec --no-startup-id i3-dmenu-desktop --dmenu='rofi -dmenu'
```

### rofi

`sudo nano /etc/apt/sources.list`

```
# sources.list

# See https://www.kali.org/docs/general-use/kali-linux-sources-list-repositories/
deb http://http.kali.org/kali kali-rolling main contrib non-free

# Additional line for source packages
# deb-src http://http.kali.org/kali kali-rolling main contrib non-free

deb http://_ftp.de.debian.org/debian_ bullseye main


```

Adding the right `armhf` packages is pretty vital after beginning to extend the environment and build stuff from scratch. 

```
sudo apt-get build-dep ..
```

## glow

https://github.com/charmbracelet/glow


### lynx


## w3m

`a` to download.

- https://w3m.sourceforge.net/MANUAL


### Python in Vim 8.2

```
sudo ./configure --with-features=huge --enable-multibyte --enable-pythoninterp=yes --with-python-config-dir=/usr/lib/python2.7/config-x86_64-linux-gnu/ --enable-python3interp=yes --with-python-config-dir=/usr/lib/python3.7/config-3.7m-x86_64-linux-gnu/ --enable-gui=gtk2 --enable-cscope --prefix=/usr/local/	
```



---
`11.7.21 PM`

[https://learn.hashicorp.com/tutorials/vagrant/getting-started-up?in=vagrant/getting-started](https://learn.hashicorp.com/tutorials/vagrant/getting-started-up?in=vagrant/getting-started)

`vagrant box add hashicorp/bionic64`

-   [https://jupyter-docker-stacks.readthedocs.io/en/latest/](https://jupyter-docker-stacks.readthedocs.io/en/latest/)
    
    -   docker run -p 8888:8888 jupyter/scipy-notebook:33add21fab64
    
-   [https://www.kali.org/blog/announcing-kali-for-vagrant/](https://www.kali.org/blog/announcing-kali-for-vagrant/)
    

Copying vim settings...

Building vim... need to build `libtool`:

-   [https://gist.github.com/byronmansfield/e972caf0f423af1c84e5b57975ccca3d](https://gist.github.com/byronmansfield/e972caf0f423af1c84e5b57975ccca3d)
    
-   [https://www.gnu.org/software/libtool/manual/libtool.html](https://www.gnu.org/software/libtool/manual/libtool.html)
    

To build libtool from source:  
​  
 1. Download the source code from https://www.gnu.org/software/libtool/.  
​  
 2. Run these commands from the root of the source code directory:  
​  
 ./configure --program-prefix=g  
 make  
 sudo make install

-   [https://github.com/vim/vim/issues/8084](https://github.com/vim/vim/issues/8084)
    

once `vim` finally compiled on MacOS, I had to find the binary in `/usr/local/bin` and deal with the existing system link.


## frameworks


### JUCE

- https://tobanteaudio.gitbook.io/juce-cookbook/setup/projucer_vs_cmake
- https://github.com/juce-framework/JUCE/blob/master/examples/CMake/AudioPlugin/CMakeLists.txt
- https://github.com/JoshMarler/react-juce 

- https://github.com/jatinchowdhury18/AnalogTapeModel/blob/master/Plugin/CMakeLists.txt
- https://developer.android.com/codelabs/android-studio-cmake#0
- https://developer.android.com/training/articles/perf-jni.html#kotlin
- https://github.com/sheldonth/ios-cmake

- https://github.com/maxwellpollack/juce-plugin-ci

- https://github.com/sudara/pamplejuce

- https://github.com/tobanteAudio/juce-cookbook/blob/master/examples/component/XYPad/Source/PluginProcessor.cpp

### openFrameworks

https://github.com/peterkolski/ofxCMake


## cloudfront

https://dash.cloudflare.com/0163c932fa68ef7090236247197a66e6/pages/view/wilderzone


## aws

https://signin.aws.amazon.com/

## Extensa, ESD-IDF, esptool.py...


## lftp


---


*Network engineering, FreeBSD, esp32, wails...*

Some tech news. You know 'trusty' C? It's so portable. Because it has a zillion implementations. Undefined behavior





## Rust in the Linux Kernel

Linus's commit:

-   [https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=8aebac82933ff1a7c8eede18cab11e1115e2062b](https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/commit/?id=8aebac82933ff1a7c8eede18cab11e1115e2062b)
-  [The Case for Writing a Kernel in Rust](https://www.cs.virginia.edu/~bjc8c/papers/levy17rustkernel.pdf)
- 

[https://www.memorysafety.org/blog/rust-in-linux-just-the-beginning/](https://www.memorysafety.org/blog/rust-in-linux-just-the-beginning/)

## FreeBSD

[https://github.com/ravynsoft/ravynos](https://github.com/ravynsoft/ravynos) - attempt at macOS parity, based on freeBSD
https://nomadbsd.org/screenshots.html - on a usb stick

[https://freebsdfoundation.org/resource/installing-freebsd-for-raspberry-pi/](https://freebsdfoundation.org/resource/installing-freebsd-for-raspberry-pi/)

[https://github.com/ravynsoft/ravynos](https://github.com/ravynsoft/ravynos)

READ THIS: https://unixsheikh.com/articles/technical-reasons-to-choose-freebsd-over-linux.html

https://docs.freebsd.org/en/books/handbook/introduction/#nutshell

### Creative Linux

#### System
- i3 (or regolith, but might be worth custom i3 this time,etc)
- zsh
- vim 

#### Audio
- KXStudio things
- JACK
- audacity
- bitwig
- renoise
- audio dev, JUCE, etc

#### Visual
- inkscape
- gimp
- blender
- https://godotengine.org worth looking into
- openshot
- VIDIVOX
- chataigne 
- creative coding: nanou, openFrameworks

## `youtube-dl`

## `ffmpeg`

## `vlc`

## Blender 

## nannou

GLSL...

### Creative Applications

```

Cyan (0,1,1) #00FFFF
Magenta(1,0,1) #FF00FF
Yellow(1,1,0) #FFFF00 
Key(1,1,1)  #000000, #FFFFFF

```


https://en.wikipedia.org/wiki/CMYK_color_model

https://thebookofshaders.com/06/

https://p5js.org/libraries/
https://github.com/patriciogonzalezvivo/glslCanvas

https://dev.to/georgedoescode/a-generative-svg-starter-kit-5cm1

https://dev.to/georgedoescode/tutorial-build-a-smooth-animated-blob-using-svg-js-3pne

https://codepen.io/shubniggurath/details/qeXRrW

https://p5js.org/examples/3d-basic-shader.html

https://github.com/nature-of-code/noc-examples-p5.js

https://github.com/ITP-xStory/p5js-shaders

https://www.geeksforgeeks.org/p5-js-createshader-method/

http://www.hardformat.org/6561/add-n-to-x-little-black-rocks-in-the-sun/

## SVG 

https://github.com/georgedoescode/generative-utils

https://dev.to/georgedoescode/a-generative-svg-starter-kit-5cm1

https://codepen.io/georgedoescode/pen/oNzamjV

## Shaders

Generative
design, graphics, c++, video, openFrameworks

https://thebookofshaders.com

VHS Effect
https://cgjam.net/old-vhs-style-movies/
https://www.youtube.com/watch?v=Ra-E0JVkXEI

http://veejayhq.github.io/screenshots/
https://vidifold.com/knowledgebase
http://lives-video.com/


https://linuxize.com/post/how-to-nvidia-drivers-on-ubuntu-20-04/

lspci | grep VGA


sudo apt install mesa-utils

glxinfo | grep "OpenGL version"

http://lives-video.com/images/livesmt.png


new england forest
https://www.youtube.com/watch?v=Vi12xaJxA5U

https://mattj.io/posts/2021-02-27-create-animated-gif-and-webp-from-videos-using-ffmpeg/

https://www-cs-faculty.stanford.edu/~knuth/organ.html

---
A pixel is not a tiny square
http://alvyray.com/Memos/CG/Microsoft/6_pixel.pdf

https://vidifold.com/manual/control
