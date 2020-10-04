import { LOGGED_USER_IN } from "../types/authTypes"

export const loginUser = (credentials) => dispatch => {
    console.log("Fetchign");
    fetch("http://troyalafg.games/account/login", {
        body: JSON.stringify(credentials),
        method: "POST",
    })
    .then(res => res.json())
    .then(user => dispatch({
        type: LOGGED_USER_IN,
        payload: user,
    }))
}