import express from 'express';
import morgan from 'morgan';
import helmet from 'helmet';
import cors from 'cors';

import * as middlewares from './middlewares';
import todo from '../src/api/index.todo'

require('dotenv').config();

const app = express();

app.use(morgan('dev'));
app.use(helmet());
app.use(cors());
app.use(express.json());

app.use('/todo',todo)

app.use(middlewares.notFound);
app.use(middlewares.errorHandler);

export default app;