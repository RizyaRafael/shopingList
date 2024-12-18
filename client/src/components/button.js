"use client"
export default function ButtonComp({text, purpose, color}) {
    const buttonStyle = `btn ${color} w-1/2 rounded-xl py-1 text-xl font-medium mt-5 mb-1`
    return <>
    <button type="button" onClick={purpose} className={buttonStyle}>{text}</button>
    </>
}