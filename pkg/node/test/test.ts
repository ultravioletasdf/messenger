import { Client } from "../mod";
const client = new Client();

(async function () {
  try {
    console.log(
      await client.users.create({ email: "abc@ga.com", password: "abc" })
    );
  } catch (error) {
    console.log(error);
  }
  try {
    console.log(await client.users.get({ id: 123 }));
  } catch (error) {
    console.log(error);
  }

  try {
    console.log(await client.users.get({ id: 50162252648448 }));
  } catch (error) {
    console.log(error);
  }
  try {
    console.log(
      await client.users.signIn({ email: "abc@ga.com", password: "abc" })
    );
  } catch (error) {
    console.log(error);
  }
})();
