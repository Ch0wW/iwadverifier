package main

const (
	AddentumMPDowngrade = `If you want to use it on multiplayer source ports, you'll need to patch it using Peter Vaskovics's tool, available below:
	• Windows binaries: http://downloads.zdaemon.org/iwadpatcher-1.2-bin.zip
	• Source code: https://github.com/petervas/iwadpatcher`

	AddentumDoomBethesda = `This IWAD won't work at all for multiplayer, as it contains several map differences.
• You will need to use the original wad instead, found in the following directory:
	- "<yoursteamfolder>\steamapps\common\Ultimate DOOM\base\DOOM.WAD" for the Steam version,
	- "<bethesdafolder>\games\Ultimate Doom\base\DOOM.WAD" for the Bethesda Launcher version.`

	AddentumDoomIIBethesda = `• This IWAD won't work at all for multiplayer, as it contains several map differences.
• You will need to use the original wad instead, found in the following directory:
	- "<yoursteamfolder>\steamapps\common\DOOM 2\base\DOOM2.WAD" for the Steam version,
	- "<bethesdafolder>\games\Doom 2\base\DOOM2.WAD" for the Bethesda Launcher version.`
)

//
//=======================
// PopulateIWADInfos
// Populate the IWAD lists with their MD5 and versions
//=======================
//
func PopulateIWADInfos() {
	Populate_Doom()
	Populate_DoomII()
	Populate_FinalDOOM()
	Populate_HereticHexen()
	Populate_Strife()
	Populate_FreeDoom()
	Populate_Misc()
}

func Populate_Doom() {

	// DOOM/UDOOM population
	IWADInfo_Doom = []WadInfo{
		WadInfo{
			MD5Hash: "90facab21eede7981be10790e3f82da2",
			Version: "DOOM 1.0 Shareware",
			Game:    GAME_SHAREWARE,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "52cbc8882f445573ce421fa5453513c1",
			Version: "DOOM 1.1 Shareware",
			Game:    GAME_SHAREWARE,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "981b03e6d1dc033301aa3095acc437ce",
			Version: "DOOM 1.1 Registered",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "30aa5beb9e5ebfbbe1e1765561c08f38",
			Version: "DOOM 1.2 Shareware",
			Game:    GAME_SHAREWARE,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "792fd1fea023d61210857089a7c1e351",
			Version: "DOOM 1.2 Registered",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "a21ae40c388cb6f2c3cc1b95589ee693",
			Version: "DOOM 1.4 Shareware Beta",
			Game:    GAME_SHAREWARE,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "e280233d533dcc28c1acd6ccdc7742d4",
			Version: "DOOM 1.5 Shareware Beta",
			Game:    GAME_SHAREWARE,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "762fd6d4b960d4b759730f01387a50a1",
			Version: "DOOM 1.6 Shareware Beta",
			Game:    GAME_SHAREWARE,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash:    "c428ea394dc52835f2580d5bfd50d76f",
			Version:    "DOOM 1.666 Shareware",
			Game:       GAME_SHAREWARE,
			Status:     NOT_FINAL,
			Additional: "Get the latest shareware version of DooM at https://www.doomworld.com/idgames/idstuff/doom/doom19s",
		},
		WadInfo{
			MD5Hash: "54978d12de87f162b9bcc011676cb3c0",
			Version: "DOOM 1.666 Registered",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash:    "5f4eb849b1af12887dec04a2a12e5e62",
			Version:    "DOOM 1.8 Shareware",
			Game:       GAME_SHAREWARE,
			Status:     NOT_FINAL,
			Additional: "Get the latest shareware version of DooM at https://www.doomworld.com/idgames/idstuff/doom/doom19s",
		},
		WadInfo{
			MD5Hash: "11e1cd216801ea2657723abc86ecb01f",
			Version: "DOOM 1.8 Registered",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "f0cefca49926d00903cf57551d901abe",
			Version: "DOOM 1.9 Shareware",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "1cd63c5ddff1bf8ce844237f580e9cf3",
			Version: "DOOM 1.9 Registered",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "dae77aff77a0491e3b7254c9c8401aa8",
			Version: "DOOM for Pocket PC (1.9)",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "c4fe9fd920207691a9f493668e0a2083",
			Version: "The Ultimate DOOM (v1.9)",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash:    "0c8758f102ccafe26a3040bee8ba5021",
			Version:    "The Ultimate DOOM (1.9) - Xbox Version",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: AddentumMPDowngrade,
		},
		WadInfo{
			MD5Hash:    "72286ddc680d47b9138053dd944b2a3d",
			Version:    "The Ultimate DOOM (1.9) - XBLA Version",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: AddentumMPDowngrade,
		},
		WadInfo{
			MD5Hash:    "fb35c4a5a9fd49ec29ab6e900572c524",
			Version:    "The Ultimate DOOM (1.9) - BFG Edition",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: AddentumMPDowngrade,
		},
		WadInfo{
			MD5Hash:    "8517c4e8f0eef90b82852667d345eb86",
			Version:    "DOOM 1.9 - Bethesda version (Unity) - 2020_08_21 Build #13735 doom",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: AddentumDoomBethesda,
		},
	}
}

