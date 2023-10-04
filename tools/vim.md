# Vim 

*Updated `2022-01-29`*

Plugins are automatically managed by `vim` in the `.vim/pack/start` folder. 

### `.vimrc`

Source a local configuration from version control:

```
so ~/.vim/.vimrc
```

### `~/.vim/.vim_baserc`

```
" ...
colorscheme nord
syntax on
set number
" CUA
" copy and paste
vmap <C-c> "+yi
vmap <C-x> "+c
vmap <C-v> c<ESC>"+p
imap <C-v> <ESC>"+pa
set mouse=a


" Airline & NERDtree 
map <C-b> :NERDTreeToggle<CR>
map <C-n> :term<CR>
map <C-h> :Hexmode<CR>
let NERDTreeShowHidden=1

let g:airline#extensions#tabline#enabled = 1
map <Tab> :bnext!<CR>

```

Undo

```
" ?? TODO
```

##### `colorscheme`

`:colorscheme fogbell-light`

- Fogbell 
- Nord
- Monochrome

#### Zen Mode: Goyo

- `:Goyo`

#### Matrix "Screen Saver"

- `:Matrix`  

### Tabs vs. Buffers

- https://github.com/vim-airline/vim-airline *updated from lightline*

```

```

*The above section returned me to a quality of life with CUA, colors and zen-mode.* 

### Syntax Highlighting

`syntax on`

- https://github.com/sheerun/vim-polyglot

#### DelmitMate

*Closing brackets*

### Code Completion

- https://github.com/neoclide/coc.nvim 

#### ALE

- https://github.com/dense-analysis/ale

#### CoC

- https://github.com/neoclide/coc.nvim

Install `node` version, Install `yarn`, build `coc.nvim` folder:

```
nvm install v14.17.0
curl --compressed -o- -L https://yarnpkg.com/install.sh | bash

# yarn install inside coc folder
cd ~/.vim/pack/plugins/start/coc.nvim && yarn install
```

Install language servers:

```
:CocInstall coc-html
:CocInstall coc-css
:CocInstall coc-json
:CocInstall coc-tsserver
```

#### vim-ccls

- https://github.com/m-pilia/vim-ccls

https://stackoverflow.com/questions/3446320/in-vimv-how-to-map-save-to-ctrl-s

#### vim-prettier
- https://github.com/prettier/vim-prettier

### Search, Find and Replace

``` :%s/findthis/replaceit/g```

- https://www.cyberciti.biz/faq/vim-text-editor-find-and-replace-all-text/
- https://stackoverflow.com/questions/38398787/how-do-i-search-in-all-files-of-my-project-using-vim

### fzf.vim

- https://github.com/junegunn/fzf.vim

#### FZF: Fuzzy Finder  && Silver Searcher

``` :Ag searchTerm ```

- https://github.com/junegunn/fzf
- https://github.com/ggreer/the_silver_searcher
- https://pragmaticpineapple.com/improving-vim-workflow-with-fzf/#speed-search-your-project
- https://github.com/phiresky/ripgrep-all


how to set up vim for 

- Javascript
- Typescript
- C++
- C: debugging with `gdb`
- Go
- Rust
- LaTeX
- Jupyter
- Python console

## Vimtex...

### vimwiki

https://vimwiki.github.io

- https://forum.obsidian.md/t/integrate-with-vim-wiki/11626

```
let g:vimwiki_list = [{'path': '~/repos/obsidian/', \ 'syntax': 'markdown', 'ext': '.md'}]

" Change `~/repos/obsidian/` to the path of your obsidian vault.

```

`~8p` New goal: get going on those little C programs for 1-bit music, the visualizer. 



### radare2 

