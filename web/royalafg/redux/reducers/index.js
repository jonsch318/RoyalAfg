import authReducer from "./authReducer";

const { combineReducers } = require("redux");

export default combineReducers({
    auth: authReducer,
})