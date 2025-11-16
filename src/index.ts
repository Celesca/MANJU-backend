import { Elysia } from "elysia";
import { UserRepository } from "./repositories/userRepository";
import { Status } from "./models/user";

const repo = new UserRepository();

const app = new Elysia()
  .get("/", () => "Hello Elysia")

  // Create
  .post("/users", async ({ body }) => {
    try {
      const data = await body as any;
      if (!data?.email || !data?.name) {
        return { status: 400, body: { error: "email and name are required" } };
      }

      const user = await repo.create({
        email: String(data.email),
        name: String(data.name),
        info: data.info ?? null,
        status: data.status ?? Status.ACTIVE,
      });

      return { status: 201, body: user };
    } catch (err: any) {
      if (err?.message === "email_already_registered") {
        return { status: 400, body: { error: "email already registered" } };
      }
      return { status: 500, body: { error: err?.message ?? String(err) } };
    }
  })

  // Read list
  .get("/users", async () => {
    const users = await repo.list();
    return users;
  })

  // Read single
  .get("/users/:id", async ({ params }) => {
    const id = params.id as string;
    const user = await repo.getById(id);
    if (!user) return { status: 404, body: { error: "not_found" } };
    return user;
  })

  // Update (partial)
  .patch("/users/:id", async ({ params, body }) => {
    const id = params.id as string;
    try {
      const payload = await body as any;
      const updated = await repo.update(id, payload);
      if (!updated) return { status: 404, body: { error: "not_found" } };
      return updated;
    } catch (err: any) {
      if (err?.message === "email_already_registered") {
        return { status: 400, body: { error: "email already registered" } };
      }
      return { status: 500, body: { error: err?.message ?? String(err) } };
    }
  })

  // Delete
  .delete("/users/:id", async ({ params }) => {
    const id = params.id as string;
    const ok = await repo.delete(id);
    if (!ok) return { status: 404, body: { error: "not_found" } };
    return { status: 204 };
  })

  .listen(3000);

console.log(`🦊 Elysia is running at ${app.server?.hostname}:${app.server?.port}`);

export { app, repo };
