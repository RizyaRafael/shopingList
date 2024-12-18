import LoginLogout from "./loginLogout";
import YourProductsComp from "./yourProducts";
export default function Navbar() {
    return (
        <>
            <div className="flex flex-rows justify-between items-center bg-violet-400 h-full w-full px-4 py-3">
                <div className="flex space-x-6">
                    <a className="btn" href="/">Home</a>
                    <YourProductsComp/>
                </div>
                <div>
                    <form action="/" method="get" className="flex justify-between space-x-4">
                        <input
                            type="text"
                            className="bg-white rounded-lg px-3 py-2"
                            name="name"
                            placeholder="Search items..."
                        />
                        <button type="submit" className="btn">Search</button>
                    </form>
                </div>
                <div className="flex space-x-4 justify-end">
                    <LoginLogout />
                </div>
            </div>
        </>
    )
}