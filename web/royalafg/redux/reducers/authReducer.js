import { LOGIN_USER_SUCCESS } from '../types/authTypes';

const intialState = {
    user: {},
    isLoggedIn: false
};

const authReducer = (state = intialState, action) => {
    switch (action.type) {
        case LOGIN_USER_SUCCESS:
            return {
                ...state,
                user: action.payload,
                isLoggedIn: true
            };
        default:
            return state;
    }
};

export default authReducer;
