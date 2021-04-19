import "./idnav";
import React, { FC } from "react";

import HeaderNavItem from "../../components/header/navItem";
import SearchInput from "../../components/search/search";

import IdNav from "./idnav";
import { useTranslation } from "next-i18next";

const NavItems: FC = () => {
    const { t } = useTranslation("common");
    return (
        <div className="md:grid  w-full" style={{ gridTemplateColumns: "auto 1fr auto" }}>
            <nav className="block md:flex md:flex-auto md:items-center w-auto">
                <HeaderNavItem href="/">{t("Home")}</HeaderNavItem>
                <HeaderNavItem href="/about">{t("About")}</HeaderNavItem>
                <HeaderNavItem href="/games">{t("Games")}</HeaderNavItem>
            </nav>
            <div className="search-container w-full flex-auto py-2">
                <SearchInput />
            </div>
            <div className="idnav md:mr-12 md:flex block my-2 w-auto">
                <IdNav />
            </div>
        </div>
    );
};

export default NavItems;
