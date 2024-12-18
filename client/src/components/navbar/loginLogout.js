'use client'
import { AuthContext } from "@/context/authContext";
import Cookies from "js-cookie";
import { useContext, useEffect, useState } from "react";
export default function LoginLogout() {
    const { isLoggedIn, logout } = useContext(AuthContext)
    const [token, setToken] = useState("")
    useEffect(() => {
        setToken(Cookies.get('Authorization'))
    },[isLoggedIn])
    if (isLoggedIn === null) {
        return <></>
    }
    return (
        <>
            {token ? (
                <div onClick={logout}>Logout</div>
            ) : (
                <>
                    <a className="btn" href="/login">Login</a>
                    <a className="btn" href="/register">Register</a>
                </>
            )}
        </>
    )
}