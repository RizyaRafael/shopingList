'use client'
import Cookies from "js-cookie"
import { createContext, useEffect, useState } from "react"

export const AuthContext = createContext()

export default function AuthProvider({children}){
    const [isLoggedIn, setIsLoggedIn] = useState(null)
    useEffect(() => {
        const loggedIn = Cookies.get('isLoggedIn') === 'true'
        setIsLoggedIn(loggedIn)
    }, [])

    const login = () => {
        setIsLoggedIn(true)
    }

    const logout = () => {
        setIsLoggedIn(false)
        Cookies.remove('userId')
        Cookies.remove('Authorization')
        window.location.reload('/')
    }

    return (
        <AuthContext.Provider value={{isLoggedIn, login, logout}}>
            {children}
        </AuthContext.Provider>
    )
}