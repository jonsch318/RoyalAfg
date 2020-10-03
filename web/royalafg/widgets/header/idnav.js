import { useSelector } from "react-redux"

export default function IdNav() {

    const isLoggedIn = useSelector(state => state.auth.isLoggedIn);


    if(!isLoggedIn){
        return (
            <nav className="flex items-center h-full w-full">
                <div className="flex items-center h-full w-full px-4">
                    <a className="id-nav-item md:bg-transparent px-4 py-1 rounded bg-gray-300 md:hover:bg-blue-700 md:mx-2 transition-colors duration-150 flex" href="/register">Register</a>
                    <a className="id-nav-item bg-blue-800 px-6 py-1 rounded hover:bg-blue-900 md:mx-2 text-white transition-colors duration-150 flex mr-0 ml-auto" href="/login">Login</a>
                </div>
                <div className="flex px-8 hidden">
                    <a href="/account">My Account</a>
                </div>
            </nav>
        )
    }

    return (
        <nav className="flex items-center h-full w-full">
            <h1>Signed </h1>
        </nav>
    )


}