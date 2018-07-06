# boris
A magic tool that uses music file's tags to automaticaly organize your music in folders, using a user provided template.

## Installation
The following commands requires a working golang instalation.
```
go get https://github.com/JosephCaillet/boris
go install https://github.com/JosephCaillet/boris
```

## Basic use
`boris -i myMusicDir -o myOrganizedMusicDir -c myConfigFile`

## Exemple

### Before boris
```
~ $ tree
.
└── zikIn
    ├── 01_02 Daddy Cool.flac
    ├── 02-01-Boney_M-Felicidad_Margherita-LLS.flac
    ├── 02_02 Gadda-Da-Vida (7_ Version) (Full Length).flac
    ├── Andy Fox - Tempus Fugit
    │   ├── Andy Fox - Tempus Fugit - 01 Tempus Fugit.flac
    │   ├── Andy Fox - Tempus Fugit - 02 Daydream (feat. Vanna).flac
    │   ├── Andy Fox - Tempus Fugit - 06 Second Life.flac
    │   ├── Andy Fox - Tempus Fugit - back.png
    │   └── cover.jpg
    ├── Annella - Perfume - Single
    │   ├── Annella - Perfume - Single - 01 Perfume.flac
    │   ├── cover.jpg
    │   └── test2.txt
    ├── Brandon - Neon Haze
    │   ├── Brandon - Neon Haze - 01 Neon Haze.flac
    │   ├── Brandon - Neon Haze - 03 Time Machine.flac
    │   ├── Brandon - Neon Haze - 04 Arcade.flac
    │   ├── cover.jpg
    │   ├── env
    │   │   ├── EnV - EnV - Shinto.flac
    │   │   └── EnV - EnV - Streetlights.flac
    │   └── test3
    │       └── test3.txt
    ├── fail.mp3
    ├── マクロスMACROSS 82-99 - SAILORWAVE
    │   ├── cover.jpg
    │   ├── マクロスMACROSS 82-99 - SAILORWAVE - 01 NEW DAWN.flac
    │   ├── マクロスMACROSS 82-99 - SAILORWAVE - 02 戦場.flac
    │   └── マクロスMACROSS 82-99 - SAILORWAVE - 10 新宿区 JAZZ POINT J.flac
    ├── Moby-Play_2014-HD_Remaster
    │   ├── 01-01-Moby-Honey_2014_Remastered_Version-SMR.flac
    │   ├── 01-02-Moby-Find_My_Baby_2014_Remastered_Version-SMR.flac
    │   └── 01-08-Moby-Natural_Blues_2014_Remastered_Version-SMR.flac
    ├── test.txt
    └── waw
        ├── 1.wav
        ├── 2.wav
        ├── 3.wav
        └── cover.jpg
```

