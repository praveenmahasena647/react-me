import * as z from 'zod'

const getAll=z.object({
    act:z.string(),
    done:z.boolean()
})

type GetAll=z.infer<typeof getAll>

export {
    GetAll,
}
