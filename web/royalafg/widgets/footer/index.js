const { default: FooterCard } = require("./card")
const { default: FooterCardItem } = require("./cardItem")

const Footer = (props) => {
    return (
        <div className="absolute bottom-0 top-auto w-full">
            <style jsx>{`
                .footer-grid{
                    grid-template-columns: auto 1fr;
                }
                .footer-grid-content{
                    grid-template-columns: auto auto;
                    grid-tempalte-rows: auto auto;
                }
            `}</style>
            <footer className="bg-blue-600 text-white font-sans px-16 py-8">
                <div className="grid footer-grid">
                    <div className="grid grid-rows-2 w-auto mr-16">
                        <div>&copy; Jonas Schneider</div>
                        <a href="/" className="font-medium font-sans text-xl cursor-pointer">Royalafg</a>
                    </div>
                    <div className="grid footer-grid-content row-span-2 gap-2">
                        <FooterCard title="legal">
                            <FooterCardItem href="/legal/terms">Terms & Conditions</FooterCardItem>
                            <FooterCardItem href="/legal/privacy">Privacy Statement</FooterCardItem>
                        </FooterCard>
                        <FooterCard title="Games">
                            <FooterCardItem href="/games">Game Selection</FooterCardItem>
                            <FooterCardItem href="/games/poker">Poker</FooterCardItem>
                            <FooterCardItem href="/games/pacman">Pacman</FooterCardItem>
                        </FooterCard>
                        <FooterCard title="legal">
                            <FooterCardItem href="/legal/terms">Terms & Conditions</FooterCardItem>
                            <FooterCardItem href="/legal/privacy">Privacy Statement</FooterCardItem>
                            <FooterCardItem >Hello</FooterCardItem>
                        </FooterCard>
                    </div>
                </div>
            </footer>
        </div>
    );
}

export default Footer