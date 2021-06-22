package main

//--------------------------
type GPatch int

const (
	GAME_NONE         GPatch = iota
	GAME_IWAD                = 1 << 0 // i.e. for Doom, DooM 2, Heretic, Hexen, HexenDD, Plutonia, TNT, Strife
	GAME_SHAREWARE           = 1 << 1 // All the Doom 1 sharewares.
	GAME_FREEDOOM            = 1 << 2 // i.e. for Freedoom Phase 1/2 and FreeDM
	GAME_HACX                = 1 << 3 // i.e. for HacX 1.0 - 1.2 (no support for 2.0 yet as it's still not released)
	GAME_CHEX_QUEST_3        = 1 << 4
	GAME_STRIFE_VE           = 1 << 5 // i.e. for Strife: Veteran Edition (heard that 1.0/1.1 are still around out there...)
	GAME_SIGIL               = 1 << 6 // SIGIL by John Romero
)

type GStatus int

const (
	NOT_FINAL GStatus = iota // is not the final version, so display errors
	IS_FINAL                 // Is final, is fine!
	IS_ALPHA                 // Special case of Alpha WADs.
)

//--------------------------

// WadInfo : all WAD data returned from the program
type WadInfo struct {
	MD5Hash      string
	Version      string
	Status       GStatus
	Game         GPatch
	PWADRequires string // If the official PWAD requires an IWAD to run
	Additional   string // If I need to display an additionnal message for this IWAD.

}

var (
	IWADInfo_Doom         []WadInfo
	IWADInfo_Doom2        []WadInfo
	IWADInfo_FinalDoom    []WadInfo
	IWADInfo_Heretic      []WadInfo
	IWADInfo_Hexen        []WadInfo
	IWADInfo_MasterLevels []WadInfo
	IWADInfo_Strife       []WadInfo
	IWADInfo_SVE          []WadInfo
	IWADInfo_FreeDoom     []WadInfo
	IWADInfo_Misc         []WadInfo // PWAD and addons

	// Patching messages
	iErrors = 0
)