### Use of boris
```
~ $ boris -i zikIn/ -o zikOut
⏺  Copy mode ⏺

🔎  Listing files...

[ 2% ][ 0s ]
[ 2% ][ 0s ]	↳ zikIn/
[ 5% ][ 0s ]		♫ 01_02 Daddy Cool.flac	➜	zikOut/Pop Rock/Boney M_/The Essential Boney M_/01-02 Daddy Cool.flac
[ 7% ][ 0s ]		♫ 02-01-Boney_M-Felicidad_Margherita-LLS.flac	➜	zikOut/Pop Rock/Boney M_/The Essential Boney M_/02-01 Felicidad (Margherita) (7_ Version).flac
[ 10% ][ 0s ]		♫ 02_02 Gadda-Da-Vida (7_ Version) (Full Length).flac	➜	zikOut/Pop Rock/Boney M_/The Essential Boney M_/02-02 Gadda-Da-Vida (7_ Version) (Full Length).flac
[ 12% ][ 0s ]		❌ error: opening tags of file "zikIn/fail.mp3": EOF
[ 15% ][ 0s ]		🖺 fail.mp3	➜	zikOut/Pop Rock/Boney M_/The Essential Boney M_/fail.mp3
[ 15% ][ 0s ]		🖺 test.txt	➜	zikOut/Pop Rock/Boney M_/The Essential Boney M_/test.txt
[ 17% ][ 0s ]
[ 17% ][ 0s ]	↳ zikIn/Andy Fox - Tempus Fugit/
[ 20% ][ 0s ]		♫ Andy Fox - Tempus Fugit - 01 Tempus Fugit.flac	➜	zikOut/Unknonw genre/Andy Fox/Tempus Fugit/01 Tempus Fugit.flac
[ 22% ][ 0s ]		♫ Andy Fox - Tempus Fugit - 02 Daydream (feat. Vanna).flac	➜	zikOut/Unknonw genre/Andy Fox/Tempus Fugit/02 Daydream (feat. Vanna).flac
[ 25% ][ 0s ]		♫ Andy Fox - Tempus Fugit - 06 Second Life.flac	➜	zikOut/Unknonw genre/Andy Fox/Tempus Fugit/06 Second Life.flac
[ 30% ][ 0s ]		🖺 Andy Fox - Tempus Fugit - back.png	➜	zikOut/Unknonw genre/Andy Fox/Tempus Fugit/Andy Fox - Tempus Fugit - back.png
[ 30% ][ 0s ]		🖺 cover.jpg	➜	zikOut/Unknonw genre/Andy Fox/Tempus Fugit/cover.jpg
[ 32% ][ 0s ]
[ 32% ][ 0s ]	↳ zikIn/Annella - Perfume - Single/
[ 35% ][ 0s ]		♫ Annella - Perfume - Single - 01 Perfume.flac	➜	zikOut/Unknonw genre/Annella/Perfume - Single/01 Perfume.flac
[ 40% ][ 0s ]		🖺 cover.jpg	➜	zikOut/Unknonw genre/Annella/Perfume - Single/cover.jpg
[ 40% ][ 0s ]		🖺 test2.txt	➜	zikOut/Unknonw genre/Annella/Perfume - Single/test2.txt
[ 42% ][ 0s ]
[ 42% ][ 0s ]	↳ zikIn/Brandon - Neon Haze/
[ 45% ][ 0s ]		♫ Brandon - Neon Haze - 01 Neon Haze.flac	➜	zikOut/Unknonw genre/Brandon/Neon Haze/01 Neon Haze.flac
[ 47% ][ 0s ]		♫ Brandon - Neon Haze - 03 Time Machine.flac	➜	zikOut/Unknonw genre/Brandon/Neon Haze/03 Time Machine.flac
[ 50% ][ 0s ]		♫ Brandon - Neon Haze - 04 Arcade.flac	➜	zikOut/Unknonw genre/Brandon/Neon Haze/04 Arcade.flac
[ 52% ][ 0s ]		🖺 cover.jpg	➜	zikOut/Unknonw genre/Brandon/Neon Haze/cover.jpg
[ 55% ][ 0s ]
[ 55% ][ 0s ]	↳ zikIn/Brandon - Neon Haze/env/
[ 57% ][ 0s ]		♫ EnV - EnV - Shinto.flac	➜	zikOut/Unknonw genre/EnV/Unknonw album/EnV - Shinto.flac
[ 60% ][ 0s ]		♫ EnV - EnV - Streetlights.flac	➜	zikOut/Unknonw genre/EnV/Unknonw album/EnV - Streetlights.flac
[ 62% ][ 0s ]
[ 62% ][ 0s ]	↳ zikIn/Brandon - Neon Haze/test3/
[ 65% ][ 0s ]		⚠ No tagged music file found, moving file(s) below to unrecognised music directory.
[ 65% ][ 0s ]		🖺 test3.txt	➜	unorganizedFiles/Brandon - Neon Haze/test3/test3.txt
[ 67% ][ 0s ]
[ 67% ][ 0s ]	↳ zikIn/Moby-Play_2014-HD_Remaster/
[ 70% ][ 0s ]		♫ 01-01-Moby-Honey_2014_Remastered_Version-SMR.flac	➜	zikOut/Electro/Moby/Play [2014 - HD Remaster]/01 Honey (2014 Remastered Version).flac
[ 72% ][ 0s ]		♫ 01-02-Moby-Find_My_Baby_2014_Remastered_Version-SMR.flac	➜	zikOut/Electro/Moby/Play [2014 - HD Remaster]/02 Find My Baby (2014 Remastered Version).flac
[ 75% ][ 0s ]		♫ 01-08-Moby-Natural_Blues_2014_Remastered_Version-SMR.flac	➜	zikOut/Electro/Moby/Play [2014 - HD Remaster]/08 Natural Blues (2014 Remastered Version).flac
[ 77% ][ 0s ]
[ 77% ][ 0s ]	↳ zikIn/waw/
[ 87% ][ 0s ]		⚠ No tagged music file found, moving file(s) below to unrecognised music directory.
[ 87% ][ 0s ]		🖺 1.wav	➜	unorganizedFiles/waw/1.wav
[ 87% ][ 1s ]		🖺 2.wav	➜	unorganizedFiles/waw/2.wav
[ 87% ][ 1s ]		🖺 3.wav	➜	unorganizedFiles/waw/3.wav
[ 87% ][ 1s ]		🖺 cover.jpg	➜	unorganizedFiles/waw/cover.jpg
[ 90% ][ 1s ]
[ 90% ][ 1s ]	↳ zikIn/マクロスMACROSS 82-99 - SAILORWAVE/
[ 95% ][ 1s ]		♫ マクロスMACROSS 82-99 - SAILORWAVE - 01 NEW DAWN.flac	➜	zikOut/Unknonw genre/マクロスMACROSS 82-99/SAILORWAVE/01 NEW DAWN.flac
[ 97% ][ 1s ]		♫ マクロスMACROSS 82-99 - SAILORWAVE - 02 戦場.flac	➜	zikOut/Unknonw genre/マクロスMACROSS 82-99/SAILORWAVE/02 戦場.flac
[ 100% ][ 1s ]		♫ マクロスMACROSS 82-99 - SAILORWAVE - 10 新宿区 JAZZ POINT J.flac	➜	zikOut/Unknonw genre/マクロスMACROSS 82-99/SAILORWAVE/10 新宿区 JAZZ POINT J.flac
[ 100% ][ 1s ]		🖺 cover.jpg	➜	zikOut/Unknonw genre/マクロスMACROSS 82-99/SAILORWAVE/cover.jpg
```

