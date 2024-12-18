"use client"

import instance from "@/config/axiosInstance"
import Cookies from "js-cookie"
import { useRouter, useSearchParams } from "next/navigation"
import { useEffect, useState } from "react"
import Swal from "sweetalert2"

export default function AddProduct() {
    const router = useRouter()
    const searchParams = useSearchParams()
    const token = Cookies.get("Authorization")
    const message = searchParams.get("message")
    const productId = searchParams.get("id")
    const [isUpdate, setIsUpdate] = useState(false);

    const [form, setForm] = useState({
        name: "",
        price: 0,
        quantity: 0,
        image_url: ""
    })

    useEffect(() => {
        if (productId) {
            setIsUpdate(true)
            const hydrateProduct = async () => {
                try {
                    const {data} = await instance({
                        url: `products/getOne/${productId}`,
                        method: "get"

                    })
                    setForm(prev => ({
                        ...prev,
                        name: data.data.Name,
                        price: data.data.Price,
                        quantity: data.data.Quantity,
                        image_url: data.data.ImageUrl,
                    }))
                    console.log(form);
                    
                    
                } catch (error) {

                    console.error("Failed to fetch product data", error);
                }
            }
            hydrateProduct()
        }
    }, [productId])




    const changeHandler = (e) => {
        e.preventDefault()
        const { name, value } = e.target
        setForm({
            ...form,
            [name]: value
        })
    }

    const submitHandler = async (e) => {
        e.preventDefault()
        try {
            let urlPage
            let instanceMethod
            if (isUpdate) {
                urlPage = `/products/update/${productId}`
                instanceMethod = "put"
            } else {
                urlPage = "/products/create"
                instanceMethod = "post"

            }
            
            const response = await instance({
                url: urlPage,
                method: instanceMethod,
                data: {
                    name: form.name,
                    price: parseInt(form.price),
                    quantity: parseInt(form.quantity),
                    imageUrl: form.image_url
                },
                headers: {
                    "Authorization": token
                }
            })
            const result = await Swal.fire({
                title: "Success!",
                text: `${form.name} created`,
                icon: "success",
                confirmButtonText: "Go back",
                customClass: {
                    popup: 'swal2-popup',
                    overlay: 'swal2-overlay'
                }
            })
            if (result.isConfirmed) {
                router.push("/yourProducts")
            }

        } catch (error) {
            const errorMessage = error.response?.data?.message
            console.log(error);
            console.log(errorMessage);
        }
    }
    return (
        <>
            <div className="bg-gray-900 w-full flex justify-center items-center" style={{ height: "calc(100vh - 60px)" }}>
                <div className="bg-gray-500 rounded-2xl w-1/4 h-auto">
                    <div className="px-2 pt-3 text-center w-full">CREATE PRODUCT</div>
                    <form onSubmit={submitHandler} className="flex flex-col p-2">
                        <label >Name</label>
                        <input type="text" name="name" value={form.name}  onChange={changeHandler}  className="px-2 py-1" required />
                        <label >Price</label>
                        <input type="number" name="price"value={form.price}  onChange={changeHandler}  className="px-2 py-1" required />
                        <label >Quantity</label>
                        <input type="number" name="quantity"value={form.quantity}  onChange={changeHandler}  className="px-2 py-1" required />
                        <label >Image Url</label>
                        <input type="text" name="image_url"value={form.image_url}  onChange={changeHandler}  className="px-2 py-1" required />
                        {message && <div className=" text-red-500 pb-3 font-bold">{message}</div>}
                        <div className="w-full flex justify-center pt-5 pb-2">
                            <button className="bg-gray-200 w-1/2">Create</button>
                        </div>
                    </form>
                </div>
            </div>
        </>
    )
}