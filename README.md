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