func Populate_DoomII() {

	// DOOM/UDOOM population
	IWADInfo_Doom2 = []WadInfo{
		WadInfo{
			MD5Hash: "30e3c2d0350b67bfbf47271970b74b2f",
			Version: "DOOM II 1.666",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash:    "d9153ced9fd5b898b36cc5844e35b520",
			Version:    "DOOM II 1.666 (German Edition)",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "Despite being the latest version for the German release, it is strongly recommended to update it to v1.9.",
		},
		WadInfo{
			MD5Hash: "ea74a47a791fdef2e9f2ea8b8a9da13b",
			Version: "DOOM II 1.7",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "d7a07e5d3f4625074312bc299d7ed33f",
			Version: "DOOM II 1.7a",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "c236745bb01d89bbb866c8fed81b6f8c",
			Version: "DOOM II 1.8",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash:    "3cb02349b3df649c86290907eed64e7b",
			Version:    "DOOM II 1.8 (French Edition)",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "It is the latest release of the French edition. However, please keep in mind this is not the latest official release widely used to play online.",
		},
		WadInfo{
			MD5Hash: "25e1459ca71d321525f84628f45ca8cd",
			Version: "DOOM II 1.9",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash:    "a793ebcdd790afad4a1f39cc39a893bd",
			Version:    "DOOM II 1.9 - Xbox Version",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: AddentumMPDowngrade,
		},
		WadInfo{
			MD5Hash:    "43c2df32dc6c740cb11f34dc5ab693fa",
			Version:    "DOOM II 1.9 - XBLA Version",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: AddentumMPDowngrade,
		},
		WadInfo{
			MD5Hash:    "c3bea40570c23e511a7ed3ebcd9865f7",
			Version:    "DOOM II 1.9 - BFG Edition",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: AddentumMPDowngrade,
		},
		WadInfo{
			MD5Hash: "f617591a6c5d07037eb716dc4863e26b",
			Version: "DOOM II 1.9 - Xbox 360 BFG Edition",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "9640fc4b2c8447bbd28f2080725d5c51",
			Version: "DOOM II 1.9 - Tapwave Zodiac Version",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash:    "8ab6d0527a29efdc1ef200e5687b5cae",
			Version:    "DOOM II 1.9 - Bethesda version (Unity) - 2020_08_21 Build #13736 doom2",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: AddentumDoomIIBethesda,
		},
	}
}

func Populate_FinalDOOM() {
	// DOOM/UDOOM population
	IWADInfo_FinalDoom = []WadInfo{
		WadInfo{
			MD5Hash: "75c8cf89566741fa9d22447604053bd7",
			Version: "Final DOOM: The Plutonia Experiment 1.9",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash:    "3493be7e1e2588bc9c8b31eab2587a04",
			Version:    "Final DOOM: The Plutonia Experiment 1.9 (id Anthology Version)",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "May not be fully compatible with online servers or demos",
		},
		WadInfo{
			MD5Hash: "4e158d9953c79ccf97bd0663244cc6b6",
			Version: "Final DOOM: TNT: Evilution 1.9",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash:    "1d39e405bf6ee3df69a8d2646c8d5c49",
			Version:    "Final DOOM: TNT: Evilution 1.9 (id Anthology Version)",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "May not be fully compatible with online servers or demos",
		},
	}
}

