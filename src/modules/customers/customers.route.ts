import Elysia from "elysia";
import { customersController } from "./customers.controller";

export const customersRoutes = new Elysia().use(customersController);
