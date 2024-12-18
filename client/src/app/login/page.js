'use client'

import instance from "@/config/axiosInstance"
import { AuthContext } from "@/context/authContext"
import Cookies from "js-cookie"
import Link from "next/link"
import Swal from "sweetalert2"

const { useSearchParams, useRouter } = require("next/navigation")
const { useState, useContext, useEffect } = require("react")

export default function login() {
    const router = useRouter()
    const { login, isLoggedIn } = useContext(AuthContext)
    const token = Cookies.get('Authorization')
    const searchParams = useSearchParams()
    const message = searchParams.get("message")
    const [form, setForm] = useState({
        password: "",
        email: ""
    })
    
    useEffect(() => {
        if (token) {
            router.push("/")
        }
    }, [isLoggedIn, router])
    if (isLoggedIn === null) {
        return <div>Redirecting...</div>
    }

    const changeHandler = async (e) => {
        e.preventDefault()
        const { name, value } = e.target
        setForm({
            ...form,
            [name]: value
        })
    }

    const loginHandler = async (e) => {
        e.preventDefault()
        try {
            const response = await instance({
                url: "/user/login",
                method: "post",
                data: {
                    email: form.email,
                    password: form.password,
                }
            })
            const { data, userId } = response.data
            Cookies.set("Authorization", data)
            Cookies.set("userId", userId)
            const result = await Swal.fire({
                title: "Welcome back!",
                text: "You successfully login",
                icon: "success",
                confirmButtonText: "Go to home",
                customClass: {
                    popup: 'swal2-popup',
                    overlay: 'swal2-overlay'
                }
            })
            login()
            if (result.isConfirmed) {
                router.push("/")
            }
        } catch (error) {
            const errorMessage = error.response?.data?.message
            router.push(`/login?message=${errorMessage}`)
        }

    }

    return (
        <>
            <div className="grid grid-cols-2 w-screen" style={{ height: "calc(100vh - 60px)" }}>
                <div className="bg-gray-900 flex flex-col justify-center items-center text-white">
                    <p className="text-8xl font-bold">WELCOME</p>
                    <p className="text-xl pt-5">do you not have a account?
                        <Link href="/register" className="font-bold"> Register here!</Link>
                    </p>
                </div>
                <div className="bg-gray-900 flex justify-center items-center">
                    <div className="bg-gray-700 p-5 rounded-xl w-1/2 h-auto text-white">
                        <p className="text-center text-4xl font-bold">Login</p>
                        <form onSubmit={loginHandler} className="flex flex-col text-xl" onChange={changeHandler}>
                            <label>email</label>
                            <input
                                name="email"
                                type="email"
                                className="form-control mb-3  text-black"
                                required

                            />
                            <label>password</label>
                            <input
                                name="password"
                                type="password"
                                className="form-control mb-3 text-black"
                                required


                            />
                            {message && <div className=" text-red-600 pb-3">{message}</div>}
                            <div className="flex justify-center items-center">
                                <button type="submit" className="btn bg-gray-500 w-1/2 rounded-xl py-2 text-xl font-medium">Login</button>
                            </div>
                        </form>
                    </div>
                </div>

            </div>
        </>
    )
}