func Populate_HereticHexen() {
	// Heretic population
	IWADInfo_Heretic = []WadInfo{
		WadInfo{
			MD5Hash:    "fc7eab659f6ee522bb57acc1a946912f",
			Version:    "Heretic Wide-Area Beta",
			Status:     IS_FINAL,
			Additional: "This is the latest Beta version of Heretic.",
		},
		WadInfo{
			MD5Hash: "023b52175d2f260c3bdc5528df5d0a8c",
			Version: "Heretic 1.0 Shareware",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash:    "ae779722390ec32fa37b0d361f7d82f8",
			Version:    "Heretic 1.2 Shareware",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "This is the latest Shareware version of Heretic.",
		},

		WadInfo{
			MD5Hash: "66d686b1ed6d35ff103f15dbd30e0341",
			Version: "Heretic: Shadow of the Serpent Riders (1.3)",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "3117e399cdb4298eaa3941625f4b2923",
			Version: "Heretic 1.0",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "1e4cb4ef075ad344dd63971637307e04",
			Version: "Heretic 1.2",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},
	}

	IWADInfo_Hexen = []WadInfo{
		WadInfo{
			MD5Hash: "abb033caf81e26f12a2103e1fa25453f",
			Version: "Hexen 1.1",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "b2543a03521365261d0a0f74d5dd90f0",
			Version: "Hexen 1.0",
			Game:    GAME_IWAD,
			Status:  NOT_FINAL,
		},

		// MAC versions
		WadInfo{
			MD5Hash:    "b68140a796f6fd7f3a5d3226a32b93be",
			Version:    "Hexen 1.1 (MAC VERSION)",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "MAC release contains differences that are not compatible with the Windows version of the game!",
		},
		WadInfo{
			MD5Hash:    "b68140a796f6fd7f3a5d3226a32b93be",
			Version:    "Hexen Demo (MAC VERSION)",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "MAC release contains differences that are not compatible with the Windows version of the game!",
		},

		// hexdd.wad
		// Steam only distributes version 1.0
		WadInfo{
			MD5Hash:    "1077432e2690d390c256ac908b5f4efa",
			Version:    "Hexen: Deathkings of the Dark Citadel 1.0",
			Game:       GAME_IWAD,
			Status:     NOT_FINAL,
			Additional: "Requires Hexen 1.1",
		},
		WadInfo{
			MD5Hash:    "78d5898e99e220e4de64edaa0e479593",
			Version:    "Hexen: Deathkings of the Dark Citadel 1.1",
			Game:       GAME_IWAD,
			Status:     IS_FINAL,
			Additional: "Requires Hexen 1.1",
		},

		// OTHER IRREVELENT THINGS
		WadInfo{
			MD5Hash: "876a5a44c7b68f04b3bb9bc7a5bd69d6",
			Version: "Hexen Demo 1.0",
			Game:    GAME_IWAD,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "9178a32a496ff5befebfe6c47dac106c",
			Version: "Hexen Demo Beta",
			Game:    GAME_IWAD,
			Status:  IS_ALPHA,
		},
		WadInfo{
			MD5Hash: "c88a2bb3d783e2ad7b599a8e301e099e",
			Version: "Hexen Beta",
			Game:    GAME_IWAD,
			Status:  IS_ALPHA,
		},
	}
}

func Populate_Strife() {
	// Strife population
	IWADInfo_Strife = []WadInfo{
		WadInfo{
			MD5Hash:    "2fed2031a5b03892106e0f117f17901f",
			Version:    "Strife 1.2 - 1.31 Registered",
			Status:     IS_FINAL,
			Additional: "Your IWAD is up-to-date. However, the latest updates of Strife modify the executable. Please make sure it is the latest version.",
		},
		WadInfo{
			MD5Hash: "bb545b9c4eca0ff92c14d466b3294023",
			Version: "Strife 1.1 Shareware",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "8f2d3a6a289f5d2f2f9c1eec02b47299",
			Version: "Strife 1.0 Registered",
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash:    "de2c8dcad7cca206292294bdab524292",
			Version:    "Strife 1.0 Shareware",
			Status:     NOT_FINAL,
			Additional: "Download the latest shareware of Strife at https://www.doomworld.com/idgames/roguestuff/strife11",
		},

		// Additionnal IWADs
		WadInfo{
			MD5Hash:      "082234d6a3f7086424856478b5aa9e95",
			Version:      "Strife voice acting samples",
			Status:       IS_FINAL,
			PWADRequires: "Strife Registered",
		},
	}

	// Strife: Veteran Edition
	IWADInfo_SVE = []WadInfo{
		WadInfo{
			MD5Hash: "47958a4fea8a54116e4b51fc155799c0",
			Version: "Strife: Veteran Edition 1.2",
			Game:    GAME_STRIFE_VE,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "2c0a712d3e39b010519c879f734d79ae",
			Version: "Strife: Veteran Edition 1.1",
			Game:    GAME_STRIFE_VE,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "06a8f99b9b756ac908917c3868b8e3bc",
			Version: "Strife: Veteran Edition 1.0",
			Game:    GAME_STRIFE_VE,
			Status:  NOT_FINAL,
		},
	}
}

