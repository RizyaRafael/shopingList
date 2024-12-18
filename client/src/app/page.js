import Card from "@/components/home/card"
import Pagination from "@/components/home/pagination"
import instance from "@/config/axiosInstance"
import { cookies } from "next/headers"
const getData = async (name, page, limit) => {
  'use server'
  const result = await instance({
    url: "/products",
    params: {
      page,
      name,
      limit
    }
  })
  return result.data
}
export default async function Home({ searchParams }) {
  const cookie = await cookies()
  const name = await searchParams.name || ""
  const page = await searchParams.page || ""
  const limit = await searchParams.limit || ""
  const datas = await getData(name, page, limit)
  const userId = cookie.get("userId")?.value || null
  const token = cookie.get("Authorization")?.value

  return (
    <>
    <div className="flex justify-center items-center h-full w-full bg-gray-900" >
      <div className="grid grid-cols-5 gap-4 p-5 text-white">
        {datas.data ? (
          datas.data.map((product) => {
            return <Card key={product.id} product={product} userId={userId} token={token}/>
          })
        ) : (
          <div>Cant find {name}</div>
        )}
      </div>
    </div> 
    <div className="bg-gray-900 pb-4">
      <Pagination page={page} total={datas.total} limit={limit}/>
    </div>
        </>
  )
}
