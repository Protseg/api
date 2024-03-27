import Elysia, { t, type Context } from "elysia";
import { searchSchema } from "./search.schema";
import { sql } from "../../db/db";

export const searchController = new Elysia()
  .use(searchSchema)
  .get(
    "/search/by-customer",
    async ({ query }) => {
      const id = query.id;

      const [results] = await sql.execute(
        "SELECT psimples.adReferencia AS product_ref, REPLACE(COALESCE(psimples.adDiscriminacao, pcomposto.adModelo), '\t', '') AS product_name FROM ppedidos pedidos LEFT JOIN ppedidosimples psimples ON pedidos.id = psimples.adIdPedido LEFT JOIN ppedidocomposto pcomposto ON pedidos.id = pcomposto.adIdPedido WHERE pedidos.adIdCliente = ? GROUP BY product_name",
        [id]
      );

      return {
        success: true,
        payload: results,
      };
    },
    {
      query: t.Object({
        id: t.String(),
      }),
    }
  )
  .get(
    "/search/by-product",
    async ({ query }) => {
      const search = query.search;
      const searchLike = `%${search}%`;

      const [results] = await sql.execute(
        "SELECT pedidos.id AS id, pedidos.timestamp as order_timestamp, pedidos.idpedido AS order_id, COALESCE(NULLIF(clientes.adRazaoSocial, ''), clientes.adNomeFantasia) AS customer_name, pedidos.adValorTotal AS order_value, pedidos.adTipoPedido AS order_type FROM ppedidos pedidos LEFT JOIN ppedidosimples psimples ON pedidos.id = psimples.adIdPedido LEFT JOIN ppedidocomposto pcomposto ON pedidos.id = pcomposto.adIdPedido INNER JOIN pclientes clientes ON pedidos.adIdCliente = clientes.id WHERE LOWER(psimples.adReferencia) = ? OR LOWER(psimples.adDiscriminacao) LIKE ? OR LOWER(pcomposto.adModelo) LIKE ? GROUP BY pedidos.id ORDER BY pedidos.id DESC",
        [search, searchLike, searchLike]
      );

      return {
        success: true,
        payload: results,
      };
    },
    {
      query: t.Object({
        search: t.String(),
      }),
    }
  )
  .get(
    "/search/by-reference",
    async ({ query }) => {
      const search = query.search.toLowerCase();

      const [results] = await sql.execute(
        "SELECT adReferencia as reference, adUnidade AS unity, adDiscriminacao AS name, adIpi AS ipi FROM ppedidosimples WHERE adReferencia = LOWER(?) LIMIT 1",
        [search]
      );

      return {
        success: true,
        payload: (results as any)?.[0] ?? {},
      };
    },
    {
      query: t.Object({
        search: t.String({
          minLength: 3,
        }),
      }),
    }
  );
