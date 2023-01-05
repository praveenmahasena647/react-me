import { Router, Request, Response } from "express";
import { Todo } from "../db/model";

const todo = Router();

todo.get("/", async (req: Request, res: Response) => {
    try {
        let todos = await Todo.find();
        res.status(200).json(todos)
    } catch (error) {
        res.status(400)    
        console.log(error)
    }
});

todo.get("/:id", async (req: Request, res: Response) => {
    let { id } = req.params;

    try {
        const todo = await Todo.findById(id);
        res.status(200).json(todo);
    } catch (error) {
        res.status(400);
        console.log(error)
    }
});

todo.post("/", async (req: Request, res: Response) => {
    let { body } = req;
    console.log(body)

    try {
        let todo = new Todo(body);
        await todo.save();
        res.status(200).end()
    } catch (error) {
        res.status(400).end()
        console.log(error);
    }
});

todo.put("/:id", async (req: Request, res: Response) => {
    let { id } = req.params;
    try {
        await Todo.findByIdAndUpdate(id, req.body);
        res.status(200).end()
    } catch (error) {
        res.status(400).end()
        console.log(error);
    }
});

todo.delete("/:id", async(req: Request, res: Response) => {
    let {id}=req.params

    try {
        await Todo.findByIdAndDelete(id) 
        res.status(200).end()
    } catch (error) {
        res.status(400).end()
        console.log(error) 
    }
});

export default todo;
