import "./idnav"
import IdNav from "./idnav"

export default function NavItems() {
    return (
        <div className="md:flex md:h-full w-full">
            <nav className="block md:flex md:flex-auto md:items-center">
                <a className="nav-item block py-4 px-4 border-gray-300 border-b-2 border-solid" href="/">Home</a>
                <a className="nav-item block py-4 px-4 border-gray-300 border-b-2 border-solid" href="/about">About</a>
            </nav>
            <div className="idnav md:mr-12 md:flex block my-2">
                <IdNav></IdNav>
            </div>
        </div>
    )
}