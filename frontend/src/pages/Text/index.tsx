import React, { useState } from 'react'
import {Button,Input} from "antd"

const { TextArea } = Input;
export default function Text() {
  const [textValue,setTextValue]=useState<string>("")
  return (
    <div className='flex flex-col w-3/5 my-10'>
      <TextArea  rows={10} placeholder="请输入需要上传的文本内容" 
      className=' mb-10'
      onChange={(e)=>{
        console.log(e.target.value)
        setTextValue(e.target.value)
      }}
      ></TextArea>
      <Button type='primary'>
        确认传输
      </Button>
    </div>
  )
}
