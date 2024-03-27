import Elysia, { t, type Context } from "elysia";
import { customersSchema } from "./customers.schema";
import { sql } from "../../db/db";

export const customersController = new Elysia()
  .use(customersSchema)
  .get("/customers", async () => {
    const [results] = await sql.execute(
      "SELECT id AS id, COALESCE(NULLIF(adRazaoSocial, ''), adNomeFantasia) AS customer_name FROM pclientes ORDER BY customer_name DESC"
    );

    return {
      success: true,
      payload: results,
    };
  })
  .get(
    "/customers/most-inactive",
    async ({ query }) => {
      const numberOfDays = +query.days;
      const providerId = +query.provider;

      let results: any;

      if (providerId === -1) {
        const [r] = await sql.execute(
          "SELECT c.id AS customer_id, c.adRazaoSocial AS customer_name, (SELECT MAX(pp.adData) FROM ppedidos pp WHERE pp.adIdCliente = c.id) AS last_order_date, (SELECT id FROM ppedidos pp WHERE pp.adIdCliente = c.id ORDER BY pp.timestamp DESC LIMIT 1) AS order_id FROM pclientes c WHERE c.id NOT IN (SELECT c1.id FROM pclientes c1, ppedidos p WHERE p.adIdCliente = c1.id AND p.adData >= (date(DATE_SUB(SYSDATE(), INTERVAL ? DAY)))) AND (SELECT MAX(pp.adData) FROM ppedidos pp WHERE pp.adIdCliente = c.id) IS NOT NULL ORDER BY last_order_date DESC",
          [numberOfDays]
        );
        results = r;
      } else {
        const [r] = await sql.execute(
          "SELECT c.id AS customer_id, c.adRazaoSocial AS customer_name, (SELECT MAX(pp.adData) FROM ppedidos pp WHERE pp.adIdCliente = c.id) AS last_order_date, (SELECT id FROM ppedidos pp WHERE pp.adIdCliente = c.id ORDER BY pp.timestamp DESC LIMIT 1) AS order_id FROM pclientes c WHERE c.id NOT IN (SELECT c1.id FROM pclientes c1, ppedidos p WHERE p.adIdCliente = c1.id AND p.adData >= (date(DATE_SUB(SYSDATE(), INTERVAL ? DAY)))) AND (SELECT MAX(pp.adData) FROM ppedidos pp WHERE pp.adIdCliente = c.id) IS NOT NULL AND (SELECT COUNT(adIdFornecedor) FROM ppedidos pp WHERE pp.adIdCliente = c.id AND pp.adIdFornecedor != ?) = 0 ORDER BY last_order_date DESC",
          [numberOfDays, providerId]
        );
        results = r;
      }

      return {
        success: true,
        payload: results,
      };
    },
    {
      query: t.Object({
        days: t.String({
          default: 0,
        }),
        provider: t.String({
          default: -1,
        }),
      }),
    }
  );
