# voy

Interact with Voyager Index from the comfort of the command line.

## Installation

```sh
go get github.com/voyager-index/voy
export PATH=$PATH:$HOME/go/bin
```

## Usage

```sh
voy search portland
```

```json
[                                                                          
  {                                  
    "city": "Portland",              
    "country": "United States",
    "lon": "-122.67",                
    "lat": "45.52",                  
    "id": 6782,                      
    "population": 1207757,           
    "mbps": "25.86",
    "beach": false,
    "airport": true,
    "elevation": 61,
    "pollution": "27.61514529",
    "palms": false,
    "...": "...",
    "...": "...",
    "...": "...",
    "uvjan": 207,
    "uvfeb": 181,
    "uvmar": 131,
    "uvapr": 77,
    "uvmay": 131,
    "uvjun": 29,
    "uvjul": 32,
    "uvaug": 51,
    "uvsep": 86,
    "uvoct": 128,
    "uvnov": 166,
    "uvdec": 195,
    "purchasingpower": "1.0894169998",
    "povertyindex": 0,
    "image": null
  }
]
```

## Development

```sh
./cross-compile.sh main.go

Building 1/7: main.go-windows-amd64.exe
Building 2/7: main.go-windows-386.exe
Building 3/7: main.go-darwin-amd64
Building 4/7: main.go-linux-amd64
Building 5/7: main.go-linux-386
Building 6/7: main.go-linux-arm
Building 7/7: main.go-linux-arm64
```
