# Nava
Nava work against SEB 3.10.1

The goal of this script is to prevent SEB to start new obfuscated desktop and put in KIOSK mode. My method is to patch all necessary function with Harmony C#. You can see all that function inside `NavaInjector/Milim/patch`

## Installation


```bash
git clone https://github.com/seynth/NavaInjector
cd NavaInjector
```

### Step 1 - Build the injector (NavaInjector.exe)

```bash
cd injector
go build
```

## Step 2 - Build the runner (Nava.dll)

```bash
cd Nava
go build -o nava.dll --buildmode=c-shared
```

## Step 3 - Build the payload (Milim.dll)

- Open solution file NavaInjector/Milim/Milim.sln in Visual Studio 2022
- Make sure to use `Debug` and `x64` platform configuration
- Right click on Milim project and select build or press `Ctrl + Shift + B`

## Usage

Usage is very simple, you just need to run NavaInjector.exe as Administrator, but make sure to turn off Real-time Protection windows defender.

Wait until it say `[Nava] Starting`, then you can start your Safe Exam Browser by double clicking .seb file configuration. After NavaInjector reaching EOF you can close your terminal




