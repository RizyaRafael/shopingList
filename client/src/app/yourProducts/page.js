import Card from "@/components/home/card"
import instance from "@/config/axiosInstance"
import { cookies } from "next/headers"
import { redirect } from "next/navigation"

const getData = async (token) => {
    'use server'
    try {
        const result = await instance({
            url: "/products/getUserProducts",
            method: "get",
            headers: {
                Authorization: token
            }
        })
        return result.data
    } catch (error) {
        return { data: null }
    }
}


export default async function YourProducts() {
    const cookie = await cookies()
    const userId = cookie.get("userId")?.value || null
    const token = cookie.get("Authorization")?.value
    if (!token) {
        redirect("/")
    }
    const datas = await getData(token)

    const hasProducts = Array.isArray(datas?.data) && datas.data.length > 0

    return (
        <>
            <div className="bg-gray-900 w-full pt-5 px-5 " >
            <a href="/addProduct" className="bg-blue-400 rounded-xl px-2 py-2">+ Add product</a>
            </div>
            <div className="flex justify-center items-center h-full w-full bg-gray-900">
                <div className="w-full p-5 text-white">
                    {hasProducts ? (
                        <div className="grid grid-cols-5 gap-4">
                            {datas.data.map((product) => (
                                <Card key={product.Id} product={product} userId={userId} token={token} />
                            ))}
                        </div>
                    ) : (
                        <div className="text-center w-full bg-red-300 p-4">
                            {datas?.data || "You don't have any products"}
                        </div>
                    )}
                </div>
            </div>
        </>
    )
}