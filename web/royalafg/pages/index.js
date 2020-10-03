import {withTranslation } from "../i18n"
import PropTypes, { string } from "prop-types"
import { initializeStore } from "../redux/store"
import { useDispatch, useSelector } from "react-redux"
import { LOGIN_USER } from "../redux/types/authTypes"

export default function Index(store) {
  
  const state = useSelector(state => state.auth.isLoggedIn)
  
  const dispatch = useDispatch()

  return (
    <div>
      <h1>Hello</h1>
      <a href="/about">About</a>
      <h1>is logged in [{state ? "signed in" : "not signed in"}]</h1>
      <button onClick={dispatch(LOGIN_USER)}>Sign in</button>
    </div>
  )
}