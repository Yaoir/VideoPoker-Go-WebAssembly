// Video Poker Game for WebAssembly/Go
//
// version 1.0
package main

// This is the game engine for the Video Poker web app

// Note to Reader:
// There are some things in this file that are non-idiomatic Go code.
// Some of them are here because this is translated from C code, and I need
// to keep things matched up to that for maintainability.
// Two things to watch for: I may sometimes use K&R comments (/* comment */) as regular
// comments, and I put empty // comments at the beginnings of blocks, to preserve
// the formatting of the original code.

// Functions that start with GUI_ (usually GUI_update_<thing>) are in main.go
// They are for modifying elements in the GUI interface to WebAssembly

// fmt.Printf() prints to the browser's Developer Tools debug console,
// allowing the game to be played in text mode.

import (
	"fmt"
	"math/rand"
	"os"
	"time"
	)

/* Replacement for C library random() function */

var randomgen *rand.Rand

func srandom() {
//
	randomgen = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func random() int {
//
	return randomgen.Int()
}

/* ASCII key codes */

// NOTE: iota is not used here because the numbers must match the keys */
// The game currently uses only a few keys, but the full list is here for
// easy development of new features

const (
	key_ctrlAT = 0
	key_ctrlA = 1
	key_ctrlB = 2
	key_ctrlC = 3
	key_ctrlD = 4
	key_ctrlE = 5
	key_ctrlF = 6
	key_ctrlG = 7
	key_ctrlH = 8
	key_ctrlI = 9
	key_ctrlJ = 10
	key_Enter = 10
	key_ctrlK = 11
	key_ctrlL = 12
	key_ctrlM = 13
	key_Return = 13
	key_ctrlN = 14
	key_ctrlO = 15
	key_ctrlP = 16
	key_ctrlQ = 17
	key_ctrlR = 18
	key_ctrlS = 19
	key_ctrlT = 20
	key_ctrlU = 21
	key_ctrlV = 22
	key_ctrlW = 23
	key_ctrlX = 24
	key_ctrlY = 25
	key_ctrlZ = 26
	key_ESC = 27
	key_ctrlBackslash = 28
	key_ctrlRightsquarebracket = 29
	key_ctrlCaret = 30
	key_ctrlUnderscore = 31
	key_space = 32
	key_exclam = 33
	key_doublequote = 34
	key_hash = 35
	key_dollar = 36
	key_percent = 37
	key_ampersand = 38
	key_singlequote = 39
	key_leftparen = 40
	key_rightparen = 41
	key_asterisk = 42
	key_plus = 43
	key_comma = 44
	key_minus = 45
	key_dot = 46
	key_slash = 47
	key_0 = 48
	key_1 = 49
	key_2 = 50
	key_3 = 51
	key_4 = 52
	key_5 = 53
	key_6 = 54
	key_7 = 55
	key_8 = 56
	key_9 = 57
	key_colon = 58
	key_semicolon = 59
	key_lessthan = 60
	key_equal = 61
	key_greaterthan = 62
	key_questionmark = 63
	key_at = 64
	key_A = 65
	key_B = 66
	key_C = 67
	key_D = 68
	key_E = 69
	key_F = 70
	key_G = 71
	key_H = 72
	key_I = 73
	key_J = 74
	key_K = 75
	key_L = 76
	key_M = 77
	key_N = 78
	key_O = 79
	key_P = 80
	key_Q = 81
	key_R = 82
	key_S = 83
	key_T = 84
	key_U = 85
	key_V = 86
	key_W = 87
	key_X = 88
	key_Y = 89
	key_Z = 90
	key_leftsquarebracket = 91
	key_backslash = 92
	key_rightsquarebracket = 93
	key_caret = 94
	key_underscore = 95
	key_backtick = 96
	key_a = 97
	key_b = 98
	key_c = 99
	key_d = 100
	key_e = 101
	key_f = 102
	key_g = 103
	key_h = 104
	key_i = 105
	key_j = 106
	key_k = 107
	key_l = 108
	key_m = 109
	key_n = 110
	key_o = 111
	key_p = 112
	key_q = 113
	key_r = 114
	key_s = 115
	key_t = 116
	key_u = 117
	key_v = 118
	key_w = 119
	key_x = 120
	key_y = 121
	key_z = 122
	key_openbrace = 123
	key_verticalbar = 124
	key_closebrace = 125
	key_tilde = 126
	key_del = 127
)

func key_action(key byte) {
//
        switch key {
                case key_ctrlJ: fallthrough
                case key_ctrlK:
                case key_ctrlL:
//                case key_ctrlM: fallthrough
                case key_Return:	// '\r'
                        if state == Draw { draw() } else { deal() }
                case key_space:
                        toggle_hold(0)
                case key_1:
                        do_bet(key)
                case key_2:
                        do_bet(key)
                case key_3:
                        do_bet(key)
                case key_4:
                        do_bet(key)
                case key_5:
                        do_bet(key)
                case key_semicolon:
                        toggle_hold(4)
                case key_A:
                        changegame(AllAmerican)
                case key_B:
                        changegame(TensOrBetter)
                case key_C:
                        changegame(BonusPoker)
                case key_D:
                        changegame(DoubleBonus)
                case key_E:
                        changegame(DoubleBonusBonus)
                case key_F:
                        changegame(JacksOrBetter)
                case key_G:
                        changegame(JacksOrBetter95)
                case key_H:
                        changegame(JacksOrBetter86)
                case key_I:
                        changegame(JacksOrBetter85)
                case key_e:
                        do_quit()
                case key_j:
                        toggle_hold(1)
                case key_k:
                        toggle_hold(2)
                case key_l:
                        toggle_hold(3)
                case key_q:
                        do_quit()
                default:
        }
}

const VERSION = "videopoker 1.0"

var hands int = 0	// number of hands played

/* The number of cards in the hand */

const CARDS = 5

/* The number of cards in the deck */

const CARDSINDECK = 52

const (
	CLUBS	= iota
	DIAMONDS
	HEARTS
	SPADES
	NUMSUITS	// 4
	)

/* one-character suit designations */

var suitname [NUMSUITS]string = [NUMSUITS]string {
	"c",
	"d",
	"h",
	"s",
	}

/* Card values. NOTE: They are one lower than number on card faces */

const (
	TWO	= iota+1	// TWO   == 1
	THREE			// THREE == 2
	FOUR			// FOUR  == 3
	FIVE			// ...
	SIX
	SEVEN
	EIGHT
	NINE
	TEN  /* needed for recognizing royal flush, or tens or better (TEN), or jacks or better (JACK) */
	JACK /* needed for recognizing royal flush, or tens or better (TEN), or jacks or better (JACK) */
	QUEEN
	KING
	ACE  /* needed for recognizing Ace-low straight (Ace, 2, 3, 4, 5) */
)

/* the card type, for holding infomation about the deck of cards */

type card struct
{
	index int	/* cards value, minus 1 */
	sym string	/* textual appearance */
	uc string	/* Unicode value for the card */
	suit int	/* card's suit (see just below) */
	gone int	/* true if it's been dealt */
}

/* The standard deck of 52 cards */

var deck [CARDSINDECK]card = [CARDSINDECK]card {
/*	index, card, filename, suit, gone */
	{ TWO,   " 2", "02-clubs.png", CLUBS, 0 },
	{ THREE, " 3", "03-clubs.png", CLUBS, 0 },
	{ FOUR,  " 4", "04-clubs.png", CLUBS, 0 },
	{ FIVE,  " 5", "05-clubs.png", CLUBS, 0 },
	{ SIX,   " 6", "06-clubs.png", CLUBS, 0 },
	{ SEVEN, " 7", "07-clubs.png", CLUBS, 0 },
	{ EIGHT, " 8", "08-clubs.png", CLUBS, 0 },
	{ NINE,  " 9", "09-clubs.png", CLUBS, 0 },
	{ TEN,   "10", "10-clubs.png", CLUBS, 0 },
	{ JACK,  " J", "11-clubs.png", CLUBS, 0 },
	{ QUEEN, " Q", "12-clubs.png", CLUBS, 0 },
	{ KING,  " K", "13-clubs.png", CLUBS, 0 },
	{ ACE,   " A", "01-clubs.png", CLUBS, 0 },

	{ TWO,   " 2", "02-diamonds.png", DIAMONDS, 0 },
	{ THREE, " 3", "03-diamonds.png", DIAMONDS, 0 },
	{ FOUR,  " 4", "04-diamonds.png", DIAMONDS, 0 },
	{ FIVE,  " 5", "05-diamonds.png", DIAMONDS, 0 },
	{ SIX,   " 6", "06-diamonds.png", DIAMONDS, 0 },
	{ SEVEN, " 7", "07-diamonds.png", DIAMONDS, 0 },
	{ EIGHT, " 8", "08-diamonds.png", DIAMONDS, 0 },
	{ NINE,  " 9", "09-diamonds.png", DIAMONDS, 0 },
	{ TEN,   "10", "10-diamonds.png", DIAMONDS, 0 },
	{ JACK,  " J", "11-diamonds.png", DIAMONDS, 0 },
	{ QUEEN, " Q", "12-diamonds.png", DIAMONDS, 0 },
	{ KING,  " K", "13-diamonds.png", DIAMONDS, 0 },
	{ ACE,   " A", "01-diamonds.png", DIAMONDS, 0 },

	{ TWO,   " 2", "02-hearts.png", HEARTS, 0 },
	{ THREE, " 3", "03-hearts.png", HEARTS, 0 },
	{ FOUR,  " 4", "04-hearts.png", HEARTS, 0 },
	{ FIVE,  " 5", "05-hearts.png", HEARTS, 0 },
	{ SIX,   " 6", "06-hearts.png", HEARTS, 0 },
	{ SEVEN, " 7", "07-hearts.png", HEARTS, 0 },
	{ EIGHT, " 8", "08-hearts.png", HEARTS, 0 },
	{ NINE,  " 9", "09-hearts.png", HEARTS, 0 },
	{ TEN,   "10", "10-hearts.png", HEARTS, 0 },
	{ JACK,  " J", "11-hearts.png", HEARTS, 0 },
	{ QUEEN, " Q", "12-hearts.png", HEARTS, 0 },
	{ KING,  " K", "13-hearts.png", HEARTS, 0 },
	{ ACE,   " A", "01-hearts.png", HEARTS, 0 },

	{ TWO,   " 2", "02-spades.png", SPADES, 0 },
	{ THREE, " 3", "03-spades.png", SPADES, 0 },
	{ FOUR,  " 4", "04-spades.png", SPADES, 0 },
	{ FIVE,  " 5", "05-spades.png", SPADES, 0 },
	{ SIX,   " 6", "06-spades.png", SPADES, 0 },
	{ SEVEN, " 7", "07-spades.png", SPADES, 0 },
	{ EIGHT, " 8", "08-spades.png", SPADES, 0 },
	{ NINE,  " 9", "09-spades.png", SPADES, 0 },
	{ TEN,   "10", "10-spades.png", SPADES, 0 },
	{ JACK,  " J", "11-spades.png", SPADES, 0 },
	{ QUEEN, " Q", "12-spades.png", SPADES, 0 },
	{ KING,  " K", "13-spades.png", SPADES, 0 },
	{ ACE,   " A", "01-spades.png", SPADES, 0 },
}

// transparent card, used at start
var transparent_card = card{ ACE, " A", "nocard.png", HEARTS, 0 }

/* state of the deal/draw button */

const (
	Deal = iota
	Draw
)

/* state is either Deal or Draw, depending on what the deal/draw button's current function is */

var state int = Deal

var msg_deal string = "To continue, click on Deal New Hand"
var msg_draw string = "Click the cards to hold, then click on Draw Cards"

/* The hand. It holds five cards. */

var hand [5]card

/* sorted hand, for internal use when recognizing winners */

var shand [5]card

/* hold[] keeps track of which cards in hand are being held */

var hold [5]int

/* initial number of chips held */

const INITCHIPS = 1000		// make sure content of HTML <... id="score"> matches this
var score int = INITCHIPS

/* minimum and maximum swing of score during this game */

var score_low int = INITCHIPS
var score_high int = INITCHIPS

/* The games starts with a bet of 10, the minimum allowed */

const INITMINBET = 10

var minbet int = INITMINBET
var bet int = INITMINBET

/* number of chips or groups of 10 chips bet */

var betmultiplier int = 1

/* The various video poker games that are supported */

const (
	AllAmerican = iota
	TensOrBetter
	BonusPoker
	DoubleBonus
	DoubleBonusBonus
	JacksOrBetter	// default
	JacksOrBetter95
	JacksOrBetter86
	JacksOrBetter85
	JacksOrBetter75
	JacksOrBetter65
	NUMGAMES
)

/*
	The game in play. Default is Jacks or Better,
	which is coded into initialization of static data
*/

var game int = JacksOrBetter

var gamenames [NUMGAMES]string = [NUMGAMES]string {
	"All American",
	"Tens or Better",
	"Bonus Poker",
	"Double Bonus",
	"Double Bonus Bonus",
	"Jacks or Better",
	"9/5 Jacks or Better",
	"8/6 Jacks or Better",
	"8/5 Jacks or Better",
	"7/5 Jacks or Better",
	"6/5 Jacks or Better",
}

/* Functions that recognize winning hands .
   These functions operate on the *sorted* hand,
   which makes it much easier */

/*
	Flush:
	returns true if the sorted hand is a flush
*/

func flush() bool {
//
	if shand[0].suit == shand[1].suit &&
	   shand[1].suit == shand[2].suit &&
	   shand[2].suit == shand[3].suit &&
	   shand[3].suit == shand[4].suit { return true }

	return false
}

/*
	Straight:
	returns true if the sorted hand is a straight
*/

func straight() bool {
//
	if shand[1].index == shand[0].index + 1 &&
	   shand[2].index == shand[1].index + 1 &&
	   shand[3].index == shand[2].index + 1 &&
	   shand[4].index == shand[3].index + 1 { return true }

	if shand[4].index == ACE   &&
	   shand[0].index == TWO   &&
	   shand[1].index == THREE &&
	   shand[2].index == FOUR  &&
	   shand[3].index == FIVE { return true }

	return false
}

/*
	Four of a kind:
	the middle 3 all match, and the first or last matches those
*/

func four() bool {
//
	if (shand[1].index == shand[2].index &&
	    shand[2].index == shand[3].index ) &&
	   ( shand[0].index == shand[2].index ||
	     shand[4].index == shand[2].index) { return true }

	return false
}

/*
	Full house:
	3 of a kind and a pair
*/

func full() bool {
//
	if shand[0].index == shand[1].index &&
	  (shand[2].index == shand[3].index &&
	   shand[3].index == shand[4].index) { return true }

	if shand[3].index == shand[4].index &&
	  (shand[0].index == shand[1].index &&
	   shand[1].index == shand[2].index) { return true }

	return false
}

/*
	Three of a kind:
	it can appear 3 ways
*/

func three() bool {
//
	if shand[0].index == shand[1].index &&
	   shand[1].index == shand[2].index { return true }

	if shand[1].index == shand[2].index &&
	   shand[2].index == shand[3].index { return true }

	if shand[2].index == shand[3].index &&
	   shand[3].index == shand[4].index { return true }

	return false
}

/*
	Two pair:
	it can appear in 3 ways
*/

func twopair() bool {
//
	if ((shand[0].index == shand[1].index) && (shand[2].index == shand[3].index)) ||
	   ((shand[0].index == shand[1].index) && (shand[3].index == shand[4].index)) ||
	   ((shand[1].index == shand[2].index) && (shand[3].index == shand[4].index)) { return true }

	return false
}

/*
	Two of a kind (pair), jacks or better
	or if the game is Tens or Better, 10s or better.
*/

func two() bool {
//
	var min int = JACK

	if game == TensOrBetter { min = TEN }

	if shand[0].index == shand[1].index && shand[1].index >= min { return true }
	if shand[1].index == shand[2].index && shand[2].index >= min { return true }
	if shand[2].index == shand[3].index && shand[3].index >= min { return true }
	if shand[3].index == shand[4].index && shand[4].index >= min { return true }

	return false
}

/* returns type of hand */

func recognize() int {
//
	var i, j, f int
	var min int = INVALID
	var tmp [CARDS]card
	var st, fl bool	/* both are auto-initialized to 0 */

	/* Sort hand into sorted hand (shand) */

	/* make copy of hand */
	for i = 0; i < CARDS; i++ { tmp[i] = hand[i] }

	/* sort it */
	for i = 0; i < CARDS; i++ {
	//
		/* put lowest card in hand into next place in shand */

		for j = 0; j < CARDS; j++ {
		//
			if tmp[j].index <= min {
			//
				min = tmp[j].index
				f = j
			}
		}

		shand[i] = tmp[f]
		tmp[f].index = INVALID	/* larger than any card */
		min = INVALID
	}

	/* royal and straight flushes, straight, and flush */

	fl = flush()
	st = straight()

	if st && fl && shand[0].index == TEN { return ROYAL }
	if st && fl { return STRFL }
	if four() { return FOURK }
	if full() { return FULL }
	if fl { return FLUSH }
	if st { return STR }
	if three() { return THREEK }
	if twopair() { return TWOPAIR }
	if two() { return PAIR }

	/* Nothing */

	return NOTHING
}

/*
	This bunch of consts is used to index into
	paytable[] and hand[], so make sure the two match.
*/

const (
	ROYAL = iota
	STRFL
	FOURK
	FULL
	FLUSH
	STR
	THREEK
	TWOPAIR
	PAIR
	NOTHING
	/* the number of the above: */
	NUMHANDTYPES
	)

var paytable [NUMHANDTYPES]int = [NUMHANDTYPES]int {
	800,	/* royal flush: 800 */
	50,	/* straight flush: 50 */
	25,	/* 4 of a kind: 25 */
	9,	/* full house: 9 */
	6,	/* flush: 6 */
	4,	/* straight: 4 */
	3,	/* 3 of a kind: 3 */
	2,	/* two pair: 2 */
	1,	/* jacks or better: 1 */
	0,	/* nothing */
}

var handname [NUMHANDTYPES]string = [NUMHANDTYPES]string {
	"Royal Flush    ",
	"Straight Flush ",
	"Four of a Kind ",
	"Full House     ",
	"Flush          ",
	"Straight       ",
	"Three of a Kind",
	"Two Pair       ",
	"Pair           ",
	"Nothing        ",
}

const INVALID = 100	/* higher than any valid card index */

func changegame(g int) {
//
        /* End this game */
        final_score()

        /* Start new game */
        game = g
        score = INITCHIPS
        score_low = INITCHIPS
        score_high = INITCHIPS
        bet = INITMINBET
        minbet = INITMINBET
        betmultiplier = 1
        setgame(game)
	GUI_update_gamename(gamenames[g])
        starting_banner()
        deal()
}

func setgame(game int) {
//
	switch game {
	//
		case JacksOrBetter95:
			paytable[FLUSH] = 5
		case JacksOrBetter86:
			paytable[FULL] = 8
		case JacksOrBetter85:
			paytable[FULL] = 8
			paytable[FLUSH] = 5
		case JacksOrBetter75:
			paytable[FULL] = 7
			paytable[FLUSH] = 5
		case JacksOrBetter65:
			paytable[FULL] = 6
			paytable[FLUSH] = 5
		case AllAmerican:
			paytable[FULL] = 8
			paytable[FLUSH] = 8
			paytable[STR] = 8
			paytable[PAIR] = 1
		case TensOrBetter:
			/* pay table same as JacksOrBetter65 */
			paytable[FULL] = 6
			paytable[FLUSH] = 5
	}
}

/* set minimum bet to 1 chip */

func bet1() {
//
        /* set minimum bet */
        minbet = 1; bet = 1
}

/* show held cards */

func showheld() {
//
        var i int
	var pm string

        for i = 0; i < CARDS; i++ {
	//
		GUI_update_hold(i)
                if(hold[i] != 0) { pm = " +" } else { pm = "  " }
                fmt.Printf("%s  ", pm)
        }
        fmt.Printf("\n")
}

/* Display the hand */

func showhand() {
//
	var i int

	GUI_update_hand()	// update card images on web page

	/* First line: show cards */

	for i = 0; i < CARDS; i++ {
	//
		fmt.Printf("%s%s ", hand[i].sym, suitname[hand[i].suit])
	}

	fmt.Printf("\n")

	/* Second line: show which cards are held */

	showheld()
}

func deal() {
//
	var i int
	var crd int

	/* initialize deck */
	for i = 0; i < CARDSINDECK; i++ { deck[i].gone = 0 }

	/* initialize hold[] */
	clear_holds()

	score -= bet
	GUI_update_score(score)

	/* To test Ace-low straights, uncomment this section and the test: label below */
/*
	hand[0] = deck[0]; deck[0].gone = 1
	hand[1] = deck[1]; deck[1].gone = 1
	hand[2] = deck[2]; deck[2].gone = 1
	hand[3] = deck[3]; deck[3].gone = 1
	hand[4] = deck[12]; deck[12].gone = 1
	goto test
*/

	/* To test royal flushes, uncomment this section and the test: label below */

/*
	hand[0] = deck[34]; deck[34].gone = 1
	hand[1] = deck[35]; deck[35].gone = 1
	hand[2] = deck[36]; deck[36].gone = 1
	hand[3] = deck[37]; deck[37].gone = 1
	hand[4] = deck[38]; deck[38].gone = 1
	// use these indices into deck[]:
	// clubs: 8-12, diamonds: 21-25, hearts: 34-38, spades: 47-51
	goto test
*/
	/* To test straight flushes, uncomment this section and the test: label below */
/*
	hand[2] = deck[3]; deck[3].gone = 1
	hand[3] = deck[4]; deck[4].gone = 1
	hand[1] = deck[5]; deck[5].gone = 1
	hand[0] = deck[6]; deck[6].gone = 1
	hand[4] = deck[7]; deck[7].gone = 1
	goto test
*/

	for i = 0; i < CARDS; i++ {
	//
		/* find a card not already dealt */

		for {
		//
			crd = random()%CARDSINDECK
			if deck[crd].gone == 0 { break }
		}

		deck[crd].gone = 1
		hand[i] = deck[crd]
	}

// test:
	// enter Draw state

	GUI_update_handname(" ")
	showhand()
	state = Draw
	GUI_update_button()
	GUI_update_message(msg_draw)
}

func starting_banner() {
//
	/* Before starting play, print the name of the game in green */

	fmt.Printf("\n%s\n\n",gamenames[game])
}

func final_score() {
//
	var msg string;

	msg = fmt.Sprintf("You quit with %d chips after playing %d hands",score,hands)
	GUI_update_message(msg)
	fmt.Printf("%s\n",msg)
	fmt.Printf("Range: %d - %d\n", score_low, score_high)
}

func do_quit() {
//
        final_score()
        os.Exit(0)
}

func do_bet(digit byte) {
//
	var s string

	// allow changing bet only before new hand is dealed
	if state == Draw { return }

        betmultiplier = int(digit) - key_0
        b := betmultiplier * minbet
	if b > score {
	//
		s = fmt.Sprintf("You don't have that many chips")
	} else {
	//
		bet = b
		s = fmt.Sprintf("Bet changed to %d chips",bet)
	}
	GUI_update_message(s)
	fmt.Printf("%s\n",s)
        showhand()
}

func clear_holds() {
//
	for i := 0; i < CARDS; i++ {
	//
		hold[i] = 0
		GUI_update_hold(i)
	}
}

func toggle_hold(i int) {
//
	if state == Deal { return }
        /* flip bit to hold/discard it */
        hold[i] ^= 1
	GUI_update_hold(i)
        /* redisplay hand */
        showhand()
}

func draw() {
//
        var i int
        var crd int
	var msg string

        /* replace cards not held */

        for i = 0; i < CARDS; i++ {
	//
                if hold[i] == 0 {
		//
			for {
			//
				crd = random()%CARDSINDECK
				if deck[crd].gone == 0 { break }
			}

                        deck[crd].gone = 1
                        hand[i] = deck[crd]
                }
        }

        /* print final hand */

        showhand()

        /* recognize and score hand */

        i = recognize()

        score += paytable[i] * bet

        fmt.Printf("%s  ",handname[i])
	GUI_update_handname(handname[i])
        fmt.Printf("%d\n\n",score)
	GUI_update_score(score)

        hands++

        if score < score_low  { score_low  = score }
        if score > score_high { score_high = score }

        if score < bet {
	//
                for ; score < bet && betmultiplier > 1; {
		//
                        betmultiplier--;
                        bet = minbet * betmultiplier
                }

                if score < bet {
		//
			msg = fmt.Sprintf("You ran out of chips after playing %d hands", hands)
			GUI_update_message(msg)
			fmt.Printf("%s\n",msg)
//			fmt.Printf("You ran out of chips after playing %d hands.\n", hands)
//			if score_high > INITCHIPS { fmt.Printf("At one point, you had %d chips.\n", score_high) }
			os.Exit(0)
                } else {
		//
// TODO: use dialog (alert) for this:
			msg = fmt.Sprintf("You are low on chips. Your bet has been reduced to %d",bet)
			GUI_update_message(msg)
			fmt.Printf("%s\n\n",msg)
//			fmt.Printf("You are low on chips. Your bet has been reduced to %d\n\n", bet)
// TODO: update bet buttons
                }
        }

	state = Deal
	GUI_update_button()
	GUI_update_message(msg_deal)
}

// The following just starts (initializes) the game
// Game play is event driven, and is handled by mouse
// and keyboard event callbacks that act through the
// key_action() function.

func videopoker() {
//
	// initialize random number generator
	srandom()

	// initialize the hand to transparent cards
	for i := 0; i < CARDS; i++ {
		hand[i] = transparent_card
	}

	// Start the game
	starting_banner()
}

// EOF
