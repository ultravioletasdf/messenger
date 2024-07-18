import { Client } from "../mod";
const client = new Client();

(async function () {
  console.log(
    await client.users.create({ email: "abc@ga.com", password: "abc" })
  );
})();
