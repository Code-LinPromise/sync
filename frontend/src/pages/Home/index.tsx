import React, { useState } from 'react'
import { Outlet } from 'react-router-dom'
import { useNavigate ,useLocation} from 'react-router-dom'


type TagType="text" | "photo" | "file"
type TagListType={
  name:string,
  tag:TagType
}
const TagList:TagListType[]=[
  {
    name:"文本传输",
    tag:"text"
  },
  {
    name:"图片传输",
    tag:"photo"
  },
  {
    name:"文件传输",
    tag:"file"
  },
];
export default function Home() {
  const navigate=useNavigate()
  const location = useLocation();
  console.log(location)
  const initTag =location.pathname.slice(1) || "text"
  const [selectTag,setSelectTag]=useState(initTag)
  return (
    <div className="container mx-auto m-20 flex-auto flex items-center flex-col">
      <ul className='flex  flex-auto  justify-evenly p-2 w-3/5  rounded-2xl bg-slate-50 '>
        {
          TagList.map((item)=>{
            return <li key={item.tag} 
            className={`p-2  rounded-2xl  cursor-pointer transition-colors whitespace-nowrap  list-none ${selectTag===item.tag? "text-sky-600" : ""}`}
            onClick={()=>{
              setSelectTag(item.tag)
              navigate(`/${item.tag}`)
            }}
            >{item.name}</li>
          })
        }
      </ul>
      <Outlet/>
    </div>
  )
}