func Populate_FreeDoom() {
	IWADInfo_FreeDoom = []WadInfo{

		//---------------
		// First, the latest IWADs
		WadInfo{
			MD5Hash: "b36aa44a23045e503c19af4b4c438a78",
			Version: "Freedoom: Phase 1 0.12.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "ca9a4159a7833544a89144c7f5053412",
			Version: "Freedoom: Phase 2 0.12.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "d40c932a9183ded919afa89f4a729668",
			Version: "FreeDM 0.12.1",
			Game:    GAME_FREEDOOM,
			Status:  IS_FINAL,
		},

		//----------------
		// Now, list the outdated ones.

		// 0.12.0
		WadInfo{
			MD5Hash: "0c5f8ff45cc3538d368a0f8d8fc11ce3",
			Version: "Freedoom: Phase 1 0.12.0",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "83560b2963424fa4a2eb971194428bf8",
			Version: "Freedoom: Phase 2 0.12.0",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "3250aad8b1d40fb7b25b7df6573eb29f",
			Version: "FreeDM 0.12.0",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},

		// 0.11.3
		WadInfo{
			MD5Hash: "ea471a3d38fcee0fb3a69bcd3221e335",
			Version: "Freedoom: Phase 1 0.11.3",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "984f99af08f085e38070f51095ab7c31",
			Version: "Freedoom: Phase 2 0.11.3",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "87ee2494d921633420ce9bdb418127c4",
			Version: "FreeDM 0.11.3",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},

		// 0.11.2
		WadInfo{
			MD5Hash: "9352b09ae878dc52c6c18aa38acda6eb",
			Version: "FreeDM 0.11.2",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "6d00c49520be26f08a6bd001814a32ab",
			Version: "Freedoom: Phase 1 0.11.2",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "90832a872b5bb0aca4ca0b20419aad5d",
			Version: "Freedoom: Phase 2 0.11.2",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},

		// 0.11.1
		WadInfo{
			MD5Hash: "77ba9c0f75c32e4a729490688bb99241",
			Version: "FreeDM 0.11.1",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "35312e99d2473297aabe0602700bee8a",
			Version: "Freedoom: Phase 1 0.11.1",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "ec5b38b30ba2b70e278205776af3fbb5",
			Version: "Freedoom: Phase 2 0.11.1",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},

		// 0.11.0
		WadInfo{
			MD5Hash: "d76d3973c075b069ecb4e16dc9eacbb4",
			Version: "FreeDM 0.11",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "21a4707fc25d29edf4b098bd400c5c42",
			Version: "Freedoom: Phase 1 0.11",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "b1018017c61b06e33c11102d8bafaad0",
			Version: "Freedoom: Phase 2 0.11",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},

		// 0.10.1
		WadInfo{
			MD5Hash: "bd4f359f1963e388beda014c5548b420",
			Version: "FreeDM 0.10.1",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "91de79621a393a08c39a9ab2c034b766",
			Version: "Freedoom: Phase 1 0.10.1",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "dd9c9e73f5f50d3778c85573cd08d9a4",
			Version: "Freedoom: Phase 2 0.10.1",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},

		// 0.10.0
		WadInfo{
			MD5Hash: "9b8d72b59fd93b2b3e116149baa1b142",
			Version: "Freedoom: Phase 1 0.10",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "c5a4f2d38d78b251d8557cb2d93e40ee",
			Version: "Freedoom: Phase 2 0.10",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "f37b8b70e1394289a7ec404f67cdec1a",
			Version: "FreeDM 0.10",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},

		// 0.9
		WadInfo{
			MD5Hash: "aca90cf5ac36e996edc58bd0329b979a",
			Version: "Freedoom: Phase 1 0.9",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "8fa57dbc7687f84528eba39dde3a20e0",
			Version: "Freedoom: Phase 2 0.9",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "cbb27c5f3c2c44d34843cf63daa627f6",
			Version: "FreeDM 0.9",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},

		// 0.8
		WadInfo{
			MD5Hash: "30095b256dd3a1566bbc30286f72bc47",
			Version: "Ultimate Freedoom 0.8",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "e3668912fc37c479b2840516c887018b",
			Version: "Freedoom 0.8",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "05859098bf191899903ef343afba369d",
			Version: "FreeDM 0.8",
			Game:    GAME_FREEDOOM,
			Status:  NOT_FINAL,
		},
	}
}

