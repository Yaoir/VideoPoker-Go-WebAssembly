### UNDER CONSTRUCTION

This README file is being working on. Please ignore most of what you see here, and try again later.

### Introduction

This is the initial public release of November 2018

To play the game:

http://jayts.com/vp

There are some odd behaviors you may notice in the app:

### Starting the game

So far, it seems to work well on either Linux or Windows, using a recent
version of Firefox, Opera, or Chrome.

Browser support on mobile devices is more limited. You may need to wait
some seconds (up to 14 seconds for Firefox on my old tablet) for the game to load.
Chrome and Opera seem to work well.

Firefox on mobile is problematic. The app may start properly and work fine the first
time the page is loaded, but reloading the page may result in the WebAssembly
app not starting. (This is a suspected bug in Go's WebAssembly support or the Firefox browser.)
At the worst, you may need to clear the browser cache and
restart Firefox to get it to work again. Before that, you can try just
restarting the browser, and please let me know if that worked for you.

One of the reasons I'm publishing this before it is more complete
is so that this issue can be demonstrated and discussed.

### Ending the game

Upon a q ("quit") or e ("exit") keypress, the game shows an end-of-game message,
then just stops and becomes completely unresponsive. That is because the Go program
exited. This is a holdover from the console version, and will work more elegantly
in a future release.

For now, reload the page to start a new game.

### How to Play


### Building from Source

### Project Description

It's a great way to practice your strategy for fun, or before going to a casino.

Many variants of video poker are included as options. (Currently accessible
only using the A-I keys.)

### Screenshots

Example of good luck:

![screenshot 1](/images/VideoPoker-StraightFlush.png)

### Manual Page

### Description

       Video Poker is a video poker game. It can be played with the mouse, and allows keyboard input for fast play.

### Disclaimer

       By default, videopoker is intended to closely match the behavior of 9/6
       Jacks  or  Better video poker machines in casinos, and an option allows
       selection of other games and pay tables. However, the author is not  an
       expert on gaming, and no guarantee whatsoever is made that videopoker´s
       behavior is an exact match to that of any  other  video  poker.  Please
       take  that  into  careful  consideration before trying out a real video
       poker machine.

### How to Play

       Start the game and rest the fingers of your right hand on the  keyboard
       as  when  touch  typing.  Your thumb will be on the space bar, and your
       index finger through little finger will be on the keys  j,  k,  l,  and
       semicolon (;).

       A five-card hand will be displayed. To hold cards, type the keys corre‐
       sponding to the cards:



```
           SPACE   Leftmost card
           j       Second card from left
           k       Middle card
           l       Second card from right
           ;       Rightmost card
```



       The keys may be typed in any order, and a key can be entered more  than
       once  to toggle the held/discarded state of the card. The backspace key
       may be used to undo mistakes.

       Then type the Enter key to deal. Discarded cards are redealt,  and  the
       final  hand  is shown, along with how it is recognized as either a win‐
       ning or losing hand, and the new score.

       The game will continue until you either quit or run out of chips.

       To quit, type either q or e

```
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


```

### Version
       This manual page is for version 1.0 of the program.

### Author
       Jay Ts
       (http://jayts.com)

### Copyright
       Copyright 2016-2018 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)
