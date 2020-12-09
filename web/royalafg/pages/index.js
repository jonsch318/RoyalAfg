import { withTranslation } from "../i18n"
import PropTypes, { string } from "prop-types"
import { initializeStore } from "../redux/store"
import { useDispatch, useSelector } from "react-redux"
import { LOGIN_USER } from "../redux/types/authTypes"
import Footer from "../widgets/footer"
import Layout from "../components/layout"

export default function Index(store) {

  const state = useSelector(state => state.auth.isLoggedIn)

  const dispatch = useDispatch()

  return (
    <Layout footerAbsolute>
      <div>
        <h1>Hello</h1>
        <a href="/about">About</a>
        <h1>is logged in [{state ? "signed in" : "not signed in"}]</h1>
        <button>Sign in</button>
      </div >
    </Layout >
  )
}