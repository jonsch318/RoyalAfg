import { LOGGED_USER_IN } from "../types/authTypes";

const intialState = {
    user: {},
    isLoggedIn: false
}

export default function(state = intialState, action) {
    switch (action.type) {
        case LOGGED_USER_IN:
            return {
                ...state,
                user: action.payload,
                isLoggedIn: true,
            }
        default:
            return state;
    }
}