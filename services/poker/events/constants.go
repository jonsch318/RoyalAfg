package events

// IN:

const JOIN = "JOIN"
const PLAYER_ACTION = "PLAYER_ACTION"

// OUT:

const JOIN_SUCCESS = "JOIN_SUCCESS"
const PLAYER_JOIN = "PLAYER_JOIN"
const PLAYER_LEAVE = "PLAYER_LEAVE"

const REQUIRED_EVENT_NAME_MISSING = "The event received does not match the required event name"
const GAME_START = "GAME_START"
const DEALER_SET = "DEALER_SET"
const HOLE_CARDS = "HOLE_CARDS"
const WAIT_FOR_PLAYER_ACTION = "WAIT_FOR_PLAYER_ACTION"
const ACTION_PROCESSED = "ACTION_PROCESSED"

const FLOP = "FLOP"
const TURN = "TURN"
const RIVER = "RIVER"

const GAME_END = "GAME_END"
