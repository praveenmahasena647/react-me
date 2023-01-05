import process from 'process'
import {connect} from 'mongoose'
import app from './app';

const port = process.env.PORT || 42069



connect("mongodb://127.0.0.1:27017/reactTodo").catch(err=>{
    console.log(err)
    process.exit()
})

app.listen(port,()=>console.log(`http://localhost:${port}`));