### After boris
```
~ $ tree
.
├── zikIn
│   └── ...
├── unorganizedFiles
│   ├── Brandon - Neon Haze
│   │   └── test3
│   │       └── test3.txt
│   └── waw
│       ├── 1.wav
│       ├── 2.wav
│       ├── 3.wav
│       └── cover.jpg
└── zikOut
    ├── Electro
    │   └── Moby
    │       └── Play [2014 - HD Remaster]
    │           ├── 01 Honey (2014 Remastered Version).flac
    │           ├── 02 Find My Baby (2014 Remastered Version).flac
    │           └── 08 Natural Blues (2014 Remastered Version).flac
    ├── Pop Rock
    │   └── Boney M_
    │       └── The Essential Boney M_
    │           ├── 01-02 Daddy Cool.flac
    │           ├── 02-01 Felicidad (Margherita) (7_ Version).flac
    │           ├── 02-02 Gadda-Da-Vida (7_ Version) (Full Length).flac
    │           ├── fail.mp3
    │           └── test.txt
    └── Unknonw genre
        ├── Andy Fox
        │   └── Tempus Fugit
        │       ├── 01 Tempus Fugit.flac
        │       ├── 02 Daydream (feat. Vanna).flac
        │       ├── 06 Second Life.flac
        │       ├── Andy Fox - Tempus Fugit - back.png
        │       └── cover.jpg
        ├── Annella
        │   └── Perfume - Single
        │       ├── 01 Perfume.flac
        │       ├── cover.jpg
        │       └── test2.txt
        ├── Brandon
        │   └── Neon Haze
        │       ├── 01 Neon Haze.flac
        │       ├── 03 Time Machine.flac
        │       ├── 04 Arcade.flac
        │       └── cover.jpg
        ├── EnV
        │   └── Unknonw album
        │       ├── EnV - Shinto.flac
        │       └── EnV - Streetlights.flac
        └── マクロスMACROSS 82-99
            └── SAILORWAVE
                ├── 01 NEW DAWN.flac
                ├── 02 戦場.flac
                ├── 10 新宿区 JAZZ POINT J.flac
                └── cover.jpg
```