git clone https://github.com/radareorg/radare2
radare2/sys/install.sh
```

`r2 ` https://github.com/radareorg/radare2

`gef` for decompile. 

https://hugsy.github.io/gef/


tmux <C-b> + ] is scroll in tmux. Finally. 
<C-b>  : " % 

```
`


---

## Building Vim

*To enable features.*

```
vi --version
```



```
sudo ./configure --with-features=huge --enable-multibyte --enable-python3interp=yes  --enable-cscope --prefix=/usr/local/	
```



```
sudo make clean
sudo make
sudo make install
```



### Python in Vim

```
sudo ./configure --with-features=huge --enable-multibyte --enable-pythoninterp=yes --with-python-config-dir=/usr/lib/python2.7/config-x86_64-linux-gnu/ --enable-python3interp=yes --with-python-config-dir=/usr/lib/python3.7/config-3.7m-x86_64-linux-gnu/ --enable-gui=gtk2 --enable-cscope --prefix=/usr/local/	
```

Building vim... need to build `libtool`:
*Vim notes 11.17*

- https://gist.github.com/byronmansfield/e972caf0f423af1c84e5b57975ccca3d
- https://www.gnu.org/software/libtool/manual/libtool.html

```
To build libtool from source:

    1. Download the source code from https://www.gnu.org/software/libtool/.

    2. Run these commands from the root of the source code directory:

           ./configure --program-prefix=g
           make
           sudo make install
```

- https://github.com/vim/vim/issues/8084

once `vim` finally compiled on MacOS, I had to find the binary in `/usr/local/bin` and deal with the existing system link. 

---
## Debugging in Vim

Download and Build `gdb`

- https://www.sourceware.org/gdb/download/

```
ftp://sourceware.org/pub/gdb/releases/ 
```

Probably need to link to Python 3.8+

```
./configure --with-python=$(which python3) 
```



#### GEF

- https://github.com/hugsy/gef

#### Vimspector

- https://github.com/puremourning/vimspector#quick-start

```
curl -L https://github.com/puremourning/vimspector/releases/download/4005121932/vimspector-macos-4005121932.tar.gz | tar -C $HOME/.vim/pack zxvf -
```

#### Voltron

- https://github.com/snare/voltron

##### Links

- https://stackoverflow.com/questions/16521071/nodejs-how-to-do-debugging-using-gdb
- https://guyinatuxedo.github.io/02-intro_tooling/gdb-gef/index.html
- https://www.dannyadam.com/blog/2019/05/debugging-in-vim/
- https://lldb.llvm.org/use/map.html

---


```py
  
  ./configure --with-features=huge --enable-cscope --enable-multibyte  --enable-python3interp   --with-python3-config-dir=/usr/lib/python3.8/config-3.8m-x86_64-linux-gnu   --enable-fail-if-missing


sudo ./configure --enable-multibyte --enable-pythoninterp=yes --with-python-config-dir=/usr/lib/python2.7/config-x86_64-linux-gnu/ --enable-python3interp=yes --with-python-config-dir=/usr/lib/python3.8/config-3.8m-x86_64-linux-gnu/ --enable-gui=gtk2 --enable-cscope --prefix=/usr/local/ --enable-gtk2-check --enable-gnome-check --with-features=huge --with-x
```


---
## History and overview

``` Vim was created by Bram Moolenaar in 1991. ```

#### Vi

https://en.wikipedia.org/wiki/Bill_Joy

#### Links

- https://dev.to/omerxx/vim-from-foe-to-friend-in-9-minutes-2np0
- http://vimawesome.com/
- https://en.wikipedia.org/wiki/Bram_Moolenaar
- https://en.wikipedia.org/wiki/Vim_(text_editor)
- https://begriffs.com/posts/2019-07-19-history-use-vim.html?hn=3
- https://blog.logrocket.com/configuring-vim-rust-development/
- https://github.com/phiresky/ripgrep-all
- https://github.com/ggreer/the_silver_searcher
- https://github.com/tpope/vim-fugitive
- https://www.vim.org/git.php
- https://shapeshed.com/vim-packages/

