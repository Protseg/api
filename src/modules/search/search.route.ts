import Elysia from "elysia";
import { searchController } from "./search.controller";

export const searchRoutes = new Elysia().use(searchController);
