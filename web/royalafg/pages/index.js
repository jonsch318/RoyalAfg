import {withTranslation } from "../i18n"
import PropTypes from "prop-types"

const Index = ({ t }) =>  {
  return (
    <div>
      <h1>Hello</h1>
      <a href="/about">About</a>
    </div>
  )
}

Index.getInitialProps = async () => ({
  namespacesRequired: ["common"],
})

Index.propTypes = {
  t: PropTypes.func.isRequired
}

export default withTranslation(Index)