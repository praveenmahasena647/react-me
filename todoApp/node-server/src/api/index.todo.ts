import { Router,Request,Response } from  'express'
import {GetAll } from '../interfaces/todo.interface'

const todo=Router()

todo.get('/',(req :Request,res :Response<GetAll[]>)=>{
    res.json([{act:'',done:false}])
})

todo.get('/:id',(req :Request,res :Response<GetAll>)=>{
    return res.json({act:'',done:false})
})

todo.post('/',(req :Request,res :Response)=>{})

todo.put('/:id',(req :Request,res :Response)=>{})

todo.delete('/:id',(req :Request,res :Response)=>{})

todo.delete('/',(req :Request,res :Response)=>{})


export default todo
