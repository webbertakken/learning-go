# Pixl

## Running the application

- Open `MSYS2 MINGW64` application (terminal).
- Change directory to `/c/Users/Webber/Repositories/learning-go/zero-to-mastery/projects/pixl`
- 

## Setup

### Installing MSYS2 and GCC

_(all questions just enter)_

Install [msys64](https://www.msys2.org/) (needed for gcc and other tools)

- Once installed, in the terminal that opened, run:

```bash
pacman -Syu
```

- The terminal closed automatically. Open MSYS2 app (again).

```bash
pacman -Su
```

- install gcc and other tools

```bash
pacman -S --needed base-devel mingw-w64-x86_64-toolchain
```

- Close the terminal and open MSYS2 MINGW64 app.
- Configure Go in Mingw64

```bash
export PATH=$PATH:/c/Program\ Files/Go/bin:~/Go/bin
```

- Verify installation

```bash
go version
gcc --version
```

Now you're ready to go back to the top.
