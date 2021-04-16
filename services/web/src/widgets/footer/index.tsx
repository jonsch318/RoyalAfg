import { useTranslation } from "next-i18next";
import React, { FC } from "react";
import FooterCard from "./card";
import FooterCardItem from "./cardItem";

type FooterProps = {
    absolute: boolean;
};

const Footer: FC<FooterProps> = ({ absolute }) => {
    const { t } = useTranslation("common");
    let containerClass = "w-full";

    if (absolute) {
        containerClass += " absolute bottom-0 top-auto";
    }

    return (
        <div className={containerClass}>
            <style jsx>{`
                .footer-grid {
                    grid-template-columns: auto 1fr;
                }
            `}</style>
            <footer className="bg-blue-600 text-white font-sans md:px-16 md:py-8 py-4 px-8">
                <div className="md:grid footer-grid">
                    <div className="md:grid md:grid-rows-2 w-auto md:mr-16 mb-2">
                        <div>&copy; Jonas Schneider</div>
                        <a href="/" className="font-medium font-sans text-xl cursor-pointer">
                            Royalafg
                        </a>
                    </div>
                    <div className="md:grid footer-grid-content row-span-2 md:gap-2 md:grid-cols-3 md:justify-items-center">
                        <FooterCard title={t("Contact")}>
                            <FooterCardItem href="/about">{t("About this project")}</FooterCardItem>
                        </FooterCard>
                        <FooterCard title={t("Games")}>
                            <FooterCardItem href="/games">{t("Game Selection")}</FooterCardItem>
                            <FooterCardItem href="/games/poker">{t("Poker")}</FooterCardItem>
                            <FooterCardItem href="/games/pacman">{t("Pacman")}</FooterCardItem>
                        </FooterCard>
                        <FooterCard title={t("Legal")}>
                            <FooterCardItem href="/legal/terms">{t("Terms & Conditions")}</FooterCardItem>
                            <FooterCardItem href="/legal/privacy">{t("Privacy Statement")}</FooterCardItem>
                        </FooterCard>
                    </div>
                </div>
            </footer>
        </div>
    );
};

export default Footer;
