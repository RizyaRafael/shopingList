'use client'

import { AuthContext } from "@/context/authContext"
import Cookies from "js-cookie"
import { useContext, useEffect, useState } from "react"

export default function YourProductsComp() {
    const [token, setToken] = useState("")
    const { isLoggedIn } = useContext(AuthContext)

    useEffect(() => {
        setToken(Cookies.get('Authorization'))
    },[isLoggedIn])
    return <>
        {token ? (
            <a className="btn" href="/yourProducts">Your Products</a>
        ) : null}
    </>
}