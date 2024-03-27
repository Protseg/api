import Elysia, { type Context } from "elysia";
import { healthCheckSchema } from "./health-check.schema";

export const healthCheckController = new Elysia()
  .use(healthCheckSchema)
  .get("/", () => {
    return {
      status: "ok",
      server_date: new Date().toISOString().substring(0, 19),
    };
  })
  .get("/ping", () => {
    return {
      o: 1,
    };
  });
