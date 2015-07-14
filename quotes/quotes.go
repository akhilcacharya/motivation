package quotes

//Quote struct associates the text to the author
type Quote struct {
	Text   string
	Author string
}

//Quotes array lists all quotes available
//I would have parsed a text file but Shia told me to JUST DO IT
var Quotes = []Quote{
	Quote{"DO IT!", "Shia LaBeouf"},
	Quote{"JUST DO IT!", "Shia LaBeouf"},
	Quote{"Don't let your dreams be dreams!", "Shia LaBeouf"},
	Quote{"Yesterday you said tomorrow. So just do it!", "Shia LaBeouf"},
	Quote{"Make your dreams come true. Just do it!", "Shia LaBeouf"},
	Quote{"Some people dream of success, while you're going to wake up and work hard at it.", "Shia LaBeouf"},
	Quote{"NOTHING is IMPOSSIBLE", "Shia LaBeouf"},
	Quote{"You should get to the point where anyone else would quit and you're not going to stop there. NO!", "Shia LaBeouf"},
	Quote{"WHAT ARE YOU WAITING FOR? JUST DO IT!", "Shia LaBeouf"},
	Quote{"YES YOU CAN! JUST DO IT!", "Shia LaBeouf"},
	Quote{"If you’re tired of starting over, stop giving up.", "Shia LaBeouf"},
}
