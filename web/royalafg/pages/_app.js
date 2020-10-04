import { Provider } from "react-redux"
import { useStore } from "../redux/store"
import "../styles/globals.css"
import "../styles/tailwind.css"
import Header from "../widgets/header"
import { appWithTranslation } from "../i18n"

function MyApp({ Component, pageProps }) {

  const store = useStore(pageProps.initialReduxState)

  return (
    <Provider store={store}>
      <div className="main-container">
        <Header></Header>
        <Component {...pageProps} />
      </div>
    </Provider>
  )
}

export default appWithTranslation(MyApp)
