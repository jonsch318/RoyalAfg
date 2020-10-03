import { withTranslation } from "../../i18n"
import "./idnav"
import IdNav from "./idnav"
import PropTypes from "prop-types"

const NavItems = ({ t }) => {
    return (
        <div className="md:flex md:h-full w-full">
            <nav className="block md:flex md:flex-auto md:items-center">
    <a className="nav-item block py-4 px-4 md:p-0 border-gray-300 border-b-2 border-solid md:border-none" href="/">{t("home")}</a>
                <a className="nav-item block py-4 px-4 md:p-0 md:border-none border-gray-300 border-b-2 border-solid" href="/about">{t("about")}</a>
            </nav>
            <div className="idnav md:mr-12 md:flex block my-2">
                <IdNav></IdNav>
            </div>
        </div>
    )
}

NavItems.getInitialProps = async () => ({
    namespacesRequired: ["common", "nav"],
})

NavItems.propTypes = {
    t: PropTypes.func.isRequired
}

export default withTranslation("nav")(NavItems)