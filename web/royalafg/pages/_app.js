import "../styles/globals.css"
import "../styles/tailwind.css"
import Header from "../widgets/header"

function MyApp({ Component, pageProps }) {
  return (
    <div className="main-container">
      <Header></Header>
      <Component {...pageProps} />
    </div>
  )
}

export default MyApp