## How does it work
boris will scan all your music files, and for each folder found, it will:
* lists all directs files at it's root
	* if a music file is found in the list, the new path for it will be determined using it's metadatas combined to a template.
	* if a non music file is found (an image cover for example):
		* if at least one music file as been found, it will be moved to the same location.
		* if none music files exist in the directory (for instance a list of waw file), they will be moved appart to a special directory.
* repeat this process for each sub folders

Please note that a file is considered to be a music file only if it supports tags. Supported files types are:
* MP3
* M4A
* M4B
* M4P
* ALAC
* FLAC
* OGG

## Configuration
boris can be configured by using a configuration file, or/and using command line arguments (that will overide configuration file directive if used). Default configuration values can bee seen using the -s flags. It can also be used to generate a base config file: `boris -s > myConf.hjson`

### Configuration file
See the conf.hjson example file for detailed explanations.

### Command line arguments
```
-c string
		config file
-d	delete input music folder after organizing
-i string
		input music directory (default ".")
-m	move file instead of copying them
-o string
		output music directory (default "organizedMusicLibrary")
-p	preview change
-r string
		replacement for <>:"\/|?* (default "_")
-s	show default config, in hjson config file format
-t string
		organizing template (default see below)
-u string
		unorganized files directory where folders will be moved if they contains
		files that does not have any tags (default "unorganizedFiles")
```

## Templating
- Templating will be used to compute the new path for each music file, based on it's metadatas/tags.
- The templating syntax is based on golang emplate syntax (https://golang.org/pkg/text/template/).
- Every tabulation and newline will be removed.
- Available function/atribute are:
	* Album
	* Artist
	* AlbumArtist
	* Composer
	* Genre
	* Year
	* Disk
	* DiskTotal
	* Track
	* TrackTotal
	* Ext (file extension)
	* OriginalFilename (with extension)
	*	golang predefined function (such as printf)

### Detailed example of the default template:
The following template wil organize music file using this patern:<br>
`genre/albumArtist/albumName/discNumber-trackNumber title.extension`
```
{{if .Genre}}     #
  {{.Genre}}      # This section will be replaced by the (in order of priority):
{{else}}          #   musical genre | default text "Unknonw genre"
  Unknonw genre   #
{{end}}           #

/ # create sub folder

{{if .AlbumArtist}} #
  {{.AlbumArtist}}  #
{{else if .Artist}} # This section will be replaced by the (in order of priority):
  {{.Artist}}       #   album artist | artist | default text "Unknonw artist"
{{else}}            #
  Unknonw artist    #
{{end}}             #

/ # create sub folder

{{if .Album}}   #
  {{.Album}}    # This section will be replaced by the (in order of priority):
{{else}}        #   album | default text "Unknonw album"
  Unknonw album #
{{end}}         #

/ # create sub folder

{{if gt .DiscTotal 1}}        # If the album is composed of more than one disc,
  {{.Disc | printf "%02d"}}-  # this section will be replaced by the disc number
{{end}}                       # on two digits, followed by a hyphen.

{{if .Track}}                         # If a track number exists, this section will be
  {{.Track | printf "%02d"}} {{/**/}} # replaced by it, on two digit, followed by a space.
{{end}}                               # (The trailling comment section is just here to make
                                      # the space more visible, and work around some text
                                      # editor that removes trailling spaces on save.)

{{if .Title}}           #
  {{.Title}}{{.Ext}}    # This section will be replaced by the (in order of priority):
{{else}}                #   track title followed by the original file extension | original filename
  {{.OriginalFilename}} #
{{end}}                 #
```

## About
boris is made by Joseph Caillet and is released under GPL-3.0.<br>
View source code, report bug, contribute here: https://github.com/JosephCaillet/boris.