// Video Poker - a single page web app in Go/WebAssembly
// build: GOOS=js GOARCH=wasm go build -o main.wasm main.go videopoker-web.go

package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

// The generalized way to change the text content of an HTML element, identified by an id property in the HTML tag

// In JavaScript, this would be
// document.getElementById(id).textContent = text

func GUI_set_text(id, text string) {
	js.Global().Get("document").Call("getElementById", id).Set("textContent", text)
}

// The following use the same method. It may look inefficient, but it's done this way to have
// code more self-documenting and simpler to build up

func GUI_update_message(msg string) {
	js.Global().Get("document").Call("getElementById", "message").Set("textContent", msg)
}

func GUI_update_gamename(name string) {
	js.Global().Get("document").Call("getElementById", "gamename").Set("textContent", name)
}

func GUI_update_handname(name string) {
	js.Global().Get("document").Call("getElementById", "hand").Set("textContent", name)
}

// The generalized way to change the CSS style of an HTML element, identified by an id property in the HTML tag

// In JavaScript, this would be
// document.getElementById(id).style = style

func GUI_set_style(id, style string) {
	js.Global().Get("document").Call("getElementById", id).Set("style", style)
}

// And the way it's actually done for now:

func GUI_button_visible() {
	js.Global().Get("document").Call("getElementById", "drawbutton").Set("style", "display: block;")
}

// Change the text in the Deal/Draw button that appears underneath the five cards of the poker hand

func GUI_update_button() {
	var label string
	if state == Draw { label = "Draw Cards" } else { label = "Deal New Hand" }
	js.Global().Get("document").Call("getElementById", "drawbutton").Set("textContent", label)
}

// Change the card images

// In JavaScript, this would be
// document.getElementById(id).src = filename
// to modify the <img src="{filename}"> property

func GUI_update_hand() {
	var i int
	for i = 0; i < 5; i++ {
		cardN := fmt.Sprintf("card%d",i+1)
		filename := fmt.Sprintf("img/%s",hand[i].uc)
		js.Global().Get("document").Call("getElementById", cardN).Set("src", filename)
	}
}

func GUI_update_score(score int) {
	score_alpha := strconv.Itoa(score)
	js.Global().Get("document").Call("getElementById", "score").Set("textContent", score_alpha)
}

// The green bar underneath each card that appears when the card is held.
// It's implemented by putting a padded border only at the bottom of the card's image, not at the other sides
// Then it can be turned on or off using these styles:

var css_card_hold string = "border-color: transparent transparent #0c0 transparent;"
var css_card_free string = "border-color: transparent transparent transparent transparent;"

// Clear the held status of all of the cards.
// This is done when dealing a new hand.

func hold_none () {
	card_style := css_card_free
	js.Global().Get("document").Call("getElementById", "card1").Set("style", card_style)
	js.Global().Get("document").Call("getElementById", "card2").Set("style", card_style)
	js.Global().Get("document").Call("getElementById", "card3").Set("style", card_style)
	js.Global().Get("document").Call("getElementById", "card4").Set("style", card_style)
	js.Global().Get("document").Call("getElementById", "card5").Set("style", card_style)
}

// Callbacks for clicking on the cards

// Hold or un-hold a card. This toggles when the card is clicked.

func GUI_update_hold(n int) {
	cardN := fmt.Sprintf("card%d",n+1)  // Card numbers in the HTML range from 1 to 5, not 0 to 4
	if hold[n] == 1 {
		// set cardN style for holding the card
		js.Global().Get("document").Call("getElementById", cardN).Set("style", css_card_hold)
	} else {
		// set cardN style for un-holding the card
		js.Global().Get("document").Call("getElementById", cardN).Set("style", css_card_free)
	}
}

// The following 5 functions are done very simplistically, and could
// also be implemented as hold(n) in the HTML, with a hold(args []js.Value)
// function to get the card number from the argument

func hold1(this js.Value, args []js.Value) interface{} {
	toggle_hold(0)
	GUI_update_hold(0)
	return nil
}

func hold2(this js.Value, args []js.Value) interface{} {
	toggle_hold(1)
	GUI_update_hold(1)
	return nil
}

func hold3(this js.Value, args []js.Value) interface{} {
	toggle_hold(2)
	GUI_update_hold(2)
	return nil
}

func hold4(this js.Value, args []js.Value) interface{} {
	toggle_hold(3)
	GUI_update_hold(3)
	return nil
}

func hold5(this js.Value, args []js.Value) interface{} {
	toggle_hold(4)
	GUI_update_hold(4)
	return nil
}

// Callback for the Deal/Draw button (the label changes based on the state variable

func deal_or_draw(this js.Value, args []js.Value) interface{} {
	// Clicking on the button causes it to have focus.
	// If the button were to retain focus, a space bar press would trigger this button,
	// resulting in an Enter/Return key event when the user is intending to hold card #1.
	// The following does a this.blur() to avoid focus on the button.

	js.Global().Get("document").Call("getElementById", "drawbutton").Call("blur")
	key_action(byte('\r'))	// process it as a press of the Enter key, which does the same thing
	return nil
}

