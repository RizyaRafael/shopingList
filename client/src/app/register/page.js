'use client'
import instance from "@/config/axiosInstance";
import { AuthContext } from "@/context/authContext";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";
import { useContext, useEffect, useState } from "react";
import Swal from "sweetalert2";

export default function Register() {
    const { isLoggedIn } = useContext(AuthContext)
    const router = useRouter()
    const searchParams = useSearchParams()
    const message = searchParams.get("message")
    const [form, setForm] = useState({
        username: "",
        password: "",
        email: ""
    })

    useEffect(() => {
        if (isLoggedIn) {
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

    const registerHandler = async (e) => {
        e.preventDefault()
        try {
            const response = await instance({
                url: "/user/register",
                method: "post",
                data: {
                    email: form.email,
                    username: form.username,
                    password: form.password,
                }
            })
            const result = await Swal.fire({
                title: "Good job!",
                text: "You successfully registered",
                icon: "success",
                confirmButtonText: "OK",
                customClass: {
                    popup: 'swal2-popup',
                    overlay: 'swal2-overlay'
                }
            })

            if (result.isConfirmed) {
                router.push("/login")
            }
        } catch (error) {
            const errorMessage = error.response?.data?.message
            router.push(`/register?message=${errorMessage}`)
        }

    }


    return (
        <>
            <div className="grid grid-cols-2 w-screen" style={{ height: "calc(100vh - 60px)" }}>
                <div className="bg-gray-900 flex justify-center items-center">
                    <div className="bg-gray-700 p-5 rounded-xl w-1/2 h-auto text-white">
                        <p className="text-center text-4xl font-bold">Register</p>
                        <form onSubmit={registerHandler} className="flex flex-col text-xl" onChange={changeHandler}>
                            <label>email</label>
                            <input
                                name="email"
                                type="email"
                                className="form-control mb-3 "
                                required
                            />
                            <label>username</label>
                            <input
                                name="username"
                                type="text"
                                className="form-control mb-3"
                                required
                            />
                            <label>password</label>
                            <input
                                name="password"
                                type="password"
                                className="form-control mb-3"
                                required
                            />
                            {message && <div className=" text-red-600 pb-3">{message}</div>}
                            <div className="flex justify-center items-center">
                                <button type="submit" className="btn bg-gray-500 w-1/2 rounded-xl py-2 text-xl font-medium">Register</button>
                            </div>
                        </form>
                    </div>
                </div>
                <div className="bg-gray-900 flex flex-col justify-center items-center text-white">
                    <p className="text-8xl font-bold">HELLO!</p>
                    <p className="text-xl pt-5">do you have a account?
                        <Link href="/login" className="font-bold"> Login here!</Link>
                    </p>
                </div>
            </div>
        </>
    )
}