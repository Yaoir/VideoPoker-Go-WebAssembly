This is the initial public release of October 2018

This code may be helpful for people learning how to write WebAssembly apps in Go.
Just please keep in mind that this is not a finished program, and support
for WebAssembly is new.

You will see code for features that are under development, and some things
may not work right. Some code may not be as concise as possible, which may
change in the future, or remain as-is for some good reason.

Also, I am not a big supporter of "idiomatic" discipline. My aim is to write
the best code I can, rather than to follow ideas that happen to be popular,
fashionable, or dogmatic.

There are some odd behaviors you may notice in the app:

Starting the game
-----------------

So far, it seems to work well on either Linux or Windows, using an up-to-date
version of Firefox, Opera, or Chrome. Browser support on mobile devices is
more limited. You may need to wait some seconds (up to 12 seconds on my
old tablet), or the game may not finish loading at all (usually when using
Firefox). This may be due to a bug in either the Go wasm compilation or
in one or more browser's support for WebAssembly (at least, the kind
produced by the Go compiler). One of the reasons I'm publishing this
before it is more complete is so that this issue can be demonstrated
and discussed.

Ending the game
---------------

Upon a q ("quit") or e ("exit") keypress, the game shows an end-of-game
message, then just stops and becomes completely unresponsive. That is
because the Go program exited!

To start a new game, the page must be reloaded. This will work
more elegantly in a future release.

### Introduction

This is the Go version of videopoker.

I did this because I was wondering how long it would take to tranlate a 1000-line C program into Go.
It took about 8 hours. It was not difficult at all, but some things had to be rewritten into the way things are done in Go: no pointer arithmetic, and different handling of Unicode, random number generation, and STDIN/STDOUT input and output.

### Quick Start

To get it working, just run

```
$ go build videopoker.go
$ videopoker
```

See the manual page below for directions on how to play.

### Project Description

A text-based video poker for Linux.

Many variants of video poker found in casinos are included as options.

The cards can be displayed as black or color characters, and also as Unicode
card faces.

It's a great way to practice your strategy for fun, or before going to a casino.
The text-based playing action allows for quick play and more rapid training than
with a GUI application.

### Screenshots

Colored Text (`-u1` option)

![screenshot 1](/images/VideoPoker-01-PlainText.png)

Unicode Suits (`-u2` option, default)

![screenshot 2](/images/VideoPoker-02-UnicodeSuits.png)

Unicode Card Faces (`-u3` option)

![screenshot 3](/images/VideoPoker-03-FullUnicode.png)

### Manual Page

