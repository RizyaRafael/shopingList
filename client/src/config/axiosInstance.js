import axios from "axios"

const url = "http://127.0.0.1:3000"
console.log(url);

const instance = axios.create({
    baseURL: url
})

export default instance