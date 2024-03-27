import mysql from "mysql2/promise";

export const sql = await mysql.createConnection({
  host: process.env.DB_IP,
  port: +(process.env.DB_PORT ?? 8989),
  user: process.env.DB_USER,
  password: process.env.DB_PASSWORD,
  database: process.env.DB_NAME,
  maxIdle: 0,
  idleTimeout: 60000,
  enableKeepAlive: true,
});
