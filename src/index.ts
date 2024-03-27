import Elysia from "elysia";
import { cors } from "@elysiajs/cors";
import { healthCheckRoutes } from "./modules/health-check/health-check.route";
import swagger from "@elysiajs/swagger";
import { searchRoutes } from "./modules/search/search.route";
import { customersRoutes } from "./modules/customers/customers.route";

export const app = new Elysia();

app

  //localhost:3000/swagger

  .use(swagger())

  // Enabling CORS
  .use(cors())

  // Routes to register
  .use(healthCheckRoutes)
  .use(searchRoutes)
  .use(customersRoutes);

app.listen(process.env.PORT || 3000);

console.log(`API running at: ${app.server?.hostname}:${app.server?.port}`);
