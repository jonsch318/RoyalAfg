import { withTranslation } from "../../i18n"
import "./idnav"
import IdNav from "./idnav"
import PropTypes from "prop-types"
import HeaderNavItem from "../../components/header/navItem"

const NavItems = ({ t }) => {
    return (
        <div className="md:flex md:h-full w-full">
            <nav className="block md:flex md:flex-auto md:items-center">
                <HeaderNavItem href="/">{t("home")}</HeaderNavItem>
                <HeaderNavItem href="/about">{t("about")}</HeaderNavItem>
                <HeaderNavItem href="/games">{t("games")}</HeaderNavItem>
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