"use client"
import instance from "@/config/axiosInstance"
import Swal from "sweetalert2"
import ButtonComp from "../button";

export default function Card({ product, userId, token }) {
    let userProduct = false
    if (userId == product.UserId) {
        userProduct = true
    }
    async function buy() {
        try {
            const response = await instance({
                url: "/products/buyProduct",
                method: "post",
                data: product,
                headers: {
                    Authorization: token
                }
            })
            const result = await Swal.fire({
                title: "Purchase Successful",
                text: `You bought ${product.Name}`,
                icon: "success",
                confirmButtonText: "Go to home",
                customClass: {
                    popup: 'swal2-popup',
                    overlay: 'swal2-overlay'
                }
            })
            if (result.isConfirmed) {
                window.location.reload()
            }
        } catch (error) {
            Swal.fire({
                title: "Error",
                text: "Something went wrong with your purchase.",
                icon: "error",
                confirmButtonText: "Try again"
            });
        }
    }

    async function deleteProduct() {
        try {
            Swal.fire({
                title: "Are you sure?",
                text: "You won't be able to revert this!",
                icon: "warning",
                showCancelButton: true,
                confirmButtonColor: "#3085d6",
                cancelButtonColor: "#d33",
                confirmButtonText: "Yes, delete it!"
            }).then(async (result) => {
                if (result.isConfirmed) {
                    const response = await instance({
                        url: `/products/delete/${product.ID}`,
                        method: "delete",
                        headers: {
                            Authorization: token
                        },
                        data: product
                    })
                    console.log(response);
                    
                    Swal.fire({
                        title: "Deleted!",
                        text: "Your file has been deleted.",
                        icon: "success"
                    }).then(() => {
                        window.location.reload()
                    })
                }
            });
        } catch (error) {
            console.log(error);
            
            Swal.fire({
                title: "Error",
                text: "Something went wrong with your purchase.",
                icon: "error",
                confirmButtonText: "Try again"
            });
        }
    }

    return (
        <>
            <div className="rounded-xl bg-gray-700 px-3 py-1 flex flex-col justify-around">
                <img
                    src={product.ImageUrl || "https://st3.depositphotos.com/17828278/33150/v/450/depositphotos_331503262-stock-illustration-no-image-vector-symbol-missing.jpg"}
                    alt="Product image"
                    className="mt-2 w-full h-1/2"
                />
                <div className="w-full h-1/4">

                    <div className="text-xl truncate hover:text-clip">{product.Name}</div>
                    <div>Rp. {product.Price}</div>
                    <div>Stock: {product.Quantity}</div>
                </div>
                <div className="mt-3 mb-2 ">
                    {token ? (
                        userId == product.UserId ? (
                            <>
                                <a href={`/addProduct?id=${product.ID}`} className="btn bg-green-500 w-1/2 rounded-xl py-1 text-xl font-medium mt-5 mb-1">Update</a>
                                <button onClick={deleteProduct} className="btn bg-red-500 w-1/2 rounded-xl py-1 text-xl font-medium mt-5 mb-1">Delete</button>
                            </>
                        ) : (
                            <ButtonComp className text="Buy" purpose={buy} color="bg-gray-400" />
                        )
                    ) : (
                        <a href="/login" className="btn w-1/2 rounded-xl  text-lg font-medium mt-6 mx-1 px-3 py-1 bg-gray-400">Login to buy</a>
                    )}
                </div>
            </div>
        </>
    )
}