func Populate_Misc() {

	// DOOM/UDOOM population
	IWADInfo_Misc = []WadInfo{
		WadInfo{
			MD5Hash:      "967d5ae23daf45196212ae1b605da3b0",
			Version:      "No Rest for the Living",
			Status:       IS_FINAL,
			PWADRequires: "DOOM II v1.9",
			Additional:   "You will need a limit-removing source port to be able to run this.",
		},
		// HACX
		WadInfo{
			MD5Hash: "65ed74d522bdf6649c2831b13b9e02b4",
			Version: "HacX 1.2",
			Game:    GAME_HACX,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "b7fd2f43f3382cf012dc6b097a3cb182",
			Version: "HacX 1.1",
			Game:    GAME_HACX,
			Status:  NOT_FINAL,
		},
		WadInfo{
			MD5Hash: "1511a7032ebc834a3884cf390d7f186e",
			Version: "HacX 1.0",
			Game:    GAME_HACX,
			Status:  NOT_FINAL,
		},

		// Chex Quest
		WadInfo{
			MD5Hash:    "25485721882b050afa96a56e5758dd52",
			Version:    "Chex Quest",
			Status:     IS_FINAL,
			Additional: "May require the DEHacked file if playing on a source port, available at https://www.doomworld.com/idgames/utils/exe_edit/patches/chexdeh",
		},

		WadInfo{
			MD5Hash:      "fdc4ffa57e1983e30912c006284a3e01",
			Version:      "Chex Quest 2",
			Status:       IS_FINAL,
			PWADRequires: "Chex Quest",
			Additional:   "May require the DEHacked file if playing on a source port: https://www.doomworld.com/idgames/utils/exe_edit/patches/chexdeh",
		},

		WadInfo{
			MD5Hash: "bce163d06521f9d15f9686786e64df13",
			Version: "Chex Quest 3 1.4",
			Game:    GAME_CHEX_QUEST_3,
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash:    "cb001c34e424687191f299cc1dff4d68",
			Version:    "Chex Quest 3 (ModDB - Unknown version)",
			Game:       GAME_CHEX_QUEST_3,
			Status:     NOT_FINAL,
			Additional: "This version of Chex Quest 3 has been released on ModDB. I have no other information about it.",
		},
		WadInfo{
			MD5Hash: "59c985995db55cd2623c1893550d82b3",
			Version: "Chex Quest 3 1.0",
			Game:    GAME_CHEX_QUEST_3,
			Status:  NOT_FINAL,
		},

		// SIGIL by John Romero
		WadInfo{
			MD5Hash:      "f53ffc4fb89e966839bb8d20c632819a",
			Version:      "SIGIL v1.0",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       NOT_FINAL,
		},
		WadInfo{
			MD5Hash:      "a775262ca0e423468196803b71a57a43",
			Version:      "SIGIL v1.0 (Compatibility WAD)",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       NOT_FINAL,
		},

		WadInfo{
			MD5Hash:      "1fe9daa0e837c7452eb2f91aac2cc983",
			Version:      "SIGIL v1.1",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       NOT_FINAL,
		},
		WadInfo{
			MD5Hash:      "c04912beab6aa82c114a19c976ec8c0d",
			Version:      "SIGIL v1.1 (Compatibility WAD)",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       NOT_FINAL,
		},

		WadInfo{
			MD5Hash:      "427ca995600970abcd2efcc684a64c88",
			Version:      "SIGIL v1.2",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       NOT_FINAL,
		},
		WadInfo{
			MD5Hash:      "9285e9cc2dbd87d238baab37d700c644",
			Version:      "SIGIL v1.2 (Compatibility WAD)",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       NOT_FINAL,
		},

		WadInfo{
			MD5Hash:      "743d6323cb2b9be24c258ff0fc350883",
			Version:      "SIGIL v1.21",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       IS_FINAL,
		},
		WadInfo{
			MD5Hash:      "573f3f178c76709f512089ed15484391",
			Version:      "SIGIL v1.21 (Compatibility WAD)",
			Game:         GAME_SIGIL,
			PWADRequires: "The Ultimate Doom v1.9",
			Status:       IS_FINAL,
		},

		WadInfo{
			MD5Hash:      "b424dcf46ae55a496c34ac37cce32646",
			Version:      "SIGIL - BucketHead soundtrack",
			Game:         GAME_SIGIL,
			PWADRequires: "SIGIL & The Ultimate DOOM v1.9",
			Status:       IS_FINAL,
		},
		WadInfo{
			MD5Hash:      "343faa815928c58faa08939a4502d5d2",
			Version:      "SIGIL - BucketHead soundtrack (Compatibility WAD)",
			Game:         GAME_SIGIL,
			PWADRequires: "SIGIL & The Ultimate DOOM v1.9",
			Status:       IS_FINAL,
		},
	}

	// Master Levels for DOOM II
	IWADInfo_MasterLevels = []WadInfo{
		// Standard PWADs
		WadInfo{
			MD5Hash: "cb03fd0cd84b10579c2b2b313199d4c1",
			Version: "Attack (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "a421ca18cea00a2696162f8d2a2beeca",
			Version: "Black Tower (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "18eb4ffb3094ddb690e62211dc6169a1",
			Version: "Bloodsea Keep (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "33493942592d764e7787fb0ad7d03044",
			Version: "Canyon (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "e7c273033376824edf95e1328261e7de",
			Version: "The Catwalk (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "77c179948df47a7a613bd1181c959892",
			Version: "The Combine (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "cbf714b499ebdef2682990eaf93fdb5f",
			Version: "The Fistula (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "f000701a3ed1f49249ee08550c03dfa5",
			Version: "The Garrison (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "a1efff02df6d873762ebac6b12358bbc",
			Version: "Geryon: 6th Canto of Inferno (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "787fa80fe9665c322f853b74838e77cc",
			Version: "Titan Manor (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "b4eaf844b135cc2a0058c6e0149b4408",
			Version: "Mephisto’s Maosoleum (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "aea597159dee96bcc58f3f9e3e586182",
			Version: "Minos’ Judgement: 4th Canto of Inferno (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "46f58580e7792f486c747cf1117c4ca1",
			Version: "Nessus: 5th Canto of Inferno (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "d560abb6d5719d46ebb47b27d7813a4b",
			Version: "Paradox (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "b572d518d564c7d7b6b259a726538cbb",
			Version: "Subspace (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "bb417f07804373415a6ed8e533242c3c",
			Version: "Subterra (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "65b4abcb74e7a386d5c024cf250d6336",
			Version: "“The Express Elevator to Hell” and “Bad Dream” (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "8474f6d663f04630de05ecac36b574d1",
			Version: "Trapped on Titan (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "a49dccebb5f32307246b7f32da121cf7",
			Version: "Vesperas: 7th Canto of Inferno (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},
		WadInfo{
			MD5Hash: "3c0874f2df3c06a002ee2a18aba0f0e8",
			Version: "Virgil’s Lead: 3rd Canto of Inferno (Master Levels for DOOM II)",
			Status:  IS_FINAL,
		},

		// Aggregate WAD from the Xbox 360 and PlayStation 3 BFG Edition
		WadInfo{
			MD5Hash:    "84cb8640f599c4a17c8eb526f90d2b7a",
			Version:    "Master Levels for DOOM II - Xbox 360 | PlayStation 3 BFG Edition",
			Status:     IS_FINAL,
			Additional: "Not generally compatible with demo files or source port multiplayer",
		},
	}
}
