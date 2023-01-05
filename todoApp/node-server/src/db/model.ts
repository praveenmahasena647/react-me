import {Schema,model} from 'mongoose'

const todoSchema=new Schema({
    act:{
        type:String,
        required:true,
    },
    done:{
        type:Boolean,
        default:false
    }
},{timestamps:true})




const Todo=model('reactTodo',todoSchema)

export {Todo}
