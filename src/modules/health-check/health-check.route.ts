import Elysia from "elysia";
import { healthCheckController } from "./health-check.controller";

export const healthCheckRoutes = new Elysia().use(healthCheckController);
