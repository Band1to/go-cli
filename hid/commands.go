package hid

var BITLOX_REPORT = []byte{0x00}
var BITLOX_MAGIC = []byte{0x23, 0x23}
var BITLOX_TERMINATOR = []byte{0x7e, 0x7e}

var COMMAND_ACK = []byte{0x00, 0x51, 0x00, 0x00, 0x00, 0x00}

var COMMAND_LIST_WALLETS = []byte{0x00, 0x10, 0x00, 0x00, 0x00, 0x00}
var COMMAND_SCAN_WALLET = []byte{0x00, 0x61, 0x00, 0x00, 0x00, 0x00}

var PREFIX_LOAD_WALLET = []byte{0x00, 0x0B, 0x00, 0x00, 0x00, 0x02, 0x08}

var PREFIX_SIGN_MESSAGE = []byte{0x00, 0x70}

var bitloxCommands = map[string][]byte{
	"magic":      []byte{0x23, 0x23},
	"terminator": []byte{0x7e, 0x7e},
	"ack":        []byte{0x00, 0x51, 0x00, 0x00, 0x00, 0x00},

	"listWallets": []byte{0x00, 0x10, 0x00, 0x00, 0x00, 0x00},

	"scanWallet": []byte{0x00, 0x61, 0x00, 0x00, 0x00, 0x00},

	// prefixes for commands that take in a variable amount of
	// data, a content size and the payload data is appended
	"newWalletPrefix":     []byte{0x00, 0x04},
	"restoreWalletPrefix": []byte{0x00, 0x18},
	"renameWalletPrefix":  []byte{0x00, 0x0F},
	"signTxPrefix":        []byte{0x00, 0x65},
	"signMessagePrefix":   []byte{0x00, 0x70},
	"otpPrefix":           []byte{0x00, 0x57},
	"qrPrefix":            []byte{0x00, 0x80},

	// these just get one byte of hex for the wallet number
	// added to them
	"deleteWalletPrefix": []byte{0x00, 0x16, 0x00, 0x00, 0x00, 0x02, 0x08},
	"loadWalletPrefix":   []byte{0x00, 0x0B, 0x00, 0x00, 0x00, 0x02, 0x08},

	// just a ping
	"ping": []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x0A, 0x05, 0x48, 0x65, 0x6C, 0x6C, 0x6F},
}

var RESPONSE_SUCCESS byte = 0x34
var RESPONSE_ERROR byte = 0x35
var RESPONSE_PLEASE_ACK byte = 0x50
var RESPONSE_MESSAGE_SIGNATURE byte = 0x71