```
VIDEOPOKER(6)                    Games Manual                    VIDEOPOKER(6)



NAME
       videopoker - a text-based video poker game

SYNOPSIS
       videopoker  [-b]  [-b1] [-g game] [-is #chips] [-k 5-char-string] [-mh]
       [-q] [-uD] [-v]

DESCRIPTION
       videopoker is a text-based video poker game. It has an  efficient  user
       interface  that  takes a little time to learn, but allows fast play and
       requires only one hand to operate.

DISCLAIMER
       By default, videopoker is intended to closely match the behavior of 9/6
       Jacks  or  Better video poker machines in casinos, and an option allows
       selection of other games and pay tables. However, the author is not  an
       expert on gaming, and no guarantee whatsoever is made that videopoker´s
       behavior is an exact match to that of any  other  video  poker.  Please
       take  that  into  careful  consideration before trying out a real video
       poker machine.

HOW TO PLAY
       Start the game and rest the fingers of your right hand on the  keyboard
       as  when  touch  typing.  Your thumb will be on the space bar, and your
       index finger through little finger will be on the keys  j,  k,  l,  and
       semicolon (;).

       A five-card hand will be displayed. To hold cards, type the keys corre‐
       sponding to the cards:



           SPACE   Leftmost card
           j       Second card from left
           k       Middle card
           l       Second card from right
           ;       Rightmost card



       The keys may be typed in any order, and a key can be entered more  than
       once  to toggle the held/discarded state of the card. The backspace key
       may be used to undo mistakes.

       Then type the Enter key to deal. Discarded cards are redealt,  and  the
       final  hand  is shown, along with how it is recognized as either a win‐
       ning or losing hand, and the new score.

       The game will continue until you either quit or run out of chips.

       To quit, type either q or e, then the Enter key.

       Changing the bet

       To increase your bet from the default, type b, followed by a digit from
       1 to 5, along with the keys to hold cards. For example,

       lb4;

       will  hold  the  rightmost  two  cards,  and increase the bet to 4x. By
       itself,

       b3

       will increase the bet to 3x, but not hold any cards.

       The increase takes effect starting with the  next  hand  and  stays  in
       effect  until changed using the same method. The only exception is when
       the number of chips is less than the bet, in  which  case  the  bet  is
       automatically  reduced  to make it equal to the number of chips remain‐
       ing, where it will stay until you change it.

OPTIONS
       -b

       ("Bold") Use boldface text output to display the cards. (When using the
       -u1 and -u2 options only.)

       -b1

       ("Bet  1")  Use a minimum bet of one chip, rather than the default ten.
       Typically used along with the -is option.

       -g name

       ("Game") Play a different variant of video poker. The  default  is  9/6
       Jacks  or Better, but you can change it to another by specifying one of
       the strings in the left column of the following list:



           aa        All American
           tens      Tens or Better
           jb95      9/5 Jacks or Better
           jb86      8/6 Jacks or Better
           jb85      8/5 Jacks or Better
           jb75      7/5 Jacks or Better
           jb65      6/5 Jacks or Better



       -is integer

       ("Initial Score") Start the game with integer chips,  rather  than  the
       default 1000. integer may be in the range of 1 to 100,000.

       -k "string"

       ("Keys")  Replace  the  standard  "  jkl;"  input keys with the ones in
       string, which must be 5 characters long and not contain a q or w  char‐
       acter.

       -mh

       ("Mark  Held")  Print  + characters under the cards that are held. This
       option is helpful for new players, and also  when  testing  the  string
       supplied with the -k option.

       -q

       ("Quiet")  Don´t  print  the  banner  at the start of the game, nor the
       final message.

       -uD

       ("Unicode") Control Unicode output. D is a digit from 0 to 3.



           -u0   As with -u1, except all text is displayed in the default
                 color. (For monochrome terminals, or any terminals that
                 do not support ANSI escape codes.)
           -u1   No Unicode. Suits are displayed with the characters C,
                 D, H, and S. Diamonds and hearts are displayed in red,
                 clubs and spades are displayed in the default text color.
           -u2   (default) Display suits using Unicode playing card
                 suit symbols.
           -u3   (Experimental) Display entire cards as Unicode characters.



       In order for the -u1 or -u2 options to work, videopoker needs to use  a
       font  that  contains  the necessary Unicode playing card characters. On
       Linux, the monospace and DejaVu fonts may be used.

       For -u2 to work, the system needs to be able to properly  display  Uni‐
       code playing card characters.

       Unicode playing card faces appear no larger than other text characters.
       When using the -u2 option, increase the font size as  necessary  (i.e.,
       24 pixels or larger).

       -v

       ("Version") Print the version number and exit.

BUGS
       When  played  with  the  (experimental)  -u3  option,  cards may not be
       equally spaced apart. This may be due  to  improper  handling  of  ANSI
       escape  codes  (used to change text colors) and/or improper handling of
       Unicode playing card characters by various virtual terminals.

VERSION
       This manual page is for version 1.0 of the program.

AUTHOR
       Jay Ts
       (http://jayts.com)

COPYRIGHT
       Copyright 2016 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)



Jay Ts                           December 2016                   VIDEOPOKER(6)
```
