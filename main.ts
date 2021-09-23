import { Application, Router } from "https://deno.land/x/oak@v9.0.1/mod.ts";
import { config } from "https://deno.land/x/dotenv@v3.0.0/mod.ts";

import { Webhooks } from "https://esm.sh/@octokit/webhooks";

const webhooks = new Webhooks({
  secret: config()["GH_WEBHOOK_SECRET"],
});

const router = new Router();
router.post("/payload", async (context) => {
  console.log("processing request");

  const payload = await context.request.body().value;
  const signature = context.request.headers.get("X-Hub-Signature") || "";

  const valid = await webhooks.verify(payload, signature);

  context.response.status = valid ? 200 : 500;
  context.response.body = valid ? "success :)" : "failure :(";
});

const app = new Application();
app.use(router.routes());
app.use(router.allowedMethods());

await app.listen({ port: 8000 });