// Callbacks for change of game
// (not implemented yet)

func jacks_or_better(this js.Value, args []js.Value) interface{} {
	// set game to Jacks or Better (default)
	return nil
}

// TODO: Callbacks for other change game functions,
// or a changegame(n) function

// Registered event handler for keypress events

// gkey() is currently unused.
// It is for alternative use with a registered event handler
// rather than the onkeypress event in the HTML <body> tag

func gkey(this js.Value, args []js.Value) interface{} {
	var c rune

	event := args[0]
	rs := []rune(event.Get("key").String())
	// IMPORTANT: The Enter/Return key shows up here as the string "Enter"
	if len(rs) == 1 {
		c = rs[0]
	} else {
		if string(rs) == "Enter" { c = '\r' }
	}

	key_action(byte(c))

	return nil
}

// hkey() is almost identical t gkey(), and is also unused at the moment

func hkey(this js.Value, args []js.Value) interface{} {
	var c rune

	rs := []rune(args[0].Get("key").String())
	// IMPORTANT: The Enter/Return key shows up here as the string "Enter",
	// so it gets converted here to a '\r'
	if len(rs) == 1 {
		c = rs[0]
	} else {
		if string(rs) == "Enter" { c = '\r' }
	}

	key_action(byte(c))

	return nil
}

// key() is the handler that is used in the current version.
// It is connected to the
// HTML <body onkeypress="return key(event);"> callback event handler for keypress events

func key(this js.Value, arg []js.Value) interface{} {
	var c rune

	// call event.preventDefault() and event.stopPropagation()
	// to avoid the keyboard events triggering other things in the browser.
	// One example is Firefox's "Search for text when you start typing".
	// Another common one is that pressing the space bar when a button is selected is the same
	// as clicking on the button, and in this app, the space bar is used for toggling selection
	// of the leftmost card.

	arg[0].Call("stopPropagation")
	arg[0].Call("preventDefault")

	rs := []rune(arg[0].Get("key").String())
	// IMPORTANT: The Enter/Return key shows up here as the string "Enter",
	// so it gets converted here to a '\r'
	if len(rs) == 1 {
		c = rs[0]
	} else {
		if string(rs) == "Enter" { c = '\r' }
	}

	key_action(byte(c))

	return nil
}

// connect events (from the JavaScript engine) to the callback functions in this file

func register_callbacks() {

	// Event handler for keyboard events

//	Use this instead of key() if you don't have a callback for keypress in the <body> tag
//	doc := js.Global().Get("document")
//	kbd_event := js.FuncOf(gkey)
//	doc.Call("addEventListener","keypress", kbd_event, "true")

	// event handler for HTML <body> keypress event callback

	// using regular callback:
//	js.Global().Set("hkey", js.FuncOf(hkey))

// Go 1.11 had NewEventCallback(), which allowed capturing events as exclusive handler:
//	js.Global().Set("key", js.NewEventCallback(js.PreventDefault|js.StopPropagation,key))
// or,
//	kbcb := js.NewEventCallback(js.PreventDefault|js.StopPropagation,key)
//	js.Global().Set("key", kbcb)

	// In Go 1.12, js.NewEventCallback() no longer exists. It was in version 1.11
	// only as a workaround because NewCallback() returned an *asynchronous* callback.
	// FuncOf() returns a *synchronous* callback, so we can call
	// (JS) event.PreventDefault() and (JS) event.StopPropagation() directly. Like this:
// TODO: update comment to reference calling event.stopPropagation() and event.preventDefault() in key()

	js.Global().Set("key", js.FuncOf(key))

	// clicks on card images, left to right
	js.Global().Set("hold1", js.FuncOf(hold1))
	js.Global().Set("hold2", js.FuncOf(hold2))
	js.Global().Set("hold3", js.FuncOf(hold3))
	js.Global().Set("hold4", js.FuncOf(hold4))
	js.Global().Set("hold5", js.FuncOf(hold5))

	// for clicks on the Deal/Draw button
	js.Global().Set("deal_or_draw", js.FuncOf(deal_or_draw))

	// click on Change Game button
	js.Global().Set("jacks_or_better", js.FuncOf(jacks_or_better))

//	Template for adding more callbacks
//	js.Global().Set("X", js.FuncOf(X))
}

func main() {
	register_callbacks()

	// startup message for the Developer Tools console
	fmt.Printf("WebAssembly program started\n")

	// Start videopoker

	// Up to here, the HTML displays a "The game is loading. Please wait." message with the Deal/Draw button hidden.
	// Now that the game is running, change those.
	GUI_button_visible()	// make Deal button visible
	GUI_update_message(msg_deal)

	videopoker()	// Initialize and start the game. See videopoker-web.go

	// Game play is event driven.
	// The event handlers in this file call key_action() in videopoker-web.go

	// To keep the app running, we need to keep main() from exiting.
	// An empty select statement is a simple way to block this goroutine.
	select{}
}
